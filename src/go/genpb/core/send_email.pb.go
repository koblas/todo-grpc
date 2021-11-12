// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: core/send_email.proto

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

// For pubsub listeners
type EmailTemplate int32

const (
	EmailTemplate_UNDEFINED         EmailTemplate = 0
	EmailTemplate_USER_REGISTERED   EmailTemplate = 1
	EmailTemplate_USER_INVITED      EmailTemplate = 2
	EmailTemplate_PASSWORD_RECOVERY EmailTemplate = 3
	EmailTemplate_PASSWORD_CHANGE   EmailTemplate = 4
)

// Enum value maps for EmailTemplate.
var (
	EmailTemplate_name = map[int32]string{
		0: "UNDEFINED",
		1: "USER_REGISTERED",
		2: "USER_INVITED",
		3: "PASSWORD_RECOVERY",
		4: "PASSWORD_CHANGE",
	}
	EmailTemplate_value = map[string]int32{
		"UNDEFINED":         0,
		"USER_REGISTERED":   1,
		"USER_INVITED":      2,
		"PASSWORD_RECOVERY": 3,
		"PASSWORD_CHANGE":   4,
	}
)

func (x EmailTemplate) Enum() *EmailTemplate {
	p := new(EmailTemplate)
	*p = x
	return p
}

func (x EmailTemplate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailTemplate) Descriptor() protoreflect.EnumDescriptor {
	return file_core_send_email_proto_enumTypes[0].Descriptor()
}

func (EmailTemplate) Type() protoreflect.EnumType {
	return &file_core_send_email_proto_enumTypes[0]
}

func (x EmailTemplate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailTemplate.Descriptor instead.
func (EmailTemplate) EnumDescriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{0}
}

// Basic user email info
type EmailUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EmailUser) Reset() {
	*x = EmailUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailUser) ProtoMessage() {}

func (x *EmailUser) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailUser.ProtoReflect.Descriptor instead.
func (*EmailUser) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{0}
}

func (x *EmailUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Basic Application information for the message
type EmailAppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlBase     string `protobuf:"bytes,1,opt,name=url_base,json=urlBase,proto3" json:"url_base,omitempty"`
	AppName     string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	SenderEmail string `protobuf:"bytes,3,opt,name=sender_email,json=senderEmail,proto3" json:"sender_email,omitempty"`
	SenderName  string `protobuf:"bytes,4,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`
}

func (x *EmailAppInfo) Reset() {
	*x = EmailAppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailAppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailAppInfo) ProtoMessage() {}

func (x *EmailAppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailAppInfo.ProtoReflect.Descriptor instead.
func (*EmailAppInfo) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{1}
}

func (x *EmailAppInfo) GetUrlBase() string {
	if x != nil {
		return x.UrlBase
	}
	return ""
}

func (x *EmailAppInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *EmailAppInfo) GetSenderEmail() string {
	if x != nil {
		return x.SenderEmail
	}
	return ""
}

func (x *EmailAppInfo) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

//
//  Specific messages
//
type EmailRegisterParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfo     *EmailAppInfo `protobuf:"bytes,1,opt,name=app_info,json=appInfo,proto3" json:"app_info,omitempty"`
	Recipient   *EmailUser    `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ReferenceId string        `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Token       string        `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *EmailRegisterParam) Reset() {
	*x = EmailRegisterParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRegisterParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRegisterParam) ProtoMessage() {}

func (x *EmailRegisterParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRegisterParam.ProtoReflect.Descriptor instead.
func (*EmailRegisterParam) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{2}
}

func (x *EmailRegisterParam) GetAppInfo() *EmailAppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

func (x *EmailRegisterParam) GetRecipient() *EmailUser {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *EmailRegisterParam) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *EmailRegisterParam) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type EmailPasswordChangeParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfo     *EmailAppInfo `protobuf:"bytes,1,opt,name=app_info,json=appInfo,proto3" json:"app_info,omitempty"`
	Recipient   *EmailUser    `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ReferenceId string        `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
}

func (x *EmailPasswordChangeParam) Reset() {
	*x = EmailPasswordChangeParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailPasswordChangeParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailPasswordChangeParam) ProtoMessage() {}

func (x *EmailPasswordChangeParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailPasswordChangeParam.ProtoReflect.Descriptor instead.
func (*EmailPasswordChangeParam) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{3}
}

func (x *EmailPasswordChangeParam) GetAppInfo() *EmailAppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

func (x *EmailPasswordChangeParam) GetRecipient() *EmailUser {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *EmailPasswordChangeParam) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

type EmailPasswordRecoveryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfo     *EmailAppInfo `protobuf:"bytes,1,opt,name=app_info,json=appInfo,proto3" json:"app_info,omitempty"`
	Recipient   *EmailUser    `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ReferenceId string        `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Token       string        `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *EmailPasswordRecoveryParam) Reset() {
	*x = EmailPasswordRecoveryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailPasswordRecoveryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailPasswordRecoveryParam) ProtoMessage() {}

func (x *EmailPasswordRecoveryParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailPasswordRecoveryParam.ProtoReflect.Descriptor instead.
func (*EmailPasswordRecoveryParam) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{4}
}

func (x *EmailPasswordRecoveryParam) GetAppInfo() *EmailAppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

func (x *EmailPasswordRecoveryParam) GetRecipient() *EmailUser {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *EmailPasswordRecoveryParam) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *EmailPasswordRecoveryParam) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type EmailInviteUserParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfo     *EmailAppInfo `protobuf:"bytes,1,opt,name=app_info,json=appInfo,proto3" json:"app_info,omitempty"`
	Sender      *EmailUser    `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient   *EmailUser    `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	ReferenceId string        `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Token       string        `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *EmailInviteUserParam) Reset() {
	*x = EmailInviteUserParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailInviteUserParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailInviteUserParam) ProtoMessage() {}

func (x *EmailInviteUserParam) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailInviteUserParam.ProtoReflect.Descriptor instead.
func (*EmailInviteUserParam) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{5}
}

func (x *EmailInviteUserParam) GetAppInfo() *EmailAppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

func (x *EmailInviteUserParam) GetSender() *EmailUser {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *EmailInviteUserParam) GetRecipient() *EmailUser {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *EmailInviteUserParam) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *EmailInviteUserParam) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type EmailOkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *EmailOkResponse) Reset() {
	*x = EmailOkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailOkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailOkResponse) ProtoMessage() {}

func (x *EmailOkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailOkResponse.ProtoReflect.Descriptor instead.
func (*EmailOkResponse) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{6}
}

func (x *EmailOkResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type EmailSentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientEmail string        `protobuf:"bytes,1,opt,name=recipient_email,json=recipientEmail,proto3" json:"recipient_email,omitempty"`
	MessageId      string        `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Template       EmailTemplate `protobuf:"varint,3,opt,name=template,proto3,enum=core.EmailTemplate" json:"template,omitempty"`
	ReferenceId    string        `protobuf:"bytes,4,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
}

func (x *EmailSentEvent) Reset() {
	*x = EmailSentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_send_email_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailSentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailSentEvent) ProtoMessage() {}

func (x *EmailSentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_core_send_email_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailSentEvent.ProtoReflect.Descriptor instead.
func (*EmailSentEvent) Descriptor() ([]byte, []int) {
	return file_core_send_email_proto_rawDescGZIP(), []int{7}
}

func (x *EmailSentEvent) GetRecipientEmail() string {
	if x != nil {
		return x.RecipientEmail
	}
	return ""
}

func (x *EmailSentEvent) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *EmailSentEvent) GetTemplate() EmailTemplate {
	if x != nil {
		return x.Template
	}
	return EmailTemplate_UNDEFINED
}

func (x *EmailSentEvent) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

var File_core_send_email_proto protoreflect.FileDescriptor

var file_core_send_email_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x35, 0x0a,
	0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x42, 0x61, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xab, 0x01, 0x0a, 0x12, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9b, 0x01,
	0x0a, 0x18, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x1a,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x14, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0xac, 0x01,
	0x0a, 0x0e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x2a, 0x71, 0x0a, 0x0d,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x04, 0x32,
	0xca, 0x02, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x15, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x17,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_send_email_proto_rawDescOnce sync.Once
	file_core_send_email_proto_rawDescData = file_core_send_email_proto_rawDesc
)

func file_core_send_email_proto_rawDescGZIP() []byte {
	file_core_send_email_proto_rawDescOnce.Do(func() {
		file_core_send_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_send_email_proto_rawDescData)
	})
	return file_core_send_email_proto_rawDescData
}

var file_core_send_email_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_send_email_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_core_send_email_proto_goTypes = []interface{}{
	(EmailTemplate)(0),                 // 0: core.EmailTemplate
	(*EmailUser)(nil),                  // 1: core.EmailUser
	(*EmailAppInfo)(nil),               // 2: core.EmailAppInfo
	(*EmailRegisterParam)(nil),         // 3: core.EmailRegisterParam
	(*EmailPasswordChangeParam)(nil),   // 4: core.EmailPasswordChangeParam
	(*EmailPasswordRecoveryParam)(nil), // 5: core.EmailPasswordRecoveryParam
	(*EmailInviteUserParam)(nil),       // 6: core.EmailInviteUserParam
	(*EmailOkResponse)(nil),            // 7: core.EmailOkResponse
	(*EmailSentEvent)(nil),             // 8: core.EmailSentEvent
}
var file_core_send_email_proto_depIdxs = []int32{
	2,  // 0: core.EmailRegisterParam.app_info:type_name -> core.EmailAppInfo
	1,  // 1: core.EmailRegisterParam.recipient:type_name -> core.EmailUser
	2,  // 2: core.EmailPasswordChangeParam.app_info:type_name -> core.EmailAppInfo
	1,  // 3: core.EmailPasswordChangeParam.recipient:type_name -> core.EmailUser
	2,  // 4: core.EmailPasswordRecoveryParam.app_info:type_name -> core.EmailAppInfo
	1,  // 5: core.EmailPasswordRecoveryParam.recipient:type_name -> core.EmailUser
	2,  // 6: core.EmailInviteUserParam.app_info:type_name -> core.EmailAppInfo
	1,  // 7: core.EmailInviteUserParam.sender:type_name -> core.EmailUser
	1,  // 8: core.EmailInviteUserParam.recipient:type_name -> core.EmailUser
	0,  // 9: core.EmailSentEvent.template:type_name -> core.EmailTemplate
	3,  // 10: core.SendEmailService.RegisterMessage:input_type -> core.EmailRegisterParam
	4,  // 11: core.SendEmailService.PasswordChangeMessage:input_type -> core.EmailPasswordChangeParam
	5,  // 12: core.SendEmailService.PasswordRecoveryMessage:input_type -> core.EmailPasswordRecoveryParam
	6,  // 13: core.SendEmailService.InviteUserMessage:input_type -> core.EmailInviteUserParam
	7,  // 14: core.SendEmailService.RegisterMessage:output_type -> core.EmailOkResponse
	7,  // 15: core.SendEmailService.PasswordChangeMessage:output_type -> core.EmailOkResponse
	7,  // 16: core.SendEmailService.PasswordRecoveryMessage:output_type -> core.EmailOkResponse
	7,  // 17: core.SendEmailService.InviteUserMessage:output_type -> core.EmailOkResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_core_send_email_proto_init() }
func file_core_send_email_proto_init() {
	if File_core_send_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_send_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailUser); i {
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
		file_core_send_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailAppInfo); i {
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
		file_core_send_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRegisterParam); i {
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
		file_core_send_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailPasswordChangeParam); i {
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
		file_core_send_email_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailPasswordRecoveryParam); i {
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
		file_core_send_email_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailInviteUserParam); i {
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
		file_core_send_email_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailOkResponse); i {
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
		file_core_send_email_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailSentEvent); i {
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
			RawDescriptor: file_core_send_email_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_send_email_proto_goTypes,
		DependencyIndexes: file_core_send_email_proto_depIdxs,
		EnumInfos:         file_core_send_email_proto_enumTypes,
		MessageInfos:      file_core_send_email_proto_msgTypes,
	}.Build()
	File_core_send_email_proto = out.File
	file_core_send_email_proto_rawDesc = nil
	file_core_send_email_proto_goTypes = nil
	file_core_send_email_proto_depIdxs = nil
}
