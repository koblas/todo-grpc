// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: core/eventbus.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventbusEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventbusEmpty) Reset() {
	*x = EventbusEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_eventbus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventbusEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventbusEmpty) ProtoMessage() {}

func (x *EventbusEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_core_eventbus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventbusEmpty.ProtoReflect.Descriptor instead.
func (*EventbusEmpty) Descriptor() ([]byte, []int) {
	return file_core_eventbus_proto_rawDescGZIP(), []int{0}
}

// This is the basis of inheritance for SNS publishing
type BaseEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdemponcyId string `protobuf:"bytes,1,opt,name=idemponcyId,proto3" json:"idemponcyId,omitempty"`
	Action      string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *BaseEvent) Reset() {
	*x = BaseEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_eventbus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseEvent) ProtoMessage() {}

func (x *BaseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_core_eventbus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseEvent.ProtoReflect.Descriptor instead.
func (*BaseEvent) Descriptor() ([]byte, []int) {
	return file_core_eventbus_proto_rawDescGZIP(), []int{1}
}

func (x *BaseEvent) GetIdemponcyId() string {
	if x != nil {
		return x.IdemponcyId
	}
	return ""
}

func (x *BaseEvent) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type TodoEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdemponcyId string      `protobuf:"bytes,1,opt,name=idemponcyId,proto3" json:"idemponcyId,omitempty"`
	Action      string      `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Current     *TodoObject `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	Previous    *TodoObject `protobuf:"bytes,4,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (x *TodoEvent) Reset() {
	*x = TodoEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_eventbus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoEvent) ProtoMessage() {}

func (x *TodoEvent) ProtoReflect() protoreflect.Message {
	mi := &file_core_eventbus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoEvent.ProtoReflect.Descriptor instead.
func (*TodoEvent) Descriptor() ([]byte, []int) {
	return file_core_eventbus_proto_rawDescGZIP(), []int{2}
}

func (x *TodoEvent) GetIdemponcyId() string {
	if x != nil {
		return x.IdemponcyId
	}
	return ""
}

func (x *TodoEvent) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *TodoEvent) GetCurrent() *TodoObject {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *TodoEvent) GetPrevious() *TodoObject {
	if x != nil {
		return x.Previous
	}
	return nil
}

var File_core_eventbus_proto protoreflect.FileDescriptor

var file_core_eventbus_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x62, 0x75, 0x73, 0x1a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75,
	0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x6e, 0x63, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f,
	0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x01,
	0x0a, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x64, 0x65, 0x6d, 0x70, 0x6f, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x32, 0x51, 0x0a, 0x0c, 0x54, 0x6f, 0x64,
	0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x62, 0x75, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_eventbus_proto_rawDescOnce sync.Once
	file_core_eventbus_proto_rawDescData = file_core_eventbus_proto_rawDesc
)

func file_core_eventbus_proto_rawDescGZIP() []byte {
	file_core_eventbus_proto_rawDescOnce.Do(func() {
		file_core_eventbus_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_eventbus_proto_rawDescData)
	})
	return file_core_eventbus_proto_rawDescData
}

var file_core_eventbus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_eventbus_proto_goTypes = []interface{}{
	(*EventbusEmpty)(nil), // 0: core.eventbus.EventbusEmpty
	(*BaseEvent)(nil),     // 1: core.eventbus.BaseEvent
	(*TodoEvent)(nil),     // 2: core.eventbus.TodoEvent
	(*TodoObject)(nil),    // 3: core.todo.TodoObject
}
var file_core_eventbus_proto_depIdxs = []int32{
	3, // 0: core.eventbus.TodoEvent.current:type_name -> core.todo.TodoObject
	3, // 1: core.eventbus.TodoEvent.previous:type_name -> core.todo.TodoObject
	2, // 2: core.eventbus.TodoEventbus.Message:input_type -> core.eventbus.TodoEvent
	0, // 3: core.eventbus.TodoEventbus.Message:output_type -> core.eventbus.EventbusEmpty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_eventbus_proto_init() }
func file_core_eventbus_proto_init() {
	if File_core_eventbus_proto != nil {
		return
	}
	file_core_todo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_eventbus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventbusEmpty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_eventbus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_eventbus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_eventbus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_eventbus_proto_goTypes,
		DependencyIndexes: file_core_eventbus_proto_depIdxs,
		MessageInfos:      file_core_eventbus_proto_msgTypes,
	}.Build()
	File_core_eventbus_proto = out.File
	file_core_eventbus_proto_rawDesc = nil
	file_core_eventbus_proto_goTypes = nil
	file_core_eventbus_proto_depIdxs = nil
}
