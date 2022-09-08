// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: core/peer.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	proto "github.com/stv0g/cunicu/pkg/proto"
	epdisc "github.com/stv0g/cunicu/pkg/proto/feat/epdisc"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public WireGuard X25519 key
	PublicKey    []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PresharedKey []byte `protobuf:"bytes,2,opt,name=preshared_key,json=presharedKey,proto3" json:"preshared_key,omitempty"`
	// List of allowed IPs
	AllowedIps                  []string `protobuf:"bytes,3,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	PersistentKeepaliveInterval uint32   `protobuf:"varint,4,opt,name=persistent_keepalive_interval,json=persistentKeepaliveInterval,proto3" json:"persistent_keepalive_interval,omitempty"`
	// Timestamps
	LastHandshakeTimestamp *proto.Timestamp `protobuf:"bytes,5,opt,name=last_handshake_timestamp,json=lastHandshakeTimestamp,proto3" json:"last_handshake_timestamp,omitempty"`
	LastReceiveTimestamp   *proto.Timestamp `protobuf:"bytes,6,opt,name=last_receive_timestamp,json=lastReceiveTimestamp,proto3" json:"last_receive_timestamp,omitempty"`
	LastTransmitTimestamp  *proto.Timestamp `protobuf:"bytes,7,opt,name=last_transmit_timestamp,json=lastTransmitTimestamp,proto3" json:"last_transmit_timestamp,omitempty"`
	// Trafic counters
	TransmitBytes int64 `protobuf:"varint,8,opt,name=transmit_bytes,json=transmitBytes,proto3" json:"transmit_bytes,omitempty"`
	ReceiveBytes  int64 `protobuf:"varint,9,opt,name=receive_bytes,json=receiveBytes,proto3" json:"receive_bytes,omitempty"`
	// WireGuard endpoint address
	Endpoint string `protobuf:"bytes,10,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// WireGuard protocol version
	ProtocolVersion uint32       `protobuf:"varint,11,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	Ice             *epdisc.Peer `protobuf:"bytes,12,opt,name=ice,proto3" json:"ice,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_core_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_core_peer_proto_rawDescGZIP(), []int{0}
}

func (x *Peer) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Peer) GetPresharedKey() []byte {
	if x != nil {
		return x.PresharedKey
	}
	return nil
}

func (x *Peer) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

func (x *Peer) GetPersistentKeepaliveInterval() uint32 {
	if x != nil {
		return x.PersistentKeepaliveInterval
	}
	return 0
}

func (x *Peer) GetLastHandshakeTimestamp() *proto.Timestamp {
	if x != nil {
		return x.LastHandshakeTimestamp
	}
	return nil
}

func (x *Peer) GetLastReceiveTimestamp() *proto.Timestamp {
	if x != nil {
		return x.LastReceiveTimestamp
	}
	return nil
}

func (x *Peer) GetLastTransmitTimestamp() *proto.Timestamp {
	if x != nil {
		return x.LastTransmitTimestamp
	}
	return nil
}

func (x *Peer) GetTransmitBytes() int64 {
	if x != nil {
		return x.TransmitBytes
	}
	return 0
}

func (x *Peer) GetReceiveBytes() int64 {
	if x != nil {
		return x.ReceiveBytes
	}
	return 0
}

func (x *Peer) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Peer) GetProtocolVersion() uint32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *Peer) GetIce() *epdisc.Peer {
	if x != nil {
		return x.Ice
	}
	return nil
}

var File_core_peer_proto protoreflect.FileDescriptor

var file_core_peer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x66, 0x65, 0x61, 0x74,
	0x2f, 0x65, 0x70, 0x64, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04,
	0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72,
	0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x12, 0x42, 0x0a, 0x1d, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x1b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x4b, 0x65,
	0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x49, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x45, 0x0a, 0x16, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x47, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x03, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x77, 0x69, 0x63,
	0x65, 0x2e, 0x65, 0x70, 0x64, 0x69, 0x73, 0x63, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x03, 0x69,
	0x63, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65, 0x75, 0x2f, 0x77,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_peer_proto_rawDescOnce sync.Once
	file_core_peer_proto_rawDescData = file_core_peer_proto_rawDesc
)

func file_core_peer_proto_rawDescGZIP() []byte {
	file_core_peer_proto_rawDescOnce.Do(func() {
		file_core_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_peer_proto_rawDescData)
	})
	return file_core_peer_proto_rawDescData
}

var file_core_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_peer_proto_goTypes = []interface{}{
	(*Peer)(nil),            // 0: cunicu.core.Peer
	(*proto.Timestamp)(nil), // 1: cunicu.Timestamp
	(*epdisc.Peer)(nil),     // 2: cunicu.epdisc.Peer
}
var file_core_peer_proto_depIdxs = []int32{
	1, // 0: cunicu.core.Peer.last_handshake_timestamp:type_name -> cunicu.Timestamp
	1, // 1: cunicu.core.Peer.last_receive_timestamp:type_name -> cunicu.Timestamp
	1, // 2: cunicu.core.Peer.last_transmit_timestamp:type_name -> cunicu.Timestamp
	2, // 3: cunicu.core.Peer.ice:type_name -> cunicu.epdisc.Peer
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_core_peer_proto_init() }
func file_core_peer_proto_init() {
	if File_core_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
			RawDescriptor: file_core_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_peer_proto_goTypes,
		DependencyIndexes: file_core_peer_proto_depIdxs,
		MessageInfos:      file_core_peer_proto_msgTypes,
	}.Build()
	File_core_peer_proto = out.File
	file_core_peer_proto_rawDesc = nil
	file_core_peer_proto_goTypes = nil
	file_core_peer_proto_depIdxs = nil
}
