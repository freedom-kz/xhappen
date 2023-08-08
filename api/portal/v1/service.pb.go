// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/portal/v1/service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "xhappen/api/basic/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SMSCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *SMSCodeRequest) Reset() {
	*x = SMSCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMSCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMSCodeRequest) ProtoMessage() {}

func (x *SMSCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMSCodeRequest.ProtoReflect.Descriptor instead.
func (*SMSCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SMSCodeRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SMSCodeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SMSCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SMSCodeReply) Reset() {
	*x = SMSCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMSCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMSCodeReply) ProtoMessage() {}

func (x *SMSCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMSCodeReply.ProtoReflect.Descriptor instead.
func (*SMSCodeReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{1}
}

type LoginByMobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	SmsCode  string `protobuf:"bytes,3,opt,name=smsCode,proto3" json:"smsCode,omitempty"`
}

func (x *LoginByMobileRequest) Reset() {
	*x = LoginByMobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByMobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByMobileRequest) ProtoMessage() {}

func (x *LoginByMobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByMobileRequest.ProtoReflect.Descriptor instead.
func (*LoginByMobileRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *LoginByMobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginByMobileRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LoginByMobileRequest) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type LoginByMobileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User  *v1.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *LoginByMobileReply) Reset() {
	*x = LoginByMobileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByMobileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByMobileReply) ProtoMessage() {}

func (x *LoginByMobileReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByMobileReply.ProtoReflect.Descriptor instead.
func (*LoginByMobileReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *LoginByMobileReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginByMobileReply) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{4}
}

type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{5}
}

type DeRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmsCode string `protobuf:"bytes,1,opt,name=smsCode,proto3" json:"smsCode,omitempty"`
}

func (x *DeRegisterRequest) Reset() {
	*x = DeRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeRegisterRequest) ProtoMessage() {}

func (x *DeRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeRegisterRequest.ProtoReflect.Descriptor instead.
func (*DeRegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeRegisterRequest) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type DeRegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeRegisterReply) Reset() {
	*x = DeRegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeRegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeRegisterReply) ProtoMessage() {}

func (x *DeRegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeRegisterReply.ProtoReflect.Descriptor instead.
func (*DeRegisterReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{7}
}

type GetUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetUserProfileRequest) Reset() {
	*x = GetUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileRequest) ProtoMessage() {}

func (x *GetUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileRequest.ProtoReflect.Descriptor instead.
func (*GetUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserProfileRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetUserProfileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users map[int64]*v1.UserProfile `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetUserProfileReply) Reset() {
	*x = GetUserProfileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileReply) ProtoMessage() {}

func (x *GetUserProfileReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileReply.ProtoReflect.Descriptor instead.
func (*GetUserProfileReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserProfileReply) GetUsers() map[int64]*v1.UserProfile {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetSelfProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSelfProfileRequest) Reset() {
	*x = GetSelfProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSelfProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSelfProfileRequest) ProtoMessage() {}

func (x *GetSelfProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSelfProfileRequest.ProtoReflect.Descriptor instead.
func (*GetSelfProfileRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{10}
}

type GetSelfProfileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *v1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetSelfProfileReply) Reset() {
	*x = GetSelfProfileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSelfProfileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSelfProfileReply) ProtoMessage() {}

func (x *GetSelfProfileReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSelfProfileReply.ProtoReflect.Descriptor instead.
func (*GetSelfProfileReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetSelfProfileReply) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_api_portal_v1_service_proto protoreflect.FileDescriptor

var file_api_portal_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x53, 0x4d,
	0x53, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x48, 0xfa, 0x42,
	0x45, 0x72, 0x43, 0x32, 0x3e, 0x5e, 0x31, 0x28, 0x33, 0x5c, 0x64, 0x7c, 0x34, 0x5b, 0x30, 0x2d,
	0x31, 0x34, 0x2d, 0x39, 0x5d, 0x7c, 0x35, 0x5b, 0x30, 0x2d, 0x33, 0x35, 0x2d, 0x39, 0x5d, 0x7c,
	0x36, 0x5b, 0x32, 0x35, 0x36, 0x37, 0x5d, 0x7c, 0x37, 0x5b, 0x30, 0x2d, 0x38, 0x5d, 0x7c, 0x38,
	0x5c, 0x64, 0x7c, 0x39, 0x5b, 0x30, 0x2d, 0x33, 0x35, 0x2d, 0x39, 0x5d, 0x29, 0x5c, 0x64, 0x7b,
	0x38, 0x7d, 0x24, 0x98, 0x01, 0x0b, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x4d,
	0x53, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x32, 0x3e, 0x5e, 0x31, 0x28, 0x33,
	0x5c, 0x64, 0x7c, 0x34, 0x5b, 0x30, 0x2d, 0x31, 0x34, 0x2d, 0x39, 0x5d, 0x7c, 0x35, 0x5b, 0x30,
	0x2d, 0x33, 0x35, 0x2d, 0x39, 0x5d, 0x7c, 0x36, 0x5b, 0x32, 0x35, 0x36, 0x37, 0x5d, 0x7c, 0x37,
	0x5b, 0x30, 0x2d, 0x38, 0x5d, 0x7c, 0x38, 0x5c, 0x64, 0x7c, 0x39, 0x5b, 0x30, 0x2d, 0x33, 0x35,
	0x2d, 0x39, 0x5d, 0x29, 0x5c, 0x64, 0x7b, 0x38, 0x7d, 0x24, 0x98, 0x01, 0x0b, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x18,
	0x18, 0x24, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x07,
	0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xfa,
	0x42, 0x0e, 0x72, 0x0c, 0x32, 0x07, 0x5e, 0x5c, 0x64, 0x7b, 0x36, 0x7d, 0x24, 0x98, 0x01, 0x06,
	0x52, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4e, 0x0a, 0x12, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x0a, 0x11, 0x44, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a,
	0x4f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6c, 0x66, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x32, 0xfb, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x5b, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4d, 0x53, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x4d, 0x53, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x73, 0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a,
	0x0a, 0x44, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x6c, 0x66, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x66,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x74, 0x73, 0x65, 0x6c, 0x66, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3a,
	0x01, 0x2a, 0x42, 0x1a, 0x5a, 0x18, 0x78, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_portal_v1_service_proto_rawDescOnce sync.Once
	file_api_portal_v1_service_proto_rawDescData = file_api_portal_v1_service_proto_rawDesc
)

func file_api_portal_v1_service_proto_rawDescGZIP() []byte {
	file_api_portal_v1_service_proto_rawDescOnce.Do(func() {
		file_api_portal_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_portal_v1_service_proto_rawDescData)
	})
	return file_api_portal_v1_service_proto_rawDescData
}

var file_api_portal_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_portal_v1_service_proto_goTypes = []interface{}{
	(*SMSCodeRequest)(nil),        // 0: portal.v1.SMSCodeRequest
	(*SMSCodeReply)(nil),          // 1: portal.v1.SMSCodeReply
	(*LoginByMobileRequest)(nil),  // 2: portal.v1.LoginByMobileRequest
	(*LoginByMobileReply)(nil),    // 3: portal.v1.LoginByMobileReply
	(*LogoutRequest)(nil),         // 4: portal.v1.LogoutRequest
	(*LogoutReply)(nil),           // 5: portal.v1.LogoutReply
	(*DeRegisterRequest)(nil),     // 6: portal.v1.DeRegisterRequest
	(*DeRegisterReply)(nil),       // 7: portal.v1.DeRegisterReply
	(*GetUserProfileRequest)(nil), // 8: portal.v1.GetUserProfileRequest
	(*GetUserProfileReply)(nil),   // 9: portal.v1.GetUserProfileReply
	(*GetSelfProfileRequest)(nil), // 10: portal.v1.GetSelfProfileRequest
	(*GetSelfProfileReply)(nil),   // 11: portal.v1.GetSelfProfileReply
	nil,                           // 12: portal.v1.GetUserProfileReply.UsersEntry
	(*v1.User)(nil),               // 13: basic.v1.User
	(*v1.UserProfile)(nil),        // 14: basic.v1.UserProfile
}
var file_api_portal_v1_service_proto_depIdxs = []int32{
	13, // 0: portal.v1.LoginByMobileReply.user:type_name -> basic.v1.User
	12, // 1: portal.v1.GetUserProfileReply.users:type_name -> portal.v1.GetUserProfileReply.UsersEntry
	13, // 2: portal.v1.GetSelfProfileReply.user:type_name -> basic.v1.User
	14, // 3: portal.v1.GetUserProfileReply.UsersEntry.value:type_name -> basic.v1.UserProfile
	0,  // 4: portal.v1.User.SendSMSCode:input_type -> portal.v1.SMSCodeRequest
	2,  // 5: portal.v1.User.LoginByMobile:input_type -> portal.v1.LoginByMobileRequest
	4,  // 6: portal.v1.User.Logout:input_type -> portal.v1.LogoutRequest
	6,  // 7: portal.v1.User.DeRegister:input_type -> portal.v1.DeRegisterRequest
	8,  // 8: portal.v1.User.GetUserProfile:input_type -> portal.v1.GetUserProfileRequest
	10, // 9: portal.v1.User.GetSelfProfile:input_type -> portal.v1.GetSelfProfileRequest
	1,  // 10: portal.v1.User.SendSMSCode:output_type -> portal.v1.SMSCodeReply
	3,  // 11: portal.v1.User.LoginByMobile:output_type -> portal.v1.LoginByMobileReply
	5,  // 12: portal.v1.User.Logout:output_type -> portal.v1.LogoutReply
	7,  // 13: portal.v1.User.DeRegister:output_type -> portal.v1.DeRegisterReply
	9,  // 14: portal.v1.User.GetUserProfile:output_type -> portal.v1.GetUserProfileReply
	11, // 15: portal.v1.User.GetSelfProfile:output_type -> portal.v1.GetSelfProfileReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_portal_v1_service_proto_init() }
func file_api_portal_v1_service_proto_init() {
	if File_api_portal_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_portal_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMSCodeRequest); i {
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
		file_api_portal_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMSCodeReply); i {
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
		file_api_portal_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByMobileRequest); i {
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
		file_api_portal_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginByMobileReply); i {
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
		file_api_portal_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_api_portal_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReply); i {
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
		file_api_portal_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeRegisterRequest); i {
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
		file_api_portal_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeRegisterReply); i {
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
		file_api_portal_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileRequest); i {
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
		file_api_portal_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileReply); i {
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
		file_api_portal_v1_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSelfProfileRequest); i {
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
		file_api_portal_v1_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSelfProfileReply); i {
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
			RawDescriptor: file_api_portal_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_portal_v1_service_proto_goTypes,
		DependencyIndexes: file_api_portal_v1_service_proto_depIdxs,
		MessageInfos:      file_api_portal_v1_service_proto_msgTypes,
	}.Build()
	File_api_portal_v1_service_proto = out.File
	file_api_portal_v1_service_proto_rawDesc = nil
	file_api_portal_v1_service_proto_goTypes = nil
	file_api_portal_v1_service_proto_depIdxs = nil
}
