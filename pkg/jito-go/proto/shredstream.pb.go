// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: shredstream.proto

package proto

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

type HeartbeatShredStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// don't trust IP:PORT from tcp header since it can be tampered over the wire
	// `socket.ip` must match incoming packet's ip. this prevents spamming an unwitting destination
	Socket *Socket `protobuf:"bytes,1,opt,name=socket,proto3" json:"socket,omitempty"`
	// regions for shredstream proxy to receive shreds from
	// list of valid regions: https://jito-labs.gitbook.io/mev/systems/connecting/mainnet
	Regions []string `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *HeartbeatShredStream) Reset() {
	*x = HeartbeatShredStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shredstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatShredStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatShredStream) ProtoMessage() {}

func (x *HeartbeatShredStream) ProtoReflect() protoreflect.Message {
	mi := &file_shredstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*HeartbeatShredStream) Descriptor() ([]byte, []int) {
	return file_shredstream_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatShredStream) GetSocket() *Socket {
	if x != nil {
		return x.Socket
	}
	return nil
}

func (x *HeartbeatShredStream) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// client must respond within `ttl_ms` to keep stream alive
	TtlMs uint32 `protobuf:"varint,1,opt,name=ttl_ms,json=ttlMs,proto3" json:"ttl_ms,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shredstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shredstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_shredstream_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatResponse) GetTtlMs() uint32 {
	if x != nil {
		return x.TtlMs
	}
	return 0
}

var File_shredstream_proto protoreflect.FileDescriptor

var file_shredstream_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x68, 0x72, 0x65, 0x64, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x68, 0x72, 0x65, 0x64, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d,
	0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2a, 0x0a,
	0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x74, 0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x74, 0x74, 0x6c, 0x4d, 0x73, 0x32, 0x58, 0x0a, 0x0b, 0x53, 0x68, 0x72,
	0x65, 0x64, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x49, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x72, 0x65,
	0x64, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x72, 0x65, 0x64, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shredstream_proto_rawDescOnce sync.Once
	file_shredstream_proto_rawDescData = file_shredstream_proto_rawDesc
)

func file_shredstream_proto_rawDescGZIP() []byte {
	file_shredstream_proto_rawDescOnce.Do(func() {
		file_shredstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_shredstream_proto_rawDescData)
	})
	return file_shredstream_proto_rawDescData
}

var file_shredstream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shredstream_proto_goTypes = []interface{}{
	(*Heartbeat)(nil),         // 0: shredstream.Heartbeat
	(*HeartbeatResponse)(nil), // 1: shredstream.HeartbeatResponse
	(*Socket)(nil),            // 2: shared.Socket
}
var file_shredstream_proto_depIdxs = []int32{
	2, // 0: shredstream.Heartbeat.socket:type_name -> shared.Socket
	0, // 1: shredstream.Shredstream.SendHeartbeat:input_type -> shredstream.Heartbeat
	1, // 2: shredstream.Shredstream.SendHeartbeat:output_type -> shredstream.HeartbeatResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shredstream_proto_init() }
func file_shredstream_proto_init() {
	if File_shredstream_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shredstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
		file_shredstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
			RawDescriptor: file_shredstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shredstream_proto_goTypes,
		DependencyIndexes: file_shredstream_proto_depIdxs,
		MessageInfos:      file_shredstream_proto_msgTypes,
	}.Build()
	File_shredstream_proto = out.File
	file_shredstream_proto_rawDesc = nil
	file_shredstream_proto_goTypes = nil
	file_shredstream_proto_depIdxs = nil
}
