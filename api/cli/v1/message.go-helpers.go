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

package cli

import (
	"google.golang.org/protobuf/proto"
)

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

// Marshal an object of type WorkflowExecutionInfo to the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionInfo from the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionInfo
	switch t := that.(type) {
	case *WorkflowExecutionInfo:
		that1 = t
	case WorkflowExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingActivityInfo to the protobuf v3 wire format
func (val *PendingActivityInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingActivityInfo from the protobuf v3 wire format
func (val *PendingActivityInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingActivityInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingActivityInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingActivityInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingActivityInfo
	switch t := that.(type) {
	case *PendingActivityInfo:
		that1 = t
	case PendingActivityInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SearchAttributes to the protobuf v3 wire format
func (val *SearchAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SearchAttributes from the protobuf v3 wire format
func (val *SearchAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SearchAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SearchAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SearchAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SearchAttributes
	switch t := that.(type) {
	case *SearchAttributes:
		that1 = t
	case SearchAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Failure to the protobuf v3 wire format
func (val *Failure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Failure from the protobuf v3 wire format
func (val *Failure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Failure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Failure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Failure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Failure
	switch t := that.(type) {
	case *Failure:
		that1 = t
	case Failure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type AddSearchAttributesResponse to the protobuf v3 wire format
func (val *AddSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AddSearchAttributesResponse from the protobuf v3 wire format
func (val *AddSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AddSearchAttributesResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AddSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AddSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AddSearchAttributesResponse
	switch t := that.(type) {
	case *AddSearchAttributesResponse:
		that1 = t
	case AddSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
