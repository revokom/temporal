// The MIT License
//
// Copyright (c) 2020 Temporal Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: temporal/server/api/enums/v1/dlq.proto

package enums

import (
	reflect "reflect"
	"strconv"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DLQOperationType int32

const (
	DLQ_OPERATION_TYPE_UNSPECIFIED DLQOperationType = 0
	DLQ_OPERATION_TYPE_MERGE       DLQOperationType = 1
	DLQ_OPERATION_TYPE_PURGE       DLQOperationType = 2
)

// Enum value maps for DLQOperationType.
var (
	DLQOperationType_name = map[int32]string{
		0: "DLQ_OPERATION_TYPE_UNSPECIFIED",
		1: "DLQ_OPERATION_TYPE_MERGE",
		2: "DLQ_OPERATION_TYPE_PURGE",
	}
	DLQOperationType_value = map[string]int32{
		"DLQ_OPERATION_TYPE_UNSPECIFIED": 0,
		"DLQ_OPERATION_TYPE_MERGE":       1,
		"DLQ_OPERATION_TYPE_PURGE":       2,
	}
)

func (x DLQOperationType) Enum() *DLQOperationType {
	p := new(DLQOperationType)
	*p = x
	return p
}

func (x DLQOperationType) String() string {
	switch x {
	case DLQ_OPERATION_TYPE_UNSPECIFIED:
		return "DlqOperationTypeUnspecified"
	case DLQ_OPERATION_TYPE_MERGE:
		return "DlqOperationTypeMerge"
	case DLQ_OPERATION_TYPE_PURGE:
		return "DlqOperationTypePurge"
	default:
		return strconv.Itoa(int(x))
	}

}

func (DLQOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_server_api_enums_v1_dlq_proto_enumTypes[0].Descriptor()
}

func (DLQOperationType) Type() protoreflect.EnumType {
	return &file_temporal_server_api_enums_v1_dlq_proto_enumTypes[0]
}

func (x DLQOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DLQOperationType.Descriptor instead.
func (DLQOperationType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_server_api_enums_v1_dlq_proto_rawDescGZIP(), []int{0}
}

type DLQOperationState int32

const (
	DLQ_OPERATION_STATE_UNSPECIFIED DLQOperationState = 0
	DLQ_OPERATION_STATE_RUNNING     DLQOperationState = 1
	DLQ_OPERATION_STATE_COMPLETED   DLQOperationState = 2
	DLQ_OPERATION_STATE_FAILED      DLQOperationState = 3
)

// Enum value maps for DLQOperationState.
var (
	DLQOperationState_name = map[int32]string{
		0: "DLQ_OPERATION_STATE_UNSPECIFIED",
		1: "DLQ_OPERATION_STATE_RUNNING",
		2: "DLQ_OPERATION_STATE_COMPLETED",
		3: "DLQ_OPERATION_STATE_FAILED",
	}
	DLQOperationState_value = map[string]int32{
		"DLQ_OPERATION_STATE_UNSPECIFIED": 0,
		"DLQ_OPERATION_STATE_RUNNING":     1,
		"DLQ_OPERATION_STATE_COMPLETED":   2,
		"DLQ_OPERATION_STATE_FAILED":      3,
	}
)

func (x DLQOperationState) Enum() *DLQOperationState {
	p := new(DLQOperationState)
	*p = x
	return p
}

func (x DLQOperationState) String() string {
	switch x {
	case DLQ_OPERATION_STATE_UNSPECIFIED:
		return "DlqOperationStateUnspecified"
	case DLQ_OPERATION_STATE_RUNNING:
		return "DlqOperationStateRunning"
	case DLQ_OPERATION_STATE_COMPLETED:
		return "DlqOperationStateCompleted"
	case DLQ_OPERATION_STATE_FAILED:
		return "DlqOperationStateFailed"
	default:
		return strconv.Itoa(int(x))
	}

}

func (DLQOperationState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_server_api_enums_v1_dlq_proto_enumTypes[1].Descriptor()
}

func (DLQOperationState) Type() protoreflect.EnumType {
	return &file_temporal_server_api_enums_v1_dlq_proto_enumTypes[1]
}

func (x DLQOperationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DLQOperationState.Descriptor instead.
func (DLQOperationState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_server_api_enums_v1_dlq_proto_rawDescGZIP(), []int{1}
}

var File_temporal_server_api_enums_v1_dlq_proto protoreflect.FileDescriptor

var file_temporal_server_api_enums_v1_dlq_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x6c, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x72, 0x0a, 0x10, 0x44, 0x4c, 0x51, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x4c,
	0x51, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c,
	0x0a, 0x18, 0x44, 0x4c, 0x51, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x44, 0x4c, 0x51, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x55, 0x52, 0x47, 0x45, 0x10, 0x02, 0x2a, 0x9c, 0x01, 0x0a, 0x11, 0x44,
	0x4c, 0x51, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4c, 0x51, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x4c, 0x51, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x4c, 0x51, 0x5f, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x4c, 0x51,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x6f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_enums_v1_dlq_proto_rawDescOnce sync.Once
	file_temporal_server_api_enums_v1_dlq_proto_rawDescData = file_temporal_server_api_enums_v1_dlq_proto_rawDesc
)

func file_temporal_server_api_enums_v1_dlq_proto_rawDescGZIP() []byte {
	file_temporal_server_api_enums_v1_dlq_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_enums_v1_dlq_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_enums_v1_dlq_proto_rawDescData)
	})
	return file_temporal_server_api_enums_v1_dlq_proto_rawDescData
}

var file_temporal_server_api_enums_v1_dlq_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_server_api_enums_v1_dlq_proto_goTypes = []interface{}{
	(DLQOperationType)(0),  // 0: temporal.server.api.enums.v1.DLQOperationType
	(DLQOperationState)(0), // 1: temporal.server.api.enums.v1.DLQOperationState
}
var file_temporal_server_api_enums_v1_dlq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_server_api_enums_v1_dlq_proto_init() }
func file_temporal_server_api_enums_v1_dlq_proto_init() {
	if File_temporal_server_api_enums_v1_dlq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_enums_v1_dlq_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_enums_v1_dlq_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_enums_v1_dlq_proto_depIdxs,
		EnumInfos:         file_temporal_server_api_enums_v1_dlq_proto_enumTypes,
	}.Build()
	File_temporal_server_api_enums_v1_dlq_proto = out.File
	file_temporal_server_api_enums_v1_dlq_proto_rawDesc = nil
	file_temporal_server_api_enums_v1_dlq_proto_goTypes = nil
	file_temporal_server_api_enums_v1_dlq_proto_depIdxs = nil
}
