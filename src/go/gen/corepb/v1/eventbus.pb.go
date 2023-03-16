// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
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

type BroadcastEventbusSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastEventbusSendResponse) Reset() {
	*x = BroadcastEventbusSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastEventbusSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastEventbusSendResponse) ProtoMessage() {}

func (x *BroadcastEventbusSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastEventbusSendResponse.ProtoReflect.Descriptor instead.
func (*BroadcastEventbusSendResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{1}
}

type UserEventbusUserChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserEventbusUserChangeResponse) Reset() {
	*x = UserEventbusUserChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEventbusUserChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEventbusUserChangeResponse) ProtoMessage() {}

func (x *UserEventbusUserChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEventbusUserChangeResponse.ProtoReflect.Descriptor instead.
func (*UserEventbusUserChangeResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{2}
}

type UserEventbusSecurityPasswordChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserEventbusSecurityPasswordChangeResponse) Reset() {
	*x = UserEventbusSecurityPasswordChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEventbusSecurityPasswordChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEventbusSecurityPasswordChangeResponse) ProtoMessage() {}

func (x *UserEventbusSecurityPasswordChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEventbusSecurityPasswordChangeResponse.ProtoReflect.Descriptor instead.
func (*UserEventbusSecurityPasswordChangeResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{3}
}

type UserEventbusSecurityForgotRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserEventbusSecurityForgotRequestResponse) Reset() {
	*x = UserEventbusSecurityForgotRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEventbusSecurityForgotRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEventbusSecurityForgotRequestResponse) ProtoMessage() {}

func (x *UserEventbusSecurityForgotRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEventbusSecurityForgotRequestResponse.ProtoReflect.Descriptor instead.
func (*UserEventbusSecurityForgotRequestResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{4}
}

type UserEventbusSecurityRegisterTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserEventbusSecurityRegisterTokenResponse) Reset() {
	*x = UserEventbusSecurityRegisterTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEventbusSecurityRegisterTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEventbusSecurityRegisterTokenResponse) ProtoMessage() {}

func (x *UserEventbusSecurityRegisterTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEventbusSecurityRegisterTokenResponse.ProtoReflect.Descriptor instead.
func (*UserEventbusSecurityRegisterTokenResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{5}
}

type UserEventbusSecurityInviteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserEventbusSecurityInviteTokenResponse) Reset() {
	*x = UserEventbusSecurityInviteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEventbusSecurityInviteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEventbusSecurityInviteTokenResponse) ProtoMessage() {}

func (x *UserEventbusSecurityInviteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEventbusSecurityInviteTokenResponse.ProtoReflect.Descriptor instead.
func (*UserEventbusSecurityInviteTokenResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{6}
}

type TodoEventbusTodoChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TodoEventbusTodoChangeResponse) Reset() {
	*x = TodoEventbusTodoChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoEventbusTodoChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoEventbusTodoChangeResponse) ProtoMessage() {}

func (x *TodoEventbusTodoChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoEventbusTodoChangeResponse.ProtoReflect.Descriptor instead.
func (*TodoEventbusTodoChangeResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{7}
}

type FileEventbusFileUploadedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FileEventbusFileUploadedResponse) Reset() {
	*x = FileEventbusFileUploadedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileEventbusFileUploadedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEventbusFileUploadedResponse) ProtoMessage() {}

func (x *FileEventbusFileUploadedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEventbusFileUploadedResponse.ProtoReflect.Descriptor instead.
func (*FileEventbusFileUploadedResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{8}
}

type FileEventbusFileCompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FileEventbusFileCompleteResponse) Reset() {
	*x = FileEventbusFileCompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_v1_eventbus_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileEventbusFileCompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEventbusFileCompleteResponse) ProtoMessage() {}

func (x *FileEventbusFileCompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_v1_eventbus_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEventbusFileCompleteResponse.ProtoReflect.Descriptor instead.
func (*FileEventbusFileCompleteResponse) Descriptor() ([]byte, []int) {
	return file_corepb_v1_eventbus_proto_rawDescGZIP(), []int{9}
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
	0x70, 0x74, 0x79, 0x22, 0x1f, 0x0a, 0x1d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x1e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x62, 0x75, 0x73, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x2a, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x29, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2b, 0x0a, 0x29, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75,
	0x73, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29,
	0x0a, 0x27, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x1e, 0x54, 0x6f, 0x64,
	0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x20, 0x46,
	0x69, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x0a, 0x20, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x67, 0x0a, 0x18, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9c, 0x04, 0x0a,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x29,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x16, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x15, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a,
	0x34, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x15, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x34, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x67, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6a, 0x0a, 0x13, 0x54,
	0x6f, 0x64, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x29, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x62, 0x75, 0x73, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd7, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5e, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12,
	0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_corepb_v1_eventbus_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_corepb_v1_eventbus_proto_goTypes = []interface{}{
	(*EventbusEmpty)(nil),                              // 0: corepb.v1.EventbusEmpty
	(*BroadcastEventbusSendResponse)(nil),              // 1: corepb.v1.BroadcastEventbusSendResponse
	(*UserEventbusUserChangeResponse)(nil),             // 2: corepb.v1.UserEventbusUserChangeResponse
	(*UserEventbusSecurityPasswordChangeResponse)(nil), // 3: corepb.v1.UserEventbusSecurityPasswordChangeResponse
	(*UserEventbusSecurityForgotRequestResponse)(nil),  // 4: corepb.v1.UserEventbusSecurityForgotRequestResponse
	(*UserEventbusSecurityRegisterTokenResponse)(nil),  // 5: corepb.v1.UserEventbusSecurityRegisterTokenResponse
	(*UserEventbusSecurityInviteTokenResponse)(nil),    // 6: corepb.v1.UserEventbusSecurityInviteTokenResponse
	(*TodoEventbusTodoChangeResponse)(nil),             // 7: corepb.v1.TodoEventbusTodoChangeResponse
	(*FileEventbusFileUploadedResponse)(nil),           // 8: corepb.v1.FileEventbusFileUploadedResponse
	(*FileEventbusFileCompleteResponse)(nil),           // 9: corepb.v1.FileEventbusFileCompleteResponse
	(*BroadcastEvent)(nil),                             // 10: corepb.v1.BroadcastEvent
	(*UserChangeEvent)(nil),                            // 11: corepb.v1.UserChangeEvent
	(*UserSecurityEvent)(nil),                          // 12: corepb.v1.UserSecurityEvent
	(*TodoChangeEvent)(nil),                            // 13: corepb.v1.TodoChangeEvent
	(*FileServiceUploadEvent)(nil),                     // 14: corepb.v1.FileServiceUploadEvent
	(*FileServiceCompleteEvent)(nil),                   // 15: corepb.v1.FileServiceCompleteEvent
}
var file_corepb_v1_eventbus_proto_depIdxs = []int32{
	10, // 0: corepb.v1.BroadcastEventbusService.Send:input_type -> corepb.v1.BroadcastEvent
	11, // 1: corepb.v1.UserEventbusService.UserChange:input_type -> corepb.v1.UserChangeEvent
	12, // 2: corepb.v1.UserEventbusService.SecurityPasswordChange:input_type -> corepb.v1.UserSecurityEvent
	12, // 3: corepb.v1.UserEventbusService.SecurityForgotRequest:input_type -> corepb.v1.UserSecurityEvent
	12, // 4: corepb.v1.UserEventbusService.SecurityRegisterToken:input_type -> corepb.v1.UserSecurityEvent
	12, // 5: corepb.v1.UserEventbusService.SecurityInviteToken:input_type -> corepb.v1.UserSecurityEvent
	13, // 6: corepb.v1.TodoEventbusService.TodoChange:input_type -> corepb.v1.TodoChangeEvent
	14, // 7: corepb.v1.FileEventbusService.FileUploaded:input_type -> corepb.v1.FileServiceUploadEvent
	15, // 8: corepb.v1.FileEventbusService.FileComplete:input_type -> corepb.v1.FileServiceCompleteEvent
	1,  // 9: corepb.v1.BroadcastEventbusService.Send:output_type -> corepb.v1.BroadcastEventbusSendResponse
	2,  // 10: corepb.v1.UserEventbusService.UserChange:output_type -> corepb.v1.UserEventbusUserChangeResponse
	3,  // 11: corepb.v1.UserEventbusService.SecurityPasswordChange:output_type -> corepb.v1.UserEventbusSecurityPasswordChangeResponse
	4,  // 12: corepb.v1.UserEventbusService.SecurityForgotRequest:output_type -> corepb.v1.UserEventbusSecurityForgotRequestResponse
	5,  // 13: corepb.v1.UserEventbusService.SecurityRegisterToken:output_type -> corepb.v1.UserEventbusSecurityRegisterTokenResponse
	6,  // 14: corepb.v1.UserEventbusService.SecurityInviteToken:output_type -> corepb.v1.UserEventbusSecurityInviteTokenResponse
	7,  // 15: corepb.v1.TodoEventbusService.TodoChange:output_type -> corepb.v1.TodoEventbusTodoChangeResponse
	8,  // 16: corepb.v1.FileEventbusService.FileUploaded:output_type -> corepb.v1.FileEventbusFileUploadedResponse
	9,  // 17: corepb.v1.FileEventbusService.FileComplete:output_type -> corepb.v1.FileEventbusFileCompleteResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
		file_corepb_v1_eventbus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastEventbusSendResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEventbusUserChangeResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEventbusSecurityPasswordChangeResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEventbusSecurityForgotRequestResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEventbusSecurityRegisterTokenResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEventbusSecurityInviteTokenResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoEventbusTodoChangeResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileEventbusFileUploadedResponse); i {
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
		file_corepb_v1_eventbus_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileEventbusFileCompleteResponse); i {
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
			NumMessages:   10,
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
