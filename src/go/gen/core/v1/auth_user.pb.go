// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: core/v1/auth_user.proto

package corev1

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

type AuthUserServiceRemoveAssociationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthUserServiceRemoveAssociationResponse) Reset() {
	*x = AuthUserServiceRemoveAssociationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceRemoveAssociationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceRemoveAssociationResponse) ProtoMessage() {}

func (x *AuthUserServiceRemoveAssociationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceRemoveAssociationResponse.ProtoReflect.Descriptor instead.
func (*AuthUserServiceRemoveAssociationResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{0}
}

type AuthUserServiceGetAuthUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider    string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	RedirectUrl string `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AuthUserServiceGetAuthUrlRequest) Reset() {
	*x = AuthUserServiceGetAuthUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceGetAuthUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceGetAuthUrlRequest) ProtoMessage() {}

func (x *AuthUserServiceGetAuthUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceGetAuthUrlRequest.ProtoReflect.Descriptor instead.
func (*AuthUserServiceGetAuthUrlRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{1}
}

func (x *AuthUserServiceGetAuthUrlRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AuthUserServiceGetAuthUrlRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *AuthUserServiceGetAuthUrlRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AuthUserServiceGetAuthUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *AuthUserServiceGetAuthUrlResponse) Reset() {
	*x = AuthUserServiceGetAuthUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceGetAuthUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceGetAuthUrlResponse) ProtoMessage() {}

func (x *AuthUserServiceGetAuthUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceGetAuthUrlResponse.ProtoReflect.Descriptor instead.
func (*AuthUserServiceGetAuthUrlResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{2}
}

func (x *AuthUserServiceGetAuthUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type AuthUserServiceRemoveAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *AuthUserServiceRemoveAssociationRequest) Reset() {
	*x = AuthUserServiceRemoveAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceRemoveAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceRemoveAssociationRequest) ProtoMessage() {}

func (x *AuthUserServiceRemoveAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceRemoveAssociationRequest.ProtoReflect.Descriptor instead.
func (*AuthUserServiceRemoveAssociationRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{3}
}

func (x *AuthUserServiceRemoveAssociationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthUserServiceRemoveAssociationRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type AuthUserServiceListAssociationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *AuthUserServiceListAssociationsRequest) Reset() {
	*x = AuthUserServiceListAssociationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceListAssociationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceListAssociationsRequest) ProtoMessage() {}

func (x *AuthUserServiceListAssociationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceListAssociationsRequest.ProtoReflect.Descriptor instead.
func (*AuthUserServiceListAssociationsRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{4}
}

func (x *AuthUserServiceListAssociationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthUserServiceListAssociationsRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type AuthOauthParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider   string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AuthOauthParams) Reset() {
	*x = AuthOauthParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthOauthParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthOauthParams) ProtoMessage() {}

func (x *AuthOauthParams) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthOauthParams.ProtoReflect.Descriptor instead.
func (*AuthOauthParams) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{5}
}

func (x *AuthOauthParams) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AuthOauthParams) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *AuthOauthParams) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AuthEmailParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthEmailParams) Reset() {
	*x = AuthEmailParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthEmailParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthEmailParams) ProtoMessage() {}

func (x *AuthEmailParams) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthEmailParams.ProtoReflect.Descriptor instead.
func (*AuthEmailParams) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{6}
}

func (x *AuthEmailParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthEmailParams) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthUserServiceListAssociationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider []string `protobuf:"bytes,1,rep,name=provider,proto3" json:"provider,omitempty"`
}

func (x *AuthUserServiceListAssociationsResponse) Reset() {
	*x = AuthUserServiceListAssociationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceListAssociationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceListAssociationsResponse) ProtoMessage() {}

func (x *AuthUserServiceListAssociationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceListAssociationsResponse.ProtoReflect.Descriptor instead.
func (*AuthUserServiceListAssociationsResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{7}
}

func (x *AuthUserServiceListAssociationsResponse) GetProvider() []string {
	if x != nil {
		return x.Provider
	}
	return nil
}

type AuthUserServiceUpsertUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If provide, then we're additing an association
	UserId      string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Oauth       *AuthOauthParams `protobuf:"bytes,2,opt,name=oauth,proto3" json:"oauth,omitempty"`
	Email       *AuthEmailParams `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	RedirectUrl string           `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	State       string           `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AuthUserServiceUpsertUserRequest) Reset() {
	*x = AuthUserServiceUpsertUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceUpsertUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceUpsertUserRequest) ProtoMessage() {}

func (x *AuthUserServiceUpsertUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceUpsertUserRequest.ProtoReflect.Descriptor instead.
func (*AuthUserServiceUpsertUserRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{8}
}

func (x *AuthUserServiceUpsertUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthUserServiceUpsertUserRequest) GetOauth() *AuthOauthParams {
	if x != nil {
		return x.Oauth
	}
	return nil
}

func (x *AuthUserServiceUpsertUserRequest) GetEmail() *AuthEmailParams {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *AuthUserServiceUpsertUserRequest) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *AuthUserServiceUpsertUserRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AuthUserServiceUpsertUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Created bool   `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *AuthUserServiceUpsertUserResponse) Reset() {
	*x = AuthUserServiceUpsertUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_auth_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserServiceUpsertUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserServiceUpsertUserResponse) ProtoMessage() {}

func (x *AuthUserServiceUpsertUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_auth_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserServiceUpsertUserResponse.ProtoReflect.Descriptor instead.
func (*AuthUserServiceUpsertUserResponse) Descriptor() ([]byte, []int) {
	return file_core_v1_auth_user_proto_rawDescGZIP(), []int{9}
}

func (x *AuthUserServiceUpsertUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthUserServiceUpsertUserResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

var File_core_v1_auth_user_proto protoreflect.FileDescriptor

var file_core_v1_auth_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x2a, 0x0a, 0x28, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77,
	0x0a, 0x20, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x21, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x5e,
	0x0a, 0x27, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x5d,
	0x0a, 0x26, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x62, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x43, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x45, 0x0a, 0x27, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xd4, 0x01,
	0x0a, 0x20, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x56, 0x0a, 0x21, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x32, 0xcc, 0x03, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x63, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x29,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x89, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x62, 0x6c, 0x61, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43,
	0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43,
	0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_auth_user_proto_rawDescOnce sync.Once
	file_core_v1_auth_user_proto_rawDescData = file_core_v1_auth_user_proto_rawDesc
)

func file_core_v1_auth_user_proto_rawDescGZIP() []byte {
	file_core_v1_auth_user_proto_rawDescOnce.Do(func() {
		file_core_v1_auth_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_auth_user_proto_rawDescData)
	})
	return file_core_v1_auth_user_proto_rawDescData
}

var file_core_v1_auth_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_core_v1_auth_user_proto_goTypes = []interface{}{
	(*AuthUserServiceRemoveAssociationResponse)(nil), // 0: core.v1.AuthUserServiceRemoveAssociationResponse
	(*AuthUserServiceGetAuthUrlRequest)(nil),         // 1: core.v1.AuthUserServiceGetAuthUrlRequest
	(*AuthUserServiceGetAuthUrlResponse)(nil),        // 2: core.v1.AuthUserServiceGetAuthUrlResponse
	(*AuthUserServiceRemoveAssociationRequest)(nil),  // 3: core.v1.AuthUserServiceRemoveAssociationRequest
	(*AuthUserServiceListAssociationsRequest)(nil),   // 4: core.v1.AuthUserServiceListAssociationsRequest
	(*AuthOauthParams)(nil),                          // 5: core.v1.AuthOauthParams
	(*AuthEmailParams)(nil),                          // 6: core.v1.AuthEmailParams
	(*AuthUserServiceListAssociationsResponse)(nil),  // 7: core.v1.AuthUserServiceListAssociationsResponse
	(*AuthUserServiceUpsertUserRequest)(nil),         // 8: core.v1.AuthUserServiceUpsertUserRequest
	(*AuthUserServiceUpsertUserResponse)(nil),        // 9: core.v1.AuthUserServiceUpsertUserResponse
}
var file_core_v1_auth_user_proto_depIdxs = []int32{
	5, // 0: core.v1.AuthUserServiceUpsertUserRequest.oauth:type_name -> core.v1.AuthOauthParams
	6, // 1: core.v1.AuthUserServiceUpsertUserRequest.email:type_name -> core.v1.AuthEmailParams
	1, // 2: core.v1.AuthUserService.GetAuthUrl:input_type -> core.v1.AuthUserServiceGetAuthUrlRequest
	4, // 3: core.v1.AuthUserService.ListAssociations:input_type -> core.v1.AuthUserServiceListAssociationsRequest
	3, // 4: core.v1.AuthUserService.RemoveAssociation:input_type -> core.v1.AuthUserServiceRemoveAssociationRequest
	8, // 5: core.v1.AuthUserService.UpsertUser:input_type -> core.v1.AuthUserServiceUpsertUserRequest
	2, // 6: core.v1.AuthUserService.GetAuthUrl:output_type -> core.v1.AuthUserServiceGetAuthUrlResponse
	7, // 7: core.v1.AuthUserService.ListAssociations:output_type -> core.v1.AuthUserServiceListAssociationsResponse
	0, // 8: core.v1.AuthUserService.RemoveAssociation:output_type -> core.v1.AuthUserServiceRemoveAssociationResponse
	9, // 9: core.v1.AuthUserService.UpsertUser:output_type -> core.v1.AuthUserServiceUpsertUserResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_v1_auth_user_proto_init() }
func file_core_v1_auth_user_proto_init() {
	if File_core_v1_auth_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1_auth_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceRemoveAssociationResponse); i {
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
		file_core_v1_auth_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceGetAuthUrlRequest); i {
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
		file_core_v1_auth_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceGetAuthUrlResponse); i {
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
		file_core_v1_auth_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceRemoveAssociationRequest); i {
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
		file_core_v1_auth_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceListAssociationsRequest); i {
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
		file_core_v1_auth_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthOauthParams); i {
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
		file_core_v1_auth_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthEmailParams); i {
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
		file_core_v1_auth_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceListAssociationsResponse); i {
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
		file_core_v1_auth_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceUpsertUserRequest); i {
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
		file_core_v1_auth_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthUserServiceUpsertUserResponse); i {
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
			RawDescriptor: file_core_v1_auth_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_v1_auth_user_proto_goTypes,
		DependencyIndexes: file_core_v1_auth_user_proto_depIdxs,
		MessageInfos:      file_core_v1_auth_user_proto_msgTypes,
	}.Build()
	File_core_v1_auth_user_proto = out.File
	file_core_v1_auth_user_proto_rawDesc = nil
	file_core_v1_auth_user_proto_goTypes = nil
	file_core_v1_auth_user_proto_depIdxs = nil
}