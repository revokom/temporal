// The MIT License
//
// Copyright (c) 2024 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/nexus-rpc/sdk-go/nexus"
	"go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	failurepb "go.temporal.io/api/failure/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/metrics"
	commonnexus "go.temporal.io/server/common/nexus"
	"go.temporal.io/server/service/history/configs"
	"go.temporal.io/server/service/history/consts"
	"go.temporal.io/server/service/history/queues"
	"go.temporal.io/server/service/history/shard"
	"go.temporal.io/server/service/history/tasks"
	"go.temporal.io/server/service/history/workflow"
	wcache "go.temporal.io/server/service/history/workflow/cache"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type callbackQueueActiveTaskExecutor struct {
	shard             shard.Context
	workflowCache     wcache.Cache
	logger            log.Logger
	metricsHandler    metrics.Handler
	config            *configs.Config
	payloadSerializer commonnexus.PayloadSerializer
	clusterName       string
}

var _ queues.Executor = &callbackQueueActiveTaskExecutor{}

func newCallbackQueueActiveTaskExecutor(
	shard shard.Context,
	workflowCache wcache.Cache,
	logger log.Logger,
	metricsHandler metrics.Handler,
	config *configs.Config,
) *callbackQueueActiveTaskExecutor {
	return &callbackQueueActiveTaskExecutor{
		shard:          shard,
		workflowCache:  workflowCache,
		logger:         logger,
		metricsHandler: metricsHandler,
		config:         config,
		clusterName:    shard.GetClusterMetadata().GetCurrentClusterName(),
	}
}

func (t *callbackQueueActiveTaskExecutor) Execute(
	ctx context.Context,
	executable queues.Executable,
) queues.ExecuteResponse {
	task := executable.GetTask()
	taskType := "Active" + task.GetType().String()
	namespaceTag, replicationState := getNamespaceTagAndReplicationStateByID(
		t.shard.GetNamespaceRegistry(),
		task.GetNamespaceID(),
	)
	metricsTags := []metrics.Tag{
		namespaceTag,
		metrics.TaskTypeTag(taskType),
	}

	if replicationState == enumspb.REPLICATION_STATE_HANDOVER {
		// TODO: exclude task types here if we believe it's safe & necessary to execute
		// them during namespace handover.
		// TODO: move this logic to queues.Executable when metrics tag doesn't need to
		// be returned from task executor
		return queues.ExecuteResponse{
			ExecutionMetricTags: metricsTags,
			ExecutedAsActive:    true,
			ExecutionErr:        consts.ErrNamespaceHandover,
		}
	}

	var err error
	switch task := task.(type) {
	case *tasks.CallbackTask:
		err = t.processCallbackTask(ctx, task)
	default:
		err = errUnknownTransferTask
	}

	return queues.ExecuteResponse{
		ExecutionMetricTags: metricsTags,
		ExecutedAsActive:    true,
		ExecutionErr:        err,
	}
}

func (t *callbackQueueActiveTaskExecutor) processCallbackTask(
	ctx context.Context,
	task *tasks.CallbackTask,
) (retErr error) {
	weContext, release, err := getWorkflowExecutionContextForTask(ctx, t.shard, t.workflowCache, task)
	if err != nil {
		return err
	}
	defer func() { release(retErr) }()

	mutableState, err := loadMutableStateForTransferTask(ctx, t.shard, weContext, task, t.metricsHandler, t.logger)
	if err != nil {
		return err
	}
	if mutableState == nil {
		release(nil) // release(nil) so that the mutable state is not unloaded from cache
		return consts.ErrWorkflowExecutionNotFound
	}

	callbacks := mutableState.GetExecutionInfo().GetCallbacks()
	// TODO: do we need something like this?
	// err = CheckTaskVersion(t.shardContext, t.logger, mutableState.GetNamespaceEntry(), ai.Version, task.Version, task)

	if len(callbacks) < task.CallbackIndex+1 {
		// TODO: think about the error returned here
		return fmt.Errorf("invalid callback index for task")
	}

	callback := callbacks[task.CallbackIndex]

	switch variant := callback.GetCallback().GetVariant().(type) {
	case *common.Callback_Nexus_:
		completion, err := t.getNexusCompletion(ctx, mutableState)
		if err != nil {
			return err
		}
		release(nil)
		return t.processNexusCallbackTask(ctx, task, variant.Nexus.GetUrl(), completion)
	default:
		return fmt.Errorf("unprocessable callback variant: %v", variant)
	}
}

func (t *callbackQueueActiveTaskExecutor) getNexusCompletion(ctx context.Context, ms workflow.MutableState) (nexus.OperationCompletion, error) {
	ce, err := ms.GetCompletionEvent(ctx)
	// TODO: not sure which errors checks are appropriate
	if err != nil {
		return nil, err
	}

	switch ms.GetExecutionState().GetStatus() {
	case enumspb.WORKFLOW_EXECUTION_STATUS_COMPLETED:
		payloads := ce.GetWorkflowExecutionCompletedEventAttributes().GetResult().GetPayloads()
		content, err := t.payloadSerializer.Serialize(payloads[0])
		if err != nil {
			// TODO: think about the error returned here
			return nil, err
		}

		// TODO: this should be made easier in the Nexus SDK
		completion := &nexus.OperationCompletionSuccessful{
			Header: make(http.Header, len(content.Header)),
			Body:   bytes.NewReader(content.Data),
		}
		for k, v := range content.Header {
			completion.Header.Set("Content-"+k, v)
		}
		return completion, nil
	}
	return nil, fmt.Errorf("invalid workflow execution status: %v", ms.GetExecutionState().GetStatus())
}

func (t *callbackQueueActiveTaskExecutor) processNexusCallbackTask(ctx context.Context, task *tasks.CallbackTask, url string, completion nexus.OperationCompletion) error {
	requestCtx, cancel := context.WithTimeout(ctx, t.config.CallbackTaskTimeout())
	defer cancel()

	request, err := nexus.NewCompletionHTTPRequest(requestCtx, url, completion)
	if err != nil {
		// TODO: think about the error returned here
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		// TODO: think about the error returned here
		return err
	}
	return t.updateCallbackState(ctx, task, func(callback *workflowpb.CallbackInfo) {
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			callback.State = enumspb.CALLBACK_STATE_SUCCEEDED
			return
		}
		// TODO: get exact non retryable vs. retryable error codes
		if response.StatusCode >= 400 && response.StatusCode < 500 {
			callback.State = enumspb.CALLBACK_STATE_BACKING_OFF
			// TODO: schedule a backoff timer
		} else {
			callback.State = enumspb.CALLBACK_STATE_FAILED
		}

		callback.LastAttemptFailure = &failurepb.Failure{
			Message: response.Status,
			FailureInfo: &failurepb.Failure_ServerFailureInfo{
				ServerFailureInfo: &failurepb.ServerFailureInfo{
					NonRetryable: callback.State == enumspb.CALLBACK_STATE_BACKING_OFF,
				},
			},
		}
	})
}

func (t *callbackQueueActiveTaskExecutor) updateCallbackState(
	ctx context.Context,
	task *tasks.CallbackTask,
	updateCallbackFn func(*workflowpb.CallbackInfo),
) (retErr error) {
	weContext, release, err := getWorkflowExecutionContextForTask(ctx, t.shard, t.workflowCache, task)
	if err != nil {
		return err
	}
	defer func() { release(retErr) }()

	mutableState, err := loadMutableStateForTransferTask(ctx, t.shard, weContext, task, t.metricsHandler, t.logger)
	if err != nil {
		return err
	}
	if mutableState == nil {
		release(nil) // release(nil) so that the mutable state is not unloaded from cache
		return consts.ErrWorkflowExecutionNotFound
	}

	// TODO: This should probably move to mutable state
	callback := mutableState.GetExecutionInfo().GetCallbacks()[task.CallbackIndex]
	callback.Attempt++
	callback.LastAttemptCompleteTime = timestamppb.New(t.shard.GetCurrentTime(t.clusterName))
	updateCallbackFn(callback)
	// TODO: replication task

	return nil
}
