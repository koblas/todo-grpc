// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: api/file/v1/file.proto

package filev1

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
type FileServiceUploadUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *FileServiceUploadUrlRequest) Reset() {
	*x = FileServiceUploadUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileServiceUploadUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileServiceUploadUrlRequest) ProtoMessage() {}

func (x *FileServiceUploadUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileServiceUploadUrlRequest.ProtoReflect.Descriptor instead.
func (*FileServiceUploadUrlRequest) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileServiceUploadUrlRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FileServiceUploadUrlRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type FileServiceUploadUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileServiceUploadUrlResponse) Reset() {
	*x = FileServiceUploadUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileServiceUploadUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileServiceUploadUrlResponse) ProtoMessage() {}

func (x *FileServiceUploadUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileServiceUploadUrlResponse.ProtoReflect.Descriptor instead.
func (*FileServiceUploadUrlResponse) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileServiceUploadUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FileServiceUploadUrlResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_file_v1_file_proto protoreflect.FileDescriptor

var file_api_file_v1_file_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x54, 0x0a, 0x1b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x40, 0x0a, 0x1c, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x70, 0x0a,
	0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0a,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x9e, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x62,
	0x6c, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x69,
	0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x46, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x46,
	0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x46, 0x69, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x46, 0x69, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_file_v1_file_proto_rawDescOnce sync.Once
	file_api_file_v1_file_proto_rawDescData = file_api_file_v1_file_proto_rawDesc
)

func file_api_file_v1_file_proto_rawDescGZIP() []byte {
	file_api_file_v1_file_proto_rawDescOnce.Do(func() {
		file_api_file_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_file_v1_file_proto_rawDescData)
	})
	return file_api_file_v1_file_proto_rawDescData
}

var file_api_file_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_file_v1_file_proto_goTypes = []interface{}{
	(*FileServiceUploadUrlRequest)(nil),  // 0: api.file.v1.FileServiceUploadUrlRequest
	(*FileServiceUploadUrlResponse)(nil), // 1: api.file.v1.FileServiceUploadUrlResponse
}
var file_api_file_v1_file_proto_depIdxs = []int32{
	0, // 0: api.file.v1.FileService.upload_url:input_type -> api.file.v1.FileServiceUploadUrlRequest
	1, // 1: api.file.v1.FileService.upload_url:output_type -> api.file.v1.FileServiceUploadUrlResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_file_v1_file_proto_init() }
func file_api_file_v1_file_proto_init() {
	if File_api_file_v1_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_file_v1_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileServiceUploadUrlRequest); i {
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
		file_api_file_v1_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileServiceUploadUrlResponse); i {
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
			RawDescriptor: file_api_file_v1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_file_v1_file_proto_goTypes,
		DependencyIndexes: file_api_file_v1_file_proto_depIdxs,
		MessageInfos:      file_api_file_v1_file_proto_msgTypes,
	}.Build()
	File_api_file_v1_file_proto = out.File
	file_api_file_v1_file_proto_rawDesc = nil
	file_api_file_v1_file_proto_goTypes = nil
	file_api_file_v1_file_proto_depIdxs = nil
}
