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

package authorization

import (
	"context"
	"fmt"

	v1 "go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/server/common/api"
)

type (
	defaultAuthorizer struct {
	}
)

var _ Authorizer = (*defaultAuthorizer)(nil)

// NewDefaultAuthorizer creates a default authorizer
func NewDefaultAuthorizer() Authorizer {
	return &defaultAuthorizer{}
}

var resultAllow = Result{Decision: DecisionAllow}
var resultDeny = Result{Decision: DecisionDeny}

// Authorize determines if an API call by given claims should be allowed or denied.
// Rules:
//
//	Health check APIs are allowed to everyone.
//	System Admin is allowed to access all APIs on all namespaces and cluster-level.
//	System Writer is allowed to access non admin APIs on all namespaces and cluster-level.
//	System Reader is allowed to access readonly APIs on all namespaces and cluster-level.
//	Namespace Admin is allowed to access all APIs on their namespaces.
//	Namespace Writer is allowed to access non admin APIs on their namespaces.
//	Namespace Reader is allowed to access non admin readonly APIs on their namespaces.
func (a *defaultAuthorizer) Authorize(_ context.Context, claims *Claims, target *CallTarget) (Result, error) {
	// APIs that are essentially read-only health checks with no sensitive information are
	// always allowed

	if IsHealthCheckAPI(target.APIName) {
		return resultAllow, nil
	}
	if claims == nil {
		return resultDeny, nil
	}

	metadata := api.GetMethodMetadata(target.APIName)

	var hasRole Role
	switch metadata.Scope {
	case api.ScopeCluster:
		hasRole = claims.System
	case api.ScopeNamespace:
		// Note: system-level claims apply across all namespaces.
		// Note: if claims.Namespace is nil or target.Namespace is not found, the lookup will return zero.
		// hasRole = claims.System | claims.Namespaces[target.Namespace]

		// check for namespace roles from CallTarget request interface{}
		// if claims.Extensions == nil {
		// 	return resultDeny, nil
		// }

		// extract namespace roles from extensions
		// namespaceRoles, ok := claims.Extensions.(NamespaceTaskQueueRoles)
		// if !ok {
		// 	return resultDeny, nil
		// }

		hasRole = claims.System | claims.Namespaces[target.Namespace]

		// if the call target api is part of task queue operation then check task queue role as well
		// if taskQueueRoles, ok := namespaceRoles[target.Namespace]; ok {
		// 	hasRole := claims.Namespaces[target.Namespace]

		// 	// check namespace role first
		// 	if hasRole >= getRequiredRole(metadata.Access) {
		// 		// then check task queue role
		// 		taskQueue := getTaskQueueFromRequest(target.Request)
		// 		hasTaskQueueRole := taskQueueRoles[taskQueue]

		// 		taskQueueMetaData := api.GetTaskQueueMetadata(taskQueue)

		// 		if hasTaskQueueRole >= getRequiredRole(taskQueueMetaData.Access) {
		// 			return resultAllow, nil
		// 		}
		// 	}
		// } else {
		// 	// then check task queue role
		// 	hasRole := claims.Namespaces[target.Namespace]

		// 	if hasRole >= getRequiredRole(metadata.Access) {
		// 		return resultAllow, nil
		// 	}
		// }
	default:
		return resultDeny, nil
	}

	if hasRole >= getRequiredRole(metadata.Access) {
		fmt.Printf("bill-authorizer-failed-for-claims %+v \n & target %+v \n & namespace %+v \n", claims, target.APIName, target.Namespace)
		return resultAllow, nil
	}

	return resultDeny, nil
}

func getTaskQueueFromRequest(request interface{}) string {
	switch req := request.(type) {
	case v1.PollWorkflowTaskQueueRequest:
		return req.TaskQueue.Name
	case v1.PollActivityTaskQueueRequest:
		return req.TaskQueue.Name
	case v1.DescribeTaskQueueRequest:
		return req.TaskQueue.Name
	case v1.StartWorkflowExecutionRequest:
		return req.TaskQueue.Name
	case v1.SignalWithStartWorkflowExecutionRequest:
		return req.TaskQueue.Name
	case v1.ListTaskQueuePartitionsRequest:
		return req.TaskQueue.Name
	default:
		return ""
	}
}

// Convert from api.Access to Role
func getRequiredRole(access api.Access) Role {
	switch access {
	case api.AccessReadOnly:
		return RoleReader
	case api.AccessWrite:
		return RoleWriter
	default:
		return RoleAdmin
	}
}
