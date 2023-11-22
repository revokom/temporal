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

package historyservice

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type StartWorkflowExecutionRequest to the protobuf v3 wire format
func (val *StartWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StartWorkflowExecutionRequest from the protobuf v3 wire format
func (val *StartWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StartWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartWorkflowExecutionRequest
	switch t := that.(type) {
	case *StartWorkflowExecutionRequest:
		that1 = t
	case StartWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StartWorkflowExecutionResponse to the protobuf v3 wire format
func (val *StartWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StartWorkflowExecutionResponse from the protobuf v3 wire format
func (val *StartWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StartWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartWorkflowExecutionResponse
	switch t := that.(type) {
	case *StartWorkflowExecutionResponse:
		that1 = t
	case StartWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetMutableStateRequest to the protobuf v3 wire format
func (val *GetMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetMutableStateRequest from the protobuf v3 wire format
func (val *GetMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetMutableStateRequest
	switch t := that.(type) {
	case *GetMutableStateRequest:
		that1 = t
	case GetMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetMutableStateResponse to the protobuf v3 wire format
func (val *GetMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetMutableStateResponse from the protobuf v3 wire format
func (val *GetMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetMutableStateResponse
	switch t := that.(type) {
	case *GetMutableStateResponse:
		that1 = t
	case GetMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollMutableStateRequest to the protobuf v3 wire format
func (val *PollMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollMutableStateRequest from the protobuf v3 wire format
func (val *PollMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollMutableStateRequest
	switch t := that.(type) {
	case *PollMutableStateRequest:
		that1 = t
	case PollMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollMutableStateResponse to the protobuf v3 wire format
func (val *PollMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollMutableStateResponse from the protobuf v3 wire format
func (val *PollMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollMutableStateResponse
	switch t := that.(type) {
	case *PollMutableStateResponse:
		that1 = t
	case PollMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetStickyTaskQueueRequest to the protobuf v3 wire format
func (val *ResetStickyTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetStickyTaskQueueRequest from the protobuf v3 wire format
func (val *ResetStickyTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetStickyTaskQueueRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetStickyTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetStickyTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetStickyTaskQueueRequest
	switch t := that.(type) {
	case *ResetStickyTaskQueueRequest:
		that1 = t
	case ResetStickyTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetStickyTaskQueueResponse to the protobuf v3 wire format
func (val *ResetStickyTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetStickyTaskQueueResponse from the protobuf v3 wire format
func (val *ResetStickyTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetStickyTaskQueueResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetStickyTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetStickyTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetStickyTaskQueueResponse
	switch t := that.(type) {
	case *ResetStickyTaskQueueResponse:
		that1 = t
	case ResetStickyTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordWorkflowTaskStartedRequest to the protobuf v3 wire format
func (val *RecordWorkflowTaskStartedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordWorkflowTaskStartedRequest from the protobuf v3 wire format
func (val *RecordWorkflowTaskStartedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordWorkflowTaskStartedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordWorkflowTaskStartedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordWorkflowTaskStartedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordWorkflowTaskStartedRequest
	switch t := that.(type) {
	case *RecordWorkflowTaskStartedRequest:
		that1 = t
	case RecordWorkflowTaskStartedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordWorkflowTaskStartedResponse to the protobuf v3 wire format
func (val *RecordWorkflowTaskStartedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordWorkflowTaskStartedResponse from the protobuf v3 wire format
func (val *RecordWorkflowTaskStartedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordWorkflowTaskStartedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordWorkflowTaskStartedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordWorkflowTaskStartedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordWorkflowTaskStartedResponse
	switch t := that.(type) {
	case *RecordWorkflowTaskStartedResponse:
		that1 = t
	case RecordWorkflowTaskStartedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordActivityTaskStartedRequest to the protobuf v3 wire format
func (val *RecordActivityTaskStartedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordActivityTaskStartedRequest from the protobuf v3 wire format
func (val *RecordActivityTaskStartedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordActivityTaskStartedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordActivityTaskStartedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskStartedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskStartedRequest
	switch t := that.(type) {
	case *RecordActivityTaskStartedRequest:
		that1 = t
	case RecordActivityTaskStartedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordActivityTaskStartedResponse to the protobuf v3 wire format
func (val *RecordActivityTaskStartedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordActivityTaskStartedResponse from the protobuf v3 wire format
func (val *RecordActivityTaskStartedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordActivityTaskStartedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordActivityTaskStartedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskStartedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskStartedResponse
	switch t := that.(type) {
	case *RecordActivityTaskStartedResponse:
		that1 = t
	case RecordActivityTaskStartedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondWorkflowTaskCompletedRequest to the protobuf v3 wire format
func (val *RespondWorkflowTaskCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondWorkflowTaskCompletedRequest from the protobuf v3 wire format
func (val *RespondWorkflowTaskCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondWorkflowTaskCompletedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondWorkflowTaskCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskCompletedRequest
	switch t := that.(type) {
	case *RespondWorkflowTaskCompletedRequest:
		that1 = t
	case RespondWorkflowTaskCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondWorkflowTaskCompletedResponse to the protobuf v3 wire format
func (val *RespondWorkflowTaskCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondWorkflowTaskCompletedResponse from the protobuf v3 wire format
func (val *RespondWorkflowTaskCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondWorkflowTaskCompletedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondWorkflowTaskCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskCompletedResponse
	switch t := that.(type) {
	case *RespondWorkflowTaskCompletedResponse:
		that1 = t
	case RespondWorkflowTaskCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondWorkflowTaskFailedRequest to the protobuf v3 wire format
func (val *RespondWorkflowTaskFailedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondWorkflowTaskFailedRequest from the protobuf v3 wire format
func (val *RespondWorkflowTaskFailedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondWorkflowTaskFailedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondWorkflowTaskFailedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskFailedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskFailedRequest
	switch t := that.(type) {
	case *RespondWorkflowTaskFailedRequest:
		that1 = t
	case RespondWorkflowTaskFailedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondWorkflowTaskFailedResponse to the protobuf v3 wire format
func (val *RespondWorkflowTaskFailedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondWorkflowTaskFailedResponse from the protobuf v3 wire format
func (val *RespondWorkflowTaskFailedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondWorkflowTaskFailedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondWorkflowTaskFailedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskFailedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskFailedResponse
	switch t := that.(type) {
	case *RespondWorkflowTaskFailedResponse:
		that1 = t
	case RespondWorkflowTaskFailedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IsWorkflowTaskValidRequest to the protobuf v3 wire format
func (val *IsWorkflowTaskValidRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IsWorkflowTaskValidRequest from the protobuf v3 wire format
func (val *IsWorkflowTaskValidRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IsWorkflowTaskValidRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IsWorkflowTaskValidRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IsWorkflowTaskValidRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IsWorkflowTaskValidRequest
	switch t := that.(type) {
	case *IsWorkflowTaskValidRequest:
		that1 = t
	case IsWorkflowTaskValidRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IsWorkflowTaskValidResponse to the protobuf v3 wire format
func (val *IsWorkflowTaskValidResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IsWorkflowTaskValidResponse from the protobuf v3 wire format
func (val *IsWorkflowTaskValidResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IsWorkflowTaskValidResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IsWorkflowTaskValidResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IsWorkflowTaskValidResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IsWorkflowTaskValidResponse
	switch t := that.(type) {
	case *IsWorkflowTaskValidResponse:
		that1 = t
	case IsWorkflowTaskValidResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordActivityTaskHeartbeatRequest to the protobuf v3 wire format
func (val *RecordActivityTaskHeartbeatRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordActivityTaskHeartbeatRequest from the protobuf v3 wire format
func (val *RecordActivityTaskHeartbeatRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordActivityTaskHeartbeatRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordActivityTaskHeartbeatRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskHeartbeatRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskHeartbeatRequest
	switch t := that.(type) {
	case *RecordActivityTaskHeartbeatRequest:
		that1 = t
	case RecordActivityTaskHeartbeatRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordActivityTaskHeartbeatResponse to the protobuf v3 wire format
func (val *RecordActivityTaskHeartbeatResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordActivityTaskHeartbeatResponse from the protobuf v3 wire format
func (val *RecordActivityTaskHeartbeatResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordActivityTaskHeartbeatResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordActivityTaskHeartbeatResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskHeartbeatResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskHeartbeatResponse
	switch t := that.(type) {
	case *RecordActivityTaskHeartbeatResponse:
		that1 = t
	case RecordActivityTaskHeartbeatResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondActivityTaskCompletedRequest to the protobuf v3 wire format
func (val *RespondActivityTaskCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondActivityTaskCompletedRequest from the protobuf v3 wire format
func (val *RespondActivityTaskCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondActivityTaskCompletedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondActivityTaskCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCompletedRequest
	switch t := that.(type) {
	case *RespondActivityTaskCompletedRequest:
		that1 = t
	case RespondActivityTaskCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondActivityTaskCompletedResponse to the protobuf v3 wire format
func (val *RespondActivityTaskCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondActivityTaskCompletedResponse from the protobuf v3 wire format
func (val *RespondActivityTaskCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondActivityTaskCompletedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondActivityTaskCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCompletedResponse
	switch t := that.(type) {
	case *RespondActivityTaskCompletedResponse:
		that1 = t
	case RespondActivityTaskCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondActivityTaskFailedRequest to the protobuf v3 wire format
func (val *RespondActivityTaskFailedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondActivityTaskFailedRequest from the protobuf v3 wire format
func (val *RespondActivityTaskFailedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondActivityTaskFailedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondActivityTaskFailedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskFailedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskFailedRequest
	switch t := that.(type) {
	case *RespondActivityTaskFailedRequest:
		that1 = t
	case RespondActivityTaskFailedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondActivityTaskFailedResponse to the protobuf v3 wire format
func (val *RespondActivityTaskFailedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondActivityTaskFailedResponse from the protobuf v3 wire format
func (val *RespondActivityTaskFailedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondActivityTaskFailedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondActivityTaskFailedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskFailedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskFailedResponse
	switch t := that.(type) {
	case *RespondActivityTaskFailedResponse:
		that1 = t
	case RespondActivityTaskFailedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondActivityTaskCanceledRequest to the protobuf v3 wire format
func (val *RespondActivityTaskCanceledRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondActivityTaskCanceledRequest from the protobuf v3 wire format
func (val *RespondActivityTaskCanceledRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondActivityTaskCanceledRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondActivityTaskCanceledRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCanceledRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCanceledRequest
	switch t := that.(type) {
	case *RespondActivityTaskCanceledRequest:
		that1 = t
	case RespondActivityTaskCanceledRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RespondActivityTaskCanceledResponse to the protobuf v3 wire format
func (val *RespondActivityTaskCanceledResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RespondActivityTaskCanceledResponse from the protobuf v3 wire format
func (val *RespondActivityTaskCanceledResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RespondActivityTaskCanceledResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RespondActivityTaskCanceledResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCanceledResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCanceledResponse
	switch t := that.(type) {
	case *RespondActivityTaskCanceledResponse:
		that1 = t
	case RespondActivityTaskCanceledResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IsActivityTaskValidRequest to the protobuf v3 wire format
func (val *IsActivityTaskValidRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IsActivityTaskValidRequest from the protobuf v3 wire format
func (val *IsActivityTaskValidRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IsActivityTaskValidRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IsActivityTaskValidRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IsActivityTaskValidRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IsActivityTaskValidRequest
	switch t := that.(type) {
	case *IsActivityTaskValidRequest:
		that1 = t
	case IsActivityTaskValidRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IsActivityTaskValidResponse to the protobuf v3 wire format
func (val *IsActivityTaskValidResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IsActivityTaskValidResponse from the protobuf v3 wire format
func (val *IsActivityTaskValidResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IsActivityTaskValidResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IsActivityTaskValidResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IsActivityTaskValidResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IsActivityTaskValidResponse
	switch t := that.(type) {
	case *IsActivityTaskValidResponse:
		that1 = t
	case IsActivityTaskValidResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalWorkflowExecutionRequest to the protobuf v3 wire format
func (val *SignalWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalWorkflowExecutionRequest from the protobuf v3 wire format
func (val *SignalWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWorkflowExecutionRequest
	switch t := that.(type) {
	case *SignalWorkflowExecutionRequest:
		that1 = t
	case SignalWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalWorkflowExecutionResponse to the protobuf v3 wire format
func (val *SignalWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalWorkflowExecutionResponse from the protobuf v3 wire format
func (val *SignalWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWorkflowExecutionResponse
	switch t := that.(type) {
	case *SignalWorkflowExecutionResponse:
		that1 = t
	case SignalWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalWithStartWorkflowExecutionRequest to the protobuf v3 wire format
func (val *SignalWithStartWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalWithStartWorkflowExecutionRequest from the protobuf v3 wire format
func (val *SignalWithStartWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalWithStartWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalWithStartWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWithStartWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWithStartWorkflowExecutionRequest
	switch t := that.(type) {
	case *SignalWithStartWorkflowExecutionRequest:
		that1 = t
	case SignalWithStartWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalWithStartWorkflowExecutionResponse to the protobuf v3 wire format
func (val *SignalWithStartWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalWithStartWorkflowExecutionResponse from the protobuf v3 wire format
func (val *SignalWithStartWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalWithStartWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalWithStartWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWithStartWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWithStartWorkflowExecutionResponse
	switch t := that.(type) {
	case *SignalWithStartWorkflowExecutionResponse:
		that1 = t
	case SignalWithStartWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveSignalMutableStateRequest to the protobuf v3 wire format
func (val *RemoveSignalMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveSignalMutableStateRequest from the protobuf v3 wire format
func (val *RemoveSignalMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveSignalMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveSignalMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveSignalMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveSignalMutableStateRequest
	switch t := that.(type) {
	case *RemoveSignalMutableStateRequest:
		that1 = t
	case RemoveSignalMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveSignalMutableStateResponse to the protobuf v3 wire format
func (val *RemoveSignalMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveSignalMutableStateResponse from the protobuf v3 wire format
func (val *RemoveSignalMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveSignalMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveSignalMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveSignalMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveSignalMutableStateResponse
	switch t := that.(type) {
	case *RemoveSignalMutableStateResponse:
		that1 = t
	case RemoveSignalMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TerminateWorkflowExecutionRequest to the protobuf v3 wire format
func (val *TerminateWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TerminateWorkflowExecutionRequest from the protobuf v3 wire format
func (val *TerminateWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TerminateWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TerminateWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TerminateWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TerminateWorkflowExecutionRequest
	switch t := that.(type) {
	case *TerminateWorkflowExecutionRequest:
		that1 = t
	case TerminateWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TerminateWorkflowExecutionResponse to the protobuf v3 wire format
func (val *TerminateWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TerminateWorkflowExecutionResponse from the protobuf v3 wire format
func (val *TerminateWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TerminateWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TerminateWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TerminateWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TerminateWorkflowExecutionResponse
	switch t := that.(type) {
	case *TerminateWorkflowExecutionResponse:
		that1 = t
	case TerminateWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteWorkflowExecutionRequest to the protobuf v3 wire format
func (val *DeleteWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteWorkflowExecutionRequest from the protobuf v3 wire format
func (val *DeleteWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowExecutionRequest
	switch t := that.(type) {
	case *DeleteWorkflowExecutionRequest:
		that1 = t
	case DeleteWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteWorkflowExecutionResponse to the protobuf v3 wire format
func (val *DeleteWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteWorkflowExecutionResponse from the protobuf v3 wire format
func (val *DeleteWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowExecutionResponse
	switch t := that.(type) {
	case *DeleteWorkflowExecutionResponse:
		that1 = t
	case DeleteWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetWorkflowExecutionRequest to the protobuf v3 wire format
func (val *ResetWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetWorkflowExecutionRequest from the protobuf v3 wire format
func (val *ResetWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetWorkflowExecutionRequest
	switch t := that.(type) {
	case *ResetWorkflowExecutionRequest:
		that1 = t
	case ResetWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetWorkflowExecutionResponse to the protobuf v3 wire format
func (val *ResetWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetWorkflowExecutionResponse from the protobuf v3 wire format
func (val *ResetWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetWorkflowExecutionResponse
	switch t := that.(type) {
	case *ResetWorkflowExecutionResponse:
		that1 = t
	case ResetWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RequestCancelWorkflowExecutionRequest to the protobuf v3 wire format
func (val *RequestCancelWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RequestCancelWorkflowExecutionRequest from the protobuf v3 wire format
func (val *RequestCancelWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RequestCancelWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelWorkflowExecutionRequest
	switch t := that.(type) {
	case *RequestCancelWorkflowExecutionRequest:
		that1 = t
	case RequestCancelWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RequestCancelWorkflowExecutionResponse to the protobuf v3 wire format
func (val *RequestCancelWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RequestCancelWorkflowExecutionResponse from the protobuf v3 wire format
func (val *RequestCancelWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RequestCancelWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelWorkflowExecutionResponse
	switch t := that.(type) {
	case *RequestCancelWorkflowExecutionResponse:
		that1 = t
	case RequestCancelWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ScheduleWorkflowTaskRequest to the protobuf v3 wire format
func (val *ScheduleWorkflowTaskRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ScheduleWorkflowTaskRequest from the protobuf v3 wire format
func (val *ScheduleWorkflowTaskRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ScheduleWorkflowTaskRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ScheduleWorkflowTaskRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleWorkflowTaskRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleWorkflowTaskRequest
	switch t := that.(type) {
	case *ScheduleWorkflowTaskRequest:
		that1 = t
	case ScheduleWorkflowTaskRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ScheduleWorkflowTaskResponse to the protobuf v3 wire format
func (val *ScheduleWorkflowTaskResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ScheduleWorkflowTaskResponse from the protobuf v3 wire format
func (val *ScheduleWorkflowTaskResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ScheduleWorkflowTaskResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ScheduleWorkflowTaskResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleWorkflowTaskResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleWorkflowTaskResponse
	switch t := that.(type) {
	case *ScheduleWorkflowTaskResponse:
		that1 = t
	case ScheduleWorkflowTaskResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VerifyFirstWorkflowTaskScheduledRequest to the protobuf v3 wire format
func (val *VerifyFirstWorkflowTaskScheduledRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VerifyFirstWorkflowTaskScheduledRequest from the protobuf v3 wire format
func (val *VerifyFirstWorkflowTaskScheduledRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VerifyFirstWorkflowTaskScheduledRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VerifyFirstWorkflowTaskScheduledRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VerifyFirstWorkflowTaskScheduledRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VerifyFirstWorkflowTaskScheduledRequest
	switch t := that.(type) {
	case *VerifyFirstWorkflowTaskScheduledRequest:
		that1 = t
	case VerifyFirstWorkflowTaskScheduledRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VerifyFirstWorkflowTaskScheduledResponse to the protobuf v3 wire format
func (val *VerifyFirstWorkflowTaskScheduledResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VerifyFirstWorkflowTaskScheduledResponse from the protobuf v3 wire format
func (val *VerifyFirstWorkflowTaskScheduledResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VerifyFirstWorkflowTaskScheduledResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VerifyFirstWorkflowTaskScheduledResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VerifyFirstWorkflowTaskScheduledResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VerifyFirstWorkflowTaskScheduledResponse
	switch t := that.(type) {
	case *VerifyFirstWorkflowTaskScheduledResponse:
		that1 = t
	case VerifyFirstWorkflowTaskScheduledResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordChildExecutionCompletedRequest to the protobuf v3 wire format
func (val *RecordChildExecutionCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordChildExecutionCompletedRequest from the protobuf v3 wire format
func (val *RecordChildExecutionCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordChildExecutionCompletedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordChildExecutionCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordChildExecutionCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordChildExecutionCompletedRequest
	switch t := that.(type) {
	case *RecordChildExecutionCompletedRequest:
		that1 = t
	case RecordChildExecutionCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RecordChildExecutionCompletedResponse to the protobuf v3 wire format
func (val *RecordChildExecutionCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RecordChildExecutionCompletedResponse from the protobuf v3 wire format
func (val *RecordChildExecutionCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RecordChildExecutionCompletedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordChildExecutionCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordChildExecutionCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordChildExecutionCompletedResponse
	switch t := that.(type) {
	case *RecordChildExecutionCompletedResponse:
		that1 = t
	case RecordChildExecutionCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VerifyChildExecutionCompletionRecordedRequest to the protobuf v3 wire format
func (val *VerifyChildExecutionCompletionRecordedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VerifyChildExecutionCompletionRecordedRequest from the protobuf v3 wire format
func (val *VerifyChildExecutionCompletionRecordedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VerifyChildExecutionCompletionRecordedRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VerifyChildExecutionCompletionRecordedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VerifyChildExecutionCompletionRecordedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VerifyChildExecutionCompletionRecordedRequest
	switch t := that.(type) {
	case *VerifyChildExecutionCompletionRecordedRequest:
		that1 = t
	case VerifyChildExecutionCompletionRecordedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VerifyChildExecutionCompletionRecordedResponse to the protobuf v3 wire format
func (val *VerifyChildExecutionCompletionRecordedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VerifyChildExecutionCompletionRecordedResponse from the protobuf v3 wire format
func (val *VerifyChildExecutionCompletionRecordedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VerifyChildExecutionCompletionRecordedResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VerifyChildExecutionCompletionRecordedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VerifyChildExecutionCompletionRecordedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VerifyChildExecutionCompletionRecordedResponse
	switch t := that.(type) {
	case *VerifyChildExecutionCompletionRecordedResponse:
		that1 = t
	case VerifyChildExecutionCompletionRecordedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeWorkflowExecutionRequest to the protobuf v3 wire format
func (val *DescribeWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeWorkflowExecutionRequest from the protobuf v3 wire format
func (val *DescribeWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeWorkflowExecutionRequest
	switch t := that.(type) {
	case *DescribeWorkflowExecutionRequest:
		that1 = t
	case DescribeWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeWorkflowExecutionResponse to the protobuf v3 wire format
func (val *DescribeWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeWorkflowExecutionResponse from the protobuf v3 wire format
func (val *DescribeWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeWorkflowExecutionResponse
	switch t := that.(type) {
	case *DescribeWorkflowExecutionResponse:
		that1 = t
	case DescribeWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicateEventsV2Request to the protobuf v3 wire format
func (val *ReplicateEventsV2Request) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicateEventsV2Request from the protobuf v3 wire format
func (val *ReplicateEventsV2Request) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicateEventsV2Request) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicateEventsV2Request values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicateEventsV2Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicateEventsV2Request
	switch t := that.(type) {
	case *ReplicateEventsV2Request:
		that1 = t
	case ReplicateEventsV2Request:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicateEventsV2Response to the protobuf v3 wire format
func (val *ReplicateEventsV2Response) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicateEventsV2Response from the protobuf v3 wire format
func (val *ReplicateEventsV2Response) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicateEventsV2Response) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicateEventsV2Response values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicateEventsV2Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicateEventsV2Response
	switch t := that.(type) {
	case *ReplicateEventsV2Response:
		that1 = t
	case ReplicateEventsV2Response:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicateWorkflowStateRequest to the protobuf v3 wire format
func (val *ReplicateWorkflowStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicateWorkflowStateRequest from the protobuf v3 wire format
func (val *ReplicateWorkflowStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicateWorkflowStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicateWorkflowStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicateWorkflowStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicateWorkflowStateRequest
	switch t := that.(type) {
	case *ReplicateWorkflowStateRequest:
		that1 = t
	case ReplicateWorkflowStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicateWorkflowStateResponse to the protobuf v3 wire format
func (val *ReplicateWorkflowStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicateWorkflowStateResponse from the protobuf v3 wire format
func (val *ReplicateWorkflowStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicateWorkflowStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicateWorkflowStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicateWorkflowStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicateWorkflowStateResponse
	switch t := that.(type) {
	case *ReplicateWorkflowStateResponse:
		that1 = t
	case ReplicateWorkflowStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SyncShardStatusRequest to the protobuf v3 wire format
func (val *SyncShardStatusRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SyncShardStatusRequest from the protobuf v3 wire format
func (val *SyncShardStatusRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SyncShardStatusRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SyncShardStatusRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SyncShardStatusRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SyncShardStatusRequest
	switch t := that.(type) {
	case *SyncShardStatusRequest:
		that1 = t
	case SyncShardStatusRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SyncShardStatusResponse to the protobuf v3 wire format
func (val *SyncShardStatusResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SyncShardStatusResponse from the protobuf v3 wire format
func (val *SyncShardStatusResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SyncShardStatusResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SyncShardStatusResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SyncShardStatusResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SyncShardStatusResponse
	switch t := that.(type) {
	case *SyncShardStatusResponse:
		that1 = t
	case SyncShardStatusResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SyncActivityRequest to the protobuf v3 wire format
func (val *SyncActivityRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SyncActivityRequest from the protobuf v3 wire format
func (val *SyncActivityRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SyncActivityRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SyncActivityRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SyncActivityRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SyncActivityRequest
	switch t := that.(type) {
	case *SyncActivityRequest:
		that1 = t
	case SyncActivityRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SyncActivityResponse to the protobuf v3 wire format
func (val *SyncActivityResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SyncActivityResponse from the protobuf v3 wire format
func (val *SyncActivityResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SyncActivityResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SyncActivityResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SyncActivityResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SyncActivityResponse
	switch t := that.(type) {
	case *SyncActivityResponse:
		that1 = t
	case SyncActivityResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeMutableStateRequest to the protobuf v3 wire format
func (val *DescribeMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeMutableStateRequest from the protobuf v3 wire format
func (val *DescribeMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeMutableStateRequest
	switch t := that.(type) {
	case *DescribeMutableStateRequest:
		that1 = t
	case DescribeMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeMutableStateResponse to the protobuf v3 wire format
func (val *DescribeMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeMutableStateResponse from the protobuf v3 wire format
func (val *DescribeMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeMutableStateResponse
	switch t := that.(type) {
	case *DescribeMutableStateResponse:
		that1 = t
	case DescribeMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeHistoryHostRequest to the protobuf v3 wire format
func (val *DescribeHistoryHostRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeHistoryHostRequest from the protobuf v3 wire format
func (val *DescribeHistoryHostRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeHistoryHostRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeHistoryHostRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeHistoryHostRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeHistoryHostRequest
	switch t := that.(type) {
	case *DescribeHistoryHostRequest:
		that1 = t
	case DescribeHistoryHostRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DescribeHistoryHostResponse to the protobuf v3 wire format
func (val *DescribeHistoryHostResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DescribeHistoryHostResponse from the protobuf v3 wire format
func (val *DescribeHistoryHostResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DescribeHistoryHostResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DescribeHistoryHostResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeHistoryHostResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeHistoryHostResponse
	switch t := that.(type) {
	case *DescribeHistoryHostResponse:
		that1 = t
	case DescribeHistoryHostResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CloseShardRequest to the protobuf v3 wire format
func (val *CloseShardRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CloseShardRequest from the protobuf v3 wire format
func (val *CloseShardRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CloseShardRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CloseShardRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CloseShardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CloseShardRequest
	switch t := that.(type) {
	case *CloseShardRequest:
		that1 = t
	case CloseShardRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CloseShardResponse to the protobuf v3 wire format
func (val *CloseShardResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CloseShardResponse from the protobuf v3 wire format
func (val *CloseShardResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CloseShardResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CloseShardResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CloseShardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CloseShardResponse
	switch t := that.(type) {
	case *CloseShardResponse:
		that1 = t
	case CloseShardResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetShardRequest to the protobuf v3 wire format
func (val *GetShardRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetShardRequest from the protobuf v3 wire format
func (val *GetShardRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetShardRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetShardRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetShardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetShardRequest
	switch t := that.(type) {
	case *GetShardRequest:
		that1 = t
	case GetShardRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetShardResponse to the protobuf v3 wire format
func (val *GetShardResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetShardResponse from the protobuf v3 wire format
func (val *GetShardResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetShardResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetShardResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetShardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetShardResponse
	switch t := that.(type) {
	case *GetShardResponse:
		that1 = t
	case GetShardResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveTaskRequest to the protobuf v3 wire format
func (val *RemoveTaskRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveTaskRequest from the protobuf v3 wire format
func (val *RemoveTaskRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveTaskRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveTaskRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveTaskRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveTaskRequest
	switch t := that.(type) {
	case *RemoveTaskRequest:
		that1 = t
	case RemoveTaskRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RemoveTaskResponse to the protobuf v3 wire format
func (val *RemoveTaskResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RemoveTaskResponse from the protobuf v3 wire format
func (val *RemoveTaskResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RemoveTaskResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RemoveTaskResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RemoveTaskResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RemoveTaskResponse
	switch t := that.(type) {
	case *RemoveTaskResponse:
		that1 = t
	case RemoveTaskResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetReplicationMessagesRequest to the protobuf v3 wire format
func (val *GetReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetReplicationMessagesRequest from the protobuf v3 wire format
func (val *GetReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetReplicationMessagesRequest
	switch t := that.(type) {
	case *GetReplicationMessagesRequest:
		that1 = t
	case GetReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetReplicationMessagesResponse to the protobuf v3 wire format
func (val *GetReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetReplicationMessagesResponse from the protobuf v3 wire format
func (val *GetReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetReplicationMessagesResponse
	switch t := that.(type) {
	case *GetReplicationMessagesResponse:
		that1 = t
	case GetReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQReplicationMessagesRequest to the protobuf v3 wire format
func (val *GetDLQReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQReplicationMessagesRequest from the protobuf v3 wire format
func (val *GetDLQReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQReplicationMessagesRequest
	switch t := that.(type) {
	case *GetDLQReplicationMessagesRequest:
		that1 = t
	case GetDLQReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQReplicationMessagesResponse to the protobuf v3 wire format
func (val *GetDLQReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQReplicationMessagesResponse from the protobuf v3 wire format
func (val *GetDLQReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQReplicationMessagesResponse
	switch t := that.(type) {
	case *GetDLQReplicationMessagesResponse:
		that1 = t
	case GetDLQReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type QueryWorkflowRequest to the protobuf v3 wire format
func (val *QueryWorkflowRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type QueryWorkflowRequest from the protobuf v3 wire format
func (val *QueryWorkflowRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *QueryWorkflowRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two QueryWorkflowRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryWorkflowRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryWorkflowRequest
	switch t := that.(type) {
	case *QueryWorkflowRequest:
		that1 = t
	case QueryWorkflowRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type QueryWorkflowResponse to the protobuf v3 wire format
func (val *QueryWorkflowResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type QueryWorkflowResponse from the protobuf v3 wire format
func (val *QueryWorkflowResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *QueryWorkflowResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two QueryWorkflowResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryWorkflowResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryWorkflowResponse
	switch t := that.(type) {
	case *QueryWorkflowResponse:
		that1 = t
	case QueryWorkflowResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReapplyEventsRequest to the protobuf v3 wire format
func (val *ReapplyEventsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReapplyEventsRequest from the protobuf v3 wire format
func (val *ReapplyEventsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReapplyEventsRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReapplyEventsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReapplyEventsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReapplyEventsRequest
	switch t := that.(type) {
	case *ReapplyEventsRequest:
		that1 = t
	case ReapplyEventsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReapplyEventsResponse to the protobuf v3 wire format
func (val *ReapplyEventsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReapplyEventsResponse from the protobuf v3 wire format
func (val *ReapplyEventsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReapplyEventsResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReapplyEventsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReapplyEventsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReapplyEventsResponse
	switch t := that.(type) {
	case *ReapplyEventsResponse:
		that1 = t
	case ReapplyEventsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQMessagesRequest to the protobuf v3 wire format
func (val *GetDLQMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQMessagesRequest from the protobuf v3 wire format
func (val *GetDLQMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQMessagesRequest
	switch t := that.(type) {
	case *GetDLQMessagesRequest:
		that1 = t
	case GetDLQMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQMessagesResponse to the protobuf v3 wire format
func (val *GetDLQMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQMessagesResponse from the protobuf v3 wire format
func (val *GetDLQMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQMessagesResponse
	switch t := that.(type) {
	case *GetDLQMessagesResponse:
		that1 = t
	case GetDLQMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PurgeDLQMessagesRequest to the protobuf v3 wire format
func (val *PurgeDLQMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PurgeDLQMessagesRequest from the protobuf v3 wire format
func (val *PurgeDLQMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PurgeDLQMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PurgeDLQMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PurgeDLQMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PurgeDLQMessagesRequest
	switch t := that.(type) {
	case *PurgeDLQMessagesRequest:
		that1 = t
	case PurgeDLQMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PurgeDLQMessagesResponse to the protobuf v3 wire format
func (val *PurgeDLQMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PurgeDLQMessagesResponse from the protobuf v3 wire format
func (val *PurgeDLQMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PurgeDLQMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PurgeDLQMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PurgeDLQMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PurgeDLQMessagesResponse
	switch t := that.(type) {
	case *PurgeDLQMessagesResponse:
		that1 = t
	case PurgeDLQMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MergeDLQMessagesRequest to the protobuf v3 wire format
func (val *MergeDLQMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MergeDLQMessagesRequest from the protobuf v3 wire format
func (val *MergeDLQMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MergeDLQMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MergeDLQMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MergeDLQMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MergeDLQMessagesRequest
	switch t := that.(type) {
	case *MergeDLQMessagesRequest:
		that1 = t
	case MergeDLQMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type MergeDLQMessagesResponse to the protobuf v3 wire format
func (val *MergeDLQMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type MergeDLQMessagesResponse from the protobuf v3 wire format
func (val *MergeDLQMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *MergeDLQMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two MergeDLQMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *MergeDLQMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *MergeDLQMessagesResponse
	switch t := that.(type) {
	case *MergeDLQMessagesResponse:
		that1 = t
	case MergeDLQMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RefreshWorkflowTasksRequest to the protobuf v3 wire format
func (val *RefreshWorkflowTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RefreshWorkflowTasksRequest from the protobuf v3 wire format
func (val *RefreshWorkflowTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RefreshWorkflowTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RefreshWorkflowTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RefreshWorkflowTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RefreshWorkflowTasksRequest
	switch t := that.(type) {
	case *RefreshWorkflowTasksRequest:
		that1 = t
	case RefreshWorkflowTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RefreshWorkflowTasksResponse to the protobuf v3 wire format
func (val *RefreshWorkflowTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RefreshWorkflowTasksResponse from the protobuf v3 wire format
func (val *RefreshWorkflowTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RefreshWorkflowTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RefreshWorkflowTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RefreshWorkflowTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RefreshWorkflowTasksResponse
	switch t := that.(type) {
	case *RefreshWorkflowTasksResponse:
		that1 = t
	case RefreshWorkflowTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GenerateLastHistoryReplicationTasksRequest to the protobuf v3 wire format
func (val *GenerateLastHistoryReplicationTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GenerateLastHistoryReplicationTasksRequest from the protobuf v3 wire format
func (val *GenerateLastHistoryReplicationTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GenerateLastHistoryReplicationTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GenerateLastHistoryReplicationTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GenerateLastHistoryReplicationTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GenerateLastHistoryReplicationTasksRequest
	switch t := that.(type) {
	case *GenerateLastHistoryReplicationTasksRequest:
		that1 = t
	case GenerateLastHistoryReplicationTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GenerateLastHistoryReplicationTasksResponse to the protobuf v3 wire format
func (val *GenerateLastHistoryReplicationTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GenerateLastHistoryReplicationTasksResponse from the protobuf v3 wire format
func (val *GenerateLastHistoryReplicationTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GenerateLastHistoryReplicationTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GenerateLastHistoryReplicationTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GenerateLastHistoryReplicationTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GenerateLastHistoryReplicationTasksResponse
	switch t := that.(type) {
	case *GenerateLastHistoryReplicationTasksResponse:
		that1 = t
	case GenerateLastHistoryReplicationTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetReplicationStatusRequest to the protobuf v3 wire format
func (val *GetReplicationStatusRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetReplicationStatusRequest from the protobuf v3 wire format
func (val *GetReplicationStatusRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetReplicationStatusRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetReplicationStatusRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetReplicationStatusRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetReplicationStatusRequest
	switch t := that.(type) {
	case *GetReplicationStatusRequest:
		that1 = t
	case GetReplicationStatusRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetReplicationStatusResponse to the protobuf v3 wire format
func (val *GetReplicationStatusResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetReplicationStatusResponse from the protobuf v3 wire format
func (val *GetReplicationStatusResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetReplicationStatusResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetReplicationStatusResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetReplicationStatusResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetReplicationStatusResponse
	switch t := that.(type) {
	case *GetReplicationStatusResponse:
		that1 = t
	case GetReplicationStatusResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ShardReplicationStatus to the protobuf v3 wire format
func (val *ShardReplicationStatus) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ShardReplicationStatus from the protobuf v3 wire format
func (val *ShardReplicationStatus) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ShardReplicationStatus) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ShardReplicationStatus values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ShardReplicationStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ShardReplicationStatus
	switch t := that.(type) {
	case *ShardReplicationStatus:
		that1 = t
	case ShardReplicationStatus:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type HandoverNamespaceInfo to the protobuf v3 wire format
func (val *HandoverNamespaceInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HandoverNamespaceInfo from the protobuf v3 wire format
func (val *HandoverNamespaceInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HandoverNamespaceInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HandoverNamespaceInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HandoverNamespaceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HandoverNamespaceInfo
	switch t := that.(type) {
	case *HandoverNamespaceInfo:
		that1 = t
	case HandoverNamespaceInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ShardReplicationStatusPerCluster to the protobuf v3 wire format
func (val *ShardReplicationStatusPerCluster) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ShardReplicationStatusPerCluster from the protobuf v3 wire format
func (val *ShardReplicationStatusPerCluster) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ShardReplicationStatusPerCluster) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ShardReplicationStatusPerCluster values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ShardReplicationStatusPerCluster) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ShardReplicationStatusPerCluster
	switch t := that.(type) {
	case *ShardReplicationStatusPerCluster:
		that1 = t
	case ShardReplicationStatusPerCluster:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RebuildMutableStateRequest to the protobuf v3 wire format
func (val *RebuildMutableStateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RebuildMutableStateRequest from the protobuf v3 wire format
func (val *RebuildMutableStateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RebuildMutableStateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RebuildMutableStateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RebuildMutableStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RebuildMutableStateRequest
	switch t := that.(type) {
	case *RebuildMutableStateRequest:
		that1 = t
	case RebuildMutableStateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RebuildMutableStateResponse to the protobuf v3 wire format
func (val *RebuildMutableStateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RebuildMutableStateResponse from the protobuf v3 wire format
func (val *RebuildMutableStateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RebuildMutableStateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RebuildMutableStateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RebuildMutableStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RebuildMutableStateResponse
	switch t := that.(type) {
	case *RebuildMutableStateResponse:
		that1 = t
	case RebuildMutableStateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ImportWorkflowExecutionRequest to the protobuf v3 wire format
func (val *ImportWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ImportWorkflowExecutionRequest from the protobuf v3 wire format
func (val *ImportWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ImportWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ImportWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ImportWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ImportWorkflowExecutionRequest
	switch t := that.(type) {
	case *ImportWorkflowExecutionRequest:
		that1 = t
	case ImportWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ImportWorkflowExecutionResponse to the protobuf v3 wire format
func (val *ImportWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ImportWorkflowExecutionResponse from the protobuf v3 wire format
func (val *ImportWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ImportWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ImportWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ImportWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ImportWorkflowExecutionResponse
	switch t := that.(type) {
	case *ImportWorkflowExecutionResponse:
		that1 = t
	case ImportWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteWorkflowVisibilityRecordRequest to the protobuf v3 wire format
func (val *DeleteWorkflowVisibilityRecordRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteWorkflowVisibilityRecordRequest from the protobuf v3 wire format
func (val *DeleteWorkflowVisibilityRecordRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteWorkflowVisibilityRecordRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteWorkflowVisibilityRecordRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowVisibilityRecordRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowVisibilityRecordRequest
	switch t := that.(type) {
	case *DeleteWorkflowVisibilityRecordRequest:
		that1 = t
	case DeleteWorkflowVisibilityRecordRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteWorkflowVisibilityRecordResponse to the protobuf v3 wire format
func (val *DeleteWorkflowVisibilityRecordResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteWorkflowVisibilityRecordResponse from the protobuf v3 wire format
func (val *DeleteWorkflowVisibilityRecordResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteWorkflowVisibilityRecordResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteWorkflowVisibilityRecordResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowVisibilityRecordResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowVisibilityRecordResponse
	switch t := that.(type) {
	case *DeleteWorkflowVisibilityRecordResponse:
		that1 = t
	case DeleteWorkflowVisibilityRecordResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateWorkflowExecutionRequest to the protobuf v3 wire format
func (val *UpdateWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateWorkflowExecutionRequest from the protobuf v3 wire format
func (val *UpdateWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkflowExecutionRequest
	switch t := that.(type) {
	case *UpdateWorkflowExecutionRequest:
		that1 = t
	case UpdateWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateWorkflowExecutionResponse to the protobuf v3 wire format
func (val *UpdateWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateWorkflowExecutionResponse from the protobuf v3 wire format
func (val *UpdateWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkflowExecutionResponse
	switch t := that.(type) {
	case *UpdateWorkflowExecutionResponse:
		that1 = t
	case UpdateWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StreamWorkflowReplicationMessagesRequest to the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StreamWorkflowReplicationMessagesRequest from the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StreamWorkflowReplicationMessagesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StreamWorkflowReplicationMessagesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StreamWorkflowReplicationMessagesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StreamWorkflowReplicationMessagesRequest
	switch t := that.(type) {
	case *StreamWorkflowReplicationMessagesRequest:
		that1 = t
	case StreamWorkflowReplicationMessagesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StreamWorkflowReplicationMessagesResponse to the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StreamWorkflowReplicationMessagesResponse from the protobuf v3 wire format
func (val *StreamWorkflowReplicationMessagesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StreamWorkflowReplicationMessagesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StreamWorkflowReplicationMessagesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StreamWorkflowReplicationMessagesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StreamWorkflowReplicationMessagesResponse
	switch t := that.(type) {
	case *StreamWorkflowReplicationMessagesResponse:
		that1 = t
	case StreamWorkflowReplicationMessagesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollWorkflowExecutionUpdateRequest to the protobuf v3 wire format
func (val *PollWorkflowExecutionUpdateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollWorkflowExecutionUpdateRequest from the protobuf v3 wire format
func (val *PollWorkflowExecutionUpdateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollWorkflowExecutionUpdateRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollWorkflowExecutionUpdateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowExecutionUpdateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowExecutionUpdateRequest
	switch t := that.(type) {
	case *PollWorkflowExecutionUpdateRequest:
		that1 = t
	case PollWorkflowExecutionUpdateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollWorkflowExecutionUpdateResponse to the protobuf v3 wire format
func (val *PollWorkflowExecutionUpdateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollWorkflowExecutionUpdateResponse from the protobuf v3 wire format
func (val *PollWorkflowExecutionUpdateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollWorkflowExecutionUpdateResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollWorkflowExecutionUpdateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowExecutionUpdateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowExecutionUpdateResponse
	switch t := that.(type) {
	case *PollWorkflowExecutionUpdateResponse:
		that1 = t
	case PollWorkflowExecutionUpdateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionHistoryRequest to the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionHistoryRequest from the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionHistoryRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionHistoryRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryRequest
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryRequest:
		that1 = t
	case GetWorkflowExecutionHistoryRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionHistoryResponse to the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionHistoryResponse from the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionHistoryResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionHistoryResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryResponse
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryResponse:
		that1 = t
	case GetWorkflowExecutionHistoryResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionHistoryReverseRequest to the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryReverseRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionHistoryReverseRequest from the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryReverseRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionHistoryReverseRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionHistoryReverseRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryReverseRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryReverseRequest
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryReverseRequest:
		that1 = t
	case GetWorkflowExecutionHistoryReverseRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionHistoryReverseResponse to the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryReverseResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionHistoryReverseResponse from the protobuf v3 wire format
func (val *GetWorkflowExecutionHistoryReverseResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionHistoryReverseResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionHistoryReverseResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryReverseResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryReverseResponse
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryReverseResponse:
		that1 = t
	case GetWorkflowExecutionHistoryReverseResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionRawHistoryV2Request to the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Request) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionRawHistoryV2Request from the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Request) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionRawHistoryV2Request) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionRawHistoryV2Request values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionRawHistoryV2Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionRawHistoryV2Request
	switch t := that.(type) {
	case *GetWorkflowExecutionRawHistoryV2Request:
		that1 = t
	case GetWorkflowExecutionRawHistoryV2Request:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetWorkflowExecutionRawHistoryV2Response to the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Response) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetWorkflowExecutionRawHistoryV2Response from the protobuf v3 wire format
func (val *GetWorkflowExecutionRawHistoryV2Response) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetWorkflowExecutionRawHistoryV2Response) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetWorkflowExecutionRawHistoryV2Response values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionRawHistoryV2Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionRawHistoryV2Response
	switch t := that.(type) {
	case *GetWorkflowExecutionRawHistoryV2Response:
		that1 = t
	case GetWorkflowExecutionRawHistoryV2Response:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ForceDeleteWorkflowExecutionRequest to the protobuf v3 wire format
func (val *ForceDeleteWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ForceDeleteWorkflowExecutionRequest from the protobuf v3 wire format
func (val *ForceDeleteWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ForceDeleteWorkflowExecutionRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ForceDeleteWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ForceDeleteWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ForceDeleteWorkflowExecutionRequest
	switch t := that.(type) {
	case *ForceDeleteWorkflowExecutionRequest:
		that1 = t
	case ForceDeleteWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ForceDeleteWorkflowExecutionResponse to the protobuf v3 wire format
func (val *ForceDeleteWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ForceDeleteWorkflowExecutionResponse from the protobuf v3 wire format
func (val *ForceDeleteWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ForceDeleteWorkflowExecutionResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ForceDeleteWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ForceDeleteWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ForceDeleteWorkflowExecutionResponse
	switch t := that.(type) {
	case *ForceDeleteWorkflowExecutionResponse:
		that1 = t
	case ForceDeleteWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQTasksRequest to the protobuf v3 wire format
func (val *GetDLQTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQTasksRequest from the protobuf v3 wire format
func (val *GetDLQTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQTasksRequest
	switch t := that.(type) {
	case *GetDLQTasksRequest:
		that1 = t
	case GetDLQTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type GetDLQTasksResponse to the protobuf v3 wire format
func (val *GetDLQTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type GetDLQTasksResponse from the protobuf v3 wire format
func (val *GetDLQTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *GetDLQTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two GetDLQTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetDLQTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetDLQTasksResponse
	switch t := that.(type) {
	case *GetDLQTasksResponse:
		that1 = t
	case GetDLQTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteDLQTasksRequest to the protobuf v3 wire format
func (val *DeleteDLQTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteDLQTasksRequest from the protobuf v3 wire format
func (val *DeleteDLQTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteDLQTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteDLQTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteDLQTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteDLQTasksRequest
	switch t := that.(type) {
	case *DeleteDLQTasksRequest:
		that1 = t
	case DeleteDLQTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeleteDLQTasksResponse to the protobuf v3 wire format
func (val *DeleteDLQTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeleteDLQTasksResponse from the protobuf v3 wire format
func (val *DeleteDLQTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeleteDLQTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeleteDLQTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteDLQTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteDLQTasksResponse
	switch t := that.(type) {
	case *DeleteDLQTasksResponse:
		that1 = t
	case DeleteDLQTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListQueuesRequest to the protobuf v3 wire format
func (val *ListQueuesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListQueuesRequest from the protobuf v3 wire format
func (val *ListQueuesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListQueuesRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListQueuesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListQueuesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListQueuesRequest
	switch t := that.(type) {
	case *ListQueuesRequest:
		that1 = t
	case ListQueuesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ListQueuesResponse to the protobuf v3 wire format
func (val *ListQueuesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ListQueuesResponse from the protobuf v3 wire format
func (val *ListQueuesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ListQueuesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ListQueuesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListQueuesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListQueuesResponse
	switch t := that.(type) {
	case *ListQueuesResponse:
		that1 = t
	case ListQueuesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddTasksRequest to the protobuf v3 wire format
func (val *AddTasksRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddTasksRequest from the protobuf v3 wire format
func (val *AddTasksRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddTasksRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddTasksRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddTasksRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddTasksRequest
	switch t := that.(type) {
	case *AddTasksRequest:
		that1 = t
	case AddTasksRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddTasksResponse to the protobuf v3 wire format
func (val *AddTasksResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddTasksResponse from the protobuf v3 wire format
func (val *AddTasksResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddTasksResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddTasksResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddTasksResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddTasksResponse
	switch t := that.(type) {
	case *AddTasksResponse:
		that1 = t
	case AddTasksResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
