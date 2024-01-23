// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

package visibility

import (
	"context"
	"time"

	"go.temporal.io/server/common/dynamicconfig"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/persistence/visibility/manager"
)

type (
	readApiRequest interface {
		*manager.ListWorkflowExecutionsRequest |
			*manager.ListWorkflowExecutionsByTypeRequest |
			*manager.ListWorkflowExecutionsByWorkflowIDRequest |
			*manager.ListClosedWorkflowExecutionsByStatusRequest |
			*manager.ListWorkflowExecutionsRequestV2 |
			*manager.CountWorkflowExecutionsRequest |
			*manager.GetWorkflowExecutionRequest
	}

	VisibilityManagerDual struct {
		visibilityManager          manager.VisibilityManager
		secondaryVisibilityManager manager.VisibilityManager
		managerSelector            managerSelector

		shadowReads dynamicconfig.BoolPropertyFnWithNamespaceFilter
	}
)

var _ manager.VisibilityManager = (*VisibilityManagerDual)(nil)

const shadowRequestTimeout = 5 * time.Second

// NewVisibilityManagerDual create a visibility manager that operate on multiple manager
// implementations based on dynamic config.
func NewVisibilityManagerDual(
	visibilityManager manager.VisibilityManager,
	secondaryVisibilityManager manager.VisibilityManager,
	managerSelector managerSelector,
	shadowReads dynamicconfig.BoolPropertyFnWithNamespaceFilter,
) *VisibilityManagerDual {
	return &VisibilityManagerDual{
		visibilityManager:          visibilityManager,
		secondaryVisibilityManager: secondaryVisibilityManager,
		managerSelector:            managerSelector,

		shadowReads: shadowReads,
	}
}

func (v *VisibilityManagerDual) GetPrimaryVisibility() manager.VisibilityManager {
	return v.visibilityManager
}

func (v *VisibilityManagerDual) GetSecondaryVisibility() manager.VisibilityManager {
	return v.secondaryVisibilityManager
}

func (v *VisibilityManagerDual) Close() {
	v.visibilityManager.Close()
	v.secondaryVisibilityManager.Close()
}

func (v *VisibilityManagerDual) GetReadStoreName(nsName namespace.Name) string {
	return v.managerSelector.readManager(nsName).GetReadStoreName(nsName)
}

func (v *VisibilityManagerDual) GetStoreNames() []string {
	return append(v.visibilityManager.GetStoreNames(), v.secondaryVisibilityManager.GetStoreNames()...)
}

func (v *VisibilityManagerDual) HasStoreName(stName string) bool {
	for _, sn := range v.GetStoreNames() {
		if sn == stName {
			return true
		}
	}
	return false
}

func (v *VisibilityManagerDual) GetIndexName() string {
	return v.visibilityManager.GetIndexName()
}

func (v *VisibilityManagerDual) ValidateCustomSearchAttributes(
	searchAttributes map[string]any,
) (map[string]any, error) {
	ms, err := v.managerSelector.writeManagers()
	if err != nil {
		return nil, err
	}
	for _, m := range ms {
		searchAttributes, err = m.ValidateCustomSearchAttributes(searchAttributes)
		if err != nil {
			return nil, err
		}
	}
	return searchAttributes, nil
}

func (v *VisibilityManagerDual) RecordWorkflowExecutionStarted(
	ctx context.Context,
	request *manager.RecordWorkflowExecutionStartedRequest,
) error {
	ms, err := v.managerSelector.writeManagers()
	if err != nil {
		return err
	}
	for _, m := range ms {
		err = m.RecordWorkflowExecutionStarted(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VisibilityManagerDual) RecordWorkflowExecutionClosed(
	ctx context.Context,
	request *manager.RecordWorkflowExecutionClosedRequest,
) error {
	ms, err := v.managerSelector.writeManagers()
	if err != nil {
		return err
	}
	for _, m := range ms {
		err = m.RecordWorkflowExecutionClosed(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VisibilityManagerDual) UpsertWorkflowExecution(
	ctx context.Context,
	request *manager.UpsertWorkflowExecutionRequest,
) error {
	ms, err := v.managerSelector.writeManagers()
	if err != nil {
		return err
	}
	for _, m := range ms {
		err = m.UpsertWorkflowExecution(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VisibilityManagerDual) DeleteWorkflowExecution(
	ctx context.Context,
	request *manager.VisibilityDeleteWorkflowExecutionRequest,
) error {
	ms, err := v.managerSelector.writeManagers()
	if err != nil {
		return err
	}
	for _, m := range ms {
		err = m.DeleteWorkflowExecution(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VisibilityManagerDual) ListOpenWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListOpenWorkflowExecutions(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListOpenWorkflowExecutions, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListClosedWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListClosedWorkflowExecutions(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListClosedWorkflowExecutions, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListOpenWorkflowExecutionsByType(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByTypeRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListOpenWorkflowExecutionsByType(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListOpenWorkflowExecutionsByType, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListClosedWorkflowExecutionsByType(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByTypeRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListClosedWorkflowExecutionsByType(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListClosedWorkflowExecutionsByType, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListOpenWorkflowExecutionsByWorkflowID(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByWorkflowIDRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListOpenWorkflowExecutionsByWorkflowID(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListOpenWorkflowExecutionsByWorkflowID, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListClosedWorkflowExecutionsByWorkflowID(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByWorkflowIDRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListClosedWorkflowExecutionsByWorkflowID(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListClosedWorkflowExecutionsByWorkflowID, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListClosedWorkflowExecutionsByStatus(
	ctx context.Context,
	request *manager.ListClosedWorkflowExecutionsByStatusRequest,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListClosedWorkflowExecutionsByStatus(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListClosedWorkflowExecutionsByStatus, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ListWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ListWorkflowExecutions(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ListWorkflowExecutions, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) ScanWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
) (*manager.ListWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].ScanWorkflowExecutions(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].ScanWorkflowExecutions, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) CountWorkflowExecutions(
	ctx context.Context,
	request *manager.CountWorkflowExecutionsRequest,
) (*manager.CountWorkflowExecutionsResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].CountWorkflowExecutions(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].CountWorkflowExecutions, request)
	}
	return resp, err
}

func (v *VisibilityManagerDual) GetWorkflowExecution(
	ctx context.Context,
	request *manager.GetWorkflowExecutionRequest,
) (*manager.GetWorkflowExecutionResponse, error) {
	ms := v.managerSelector.readManagers(request.Namespace)
	resp, err := ms[0].GetWorkflowExecution(ctx, request)
	if len(ms) > 1 && v.shadowReads(request.Namespace.String()) {
		go sendShadowRequest(ms[1].GetWorkflowExecution, request)
	}
	return resp, err
}

func sendShadowRequest[Request readApiRequest, Response any](
	fn func(context.Context, Request) (Response, error),
	request Request,
) {
	ctx, cancel := context.WithTimeout(context.Background(), shadowRequestTimeout)
	defer cancel()
	_, _ = fn(ctx, request)
}
