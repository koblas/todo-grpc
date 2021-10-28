// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: core/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Find by email address
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindParam) Reset() {
	*x = FindParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindParam) ProtoMessage() {}

func (x *FindParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindParam.ProtoReflect.Descriptor instead.
func (*FindParam) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{1}
}

func (x *FindParam) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AuthenticateParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User Identifier
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticateParam) Reset() {
	*x = AuthenticateParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateParam) ProtoMessage() {}

func (x *AuthenticateParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateParam.ProtoReflect.Descriptor instead.
func (*AuthenticateParam) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticateParam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthenticateParam) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is required
	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateParam) Reset() {
	*x = CreateParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateParam) ProtoMessage() {}

func (x *CreateParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateParam.ProtoReflect.Descriptor instead.
func (*CreateParam) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{3}
}

func (x *CreateParam) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateParam) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is required
	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdateParam) Reset() {
	*x = UpdateParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateParam) ProtoMessage() {}

func (x *UpdateParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateParam.ProtoReflect.Descriptor instead.
func (*UpdateParam) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateParam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateParam) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateParam) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserEither struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	User  *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserEither) Reset() {
	*x = UserEither{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEither) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEither) ProtoMessage() {}

func (x *UserEither) ProtoReflect() protoreflect.Message {
	mi := &file_core_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEither.ProtoReflect.Descriptor instead.
func (*UserEither) Descriptor() ([]byte, []int) {
	return file_core_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserEither) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UserEither) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_core_user_proto protoreflect.FileDescriptor

var file_core_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x40, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x21, 0x0a, 0x09, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x48, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x53, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x69, 0x74, 0x68, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xe0, 0x01, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x66, 0x69, 0x6e,
	0x64, 0x5f, 0x62, 0x79, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x69, 0x74, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x69, 0x74, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x69, 0x74, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x69, 0x74, 0x68, 0x65, 0x72, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_user_proto_rawDescOnce sync.Once
	file_core_user_proto_rawDescData = file_core_user_proto_rawDesc
)

func file_core_user_proto_rawDescGZIP() []byte {
	file_core_user_proto_rawDescOnce.Do(func() {
		file_core_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_user_proto_rawDescData)
	})
	return file_core_user_proto_rawDescData
}

var file_core_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_core_user_proto_goTypes = []interface{}{
	(*User)(nil),              // 0: core.User
	(*FindParam)(nil),         // 1: core.FindParam
	(*AuthenticateParam)(nil), // 2: core.AuthenticateParam
	(*CreateParam)(nil),       // 3: core.CreateParam
	(*UpdateParam)(nil),       // 4: core.UpdateParam
	(*Error)(nil),             // 5: core.Error
	(*UserEither)(nil),        // 6: core.UserEither
}
var file_core_user_proto_depIdxs = []int32{
	5, // 0: core.UserEither.error:type_name -> core.Error
	0, // 1: core.UserEither.user:type_name -> core.User
	1, // 2: core.userService.find_by:input_type -> core.FindParam
	3, // 3: core.userService.create:input_type -> core.CreateParam
	4, // 4: core.userService.update:input_type -> core.UpdateParam
	2, // 5: core.userService.compare_password:input_type -> core.AuthenticateParam
	6, // 6: core.userService.find_by:output_type -> core.UserEither
	6, // 7: core.userService.create:output_type -> core.UserEither
	6, // 8: core.userService.update:output_type -> core.UserEither
	6, // 9: core.userService.compare_password:output_type -> core.UserEither
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_user_proto_init() }
func file_core_user_proto_init() {
	if File_core_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_core_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindParam); i {
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
		file_core_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateParam); i {
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
		file_core_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateParam); i {
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
		file_core_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateParam); i {
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
		file_core_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_core_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEither); i {
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
			RawDescriptor: file_core_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_user_proto_goTypes,
		DependencyIndexes: file_core_user_proto_depIdxs,
		MessageInfos:      file_core_user_proto_msgTypes,
	}.Build()
	File_core_user_proto = out.File
	file_core_user_proto_rawDesc = nil
	file_core_user_proto_goTypes = nil
	file_core_user_proto_depIdxs = nil
}