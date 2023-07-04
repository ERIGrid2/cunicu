// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: signaling/signaling.proto

package signaling

import (
	proto "github.com/stv0g/cunicu/pkg/proto"
	epdisc "github.com/stv0g/cunicu/pkg/proto/feature/epdisc"
	pdisc "github.com/stv0g/cunicu/pkg/proto/feature/pdisc"
	pske "github.com/stv0g/cunicu/pkg/proto/feature/pske"
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

type BackendType int32

const (
	BackendType_MULTI     BackendType = 0
	BackendType_GRPC      BackendType = 1
	BackendType_INPROCESS BackendType = 2
)

// Enum value maps for BackendType.
var (
	BackendType_name = map[int32]string{
		0: "MULTI",
		1: "GRPC",
		2: "INPROCESS",
	}
	BackendType_value = map[string]int32{
		"MULTI":     0,
		"GRPC":      1,
		"INPROCESS": 2,
	}
)

func (x BackendType) Enum() *BackendType {
	p := new(BackendType)
	*p = x
	return p
}

func (x BackendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackendType) Descriptor() protoreflect.EnumDescriptor {
	return file_signaling_signaling_proto_enumTypes[0].Descriptor()
}

func (BackendType) Type() protoreflect.EnumType {
	return &file_signaling_signaling_proto_enumTypes[0]
}

func (x BackendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackendType.Descriptor instead.
func (BackendType) EnumDescriptor() ([]byte, []int) {
	return file_signaling_signaling_proto_rawDescGZIP(), []int{0}
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    []byte            `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient []byte            `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Contents  *EncryptedMessage `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"` // of type SignalingMessage
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signaling_signaling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_signaling_signaling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_signaling_signaling_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Envelope) GetRecipient() []byte {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *Envelope) GetContents() *EncryptedMessage {
	if x != nil {
		return x.Contents
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials *epdisc.Credentials             `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Candidate   *epdisc.Candidate               `protobuf:"bytes,2,opt,name=candidate,proto3" json:"candidate,omitempty"`
	Peer        *pdisc.PeerDescription          `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`
	Pske        *pske.PresharedKeyEstablishment `protobuf:"bytes,4,opt,name=pske,proto3" json:"pske,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signaling_signaling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_signaling_signaling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_signaling_signaling_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetCredentials() *epdisc.Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *Message) GetCandidate() *epdisc.Candidate {
	if x != nil {
		return x.Candidate
	}
	return nil
}

func (x *Message) GetPeer() *pdisc.PeerDescription {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *Message) GetPske() *pske.PresharedKeyEstablishment {
	if x != nil {
		return x.Pske
	}
	return nil
}

// A container for an encrypted protobuf message
type EncryptedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body  []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *EncryptedMessage) Reset() {
	*x = EncryptedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signaling_signaling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedMessage) ProtoMessage() {}

func (x *EncryptedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_signaling_signaling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedMessage.ProtoReflect.Descriptor instead.
func (*EncryptedMessage) Descriptor() ([]byte, []int) {
	return file_signaling_signaling_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptedMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *EncryptedMessage) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

type SubscribeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SubscribeParams) Reset() {
	*x = SubscribeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signaling_signaling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeParams) ProtoMessage() {}

func (x *SubscribeParams) ProtoReflect() protoreflect.Message {
	mi := &file_signaling_signaling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeParams.ProtoReflect.Descriptor instead.
func (*SubscribeParams) Descriptor() ([]byte, []int) {
	return file_signaling_signaling_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeParams) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_signaling_signaling_proto protoreflect.FileDescriptor

var file_signaling_signaling_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x75, 0x6e,
	0x69, 0x63, 0x75, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x64, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x73, 0x6b, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x65, 0x70,
	0x64, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x65, 0x70, 0x64, 0x69, 0x73, 0x63, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xee, 0x01,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x65, 0x70, 0x64, 0x69, 0x73, 0x63, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x75, 0x6e,
	0x69, 0x63, 0x75, 0x2e, 0x65, 0x70, 0x64, 0x69, 0x73, 0x63, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x70, 0x64, 0x69, 0x73, 0x63, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x65,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x73, 0x6b, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x70, 0x73, 0x6b, 0x65, 0x2e, 0x50,
	0x72, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x74, 0x61, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x70, 0x73, 0x6b, 0x65, 0x22, 0x3c,
	0x0a, 0x10, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x0f,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x2a, 0x31, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x52, 0x50, 0x43, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x02, 0x32, 0xc7, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x11, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x1a, 0x2e, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x0d, 0x2e,
	0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x76,
	0x30, 0x67, 0x2f, 0x63, 0x75, 0x6e, 0x69, 0x63, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signaling_signaling_proto_rawDescOnce sync.Once
	file_signaling_signaling_proto_rawDescData = file_signaling_signaling_proto_rawDesc
)

func file_signaling_signaling_proto_rawDescGZIP() []byte {
	file_signaling_signaling_proto_rawDescOnce.Do(func() {
		file_signaling_signaling_proto_rawDescData = protoimpl.X.CompressGZIP(file_signaling_signaling_proto_rawDescData)
	})
	return file_signaling_signaling_proto_rawDescData
}

var file_signaling_signaling_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_signaling_signaling_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_signaling_signaling_proto_goTypes = []interface{}{
	(BackendType)(0),                       // 0: cunicu.signaling.BackendType
	(*Envelope)(nil),                       // 1: cunicu.signaling.Envelope
	(*Message)(nil),                        // 2: cunicu.signaling.Message
	(*EncryptedMessage)(nil),               // 3: cunicu.signaling.EncryptedMessage
	(*SubscribeParams)(nil),                // 4: cunicu.signaling.SubscribeParams
	(*epdisc.Credentials)(nil),             // 5: cunicu.epdisc.Credentials
	(*epdisc.Candidate)(nil),               // 6: cunicu.epdisc.Candidate
	(*pdisc.PeerDescription)(nil),          // 7: cunicu.pdisc.PeerDescription
	(*pske.PresharedKeyEstablishment)(nil), // 8: cunicu.pske.PresharedKeyEstablishment
	(*proto.Empty)(nil),                    // 9: cunicu.Empty
	(*proto.BuildInfo)(nil),                // 10: cunicu.BuildInfo
}
var file_signaling_signaling_proto_depIdxs = []int32{
	3,  // 0: cunicu.signaling.Envelope.contents:type_name -> cunicu.signaling.EncryptedMessage
	5,  // 1: cunicu.signaling.Message.credentials:type_name -> cunicu.epdisc.Credentials
	6,  // 2: cunicu.signaling.Message.candidate:type_name -> cunicu.epdisc.Candidate
	7,  // 3: cunicu.signaling.Message.peer:type_name -> cunicu.pdisc.PeerDescription
	8,  // 4: cunicu.signaling.Message.pske:type_name -> cunicu.pske.PresharedKeyEstablishment
	9,  // 5: cunicu.signaling.Signaling.GetBuildInfo:input_type -> cunicu.Empty
	4,  // 6: cunicu.signaling.Signaling.Subscribe:input_type -> cunicu.signaling.SubscribeParams
	1,  // 7: cunicu.signaling.Signaling.Publish:input_type -> cunicu.signaling.Envelope
	10, // 8: cunicu.signaling.Signaling.GetBuildInfo:output_type -> cunicu.BuildInfo
	1,  // 9: cunicu.signaling.Signaling.Subscribe:output_type -> cunicu.signaling.Envelope
	9,  // 10: cunicu.signaling.Signaling.Publish:output_type -> cunicu.Empty
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_signaling_signaling_proto_init() }
func file_signaling_signaling_proto_init() {
	if File_signaling_signaling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signaling_signaling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_signaling_signaling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_signaling_signaling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedMessage); i {
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
		file_signaling_signaling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeParams); i {
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
			RawDescriptor: file_signaling_signaling_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_signaling_signaling_proto_goTypes,
		DependencyIndexes: file_signaling_signaling_proto_depIdxs,
		EnumInfos:         file_signaling_signaling_proto_enumTypes,
		MessageInfos:      file_signaling_signaling_proto_msgTypes,
	}.Build()
	File_signaling_signaling_proto = out.File
	file_signaling_signaling_proto_rawDesc = nil
	file_signaling_signaling_proto_goTypes = nil
	file_signaling_signaling_proto_depIdxs = nil
}
