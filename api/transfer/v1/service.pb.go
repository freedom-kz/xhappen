// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.15.6
// source: api/transfer/v1/service.proto

package v1

import (
	errors "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	v1 "xhappen/api/protocol/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BindInfo *v1.Bind `protobuf:"bytes,1,opt,name=bindInfo,proto3" json:"bindInfo,omitempty"`
	ServerID string   `protobuf:"bytes,2,opt,name=serverID,proto3" json:"serverID,omitempty"`
}

func (x *BindRequest) Reset() {
	*x = BindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRequest) ProtoMessage() {}

func (x *BindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRequest.ProtoReflect.Descriptor instead.
func (*BindRequest) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *BindRequest) GetBindInfo() *v1.Bind {
	if x != nil {
		return x.BindInfo
	}
	return nil
}

func (x *BindRequest) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

type BindReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Err *errors.Status `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"` //kratos通用错误
}

func (x *BindReply) Reset() {
	*x = BindReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindReply) ProtoMessage() {}

func (x *BindReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindReply.ProtoReflect.Descriptor instead.
func (*BindReply) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *BindReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *BindReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID       string        `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	ServerID       string        `protobuf:"bytes,2,opt,name=serverID,proto3" json:"serverID,omitempty"`
	ConnectSequece uint64        `protobuf:"varint,3,opt,name=connectSequece,proto3" json:"connectSequece,omitempty"`
	LoginType      v1.LoginType  `protobuf:"varint,4,opt,name=loginType,proto3,enum=protocol.v1.LoginType" json:"loginType,omitempty"`
	DeviceType     v1.DeviceType `protobuf:"varint,5,opt,name=deviceType,proto3,enum=protocol.v1.DeviceType" json:"deviceType,omitempty"`
	CurVersion     int32         `protobuf:"varint,6,opt,name=curVersion,proto3" json:"curVersion,omitempty"`
	AuthInfo       *v1.Auth      `protobuf:"bytes,7,opt,name=authInfo,proto3" json:"authInfo,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRequest) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *AuthRequest) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *AuthRequest) GetConnectSequece() uint64 {
	if x != nil {
		return x.ConnectSequece
	}
	return 0
}

func (x *AuthRequest) GetLoginType() v1.LoginType {
	if x != nil {
		return x.LoginType
	}
	return v1.LoginType(0)
}

func (x *AuthRequest) GetDeviceType() v1.DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return v1.DeviceType(0)
}

func (x *AuthRequest) GetCurVersion() int32 {
	if x != nil {
		return x.CurVersion
	}
	return 0
}

func (x *AuthRequest) GetAuthInfo() *v1.Auth {
	if x != nil {
		return x.AuthInfo
	}
	return nil
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret         bool                   `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	UID         string                 `protobuf:"bytes,2,opt,name=UID,proto3" json:"UID,omitempty"`
	TokenExpire *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=token_expire,json=tokenExpire,proto3" json:"token_expire,omitempty"`
	Sessions    []uint64               `protobuf:"varint,4,rep,packed,name=sessions,proto3" json:"sessions,omitempty"`
	UType       v1.UserType            `protobuf:"varint,5,opt,name=uType,proto3,enum=protocol.v1.UserType" json:"uType,omitempty"`
	Err         *errors.Status         `protobuf:"bytes,6,opt,name=err,proto3" json:"err,omitempty"` //kratos通用错误
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *AuthReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *AuthReply) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *AuthReply) GetTokenExpire() *timestamppb.Timestamp {
	if x != nil {
		return x.TokenExpire
	}
	return nil
}

func (x *AuthReply) GetSessions() []uint64 {
	if x != nil {
		return x.Sessions
	}
	return nil
}

func (x *AuthReply) GetUType() v1.UserType {
	if x != nil {
		return x.UType
	}
	return v1.UserType(0)
}

func (x *AuthReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string     `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`     //用户ID
	DeviceID string     `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"` //设备ID
	Submit   *v1.Submit `protobuf:"bytes,3,opt,name=submit,proto3" json:"submit,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SubmitRequest) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *SubmitRequest) GetSubmit() *v1.Submit {
	if x != nil {
		return x.Submit
	}
	return nil
}

type SubmitReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret       bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Err       *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"` //kratos通用错误
	SessionID uint64         `protobuf:"varint,3,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Sequence  uint64         `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Timestamp uint64         `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SubmitReply) Reset() {
	*x = SubmitReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitReply) ProtoMessage() {}

func (x *SubmitReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitReply.ProtoReflect.Descriptor instead.
func (*SubmitReply) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *SubmitReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *SubmitReply) GetSessionID() uint64 {
	if x != nil {
		return x.SessionID
	}
	return 0
}

func (x *SubmitReply) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *SubmitReply) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID      string     `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	DeviceId string     `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Action   *v1.Action `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ActionRequest) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *ActionRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ActionRequest) GetAction() *v1.Action {
	if x != nil {
		return x.Action
	}
	return nil
}

type ActionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret       bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Err       *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"` //kratos通用错误
	Timestamp uint64         `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Payload   []byte         `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ActionReply) Reset() {
	*x = ActionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionReply) ProtoMessage() {}

func (x *ActionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionReply.ProtoReflect.Descriptor instead.
func (*ActionReply) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *ActionReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *ActionReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *ActionReply) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ActionReply) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type QuitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeviceID   string        `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	DeviceType v1.DeviceType `protobuf:"varint,3,opt,name=deviceType,proto3,enum=protocol.v1.DeviceType" json:"deviceType,omitempty"`
}

func (x *QuitRequest) Reset() {
	*x = QuitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitRequest) ProtoMessage() {}

func (x *QuitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitRequest.ProtoReflect.Descriptor instead.
func (*QuitRequest) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *QuitRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *QuitRequest) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *QuitRequest) GetDeviceType() v1.DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return v1.DeviceType(0)
}

type QuitReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Err *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"` //kratos通用错误
}

func (x *QuitReply) Reset() {
	*x = QuitReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_transfer_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitReply) ProtoMessage() {}

func (x *QuitReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_transfer_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitReply.ProtoReflect.Descriptor instead.
func (*QuitReply) Descriptor() ([]byte, []int) {
	return file_api_transfer_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *QuitReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *QuitReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

var File_api_transfer_v1_service_proto protoreflect.FileDescriptor

var file_api_transfer_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x58, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x69, 0x6e, 0x64, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x09, 0x42, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0xab, 0x02, 0x0a, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x72,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x75, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x75, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x22, 0x70, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x06,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x6a, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x79,
	0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12,
	0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x7a, 0x0a, 0x0b, 0x51, 0x75, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3f, 0x0a, 0x09, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0xbe, 0x02, 0x0a, 0x04, 0x50, 0x61, 0x73, 0x73, 0x12,
	0x3a, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x51,
	0x75, 0x69, 0x74, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x69, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x38, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x78, 0x68, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_transfer_v1_service_proto_rawDescOnce sync.Once
	file_api_transfer_v1_service_proto_rawDescData = file_api_transfer_v1_service_proto_rawDesc
)

func file_api_transfer_v1_service_proto_rawDescGZIP() []byte {
	file_api_transfer_v1_service_proto_rawDescOnce.Do(func() {
		file_api_transfer_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_transfer_v1_service_proto_rawDescData)
	})
	return file_api_transfer_v1_service_proto_rawDescData
}

var file_api_transfer_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_transfer_v1_service_proto_goTypes = []interface{}{
	(*BindRequest)(nil),           // 0: transfer.v1.BindRequest
	(*BindReply)(nil),             // 1: transfer.v1.BindReply
	(*AuthRequest)(nil),           // 2: transfer.v1.AuthRequest
	(*AuthReply)(nil),             // 3: transfer.v1.AuthReply
	(*SubmitRequest)(nil),         // 4: transfer.v1.SubmitRequest
	(*SubmitReply)(nil),           // 5: transfer.v1.SubmitReply
	(*ActionRequest)(nil),         // 6: transfer.v1.ActionRequest
	(*ActionReply)(nil),           // 7: transfer.v1.ActionReply
	(*QuitRequest)(nil),           // 8: transfer.v1.QuitRequest
	(*QuitReply)(nil),             // 9: transfer.v1.QuitReply
	(*v1.Bind)(nil),               // 10: protocol.v1.Bind
	(*errors.Status)(nil),         // 11: errors.Status
	(v1.LoginType)(0),             // 12: protocol.v1.LoginType
	(v1.DeviceType)(0),            // 13: protocol.v1.DeviceType
	(*v1.Auth)(nil),               // 14: protocol.v1.Auth
	(*timestamppb.Timestamp)(nil), // 15: google.protobuf.Timestamp
	(v1.UserType)(0),              // 16: protocol.v1.UserType
	(*v1.Submit)(nil),             // 17: protocol.v1.Submit
	(*v1.Action)(nil),             // 18: protocol.v1.Action
}
var file_api_transfer_v1_service_proto_depIdxs = []int32{
	10, // 0: transfer.v1.BindRequest.bindInfo:type_name -> protocol.v1.Bind
	11, // 1: transfer.v1.BindReply.err:type_name -> errors.Status
	12, // 2: transfer.v1.AuthRequest.loginType:type_name -> protocol.v1.LoginType
	13, // 3: transfer.v1.AuthRequest.deviceType:type_name -> protocol.v1.DeviceType
	14, // 4: transfer.v1.AuthRequest.authInfo:type_name -> protocol.v1.Auth
	15, // 5: transfer.v1.AuthReply.token_expire:type_name -> google.protobuf.Timestamp
	16, // 6: transfer.v1.AuthReply.uType:type_name -> protocol.v1.UserType
	11, // 7: transfer.v1.AuthReply.err:type_name -> errors.Status
	17, // 8: transfer.v1.SubmitRequest.submit:type_name -> protocol.v1.Submit
	11, // 9: transfer.v1.SubmitReply.err:type_name -> errors.Status
	18, // 10: transfer.v1.ActionRequest.action:type_name -> protocol.v1.Action
	11, // 11: transfer.v1.ActionReply.err:type_name -> errors.Status
	13, // 12: transfer.v1.QuitRequest.deviceType:type_name -> protocol.v1.DeviceType
	11, // 13: transfer.v1.QuitReply.err:type_name -> errors.Status
	0,  // 14: transfer.v1.Pass.Bind:input_type -> transfer.v1.BindRequest
	2,  // 15: transfer.v1.Pass.Auth:input_type -> transfer.v1.AuthRequest
	4,  // 16: transfer.v1.Pass.Submit:input_type -> transfer.v1.SubmitRequest
	6,  // 17: transfer.v1.Pass.Action:input_type -> transfer.v1.ActionRequest
	8,  // 18: transfer.v1.Pass.Quit:input_type -> transfer.v1.QuitRequest
	1,  // 19: transfer.v1.Pass.Bind:output_type -> transfer.v1.BindReply
	3,  // 20: transfer.v1.Pass.Auth:output_type -> transfer.v1.AuthReply
	5,  // 21: transfer.v1.Pass.Submit:output_type -> transfer.v1.SubmitReply
	7,  // 22: transfer.v1.Pass.Action:output_type -> transfer.v1.ActionReply
	9,  // 23: transfer.v1.Pass.Quit:output_type -> transfer.v1.QuitReply
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_api_transfer_v1_service_proto_init() }
func file_api_transfer_v1_service_proto_init() {
	if File_api_transfer_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_transfer_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindRequest); i {
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
		file_api_transfer_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindReply); i {
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
		file_api_transfer_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_api_transfer_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_api_transfer_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_api_transfer_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitReply); i {
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
		file_api_transfer_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_api_transfer_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionReply); i {
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
		file_api_transfer_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitRequest); i {
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
		file_api_transfer_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitReply); i {
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
			RawDescriptor: file_api_transfer_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_transfer_v1_service_proto_goTypes,
		DependencyIndexes: file_api_transfer_v1_service_proto_depIdxs,
		MessageInfos:      file_api_transfer_v1_service_proto_msgTypes,
	}.Build()
	File_api_transfer_v1_service_proto = out.File
	file_api_transfer_v1_service_proto_rawDesc = nil
	file_api_transfer_v1_service_proto_goTypes = nil
	file_api_transfer_v1_service_proto_depIdxs = nil
}
