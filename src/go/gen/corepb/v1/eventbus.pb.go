// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: corepb/v1/eventbus.proto

package corepbv1

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
		mi := &file_corepb_v1_eventbus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventbusEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventbusEmpty) ProtoMessage() {}

func (x *EventbusEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[0]
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
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{0}
}

var File_corepb_v1_eventbus_proto protoreflect.FileDescriptor

var file_corepb_v1_eventbus_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x62, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x72, 0x65,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x72,
	0x65, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x50, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x95, 0x03, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x16, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x15,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a,
	0x15, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d,
	0x0a, 0x13, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x52, 0x0a,
	0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x12, 0x42, 0x0a,
	0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0xaa, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62,
	0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x4d, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_corepb_v1_eventbus_proto_rawDescOnce sync.Once
	file_corepb_v1_eventbus_proto_rawDescData = file_corepb_v1_eventbus_proto_rawDesc
)

func file_corepb_v1_eventbus_proto_rawDescGZIP() []byte {
	file_corepb_v1_eventbus_proto_rawDescOnce.Do(func() {
		file_corepb_v1_eventbus_proto_rawDescData = protoimpl.X.CompressGZIP(file_corepb_v1_eventbus_proto_rawDescData)
	})
	return file_corepb_v1_eventbus_proto_rawDescData
}

var file_corepb_v1_eventbus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_corepb_v1_eventbus_proto_goTypes = []interface{}{
	(*EventbusEmpty)(nil),            // 0: corepb.v1.EventbusEmpty
	(*BroadcastEvent)(nil),           // 1: corepb.v1.BroadcastEvent
	(*UserChangeEvent)(nil),          // 2: corepb.v1.UserChangeEvent
	(*UserSecurityEvent)(nil),        // 3: corepb.v1.UserSecurityEvent
	(*TodoChangeEvent)(nil),          // 4: corepb.v1.TodoChangeEvent
	(*FileServiceUploadEvent)(nil),   // 5: corepb.v1.FileServiceUploadEvent
	(*FileServiceCompleteEvent)(nil), // 6: corepb.v1.FileServiceCompleteEvent
}
var file_corepb_v1_eventbus_proto_depIdxs = []int32{
	1, // 0: corepb.v1.BroadcastEventbus.Send:input_type -> corepb.v1.BroadcastEvent
	2, // 1: corepb.v1.UserEventbus.UserChange:input_type -> corepb.v1.UserChangeEvent
	3, // 2: corepb.v1.UserEventbus.SecurityPasswordChange:input_type -> corepb.v1.UserSecurityEvent
	3, // 3: corepb.v1.UserEventbus.SecurityForgotRequest:input_type -> corepb.v1.UserSecurityEvent
	3, // 4: corepb.v1.UserEventbus.SecurityRegisterToken:input_type -> corepb.v1.UserSecurityEvent
	3, // 5: corepb.v1.UserEventbus.SecurityInviteToken:input_type -> corepb.v1.UserSecurityEvent
	4, // 6: corepb.v1.TodoEventbus.TodoChange:input_type -> corepb.v1.TodoChangeEvent
	5, // 7: corepb.v1.FileEventbus.FileUploaded:input_type -> corepb.v1.FileServiceUploadEvent
	6, // 8: corepb.v1.FileEventbus.FileComplete:input_type -> corepb.v1.FileServiceCompleteEvent
	0, // 9: corepb.v1.BroadcastEventbus.Send:output_type -> corepb.v1.EventbusEmpty
	0, // 10: corepb.v1.UserEventbus.UserChange:output_type -> corepb.v1.EventbusEmpty
	0, // 11: corepb.v1.UserEventbus.SecurityPasswordChange:output_type -> corepb.v1.EventbusEmpty
	0, // 12: corepb.v1.UserEventbus.SecurityForgotRequest:output_type -> corepb.v1.EventbusEmpty
	0, // 13: corepb.v1.UserEventbus.SecurityRegisterToken:output_type -> corepb.v1.EventbusEmpty
	0, // 14: corepb.v1.UserEventbus.SecurityInviteToken:output_type -> corepb.v1.EventbusEmpty
	0, // 15: corepb.v1.TodoEventbus.TodoChange:output_type -> corepb.v1.EventbusEmpty
	0, // 16: corepb.v1.FileEventbus.FileUploaded:output_type -> corepb.v1.EventbusEmpty
	0, // 17: corepb.v1.FileEventbus.FileComplete:output_type -> corepb.v1.EventbusEmpty
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_corepb_v1_eventbus_proto_init() }
func file_corepb_v1_eventbus_proto_init() {
	if File_corepb_v1_eventbus_proto != nil {
		return
	}
	file_corepb_v1_user_proto_init()
	file_corepb_v1_todo_proto_init()
	file_corepb_v1_file_proto_init()
	file_corepb_v1_websocket_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_corepb_v1_eventbus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_corepb_v1_eventbus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_corepb_v1_eventbus_proto_goTypes,
		DependencyIndexes: file_corepb_v1_eventbus_proto_depIdxs,
		MessageInfos:      file_corepb_v1_eventbus_proto_msgTypes,
	}.Build()
	File_corepb_v1_eventbus_proto = out.File
	file_corepb_v1_eventbus_proto_rawDesc = nil
	file_corepb_v1_eventbus_proto_goTypes = nil
	file_corepb_v1_eventbus_proto_depIdxs = nil
}
