// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: internal/server/api/v1/log.proto

package log_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Record struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []byte                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset        uint64                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_internal_server_api_v1_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_api_v1_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_internal_server_api_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Record) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_internal_server_api_v1_log_proto protoreflect.FileDescriptor

const file_internal_server_api_v1_log_proto_rawDesc = "" +
	"\n" +
	" internal/server/api/v1/log.proto\x12\x06log.v1\"6\n" +
	"\x06Record\x12\x14\n" +
	"\x05value\x18\x01 \x01(\fR\x05value\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x04R\x06offsetB#Z!github.com/Dhairya3124/api/log_v1b\x06proto3"

var (
	file_internal_server_api_v1_log_proto_rawDescOnce sync.Once
	file_internal_server_api_v1_log_proto_rawDescData []byte
)

func file_internal_server_api_v1_log_proto_rawDescGZIP() []byte {
	file_internal_server_api_v1_log_proto_rawDescOnce.Do(func() {
		file_internal_server_api_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_server_api_v1_log_proto_rawDesc), len(file_internal_server_api_v1_log_proto_rawDesc)))
	})
	return file_internal_server_api_v1_log_proto_rawDescData
}

var file_internal_server_api_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_server_api_v1_log_proto_goTypes = []any{
	(*Record)(nil), // 0: log.v1.Record
}
var file_internal_server_api_v1_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_server_api_v1_log_proto_init() }
func file_internal_server_api_v1_log_proto_init() {
	if File_internal_server_api_v1_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_server_api_v1_log_proto_rawDesc), len(file_internal_server_api_v1_log_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_server_api_v1_log_proto_goTypes,
		DependencyIndexes: file_internal_server_api_v1_log_proto_depIdxs,
		MessageInfos:      file_internal_server_api_v1_log_proto_msgTypes,
	}.Build()
	File_internal_server_api_v1_log_proto = out.File
	file_internal_server_api_v1_log_proto_goTypes = nil
	file_internal_server_api_v1_log_proto_depIdxs = nil
}
