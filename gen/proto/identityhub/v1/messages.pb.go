// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/identityhub/v1/messages.proto

package v1

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        string             `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Descriptor_ *MessageDescriptor `protobuf:"bytes,2,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identityhub_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identityhub_v1_messages_proto_msgTypes[0]
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
	return file_proto_identityhub_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Message) GetDescriptor_() *MessageDescriptor {
	if x != nil {
		return x.Descriptor_
	}
	return nil
}

type MessageDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method     string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Cid        string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	DataFormat string `protobuf:"bytes,3,opt,name=dataFormat,proto3" json:"dataFormat,omitempty"`
}

func (x *MessageDescriptor) Reset() {
	*x = MessageDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identityhub_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDescriptor) ProtoMessage() {}

func (x *MessageDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identityhub_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDescriptor.ProtoReflect.Descriptor instead.
func (*MessageDescriptor) Descriptor() ([]byte, []int) {
	return file_proto_identityhub_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *MessageDescriptor) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MessageDescriptor) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *MessageDescriptor) GetDataFormat() string {
	if x != nil {
		return x.DataFormat
	}
	return ""
}

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protected *AttestationProtected `protobuf:"bytes,1,opt,name=protected,proto3" json:"protected,omitempty"`
	Payload   string                `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature string                `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identityhub_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identityhub_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_proto_identityhub_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Attestation) GetProtected() *AttestationProtected {
	if x != nil {
		return x.Protected
	}
	return nil
}

func (x *Attestation) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Attestation) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type AttestationProtected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alg string `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	Kid string `protobuf:"bytes,2,opt,name=kid,proto3" json:"kid,omitempty"`
}

func (x *AttestationProtected) Reset() {
	*x = AttestationProtected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_identityhub_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationProtected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationProtected) ProtoMessage() {}

func (x *AttestationProtected) ProtoReflect() protoreflect.Message {
	mi := &file_proto_identityhub_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationProtected.ProtoReflect.Descriptor instead.
func (*AttestationProtected) Descriptor() ([]byte, []int) {
	return file_proto_identityhub_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *AttestationProtected) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *AttestationProtected) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

var File_proto_identityhub_v1_messages_proto protoreflect.FileDescriptor

var file_proto_identityhub_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x66, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0a, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x22, 0x5d, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x6c, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x65, 0x74, 0x7a, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_identityhub_v1_messages_proto_rawDescOnce sync.Once
	file_proto_identityhub_v1_messages_proto_rawDescData = file_proto_identityhub_v1_messages_proto_rawDesc
)

func file_proto_identityhub_v1_messages_proto_rawDescGZIP() []byte {
	file_proto_identityhub_v1_messages_proto_rawDescOnce.Do(func() {
		file_proto_identityhub_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_identityhub_v1_messages_proto_rawDescData)
	})
	return file_proto_identityhub_v1_messages_proto_rawDescData
}

var file_proto_identityhub_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_identityhub_v1_messages_proto_goTypes = []interface{}{
	(*Message)(nil),              // 0: proto.identityhub.v1.Message
	(*MessageDescriptor)(nil),    // 1: proto.identityhub.v1.MessageDescriptor
	(*Attestation)(nil),          // 2: proto.identityhub.v1.Attestation
	(*AttestationProtected)(nil), // 3: proto.identityhub.v1.AttestationProtected
}
var file_proto_identityhub_v1_messages_proto_depIdxs = []int32{
	1, // 0: proto.identityhub.v1.Message.descriptor:type_name -> proto.identityhub.v1.MessageDescriptor
	3, // 1: proto.identityhub.v1.Attestation.protected:type_name -> proto.identityhub.v1.AttestationProtected
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_identityhub_v1_messages_proto_init() }
func file_proto_identityhub_v1_messages_proto_init() {
	if File_proto_identityhub_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_identityhub_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_identityhub_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDescriptor); i {
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
		file_proto_identityhub_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
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
		file_proto_identityhub_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationProtected); i {
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
			RawDescriptor: file_proto_identityhub_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_identityhub_v1_messages_proto_goTypes,
		DependencyIndexes: file_proto_identityhub_v1_messages_proto_depIdxs,
		MessageInfos:      file_proto_identityhub_v1_messages_proto_msgTypes,
	}.Build()
	File_proto_identityhub_v1_messages_proto = out.File
	file_proto_identityhub_v1_messages_proto_rawDesc = nil
	file_proto_identityhub_v1_messages_proto_goTypes = nil
	file_proto_identityhub_v1_messages_proto_depIdxs = nil
}