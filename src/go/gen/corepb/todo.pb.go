// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: corepb/todo.proto

package corepb

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

type TodoGetParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TodoGetParams) Reset() {
	*x = TodoGetParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoGetParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoGetParams) ProtoMessage() {}

func (x *TodoGetParams) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoGetParams.ProtoReflect.Descriptor instead.
func (*TodoGetParams) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoGetParams) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TodoAddParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Task   string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TodoAddParams) Reset() {
	*x = TodoAddParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoAddParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoAddParams) ProtoMessage() {}

func (x *TodoAddParams) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoAddParams.ProtoReflect.Descriptor instead.
func (*TodoAddParams) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoAddParams) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TodoAddParams) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type TodoDeleteParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoDeleteParams) Reset() {
	*x = TodoDeleteParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoDeleteParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoDeleteParams) ProtoMessage() {}

func (x *TodoDeleteParams) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoDeleteParams.ProtoReflect.Descriptor instead.
func (*TodoDeleteParams) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoDeleteParams) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TodoDeleteParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TodoObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Task   string `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TodoObject) Reset() {
	*x = TodoObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoObject) ProtoMessage() {}

func (x *TodoObject) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoObject.ProtoReflect.Descriptor instead.
func (*TodoObject) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoObject) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TodoObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoObject) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type TodoChangeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdemponcyId string      `protobuf:"bytes,1,opt,name=idemponcyId,proto3" json:"idemponcyId,omitempty"`
	Current     *TodoObject `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	Original    *TodoObject `protobuf:"bytes,4,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *TodoChangeEvent) Reset() {
	*x = TodoChangeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoChangeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoChangeEvent) ProtoMessage() {}

func (x *TodoChangeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoChangeEvent.ProtoReflect.Descriptor instead.
func (*TodoChangeEvent) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{4}
}

func (x *TodoChangeEvent) GetIdemponcyId() string {
	if x != nil {
		return x.IdemponcyId
	}
	return ""
}

func (x *TodoChangeEvent) GetCurrent() *TodoObject {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *TodoChangeEvent) GetOriginal() *TodoObject {
	if x != nil {
		return x.Original
	}
	return nil
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*TodoObject `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{5}
}

func (x *TodoResponse) GetTodos() []*TodoObject {
	if x != nil {
		return x.Todos
	}
	return nil
}

type TodoDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TodoDeleteResponse) Reset() {
	*x = TodoDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corepb_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoDeleteResponse) ProtoMessage() {}

func (x *TodoDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_corepb_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoDeleteResponse.ProtoReflect.Descriptor instead.
func (*TodoDeleteResponse) Descriptor() ([]byte, []int) {
	return file_corepb_todo_proto_rawDescGZIP(), []int{6}
}

func (x *TodoDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_corepb_todo_proto protoreflect.FileDescriptor

var file_corepb_todo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x22, 0x28, 0x0a, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x54, 0x6f,
	0x64, 0x6f, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x3b, 0x0a, 0x10, 0x54, 0x6f, 0x64, 0x6f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x6e, 0x63,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6d, 0x70,
	0x6f, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x3d,
	0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x2e, 0x0a,
	0x12, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xde, 0x01,
	0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70,
	0x62, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x41, 0x64, 0x64, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4c, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_corepb_todo_proto_rawDescOnce sync.Once
	file_corepb_todo_proto_rawDescData = file_corepb_todo_proto_rawDesc
)

func file_corepb_todo_proto_rawDescGZIP() []byte {
	file_corepb_todo_proto_rawDescOnce.Do(func() {
		file_corepb_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_corepb_todo_proto_rawDescData)
	})
	return file_corepb_todo_proto_rawDescData
}

var file_corepb_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_corepb_todo_proto_goTypes = []interface{}{
	(*TodoGetParams)(nil),      // 0: corepb.todo.TodoGetParams
	(*TodoAddParams)(nil),      // 1: corepb.todo.TodoAddParams
	(*TodoDeleteParams)(nil),   // 2: corepb.todo.TodoDeleteParams
	(*TodoObject)(nil),         // 3: corepb.todo.TodoObject
	(*TodoChangeEvent)(nil),    // 4: corepb.todo.TodoChangeEvent
	(*TodoResponse)(nil),       // 5: corepb.todo.TodoResponse
	(*TodoDeleteResponse)(nil), // 6: corepb.todo.TodoDeleteResponse
}
var file_corepb_todo_proto_depIdxs = []int32{
	3, // 0: corepb.todo.TodoChangeEvent.current:type_name -> corepb.todo.TodoObject
	3, // 1: corepb.todo.TodoChangeEvent.original:type_name -> corepb.todo.TodoObject
	3, // 2: corepb.todo.TodoResponse.todos:type_name -> corepb.todo.TodoObject
	1, // 3: corepb.todo.TodoService.AddTodo:input_type -> corepb.todo.TodoAddParams
	2, // 4: corepb.todo.TodoService.DeleteTodo:input_type -> corepb.todo.TodoDeleteParams
	0, // 5: corepb.todo.TodoService.GetTodos:input_type -> corepb.todo.TodoGetParams
	3, // 6: corepb.todo.TodoService.AddTodo:output_type -> corepb.todo.TodoObject
	6, // 7: corepb.todo.TodoService.DeleteTodo:output_type -> corepb.todo.TodoDeleteResponse
	5, // 8: corepb.todo.TodoService.GetTodos:output_type -> corepb.todo.TodoResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_corepb_todo_proto_init() }
func file_corepb_todo_proto_init() {
	if File_corepb_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_corepb_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoGetParams); i {
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
		file_corepb_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoAddParams); i {
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
		file_corepb_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoDeleteParams); i {
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
		file_corepb_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoObject); i {
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
		file_corepb_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoChangeEvent); i {
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
		file_corepb_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoResponse); i {
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
		file_corepb_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoDeleteResponse); i {
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
			RawDescriptor: file_corepb_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_corepb_todo_proto_goTypes,
		DependencyIndexes: file_corepb_todo_proto_depIdxs,
		MessageInfos:      file_corepb_todo_proto_msgTypes,
	}.Build()
	File_corepb_todo_proto = out.File
	file_corepb_todo_proto_rawDesc = nil
	file_corepb_todo_proto_goTypes = nil
	file_corepb_todo_proto_depIdxs = nil
}
