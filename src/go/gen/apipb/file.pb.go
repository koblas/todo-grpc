// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: apipb/file.proto

package apipb

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

// Upload URL RPC
type UploadUrlParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UploadUrlParams) Reset() {
	*x = UploadUrlParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apipb_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadUrlParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUrlParams) ProtoMessage() {}

func (x *UploadUrlParams) ProtoReflect() protoreflect.Message {
	mi := &file_apipb_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUrlParams.ProtoReflect.Descriptor instead.
func (*UploadUrlParams) Descriptor() ([]byte, []int) {
	return file_apipb_file_proto_rawDescGZIP(), []int{0}
}

func (x *UploadUrlParams) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UploadUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadUrlResponse) Reset() {
	*x = UploadUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apipb_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUrlResponse) ProtoMessage() {}

func (x *UploadUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apipb_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUrlResponse.ProtoReflect.Descriptor instead.
func (*UploadUrlResponse) Descriptor() ([]byte, []int) {
	return file_apipb_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_apipb_file_proto protoreflect.FileDescriptor

var file_apipb_file_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x25,
	0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x57, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x70,
	0x62, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x70, 0x62, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apipb_file_proto_rawDescOnce sync.Once
	file_apipb_file_proto_rawDescData = file_apipb_file_proto_rawDesc
)

func file_apipb_file_proto_rawDescGZIP() []byte {
	file_apipb_file_proto_rawDescOnce.Do(func() {
		file_apipb_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_apipb_file_proto_rawDescData)
	})
	return file_apipb_file_proto_rawDescData
}

var file_apipb_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apipb_file_proto_goTypes = []interface{}{
	(*UploadUrlParams)(nil),   // 0: apipb.file.UploadUrlParams
	(*UploadUrlResponse)(nil), // 1: apipb.file.UploadUrlResponse
}
var file_apipb_file_proto_depIdxs = []int32{
	0, // 0: apipb.file.FileService.upload_url:input_type -> apipb.file.UploadUrlParams
	1, // 1: apipb.file.FileService.upload_url:output_type -> apipb.file.UploadUrlResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apipb_file_proto_init() }
func file_apipb_file_proto_init() {
	if File_apipb_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apipb_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadUrlParams); i {
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
		file_apipb_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadUrlResponse); i {
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
			RawDescriptor: file_apipb_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apipb_file_proto_goTypes,
		DependencyIndexes: file_apipb_file_proto_depIdxs,
		MessageInfos:      file_apipb_file_proto_msgTypes,
	}.Build()
	File_apipb_file_proto = out.File
	file_apipb_file_proto_rawDesc = nil
	file_apipb_file_proto_goTypes = nil
	file_apipb_file_proto_depIdxs = nil
}
