// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: publicapi/auth.proto

package publicapi

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RegisterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Required
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Username
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// invite code
	Invite string `protobuf:"bytes,5,opt,name=invite,proto3" json:"invite,omitempty"`
}

func (x *RegisterParams) Reset() {
	*x = RegisterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterParams) ProtoMessage() {}

func (x *RegisterParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterParams.ProtoReflect.Descriptor instead.
func (*RegisterParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterParams) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterParams) GetInvite() string {
	if x != nil {
		return x.Invite
	}
	return ""
}

type LoginParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Required
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// TFA One Time Token
	TfaOtp string `protobuf:"bytes,4,opt,name=tfa_otp,json=tfaOtp,proto3" json:"tfa_otp,omitempty"`
	// TFA Type
	TfaType string `protobuf:"bytes,5,opt,name=tfa_type,json=tfaType,proto3" json:"tfa_type,omitempty"`
}

func (x *LoginParams) Reset() {
	*x = LoginParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginParams) ProtoMessage() {}

func (x *LoginParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginParams.ProtoReflect.Descriptor instead.
func (*LoginParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginParams) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginParams) GetTfaOtp() string {
	if x != nil {
		return x.TfaOtp
	}
	return ""
}

func (x *LoginParams) GetTfaType() string {
	if x != nil {
		return x.TfaType
	}
	return ""
}

type ConfirmParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: Verify and Update
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Email confirmation token
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ConfirmParams) Reset() {
	*x = ConfirmParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmParams) ProtoMessage() {}

func (x *ConfirmParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmParams.ProtoReflect.Descriptor instead.
func (*ConfirmParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmParams) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConfirmParams) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RecoverySendParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: Send
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RecoverySendParams) Reset() {
	*x = RecoverySendParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverySendParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverySendParams) ProtoMessage() {}

func (x *RecoverySendParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverySendParams.ProtoReflect.Descriptor instead.
func (*RecoverySendParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RecoverySendParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RecoveryUpdateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: Verify and Update
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Required: Verify and Update
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// Required: Update
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RecoveryUpdateParams) Reset() {
	*x = RecoveryUpdateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoveryUpdateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoveryUpdateParams) ProtoMessage() {}

func (x *RecoveryUpdateParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoveryUpdateParams.ProtoReflect.Descriptor instead.
func (*RecoveryUpdateParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RecoveryUpdateParams) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RecoveryUpdateParams) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RecoveryUpdateParams) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Oauth
type OauthUrlParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider    string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	RedirectUrl string `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *OauthUrlParams) Reset() {
	*x = OauthUrlParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthUrlParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthUrlParams) ProtoMessage() {}

func (x *OauthUrlParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthUrlParams.ProtoReflect.Descriptor instead.
func (*OauthUrlParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{5}
}

func (x *OauthUrlParams) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OauthUrlParams) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *OauthUrlParams) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type OauthUrlResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OauthUrlResult) Reset() {
	*x = OauthUrlResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthUrlResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthUrlResult) ProtoMessage() {}

func (x *OauthUrlResult) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthUrlResult.ProtoReflect.Descriptor instead.
func (*OauthUrlResult) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{6}
}

func (x *OauthUrlResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type OauthAssociateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider    string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	RedirectUrl string `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	State       string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *OauthAssociateParams) Reset() {
	*x = OauthAssociateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthAssociateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthAssociateParams) ProtoMessage() {}

func (x *OauthAssociateParams) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthAssociateParams.ProtoReflect.Descriptor instead.
func (*OauthAssociateParams) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{7}
}

func (x *OauthAssociateParams) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OauthAssociateParams) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *OauthAssociateParams) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OauthAssociateParams) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type ValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field name
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// Human readable message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ValidationError) Reset() {
	*x = ValidationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationError) ProtoMessage() {}

func (x *ValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationError.ProtoReflect.Descriptor instead.
func (*ValidationError) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{8}
}

func (x *ValidationError) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *ValidationError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{9}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	ExpiresIn    int32  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{10}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *Token) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type TokenRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// More for the OAuth case to distingish between new account and old
	Created bool `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *TokenRegister) Reset() {
	*x = TokenRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicapi_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRegister) ProtoMessage() {}

func (x *TokenRegister) ProtoReflect() protoreflect.Message {
	mi := &file_publicapi_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRegister.ProtoReflect.Descriptor instead.
func (*TokenRegister) Descriptor() ([]byte, []int) {
	return file_publicapi_auth_proto_rawDescGZIP(), []int{11}
}

func (x *TokenRegister) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokenRegister) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

var File_publicapi_auth_proto protoreflect.FileDescriptor

var file_publicapi_auth_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x22, 0x73,
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x66, 0x61, 0x5f, 0x6f, 0x74, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x66, 0x61, 0x4f, 0x74, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x66, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x66, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x61, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x65, 0x0a, 0x0e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x7f, 0x0a,
	0x14, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x41,
	0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x50, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x32, 0x90, 0x06, 0x0a, 0x15, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x5e,
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x64,
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0b, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x5b,
	0x0a, 0x09, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x72, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_publicapi_auth_proto_rawDescOnce sync.Once
	file_publicapi_auth_proto_rawDescData = file_publicapi_auth_proto_rawDesc
)

func file_publicapi_auth_proto_rawDescGZIP() []byte {
	file_publicapi_auth_proto_rawDescOnce.Do(func() {
		file_publicapi_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicapi_auth_proto_rawDescData)
	})
	return file_publicapi_auth_proto_rawDescData
}

var file_publicapi_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_publicapi_auth_proto_goTypes = []interface{}{
	(*RegisterParams)(nil),       // 0: api.auth.RegisterParams
	(*LoginParams)(nil),          // 1: api.auth.LoginParams
	(*ConfirmParams)(nil),        // 2: api.auth.ConfirmParams
	(*RecoverySendParams)(nil),   // 3: api.auth.RecoverySendParams
	(*RecoveryUpdateParams)(nil), // 4: api.auth.RecoveryUpdateParams
	(*OauthUrlParams)(nil),       // 5: api.auth.OauthUrlParams
	(*OauthUrlResult)(nil),       // 6: api.auth.OauthUrlResult
	(*OauthAssociateParams)(nil), // 7: api.auth.OauthAssociateParams
	(*ValidationError)(nil),      // 8: api.auth.ValidationError
	(*Success)(nil),              // 9: api.auth.Success
	(*Token)(nil),                // 10: api.auth.Token
	(*TokenRegister)(nil),        // 11: api.auth.TokenRegister
}
var file_publicapi_auth_proto_depIdxs = []int32{
	10, // 0: api.auth.TokenRegister.token:type_name -> api.auth.Token
	0,  // 1: api.auth.AuthenticationService.register:input_type -> api.auth.RegisterParams
	1,  // 2: api.auth.AuthenticationService.authenticate:input_type -> api.auth.LoginParams
	2,  // 3: api.auth.AuthenticationService.verify_email:input_type -> api.auth.ConfirmParams
	3,  // 4: api.auth.AuthenticationService.recover_send:input_type -> api.auth.RecoverySendParams
	4,  // 5: api.auth.AuthenticationService.recover_verify:input_type -> api.auth.RecoveryUpdateParams
	4,  // 6: api.auth.AuthenticationService.recover_update:input_type -> api.auth.RecoveryUpdateParams
	7,  // 7: api.auth.AuthenticationService.oauth_login:input_type -> api.auth.OauthAssociateParams
	5,  // 8: api.auth.AuthenticationService.oauth_url:input_type -> api.auth.OauthUrlParams
	11, // 9: api.auth.AuthenticationService.register:output_type -> api.auth.TokenRegister
	10, // 10: api.auth.AuthenticationService.authenticate:output_type -> api.auth.Token
	9,  // 11: api.auth.AuthenticationService.verify_email:output_type -> api.auth.Success
	9,  // 12: api.auth.AuthenticationService.recover_send:output_type -> api.auth.Success
	9,  // 13: api.auth.AuthenticationService.recover_verify:output_type -> api.auth.Success
	10, // 14: api.auth.AuthenticationService.recover_update:output_type -> api.auth.Token
	11, // 15: api.auth.AuthenticationService.oauth_login:output_type -> api.auth.TokenRegister
	6,  // 16: api.auth.AuthenticationService.oauth_url:output_type -> api.auth.OauthUrlResult
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_publicapi_auth_proto_init() }
func file_publicapi_auth_proto_init() {
	if File_publicapi_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publicapi_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterParams); i {
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
		file_publicapi_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginParams); i {
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
		file_publicapi_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmParams); i {
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
		file_publicapi_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverySendParams); i {
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
		file_publicapi_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryUpdateParams); i {
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
		file_publicapi_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthUrlParams); i {
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
		file_publicapi_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthUrlResult); i {
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
		file_publicapi_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthAssociateParams); i {
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
		file_publicapi_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationError); i {
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
		file_publicapi_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_publicapi_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_publicapi_auth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRegister); i {
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
			RawDescriptor: file_publicapi_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publicapi_auth_proto_goTypes,
		DependencyIndexes: file_publicapi_auth_proto_depIdxs,
		MessageInfos:      file_publicapi_auth_proto_msgTypes,
	}.Build()
	File_publicapi_auth_proto = out.File
	file_publicapi_auth_proto_rawDesc = nil
	file_publicapi_auth_proto_goTypes = nil
	file_publicapi_auth_proto_depIdxs = nil
}
