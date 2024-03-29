// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: core/v1/secure_value.proto

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

type SecureValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g.
	// "arn:aws:kms:us-east-1:111111111:key/abcdef00-0123-4567-7799-012345678990"
	KeyUri string `protobuf:"bytes,1,opt,name=key_uri,json=keyUri,proto3" json:"key_uri,omitempty"`
	// opaque value that key encrypted
	DataKey string `protobuf:"bytes,2,opt,name=data_key,json=dataKey,proto3" json:"data_key,omitempty"`
	// encrypted value
	DataValue string `protobuf:"bytes,3,opt,name=data_value,json=dataValue,proto3" json:"data_value,omitempty"`
}

func (x *SecureValue) Reset() {
	*x = SecureValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_secure_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecureValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecureValue) ProtoMessage() {}

func (x *SecureValue) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_secure_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecureValue.ProtoReflect.Descriptor instead.
func (*SecureValue) Descriptor() ([]byte, []int) {
	return file_core_v1_secure_value_proto_rawDescGZIP(), []int{0}
}

func (x *SecureValue) GetKeyUri() string {
	if x != nil {
		return x.KeyUri
	}
	return ""
}

func (x *SecureValue) GetDataKey() string {
	if x != nil {
		return x.DataKey
	}
	return ""
}

func (x *SecureValue) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

var File_core_v1_secure_value_proto protoreflect.FileDescriptor

var file_core_v1_secure_value_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x60, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x55, 0x72, 0x69, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x8c, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x62, 0x6c, 0x61, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43, 0x6f,
	0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43, 0x6f,
	0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_secure_value_proto_rawDescOnce sync.Once
	file_core_v1_secure_value_proto_rawDescData = file_core_v1_secure_value_proto_rawDesc
)

func file_core_v1_secure_value_proto_rawDescGZIP() []byte {
	file_core_v1_secure_value_proto_rawDescOnce.Do(func() {
		file_core_v1_secure_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_secure_value_proto_rawDescData)
	})
	return file_core_v1_secure_value_proto_rawDescData
}

var file_core_v1_secure_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_v1_secure_value_proto_goTypes = []interface{}{
	(*SecureValue)(nil), // 0: core.v1.SecureValue
}
var file_core_v1_secure_value_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_v1_secure_value_proto_init() }
func file_core_v1_secure_value_proto_init() {
	if File_core_v1_secure_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1_secure_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecureValue); i {
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
			RawDescriptor: file_core_v1_secure_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_v1_secure_value_proto_goTypes,
		DependencyIndexes: file_core_v1_secure_value_proto_depIdxs,
		MessageInfos:      file_core_v1_secure_value_proto_msgTypes,
	}.Build()
	File_core_v1_secure_value_proto = out.File
	file_core_v1_secure_value_proto_rawDesc = nil
	file_core_v1_secure_value_proto_goTypes = nil
	file_core_v1_secure_value_proto_depIdxs = nil
}
