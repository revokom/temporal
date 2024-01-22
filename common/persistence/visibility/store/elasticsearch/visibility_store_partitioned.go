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

package elasticsearch

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/olivere/elastic/v7"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"

	"go.temporal.io/server/common/dynamicconfig"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/persistence/visibility/manager"
	"go.temporal.io/server/common/persistence/visibility/store"
	"go.temporal.io/server/common/persistence/visibility/store/elasticsearch/client"
	"go.temporal.io/server/common/persistence/visibility/store/query"
	"go.temporal.io/server/common/searchattribute"
)

const (
	PersistenceNamePartitioned = "elasticsearch_partitioned"
)

type (
	visibilityStorePartitioned struct {
		esClient                       client.Client
		baseIndexName                  string
		searchAttributesProvider       searchattribute.Provider
		searchAttributesMapperProvider searchattribute.MapperProvider
		namespaceRegistry              namespace.Registry
		processor                      Processor
		processorAckTimeout            dynamicconfig.DurationPropertyFn
		disableOrderByClause           dynamicconfig.BoolPropertyFnWithNamespaceFilter
		enableManualPagination         dynamicconfig.BoolPropertyFnWithNamespaceFilter
		metricsHandler                 metrics.Handler

		indexNameCache     map[string]bool
		lockIndexNameCache sync.RWMutex

		done chan bool
	}
)

var _ store.VisibilityStore = (*visibilityStorePartitioned)(nil)

// NewVisibilityStore create a visibility store connecting to ElasticSearch
func NewVisibilityStorePartitioned(
	esClient client.Client,
	baseIndexName string,
	searchAttributesProvider searchattribute.Provider,
	searchAttributesMapperProvider searchattribute.MapperProvider,
	namespaceRegistry namespace.Registry,
	processor Processor,
	processorAckTimeout dynamicconfig.DurationPropertyFn,
	disableOrderByClause dynamicconfig.BoolPropertyFnWithNamespaceFilter,
	enableManualPagination dynamicconfig.BoolPropertyFnWithNamespaceFilter,
	metricsHandler metrics.Handler,
) *visibilityStorePartitioned {
	store := &visibilityStorePartitioned{
		esClient:                       esClient,
		baseIndexName:                  baseIndexName,
		searchAttributesProvider:       searchAttributesProvider,
		searchAttributesMapperProvider: searchAttributesMapperProvider,
		namespaceRegistry:              namespaceRegistry,
		processor:                      processor,
		processorAckTimeout:            processorAckTimeout,
		disableOrderByClause:           disableOrderByClause,
		enableManualPagination:         enableManualPagination,
		metricsHandler:                 metricsHandler.WithTags(metrics.OperationTag(metrics.ElasticsearchVisibility)),
		indexNameCache:                 make(map[string]bool),
		done:                           make(chan bool),
	}

	if err := store.loadIndices(); err != nil {
		panic(fmt.Errorf("failed to load indices: %w", err))
	}

	go store.runCleanup()

	return store
}

func (s *visibilityStorePartitioned) loadIndices() error {
	resp, err := s.esClient.CatIndices(context.Background())
	if err != nil {
		return err
	}
	for _, item := range resp {
		s.indexNameCache[item.Index] = true
	}
	return nil
}

func (s *visibilityStorePartitioned) runCleanup() {
	period := 24 * time.Hour
	s.cleanupIndices(context.Background())
	for {
		now := time.Now()
		ch := time.After(now.Truncate(period).Add(period).Sub(now))
		select {
		case <-ch:
			s.cleanupIndices(context.Background())
		case <-s.done:
			return
		}
	}
}

func (s *visibilityStorePartitioned) cleanupIndices(ctx context.Context) error {
	resp, err := s.esClient.CatIndices(ctx)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(fmt.Sprintf(`^%s-closed-\d{4}-\d{2}-\d{2}$`, s.GetIndexName()))
	deletedIndices := []string{}
	today := time.Now().Truncate(24 * time.Hour)
	for _, item := range resp {
		if re.MatchString(item.Index) {
			tm, err := time.Parse(time.DateOnly, item.Index[len(item.Index)-len(time.DateOnly):])
			if err != nil {
				return err
			}
			if tm.Before(today) {
				if _, err := s.esClient.DeleteIndex(ctx, item.Index); err != nil {
					return err
				}
				deletedIndices = append(deletedIndices, item.Index)
			}
		}
	}
	s.lockIndexNameCache.Lock()
	defer s.lockIndexNameCache.Unlock()
	for _, indexName := range deletedIndices {
		delete(s.indexNameCache, indexName)
	}
	return nil
}

func (s *visibilityStorePartitioned) Close() {
	// TODO (alex): visibilityStore shouldn't Stop processor. Processor should be stopped where it is created.
	s.done <- true
	if s.processor != nil {
		s.processor.Stop()
	}
}

func (s *visibilityStorePartitioned) GetName() string {
	return PersistenceNamePartitioned
}

func (s *visibilityStorePartitioned) GetIndexName() string {
	return s.baseIndexName
}

func (s *visibilityStorePartitioned) ValidateCustomSearchAttributes(
	searchAttributes map[string]any,
) (map[string]any, error) {
	validatedSearchAttributes := make(map[string]any, len(searchAttributes))
	var invalidValueErrs []error
	for saName, saValue := range searchAttributes {
		var err error
		switch value := saValue.(type) {
		case time.Time:
			err = validateDatetime(value)
		case []time.Time:
			for _, item := range value {
				if err = validateDatetime(item); err != nil {
					break
				}
			}
		case string:
			err = validateString(value)
		case []string:
			for _, item := range value {
				if err = validateString(item); err != nil {
					break
				}
			}
		}
		if err != nil {
			invalidValueErrs = append(invalidValueErrs, err)
			continue
		}
		validatedSearchAttributes[saName] = saValue
	}
	var retError error
	if len(invalidValueErrs) > 0 {
		retError = store.NewVisibilityStoreInvalidValuesError(invalidValueErrs)
	}
	return validatedSearchAttributes, retError
}

func (s *visibilityStorePartitioned) checkIndexName(ctx context.Context, indexName string) error {
	s.lockIndexNameCache.RLock()
	cached := s.indexNameCache[indexName]
	s.lockIndexNameCache.RUnlock()
	if cached {
		return nil
	}
	exists, err := s.esClient.IndexExists(ctx, indexName)
	if err != nil {
		return err
	}
	if !exists {
		alias := map[string]any{
			"aliases": map[string]any{
				s.GetIndexName(): map[string]any{},
			},
		}
		if _, err := s.esClient.CreateIndex(ctx, indexName, alias); err != nil {
			return err
		}
		if _, err := s.esClient.WaitForYellowStatus(ctx, indexName); err != nil {
			return err
		}
	}
	s.lockIndexNameCache.Lock()
	defer s.lockIndexNameCache.Unlock()
	s.indexNameCache[indexName] = true
	return nil
}

func (s *visibilityStorePartitioned) getIndexForOpenWorkflow(
	ctx context.Context,
	namespaceID string,
) (string, error) {
	indexName := s.baseIndexName + "-open"
	if err := s.checkIndexName(ctx, indexName); err != nil {
		return "", err
	}
	return indexName, nil
}

func (s *visibilityStorePartitioned) getIndexForClosedWorkflow(
	ctx context.Context,
	namespaceID string,
	closeTime time.Time,
) (string, error) {
	ns, err := s.namespaceRegistry.GetNamespaceByID(namespace.ID(namespaceID))
	if err != nil {
		return "", err
	}
	indexName := s.baseIndexName + "-closed"
	if ns.Retention() > 0 {
		expire := closeTime.Add(ns.Retention())
		indexName = indexName + "-" + expire.Format(time.DateOnly)
	}
	if err := s.checkIndexName(ctx, indexName); err != nil {
		return "", err
	}
	return indexName, nil
}

func (s *visibilityStorePartitioned) RecordWorkflowExecutionStarted(
	ctx context.Context,
	request *store.InternalRecordWorkflowExecutionStartedRequest,
) error {
	index, err := s.getIndexForOpenWorkflow(ctx, request.NamespaceID)
	if err != nil {
		return err
	}

	visibilityTaskKey := getVisibilityTaskKey(request.ShardID, request.TaskID)
	doc, err := s.generateESDoc(request.InternalVisibilityRequestBase, visibilityTaskKey)
	if err != nil {
		return err
	}

	return s.addBulkIndexRequestAndWait(
		ctx,
		request.InternalVisibilityRequestBase,
		index,
		doc,
		visibilityTaskKey,
	)
}

func (s *visibilityStorePartitioned) RecordWorkflowExecutionClosed(
	ctx context.Context,
	request *store.InternalRecordWorkflowExecutionClosedRequest,
) error {
	indexOpenWf, err := s.getIndexForOpenWorkflow(ctx, request.NamespaceID)
	if err != nil {
		return err
	}
	indexClosedWf, err := s.getIndexForClosedWorkflow(ctx, request.NamespaceID, request.CloseTime)
	if err != nil {
		return err
	}

	visibilityTaskKey := getVisibilityTaskKey(request.ShardID, request.TaskID)
	doc, err := s.generateESDoc(request.InternalVisibilityRequestBase, visibilityTaskKey)
	if err != nil {
		return err
	}

	doc[searchattribute.CloseTime] = request.CloseTime
	doc[searchattribute.ExecutionDuration] = request.ExecutionDuration
	doc[searchattribute.HistoryLength] = request.HistoryLength
	doc[searchattribute.StateTransitionCount] = request.StateTransitionCount
	doc[searchattribute.HistorySizeBytes] = request.HistorySizeBytes

	err = s.addBulkIndexRequestAndWait(
		ctx,
		request.InternalVisibilityRequestBase,
		indexOpenWf,
		doc,
		visibilityTaskKey,
	)
	if err != nil {
		return err
	}

	err = s.addBulkIndexRequestAndWait(
		ctx,
		request.InternalVisibilityRequestBase,
		indexClosedWf,
		doc,
		visibilityTaskKey,
	)
	if err != nil {
		return err
	}

	return s.deleteWorkflowExecution(
		ctx,
		indexOpenWf,
		request.WorkflowID,
		request.RunID,
		request.TaskID+1, // need add one because of the update above
	)
}

func (s *visibilityStorePartitioned) UpsertWorkflowExecution(
	ctx context.Context,
	request *store.InternalUpsertWorkflowExecutionRequest,
) error {
	index, err := s.getIndexForOpenWorkflow(ctx, request.NamespaceID)
	if err != nil {
		return err
	}

	visibilityTaskKey := getVisibilityTaskKey(request.ShardID, request.TaskID)
	doc, err := s.generateESDoc(request.InternalVisibilityRequestBase, visibilityTaskKey)
	if err != nil {
		return err
	}

	return s.addBulkIndexRequestAndWait(
		ctx,
		request.InternalVisibilityRequestBase,
		index,
		doc,
		visibilityTaskKey,
	)
}

func (s *visibilityStorePartitioned) DeleteWorkflowExecution(
	ctx context.Context,
	request *manager.VisibilityDeleteWorkflowExecutionRequest,
) error {
	indexOpenWf, err := s.getIndexForOpenWorkflow(ctx, request.NamespaceID.String())
	if err != nil {
		return err
	}
	indexClosedWf, err := s.getIndexForClosedWorkflow(ctx, request.NamespaceID.String(), request.CloseTime)
	if err != nil {
		return err
	}

	err = s.deleteWorkflowExecution(
		ctx,
		indexClosedWf,
		request.WorkflowID,
		request.RunID,
		request.TaskID,
	)
	if err != nil {
		return err
	}

	return s.deleteWorkflowExecution(
		ctx,
		indexOpenWf,
		request.WorkflowID,
		request.RunID,
		request.TaskID,
	)
}

func (s *visibilityStorePartitioned) deleteWorkflowExecution(
	ctx context.Context,
	index string,
	workflowID string,
	runID string,
	taskID int64,
) error {
	docID := getDocID(workflowID, runID)
	bulkDeleteRequest := &client.BulkableRequest{
		Index:       index,
		ID:          docID,
		Version:     taskID,
		RequestType: client.BulkableRequestTypeDelete,
	}
	return s.addBulkRequestAndWait(ctx, bulkDeleteRequest, docID)
}

func (s *visibilityStorePartitioned) addBulkIndexRequestAndWait(
	ctx context.Context,
	request *store.InternalVisibilityRequestBase,
	index string,
	esDoc map[string]interface{},
	visibilityTaskKey string,
) error {
	bulkIndexRequest := &client.BulkableRequest{
		Index:       index,
		ID:          getDocID(request.WorkflowID, request.RunID),
		Version:     request.TaskID,
		RequestType: client.BulkableRequestTypeIndex,
		Doc:         esDoc,
	}

	return s.addBulkRequestAndWait(ctx, bulkIndexRequest, visibilityTaskKey)
}

func (s *visibilityStorePartitioned) addBulkRequestAndWait(
	_ context.Context,
	bulkRequest *client.BulkableRequest,
	visibilityTaskKey string,
) error {
	s.checkProcessor()

	// Add method is blocking. If bulk processor is busy flushing previous bulk, request will wait here.
	ackF := s.processor.Add(bulkRequest, visibilityTaskKey)

	// processorAckTimeout is a maximum duration for bulk processor to commit the bulk and unblock the `ackF`.
	// Default value is 30s and this timeout should never have happened,
	// because Elasticsearch must process a bulk within 30s.
	// Parent context is not respected here because it has shorter timeout (3s),
	// which might already expired here due to wait at Add method above.
	ctx, cancel := context.WithTimeout(context.Background(), s.processorAckTimeout())
	defer cancel()
	ack, err := ackF.Get(ctx)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return &persistence.TimeoutError{Msg: fmt.Sprintf("visibility task %s timed out waiting for ACK after %v", visibilityTaskKey, s.processorAckTimeout())}
		}
		// Returns non-retryable Internal error here because these errors are unexpected.
		// Visibility task processor retries all errors though, therefore new request will be generated for the same visibility task.
		return serviceerror.NewInternal(fmt.Sprintf("visibility task %s received error %v", visibilityTaskKey, err))
	}

	if !ack {
		// Returns retryable Unavailable error here because NACK from bulk processor
		// means that this request wasn't processed successfully and needs to be retried.
		// Visibility task processor retries all errors anyway, therefore new request will be generated for the same visibility task.
		return serviceerror.NewUnavailable(fmt.Sprintf("visibility task %s received NACK", visibilityTaskKey))
	}
	return nil
}

func (s *visibilityStorePartitioned) checkProcessor() {
	if s.processor == nil {
		// must be bug, check history setup
		panic("Elasticsearch processor is nil")
	}
	if s.processorAckTimeout == nil {
		// must be bug, check history setup
		panic("config.ESProcessorAckTimeout is nil")
	}
}

func (s *visibilityStorePartitioned) ListOpenWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf(
			"%s = '%s'",
			searchattribute.ExecutionStatus,
			enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		),
	}
	query := s.buildSearchQuery(request, baseConditions, true)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListClosedWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf(
			"%s != '%s'",
			searchattribute.ExecutionStatus,
			enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		),
	}
	query := s.buildSearchQuery(request, baseConditions, false)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListOpenWorkflowExecutionsByType(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByTypeRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf(
			"%s = '%s'",
			searchattribute.ExecutionStatus,
			enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		),
		fmt.Sprintf("%s = '%s'", searchattribute.WorkflowType, request.WorkflowTypeName),
	}
	query := s.buildSearchQuery(request.ListWorkflowExecutionsRequest, baseConditions, true)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListClosedWorkflowExecutionsByType(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByTypeRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf(
			"%s != '%s'",
			searchattribute.ExecutionStatus,
			enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		),
		fmt.Sprintf("%s = '%s'", searchattribute.WorkflowType, request.WorkflowTypeName),
	}
	query := s.buildSearchQuery(request.ListWorkflowExecutionsRequest, baseConditions, false)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListOpenWorkflowExecutionsByWorkflowID(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByWorkflowIDRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf(
			"%s = '%s'",
			searchattribute.ExecutionStatus,
			enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		),
		fmt.Sprintf("%s = '%s'", searchattribute.WorkflowID, request.WorkflowID),
	}
	query := s.buildSearchQuery(request.ListWorkflowExecutionsRequest, baseConditions, true)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListClosedWorkflowExecutionsByWorkflowID(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsByWorkflowIDRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf(
			"%s != '%s'",
			searchattribute.ExecutionStatus,
			enumspb.WORKFLOW_EXECUTION_STATUS_RUNNING,
		),
		fmt.Sprintf("%s = '%s'", searchattribute.WorkflowID, request.WorkflowID),
	}
	query := s.buildSearchQuery(request.ListWorkflowExecutionsRequest, baseConditions, false)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListClosedWorkflowExecutionsByStatus(
	ctx context.Context,
	request *manager.ListClosedWorkflowExecutionsByStatusRequest,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	baseConditions := []string{
		fmt.Sprintf("%s = '%s'", searchattribute.ExecutionStatus, request.Status),
	}
	query := s.buildSearchQuery(request.ListWorkflowExecutionsRequest, baseConditions, false)
	return s.ListWorkflowExecutions(
		ctx,
		&manager.ListWorkflowExecutionsRequestV2{
			NamespaceID:   request.NamespaceID,
			Namespace:     request.Namespace,
			PageSize:      request.PageSize,
			NextPageToken: request.NextPageToken,
			Query:         query,
		},
	)
}

func (s *visibilityStorePartitioned) ListWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	p, err := s.buildSearchParametersV2(ctx, request, s.getListFieldSorter)
	if err != nil {
		return nil, err
	}

	searchResult, err := s.esClient.Search(ctx, p)
	if err != nil {
		return nil, convertElasticsearchClientError("ListWorkflowExecutions failed", err)
	}

	return s.getListWorkflowExecutionsResponse(searchResult, request.Namespace, request.PageSize)
}

func (s *visibilityStorePartitioned) ScanWorkflowExecutions(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	// Point in time is only supported in Elasticsearch 7.10+ in default flavor.
	if s.esClient.IsPointInTimeSupported(ctx) {
		return s.scanWorkflowExecutionsWithPit(ctx, request)
	}
	return s.scanWorkflowExecutionsWithScroll(ctx, request)
}

func (s *visibilityStorePartitioned) scanWorkflowExecutionsWithScroll(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	var (
		searchResult *elastic.SearchResult
		scrollErr    error
	)

	p, err := s.buildSearchParametersV2(ctx, request, s.getScanFieldSorter)
	if err != nil {
		return nil, err
	}

	if len(request.NextPageToken) == 0 {
		searchResult, scrollErr = s.esClient.OpenScroll(ctx, p, scrollKeepAliveInterval)
	} else if p.ScrollID != "" {
		searchResult, scrollErr = s.esClient.Scroll(ctx, p.ScrollID, scrollKeepAliveInterval)
	} else {
		return nil, serviceerror.NewInvalidArgument("scrollId must present in pagination token")
	}

	if scrollErr != nil && scrollErr != io.EOF {
		return nil, convertElasticsearchClientError("ScanWorkflowExecutions failed", scrollErr)
	}

	// Both io.IOF and empty hits list indicate that this is a last page.
	if (searchResult.Hits != nil && len(searchResult.Hits.Hits) < request.PageSize) ||
		scrollErr == io.EOF {
		err := s.esClient.CloseScroll(ctx, searchResult.ScrollId)
		if err != nil {
			return nil, convertElasticsearchClientError("Unable to close scroll", err)
		}
	}

	return s.getListWorkflowExecutionsResponse(searchResult, request.Namespace, request.PageSize)
}

func (s *visibilityStorePartitioned) scanWorkflowExecutionsWithPit(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
) (*store.InternalListWorkflowExecutionsResponse, error) {
	p, err := s.buildSearchParametersV2(ctx, request, s.getScanFieldSorter)
	if err != nil {
		return nil, err
	}

	// First call doesn't have token with PointInTimeID.
	if len(request.NextPageToken) == 0 {
		pitID, err := s.esClient.OpenPointInTime(ctx, s.GetIndexName(), pointInTimeKeepAliveInterval)
		if err != nil {
			return nil, convertElasticsearchClientError("Unable to create point in time", err)
		}
		p.PointInTime = elastic.NewPointInTimeWithKeepAlive(pitID, pointInTimeKeepAliveInterval)
	} else if p.PointInTime == nil {
		return nil, serviceerror.NewInvalidArgument("pointInTimeId must present in pagination token")
	}

	searchResult, err := s.esClient.Search(ctx, p)
	if err != nil {
		return nil, convertElasticsearchClientError("ScanWorkflowExecutions failed", err)
	}

	// Number hits smaller than the page size indicate that this is the last page.
	if searchResult.Hits != nil && len(searchResult.Hits.Hits) < request.PageSize {
		_, err := s.esClient.ClosePointInTime(ctx, searchResult.PitId)
		if err != nil {
			return nil, convertElasticsearchClientError("Unable to close point in time", err)
		}
	}

	return s.getListWorkflowExecutionsResponse(searchResult, request.Namespace, request.PageSize)
}

func (s *visibilityStorePartitioned) CountWorkflowExecutions(
	ctx context.Context,
	request *manager.CountWorkflowExecutionsRequest,
) (*manager.CountWorkflowExecutionsResponse, error) {
	queryParams, err := s.convertQuery(request.Namespace, request.NamespaceID, request.Query)
	if err != nil {
		return nil, err
	}

	if len(queryParams.GroupBy) > 0 {
		return s.countGroupByWorkflowExecutions(ctx, request.NamespaceID, queryParams)
	}

	count, err := s.esClient.Count(ctx, s.GetIndexName(), queryParams.Query)
	if err != nil {
		return nil, convertElasticsearchClientError("CountWorkflowExecutions failed", err)
	}

	response := &manager.CountWorkflowExecutionsResponse{Count: count}
	return response, nil
}

func (s *visibilityStorePartitioned) countGroupByWorkflowExecutions(
	ctx context.Context,
	namespaceID namespace.ID,
	queryParams *query.QueryParams,
) (*manager.CountWorkflowExecutionsResponse, error) {
	groupByFields := queryParams.GroupBy

	// Elasticsearch aggregation is nested. so need to loop backwards to build it.
	// Example: when grouping by (field1, field2), the object looks like
	// {
	//   "aggs": {
	//     "field1": {
	//       "terms": {
	//         "field": "field1"
	//       },
	//       "aggs": {
	//         "field2": {
	//           "terms": {
	//             "field": "field2"
	//           }
	//         }
	//       }
	//     }
	//   }
	// }
	termsAgg := elastic.NewTermsAggregation().Field(groupByFields[len(groupByFields)-1])
	for i := len(groupByFields) - 2; i >= 0; i-- {
		termsAgg = elastic.NewTermsAggregation().
			Field(groupByFields[i]).
			SubAggregation(groupByFields[i+1], termsAgg)
	}
	esResponse, err := s.esClient.CountGroupBy(
		ctx,
		s.GetIndexName(),
		queryParams.Query,
		groupByFields[0],
		termsAgg,
	)
	if err != nil {
		return nil, err
	}
	return s.parseCountGroupByResponse(esResponse, groupByFields)
}

func (s *visibilityStorePartitioned) GetWorkflowExecution(
	ctx context.Context,
	request *manager.GetWorkflowExecutionRequest,
) (*store.InternalGetWorkflowExecutionResponse, error) {
	docID := getDocID(request.WorkflowID, request.RunID)
	result, err := s.esClient.Get(ctx, s.GetIndexName(), docID)
	if err != nil {
		return nil, convertElasticsearchClientError("GetWorkflowExecution failed", err)
	}

	typeMap, err := s.searchAttributesProvider.GetSearchAttributes(s.baseIndexName, false)
	if err != nil {
		return nil, serviceerror.NewUnavailable(
			fmt.Sprintf("Unable to read search attribute types: %v", err),
		)
	}

	if !result.Found {
		return nil, serviceerror.NewNotFound(
			fmt.Sprintf("Workflow execution with run id %s not found.", request.RunID),
		)
	}

	workflowExecutionInfo, err := s.parseESDoc(result.Id, result.Source, typeMap, request.Namespace)
	if err != nil {
		return nil, err
	}

	return &store.InternalGetWorkflowExecutionResponse{
		Execution: workflowExecutionInfo,
	}, nil
}

func (s *visibilityStorePartitioned) buildSearchQuery(
	request *manager.ListWorkflowExecutionsRequest,
	baseConditions []string,
	overStartTime bool,
) string {
	var conditions []string
	conditions = append(conditions, baseConditions...)

	if request.NamespaceDivision != "" {
		conditions = append(
			conditions,
			fmt.Sprintf(
				"%s = '%s'",
				searchattribute.TemporalNamespaceDivision,
				request.NamespaceDivision,
			),
		)
	}

	if !request.EarliestStartTime.IsZero() || !request.LatestStartTime.IsZero() {
		var tmSearchAttribute string
		if overStartTime {
			tmSearchAttribute = searchattribute.StartTime
		} else {
			tmSearchAttribute = searchattribute.CloseTime
		}
		if !request.EarliestStartTime.IsZero() && !request.LatestStartTime.IsZero() {
			conditions = append(
				conditions,
				fmt.Sprintf(
					"%s BETWEEN '%s' AND '%s'",
					tmSearchAttribute,
					request.EarliestStartTime.Format(time.RFC3339Nano),
					request.LatestStartTime.Format(time.RFC3339Nano),
				),
			)
		} else if !request.EarliestStartTime.IsZero() {
			conditions = append(
				conditions,
				fmt.Sprintf(
					"%s >= '%s'",
					tmSearchAttribute,
					request.EarliestStartTime.Format(time.RFC3339Nano),
				),
			)
		} else {
			conditions = append(
				conditions,
				fmt.Sprintf(
					"%s <= '%s'",
					tmSearchAttribute,
					request.LatestStartTime.Format(time.RFC3339Nano),
				),
			)
		}
	}

	return strings.Join(conditions, " AND ")
}

func (s *visibilityStorePartitioned) buildSearchParametersV2(
	ctx context.Context,
	request *manager.ListWorkflowExecutionsRequestV2,
	getFieldSorter func([]elastic.Sorter) ([]elastic.Sorter, error),
) (*client.SearchParameters, error) {
	queryParams, err := s.convertQuery(
		request.Namespace,
		request.NamespaceID,
		request.Query,
	)
	if err != nil {
		return nil, err
	}

	searchParams := &client.SearchParameters{
		Index:    s.GetIndexName(),
		PageSize: request.PageSize,
		Query:    queryParams.Query,
	}

	if len(queryParams.GroupBy) > 0 {
		return nil, serviceerror.NewInvalidArgument("GROUP BY clause is not supported")
	}

	// TODO(rodrigozhou): investigate possible solutions to slow ORDER BY.
	// ORDER BY clause can be slow if there is a large number of documents and
	// using a field that was not indexed by ES. Since slow queries can block
	// writes for unreasonably long, this option forbids the usage of ORDER BY
	// clause to prevent slow down issues.
	if s.disableOrderByClause(request.Namespace.String()) && len(queryParams.Sorter) > 0 {
		return nil, serviceerror.NewInvalidArgument("ORDER BY clause is not supported")
	}

	if len(queryParams.Sorter) > 0 {
		// If params.Sorter is not empty, then it's using custom order by.
		s.metricsHandler.WithTags(metrics.NamespaceTag(request.Namespace.String())).
			Counter(metrics.ElasticsearchCustomOrderByClauseCount.Name()).Record(1)
	}

	searchParams.Sorter, err = getFieldSorter(queryParams.Sorter)
	if err != nil {
		return nil, err
	}

	pageToken, err := s.deserializePageToken(request.NextPageToken)
	if err != nil {
		return nil, err
	}
	err = s.processPageToken(searchParams, pageToken, request.Namespace)
	if err != nil {
		return nil, err
	}

	return searchParams, nil
}

func (s *visibilityStorePartitioned) processPageToken(
	params *client.SearchParameters,
	pageToken *visibilityPageToken,
	namespaceName namespace.Name,
) error {
	if pageToken == nil {
		return nil
	}
	if pageToken.ScrollID != "" {
		params.ScrollID = pageToken.ScrollID
		return nil
	}
	if len(pageToken.SearchAfter) == 0 {
		return nil
	}
	if pageToken.PointInTimeID != "" {
		params.SearchAfter = pageToken.SearchAfter
		params.PointInTime = elastic.NewPointInTimeWithKeepAlive(
			pageToken.PointInTimeID,
			pointInTimeKeepAliveInterval,
		)
		return nil
	}
	if len(pageToken.SearchAfter) != len(params.Sorter) {
		return serviceerror.NewInvalidArgument(fmt.Sprintf(
			"Invalid page token for given sort fields: expected %d fields, got %d",
			len(params.Sorter),
			len(pageToken.SearchAfter),
		))
	}
	if !s.enableManualPagination(namespaceName.String()) || !isDefaultSorter(params.Sorter) {
		params.SearchAfter = pageToken.SearchAfter
		return nil
	}

	boolQuery, ok := params.Query.(*elastic.BoolQuery)
	if !ok {
		return serviceerror.NewInternal(fmt.Sprintf(
			"Unexpected query type: expected *elastic.BoolQuery, got %T",
			params.Query,
		))
	}

	saTypeMap, err := s.searchAttributesProvider.GetSearchAttributes(s.baseIndexName, false)
	if err != nil {
		return serviceerror.NewUnavailable(
			fmt.Sprintf("Unable to read search attribute types: %v", err),
		)
	}

	// build pagination search query for default sorter
	shouldQueries, err := buildPaginationQuery(defaultSorterFields, pageToken.SearchAfter, saTypeMap)
	if err != nil {
		return err
	}

	boolQuery.Should(shouldQueries...)
	boolQuery.MinimumNumberShouldMatch(1)
	return nil
}

func (s *visibilityStorePartitioned) convertQuery(
	namespace namespace.Name,
	namespaceID namespace.ID,
	requestQueryStr string,
) (*query.QueryParams, error) {
	saTypeMap, err := s.searchAttributesProvider.GetSearchAttributes(s.baseIndexName, false)
	if err != nil {
		return nil, serviceerror.NewUnavailable(fmt.Sprintf("Unable to read search attribute types: %v", err))
	}
	nameInterceptor := newNameInterceptor(namespace, saTypeMap, s.searchAttributesMapperProvider)
	queryConverter := newQueryConverter(
		nameInterceptor,
		NewValuesInterceptor(namespace, saTypeMap, s.searchAttributesMapperProvider),
	)
	queryParams, err := queryConverter.ConvertWhereOrderBy(requestQueryStr)
	if err != nil {
		// Convert ConverterError to InvalidArgument and pass through all other errors (which should be only mapper errors).
		var converterErr *query.ConverterError
		if errors.As(err, &converterErr) {
			return nil, converterErr.ToInvalidArgument()
		}
		return nil, err
	}

	filterQuery := elastic.NewBoolQuery()
	if queryParams.Query != nil {
		filterQuery.Filter(queryParams.Query)
	}

	// If the query did not explicitly filter on TemporalNamespaceDivision somehow, then add a
	// "must not exist" (i.e. "is null") query for it.
	if !nameInterceptor.seenNamespaceDivision {
		filterQuery.MustNot(elastic.NewExistsQuery(searchattribute.TemporalNamespaceDivision))
	}

	queryParams.Query = filterQuery
	return queryParams, nil
}

func (s *visibilityStorePartitioned) getScanFieldSorter(fieldSorts []elastic.Sorter) ([]elastic.Sorter, error) {
	// custom order is not supported by Scan API
	if len(fieldSorts) > 0 {
		return nil, serviceerror.NewInvalidArgument("ORDER BY clause is not supported")
	}

	return docSorter, nil
}

func (s *visibilityStorePartitioned) getListFieldSorter(fieldSorts []elastic.Sorter) ([]elastic.Sorter, error) {
	if len(fieldSorts) == 0 {
		return defaultSorter, nil
	}
	res := make([]elastic.Sorter, len(fieldSorts)+1)
	for i, fs := range fieldSorts {
		res[i] = fs
	}
	// RunID is explicit tiebreaker.
	res[len(res)-1] = elastic.NewFieldSort(searchattribute.RunID).Desc()

	return res, nil
}

func (s *visibilityStorePartitioned) getListWorkflowExecutionsResponse(
	searchResult *elastic.SearchResult,
	namespace namespace.Name,
	pageSize int,
) (*store.InternalListWorkflowExecutionsResponse, error) {

	if searchResult.Hits == nil || len(searchResult.Hits.Hits) == 0 {
		return &store.InternalListWorkflowExecutionsResponse{}, nil
	}

	typeMap, err := s.searchAttributesProvider.GetSearchAttributes(s.baseIndexName, false)
	if err != nil {
		return nil, serviceerror.NewUnavailable(fmt.Sprintf("Unable to read search attribute types: %v", err))
	}

	response := &store.InternalListWorkflowExecutionsResponse{
		Executions: make([]*store.InternalWorkflowExecutionInfo, 0, len(searchResult.Hits.Hits)),
	}
	var lastHitID string
	var lastHitSort []interface{}
	for _, hit := range searchResult.Hits.Hits {
		if hit.Id == lastHitID {
			continue
		}
		workflowExecutionInfo, err := s.parseESDoc(hit.Id, hit.Source, typeMap, namespace)
		if err != nil {
			return nil, err
		}
		response.Executions = append(response.Executions, workflowExecutionInfo)
		lastHitSort = hit.Sort
		lastHitID = hit.Id
	}

	if len(searchResult.Hits.Hits) == pageSize { // this means the response might not the last page
		response.NextPageToken, err = s.serializePageToken(&visibilityPageToken{
			SearchAfter:   lastHitSort,
			ScrollID:      searchResult.ScrollId,
			PointInTimeID: searchResult.PitId,
		})
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (s *visibilityStorePartitioned) deserializePageToken(data []byte) (*visibilityPageToken, error) {
	if len(data) == 0 {
		return nil, nil
	}

	var token *visibilityPageToken
	dec := json.NewDecoder(bytes.NewReader(data))
	// UseNumber will not lose precision on big int64.
	dec.UseNumber()
	err := dec.Decode(&token)
	if err != nil {
		return nil, serviceerror.NewInvalidArgument(fmt.Sprintf("unable to deserialize page token: %v", err))
	}
	return token, nil
}

func (s *visibilityStorePartitioned) serializePageToken(token *visibilityPageToken) ([]byte, error) {
	if token == nil {
		return nil, nil
	}

	data, err := json.Marshal(token)
	if err != nil {
		return nil, serviceerror.NewInternal(fmt.Sprintf("unable to serialize page token: %v", err))
	}
	return data, nil
}

func (s *visibilityStorePartitioned) generateESDoc(
	request *store.InternalVisibilityRequestBase,
	visibilityTaskKey string,
) (map[string]interface{}, error) {
	doc := map[string]interface{}{
		searchattribute.VisibilityTaskKey: visibilityTaskKey,
		searchattribute.NamespaceID:       request.NamespaceID,
		searchattribute.WorkflowID:        request.WorkflowID,
		searchattribute.RunID:             request.RunID,
		searchattribute.WorkflowType:      request.WorkflowTypeName,
		searchattribute.StartTime:         request.StartTime,
		searchattribute.ExecutionTime:     request.ExecutionTime,
		searchattribute.ExecutionStatus:   request.Status.String(),
		searchattribute.TaskQueue:         request.TaskQueue,
	}

	if request.ParentWorkflowID != nil {
		doc[searchattribute.ParentWorkflowID] = *request.ParentWorkflowID
	}
	if request.ParentRunID != nil {
		doc[searchattribute.ParentRunID] = *request.ParentRunID
	}

	if len(request.Memo.GetData()) > 0 {
		doc[searchattribute.Memo] = request.Memo.GetData()
		doc[searchattribute.MemoEncoding] = request.Memo.GetEncodingType().String()
	}

	typeMap, err := s.searchAttributesProvider.GetSearchAttributes(s.baseIndexName, false)
	if err != nil {
		s.metricsHandler.Counter(metrics.ElasticsearchDocumentGenerateFailuresCount.Name()).Record(1)
		return nil, serviceerror.NewUnavailable(fmt.Sprintf("Unable to read search attribute types: %v", err))
	}

	searchAttributes, err := searchattribute.Decode(request.SearchAttributes, &typeMap, true)
	if err != nil {
		s.metricsHandler.Counter(metrics.ElasticsearchDocumentGenerateFailuresCount.Name()).Record(1)
		return nil, serviceerror.NewInternal(fmt.Sprintf("Unable to decode search attributes: %v", err))
	}
	// This is to prevent existing tasks to fail indefinitely.
	// If it's only invalid values error, then silently continue without them.
	searchAttributes, err = s.ValidateCustomSearchAttributes(searchAttributes)
	if err != nil {
		if _, ok := err.(*store.VisibilityStoreInvalidValuesError); !ok {
			return nil, err
		}
	}
	for saName, saValue := range searchAttributes {
		if saValue == nil {
			// If search attribute value is `nil`, it means that it shouldn't be added to the document.
			// Empty slices are converted to `nil` while decoding.
			continue
		}
		doc[saName] = saValue
	}

	return doc, nil
}

func (s *visibilityStorePartitioned) parseESDoc(docID string, docSource json.RawMessage, saTypeMap searchattribute.NameTypeMap, namespace namespace.Name) (*store.InternalWorkflowExecutionInfo, error) {
	logParseError := func(fieldName string, fieldValue interface{}, err error, docID string) error {
		s.metricsHandler.Counter(metrics.ElasticsearchDocumentParseFailuresCount.Name()).Record(1)
		return serviceerror.NewInternal(fmt.Sprintf("Unable to parse Elasticsearch document(%s) %q field value %q: %v", docID, fieldName, fieldValue, err))
	}

	var sourceMap map[string]interface{}
	d := json.NewDecoder(bytes.NewReader(docSource))
	// Very important line. See finishParseJSONValue bellow.
	d.UseNumber()
	if err := d.Decode(&sourceMap); err != nil {
		s.metricsHandler.Counter(metrics.ElasticsearchDocumentParseFailuresCount.Name()).Record(1)
		return nil, serviceerror.NewInternal(fmt.Sprintf("Unable to unmarshal JSON from Elasticsearch document(%s): %v", docID, err))
	}

	var (
		isValidType            bool
		memo                   []byte
		memoEncoding           string
		customSearchAttributes map[string]interface{}
	)
	record := &store.InternalWorkflowExecutionInfo{}
	for fieldName, fieldValue := range sourceMap {
		switch fieldName {
		case searchattribute.NamespaceID,
			searchattribute.ExecutionDuration,
			searchattribute.VisibilityTaskKey:
			// Ignore these fields.
			continue
		case searchattribute.Memo:
			var memoStr string
			if memoStr, isValidType = fieldValue.(string); !isValidType {
				return nil, logParseError(fieldName, fieldValue, fmt.Errorf("%w: expected string got %T", errUnexpectedJSONFieldType, fieldValue), docID)
			}
			var err error
			if memo, err = base64.StdEncoding.DecodeString(memoStr); err != nil {
				return nil, logParseError(fieldName, memoStr[:10], err, docID)
			}
			continue
		case searchattribute.MemoEncoding:
			if memoEncoding, isValidType = fieldValue.(string); !isValidType {
				return nil, logParseError(fieldName, fieldValue, fmt.Errorf("%w: expected string got %T", errUnexpectedJSONFieldType, fieldValue), docID)
			}
			continue
		}

		fieldType, err := saTypeMap.GetType(fieldName)
		if err != nil {
			// Silently ignore ErrInvalidName because it indicates unknown field in Elasticsearch document.
			if errors.Is(err, searchattribute.ErrInvalidName) {
				continue
			}
			s.metricsHandler.Counter(metrics.ElasticsearchDocumentParseFailuresCount.Name()).Record(1)
			return nil, serviceerror.NewInternal(fmt.Sprintf("Unable to get type for Elasticsearch document(%s) field %q: %v", docID, fieldName, err))
		}

		fieldValueParsed, err := finishParseJSONValue(fieldValue, fieldType)
		if err != nil {
			return nil, logParseError(fieldName, fieldValue, err, docID)
		}

		switch fieldName {
		case searchattribute.WorkflowID:
			record.WorkflowID = fieldValueParsed.(string)
		case searchattribute.RunID:
			record.RunID = fieldValueParsed.(string)
		case searchattribute.WorkflowType:
			record.TypeName = fieldValue.(string)
		case searchattribute.StartTime:
			record.StartTime = fieldValueParsed.(time.Time)
		case searchattribute.ExecutionTime:
			record.ExecutionTime = fieldValueParsed.(time.Time)
		case searchattribute.CloseTime:
			record.CloseTime = fieldValueParsed.(time.Time)
		case searchattribute.TaskQueue:
			record.TaskQueue = fieldValueParsed.(string)
		case searchattribute.ExecutionStatus:
			status, err := enumspb.WorkflowExecutionStatusFromString(fieldValueParsed.(string))
			if err != nil {
				return nil, logParseError(fieldName, fieldValueParsed.(string), err, docID)
			}
			record.Status = status
		case searchattribute.HistoryLength:
			record.HistoryLength = fieldValueParsed.(int64)
		case searchattribute.StateTransitionCount:
			record.StateTransitionCount = fieldValueParsed.(int64)
		case searchattribute.HistorySizeBytes:
			record.HistorySizeBytes = fieldValueParsed.(int64)
		case searchattribute.ParentWorkflowID:
			record.ParentWorkflowID = fieldValueParsed.(string)
		case searchattribute.ParentRunID:
			record.ParentRunID = fieldValueParsed.(string)
		default:
			// All custom and predefined search attributes are handled here.
			if customSearchAttributes == nil {
				customSearchAttributes = map[string]interface{}{}
			}
			customSearchAttributes[fieldName] = fieldValueParsed
		}
	}

	if customSearchAttributes != nil {
		var err error
		record.SearchAttributes, err = searchattribute.Encode(customSearchAttributes, &saTypeMap)
		if err != nil {
			s.metricsHandler.Counter(metrics.ElasticsearchDocumentParseFailuresCount.Name()).Record(1)
			return nil, serviceerror.NewInternal(fmt.Sprintf("Unable to encode custom search attributes of Elasticsearch document(%s): %v", docID, err))
		}
		aliasedSas, err := searchattribute.AliasFields(s.searchAttributesMapperProvider, record.SearchAttributes, namespace.String())
		if err != nil {
			return nil, err
		}

		if aliasedSas != nil {
			record.SearchAttributes = aliasedSas
		}
	}

	if memoEncoding != "" {
		record.Memo = persistence.NewDataBlob(memo, memoEncoding)
	} else if memo != nil {
		s.metricsHandler.Counter(metrics.ElasticsearchDocumentParseFailuresCount.Name()).Record(1)
		return nil, serviceerror.NewInternal(fmt.Sprintf("%q field is missing in Elasticsearch document(%s)", searchattribute.MemoEncoding, docID))
	}

	return record, nil
}

// Elasticsearch aggregation groups are returned as nested object.
// This function flattens the response into rows.
//
//nolint:revive // cognitive complexity 27 (> max enabled 25)
func (s *visibilityStorePartitioned) parseCountGroupByResponse(
	searchResult *elastic.SearchResult,
	groupByFields []string,
) (*manager.CountWorkflowExecutionsResponse, error) {
	response := &manager.CountWorkflowExecutionsResponse{}
	typeMap, err := s.searchAttributesProvider.GetSearchAttributes(s.baseIndexName, false)
	if err != nil {
		return nil, serviceerror.NewUnavailable(
			fmt.Sprintf("Unable to read search attribute types: %v", err),
		)
	}
	groupByTypes := make([]enumspb.IndexedValueType, len(groupByFields))
	for i, saName := range groupByFields {
		tp, err := typeMap.GetType(saName)
		if err != nil {
			return nil, err
		}
		groupByTypes[i] = tp
	}

	parseJsonNumber := func(val any) (int64, error) {
		numberVal, isNumber := val.(json.Number)
		if !isNumber {
			return 0, fmt.Errorf("%w: expected json.Number, got %T", errUnexpectedJSONFieldType, val)
		}
		return numberVal.Int64()
	}

	var parseInternal func(map[string]any, []*commonpb.Payload) error
	parseInternal = func(aggs map[string]any, bucketValues []*commonpb.Payload) error {
		if len(bucketValues) == len(groupByFields) {
			cnt, err := parseJsonNumber(aggs["doc_count"])
			if err != nil {
				return fmt.Errorf("Unable to parse 'doc_count' field: %w", err)
			}
			groupValues := make([]*commonpb.Payload, len(groupByFields))
			for i := range bucketValues {
				groupValues[i] = bucketValues[i]
			}
			response.Groups = append(
				response.Groups,
				&workflowservice.CountWorkflowExecutionsResponse_AggregationGroup{
					GroupValues: groupValues,
					Count:       cnt,
				},
			)
			response.Count += cnt
			return nil
		}

		index := len(bucketValues)
		fieldName := groupByFields[index]
		buckets := aggs[fieldName].(map[string]any)["buckets"].([]any)
		for i := range buckets {
			bucket := buckets[i].(map[string]any)
			value, err := finishParseJSONValue(bucket["key"], groupByTypes[index])
			if err != nil {
				return fmt.Errorf("Failed to parse value %v: %w", bucket["key"], err)
			}
			payload, err := searchattribute.EncodeValue(value, groupByTypes[index])
			if err != nil {
				return fmt.Errorf("Failed to encode value %v: %w", value, err)
			}
			err = parseInternal(bucket, append(bucketValues, payload))
			if err != nil {
				return err
			}
		}
		return nil
	}

	var bucketsJson map[string]any
	dec := json.NewDecoder(bytes.NewReader(searchResult.Aggregations[groupByFields[0]]))
	dec.UseNumber()
	if err := dec.Decode(&bucketsJson); err != nil {
		return nil, serviceerror.NewInternal(fmt.Sprintf("unable to unmarshal json response: %v", err))
	}
	if err := parseInternal(map[string]any{groupByFields[0]: bucketsJson}, nil); err != nil {
		return nil, err
	}
	return response, nil
}
