// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: offer.proto

package pb

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

type Offer_Role int32

const (
	Offer_CONTROLLED  Offer_Role = 0
	Offer_CONTROLLING Offer_Role = 1
)

// Enum value maps for Offer_Role.
var (
	Offer_Role_name = map[int32]string{
		0: "CONTROLLED",
		1: "CONTROLLING",
	}
	Offer_Role_value = map[string]int32{
		"CONTROLLED":  0,
		"CONTROLLING": 1,
	}
)

func (x Offer_Role) Enum() *Offer_Role {
	p := new(Offer_Role)
	*p = x
	return p
}

func (x Offer_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Offer_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_offer_proto_enumTypes[0].Descriptor()
}

func (Offer_Role) Type() protoreflect.EnumType {
	return &file_offer_proto_enumTypes[0]
}

func (x Offer_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Offer_Role.Descriptor instead.
func (Offer_Role) EnumDescriptor() ([]byte, []int) {
	return file_offer_proto_rawDescGZIP(), []int{1, 0}
}

type Offer_Implementation int32

const (
	Offer_FULL Offer_Implementation = 0
	Offer_LITE Offer_Implementation = 1
)

// Enum value maps for Offer_Implementation.
var (
	Offer_Implementation_name = map[int32]string{
		0: "FULL",
		1: "LITE",
	}
	Offer_Implementation_value = map[string]int32{
		"FULL": 0,
		"LITE": 1,
	}
)

func (x Offer_Implementation) Enum() *Offer_Implementation {
	p := new(Offer_Implementation)
	*p = x
	return p
}

func (x Offer_Implementation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Offer_Implementation) Descriptor() protoreflect.EnumDescriptor {
	return file_offer_proto_enumTypes[1].Descriptor()
}

func (Offer_Implementation) Type() protoreflect.EnumType {
	return &file_offer_proto_enumTypes[1]
}

func (x Offer_Implementation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Offer_Implementation.Descriptor instead.
func (Offer_Implementation) EnumDescriptor() ([]byte, []int) {
	return file_offer_proto_rawDescGZIP(), []int{1, 1}
}

type SignedOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offer    []byte `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
	Sigature []byte `protobuf:"bytes,2,opt,name=sigature,proto3" json:"sigature,omitempty"`
}

func (x *SignedOffer) Reset() {
	*x = SignedOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedOffer) ProtoMessage() {}

func (x *SignedOffer) ProtoReflect() protoreflect.Message {
	mi := &file_offer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedOffer.ProtoReflect.Descriptor instead.
func (*SignedOffer) Descriptor() ([]byte, []int) {
	return file_offer_proto_rawDescGZIP(), []int{0}
}

func (x *SignedOffer) GetOffer() []byte {
	if x != nil {
		return x.Offer
	}
	return nil
}

func (x *SignedOffer) GetSigature() []byte {
	if x != nil {
		return x.Sigature
	}
	return nil
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        int64                `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Epoch          int64                `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Role           Offer_Role           `protobuf:"varint,3,opt,name=role,proto3,enum=wice.Offer_Role" json:"role,omitempty"`
	Implementation Offer_Implementation `protobuf:"varint,4,opt,name=implementation,proto3,enum=wice.Offer_Implementation" json:"implementation,omitempty"`
	Candidates     []*Candidate         `protobuf:"bytes,5,rep,name=candidates,proto3" json:"candidates,omitempty"`
	Ufrag          string               `protobuf:"bytes,6,opt,name=Ufrag,proto3" json:"Ufrag,omitempty"`
	Pwd            string               `protobuf:"bytes,7,opt,name=Pwd,proto3" json:"Pwd,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_offer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_offer_proto_rawDescGZIP(), []int{1}
}

func (x *Offer) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Offer) GetEpoch() int64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *Offer) GetRole() Offer_Role {
	if x != nil {
		return x.Role
	}
	return Offer_CONTROLLED
}

func (x *Offer) GetImplementation() Offer_Implementation {
	if x != nil {
		return x.Implementation
	}
	return Offer_FULL
}

func (x *Offer) GetCandidates() []*Candidate {
	if x != nil {
		return x.Candidates
	}
	return nil
}

func (x *Offer) GetUfrag() string {
	if x != nil {
		return x.Ufrag
	}
	return ""
}

func (x *Offer) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_offer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_offer_proto_rawDescGZIP(), []int{2}
}

var File_offer_proto protoreflect.FileDescriptor

var file_offer_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77,
	0x69, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x67, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0xc9, 0x02, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x24,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x66, 0x72,
	0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x55, 0x66, 0x72, 0x61, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x50, 0x77, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x77,
	0x64, 0x22, 0x27, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x22, 0x24, 0x0a, 0x0e, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x54, 0x45, 0x10, 0x01,
	0x22, 0x0b, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x16, 0x5a,
	0x14, 0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65, 0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_offer_proto_rawDescOnce sync.Once
	file_offer_proto_rawDescData = file_offer_proto_rawDesc
)

func file_offer_proto_rawDescGZIP() []byte {
	file_offer_proto_rawDescOnce.Do(func() {
		file_offer_proto_rawDescData = protoimpl.X.CompressGZIP(file_offer_proto_rawDescData)
	})
	return file_offer_proto_rawDescData
}

var file_offer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_offer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_offer_proto_goTypes = []interface{}{
	(Offer_Role)(0),           // 0: wice.Offer.Role
	(Offer_Implementation)(0), // 1: wice.Offer.Implementation
	(*SignedOffer)(nil),       // 2: wice.SignedOffer
	(*Offer)(nil),             // 3: wice.Offer
	(*Candidate)(nil),         // 4: wice.Candidate
}
var file_offer_proto_depIdxs = []int32{
	0, // 0: wice.Offer.role:type_name -> wice.Offer.Role
	1, // 1: wice.Offer.implementation:type_name -> wice.Offer.Implementation
	4, // 2: wice.Offer.candidates:type_name -> wice.Candidate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_offer_proto_init() }
func file_offer_proto_init() {
	if File_offer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_offer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedOffer); i {
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
		file_offer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_offer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
			RawDescriptor: file_offer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_offer_proto_goTypes,
		DependencyIndexes: file_offer_proto_depIdxs,
		EnumInfos:         file_offer_proto_enumTypes,
		MessageInfos:      file_offer_proto_msgTypes,
	}.Build()
	File_offer_proto = out.File
	file_offer_proto_rawDesc = nil
	file_offer_proto_goTypes = nil
	file_offer_proto_depIdxs = nil
}
