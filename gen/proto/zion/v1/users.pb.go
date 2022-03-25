// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/zion/v1/users.proto

package v1

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User primary ID & key
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// User's decentralized identifier (DID), e.g. did:ion:[fingerprint] - Required; unique
	Did string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	// User's username, unique in Zion - Required; unique
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// User's display name - Optional
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// User's email address for opt-in marketing updates - Optional
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// User's personal biography that shows on their profile - Optional
	Bio string `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio,omitempty"`
	// URL for user's profile picture - Optional
	Img string `protobuf:"bytes,7,opt,name=img,proto3" json:"img,omitempty"`
	// How many sats it costs to direct-message this user - Optional - Default to 0
	PriceToMessage int64 `protobuf:"varint,8,opt,name=price_to_message,json=priceToMessage,proto3" json:"price_to_message,omitempty"`
	// Created when? - Required
	Created int64 `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	// Updated when? - Optional
	Updated int64 `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_zion_v1_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_zion_v1_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_zion_v1_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *User) GetPriceToMessage() int64 {
	if x != nil {
		return x.PriceToMessage
	}
	return 0
}

func (x *User) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *User) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_proto_zion_v1_users_proto protoreflect.FileDescriptor

var file_proto_zion_v1_users_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x7a, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x7a, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72,
	0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a,
	0x04, 0x28, 0x01, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x03, 0x64, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x30, 0x01,
	0x40, 0x01, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a,
	0x04, 0x30, 0x01, 0x40, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x7a,
	0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x7a, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_zion_v1_users_proto_rawDescOnce sync.Once
	file_proto_zion_v1_users_proto_rawDescData = file_proto_zion_v1_users_proto_rawDesc
)

func file_proto_zion_v1_users_proto_rawDescGZIP() []byte {
	file_proto_zion_v1_users_proto_rawDescOnce.Do(func() {
		file_proto_zion_v1_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_zion_v1_users_proto_rawDescData)
	})
	return file_proto_zion_v1_users_proto_rawDescData
}

var file_proto_zion_v1_users_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_zion_v1_users_proto_goTypes = []interface{}{
	(*User)(nil), // 0: proto.zion.v1.User
}
var file_proto_zion_v1_users_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_zion_v1_users_proto_init() }
func file_proto_zion_v1_users_proto_init() {
	if File_proto_zion_v1_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_zion_v1_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_proto_zion_v1_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_zion_v1_users_proto_goTypes,
		DependencyIndexes: file_proto_zion_v1_users_proto_depIdxs,
		MessageInfos:      file_proto_zion_v1_users_proto_msgTypes,
	}.Build()
	File_proto_zion_v1_users_proto = out.File
	file_proto_zion_v1_users_proto_rawDesc = nil
	file_proto_zion_v1_users_proto_goTypes = nil
	file_proto_zion_v1_users_proto_depIdxs = nil
}
