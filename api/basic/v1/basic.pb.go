// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.15.6
// source: api/basic/v1/basic.proto

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

// 用户个人信息
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	HID      string   `protobuf:"bytes,2,opt,name=HID,proto3" json:"HID,omitempty"`
	Phone    string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Nick     string   `protobuf:"bytes,4,opt,name=nick,proto3" json:"nick,omitempty"`
	Birth    string   `protobuf:"bytes,5,opt,name=birth,proto3" json:"birth,omitempty"`
	Icon     string   `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	Gender   int32    `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Sign     string   `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
	State    int32    `protobuf:"varint,9,opt,name=state,proto3" json:"state,omitempty"`
	Roles    []string `protobuf:"bytes,10,rep,name=roles,proto3" json:"roles,omitempty"`
	UpdateAt int64    `protobuf:"varint,11,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	CreateAt int64    `protobuf:"varint,12,opt,name=createAt,proto3" json:"createAt,omitempty"`
	DeleteAt int64    `protobuf:"varint,13,opt,name=deleteAt,proto3" json:"deleteAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_basic_v1_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_basic_v1_basic_proto_msgTypes[0]
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
	return file_api_basic_v1_basic_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetHID() string {
	if x != nil {
		return x.HID
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *User) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *User) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *User) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *User) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *User) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *User) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *User) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *User) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *User) GetDeleteAt() int64 {
	if x != nil {
		return x.DeleteAt
	}
	return 0
}

// 用户个人公开信息
type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Nick     string `protobuf:"bytes,2,opt,name=nick,proto3" json:"nick,omitempty"`
	Icon     string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	UpdateAt int64  `protobuf:"varint,4,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	DeleteAt int64  `protobuf:"varint,5,opt,name=deleteAt,proto3" json:"deleteAt,omitempty"`
	Birth    string `protobuf:"bytes,6,opt,name=birth,proto3" json:"birth,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_basic_v1_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_api_basic_v1_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_api_basic_v1_basic_proto_rawDescGZIP(), []int{1}
}

func (x *UserProfile) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserProfile) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *UserProfile) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UserProfile) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *UserProfile) GetDeleteAt() int64 {
	if x != nil {
		return x.DeleteAt
	}
	return 0
}

func (x *UserProfile) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

var File_api_basic_v1_basic_proto protoreflect.FileDescriptor

var file_api_basic_v1_basic_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x22, 0xa8, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x48, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x48, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x22,
	0x93, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x69, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x42, 0x30, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x17, 0x78, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x08, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x6c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_basic_v1_basic_proto_rawDescOnce sync.Once
	file_api_basic_v1_basic_proto_rawDescData = file_api_basic_v1_basic_proto_rawDesc
)

func file_api_basic_v1_basic_proto_rawDescGZIP() []byte {
	file_api_basic_v1_basic_proto_rawDescOnce.Do(func() {
		file_api_basic_v1_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_basic_v1_basic_proto_rawDescData)
	})
	return file_api_basic_v1_basic_proto_rawDescData
}

var file_api_basic_v1_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_basic_v1_basic_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: basic.v1.User
	(*UserProfile)(nil), // 1: basic.v1.UserProfile
}
var file_api_basic_v1_basic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_basic_v1_basic_proto_init() }
func file_api_basic_v1_basic_proto_init() {
	if File_api_basic_v1_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_basic_v1_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_basic_v1_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
			RawDescriptor: file_api_basic_v1_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_basic_v1_basic_proto_goTypes,
		DependencyIndexes: file_api_basic_v1_basic_proto_depIdxs,
		MessageInfos:      file_api_basic_v1_basic_proto_msgTypes,
	}.Build()
	File_api_basic_v1_basic_proto = out.File
	file_api_basic_v1_basic_proto_rawDesc = nil
	file_api_basic_v1_basic_proto_goTypes = nil
	file_api_basic_v1_basic_proto_depIdxs = nil
}
