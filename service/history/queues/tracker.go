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

package queues

import (
	"fmt"
	"maps"

	ctasks "go.temporal.io/server/common/tasks"
	"go.temporal.io/server/service/history/tasks"
)

type (
	// TODO: make task tracking a standalone component
	// currently it's used as a implementation detail in SliceImpl
	executableTracker struct {
		pendingExecutables map[tasks.Key]Executable
		groupClassifiers   map[string]func(tasks.Task) (any, bool)
		pendingPerGroup    map[string]map[any]int
	}
)

func newExecutableTracker(
	groupClassifiers map[string]func(tasks.Task) (any, bool),
) *executableTracker {
	pendingPerGroup := make(map[string]map[any]int, len(groupClassifiers))
	for k := range groupClassifiers {
		pendingPerGroup[k] = make(map[any]int, 0)
	}
	return &executableTracker{
		pendingExecutables: make(map[tasks.Key]Executable),
		groupClassifiers:   groupClassifiers,
		pendingPerGroup:    pendingPerGroup,
	}
}

func (t *executableTracker) split(
	thisScope Scope,
	thatScope Scope,
) (*executableTracker, *executableTracker) {
	that := executableTracker{
		pendingExecutables: make(map[tasks.Key]Executable, len(t.pendingExecutables)/2),
		groupClassifiers:   t.groupClassifiers,
		pendingPerGroup:    make(map[string]map[any]int, len(t.pendingPerGroup)),
	}
	for k := range t.pendingPerGroup {
		that.pendingPerGroup[k] = make(map[any]int, 0)
	}

	for key, executable := range t.pendingExecutables {
		if thisScope.Contains(executable) {
			continue
		}

		if !thatScope.Contains(executable) {
			panic(fmt.Sprintf("Queue slice encountered task doesn't belong to either scopes during split, scope: %v and %v, task: %v, task type: %v",
				thisScope, thatScope, executable.GetTask(), executable.GetType()))
		}

		delete(t.pendingExecutables, key)
		that.pendingExecutables[key] = executable

		for groupName, classifier := range t.groupClassifiers {
			groupKey, ok := classifier(executable)
			if !ok {
				continue
			}
			t.pendingPerGroup[groupName][groupKey]--
			that.pendingPerGroup[groupName][groupKey]++
		}
	}

	return t, &that
}

func (t *executableTracker) merge(incomingTracker *executableTracker) *executableTracker {
	thisExecutables, thisPendingTasks := t.pendingExecutables, t.pendingPerGroup
	thatExecutables, thatPendingTasks := incomingTracker.pendingExecutables, incomingTracker.pendingPerGroup
	if len(thisExecutables) < len(thatExecutables) {
		thisExecutables, thatExecutables = thatExecutables, thisExecutables
		thisPendingTasks = thatPendingTasks
	}

	maps.Copy(t.groupClassifiers, incomingTracker.groupClassifiers)
	maps.Copy(t.pendingExecutables, incomingTracker.pendingExecutables)

	for groupName := range t.groupClassifiers {
		pending, ok := thisPendingTasks[groupName]
		if !ok {
			thisPendingTasks[groupName] = thatPendingTasks[groupName]
			continue
		}
		maps.Copy(pending, thatPendingTasks[groupName])
	}

	// t.pendingExecutables = thisExecutables
	t.pendingPerGroup = thisPendingTasks
	return t
}

func (t *executableTracker) add(
	executable Executable,
) {
	t.pendingExecutables[executable.GetKey()] = executable
	for groupName, classifier := range t.groupClassifiers {
		groupKey, ok := classifier(executable)
		if !ok {
			continue
		}
		t.pendingPerGroup[groupName][groupKey]++
	}
}

func (t *executableTracker) shrink() tasks.Key {
	minPendingTaskKey := tasks.MaximumKey
	for key, executable := range t.pendingExecutables {
		if executable.State() == ctasks.TaskStateAcked {
			for groupName, classifier := range t.groupClassifiers {
				groupKey, ok := classifier(executable)
				if !ok {
					continue
				}
				t.pendingPerGroup[groupName][groupKey]--
			}
			delete(t.pendingExecutables, key)
			continue
		}

		minPendingTaskKey = tasks.MinKey(minPendingTaskKey, key)
	}

	for _, group := range t.pendingPerGroup {
		for groupKey, numPending := range group {
			if numPending == 0 {
				delete(group, groupKey)
			}
		}
	}

	return minPendingTaskKey
}

func (t *executableTracker) clear() {
	for _, executable := range t.pendingExecutables {
		executable.Cancel()
	}

	t.pendingExecutables = make(map[tasks.Key]Executable)
	t.pendingPerGroup = make(map[string]map[any]int, len(t.groupClassifiers))
}
