// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/gateway/v1/service.proto

package v1

import (
	errors "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid   string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`     //接收用户ID
	Clientid string   `protobuf:"bytes,2,opt,name=clientid,proto3" json:"clientid,omitempty"` //接收的客户端标志
	Sync     *v1.Sync `protobuf:"bytes,3,opt,name=sync,proto3" json:"sync,omitempty"`         //数据包
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *SyncRequest) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

func (x *SyncRequest) GetSync() *v1.Sync {
	if x != nil {
		return x.Sync
	}
	return nil
}

type SyncReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"` //结果
	Err *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`  //kratos通用错误
}

func (x *SyncReply) Reset() {
	*x = SyncReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncReply) ProtoMessage() {}

func (x *SyncReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncReply.ProtoReflect.Descriptor instead.
func (*SyncReply) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SyncReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *SyncReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OmitClientids []string    `protobuf:"bytes,1,rep,name=omitClientids,proto3" json:"omitClientids,omitempty"` //忽略设备ID
	OmitUserIds   []string    `protobuf:"bytes,2,rep,name=omitUserIds,proto3" json:"omitUserIds,omitempty"`     //忽略设备ID
	Deliver       *v1.Deliver `protobuf:"bytes,3,opt,name=deliver,proto3" json:"deliver,omitempty"`             //消息
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastRequest) GetOmitClientids() []string {
	if x != nil {
		return x.OmitClientids
	}
	return nil
}

func (x *BroadcastRequest) GetOmitUserIds() []string {
	if x != nil {
		return x.OmitUserIds
	}
	return nil
}

func (x *BroadcastRequest) GetDeliver() *v1.Deliver {
	if x != nil {
		return x.Deliver
	}
	return nil
}

type BroadcastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"` //结果
	Err *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`  //kratos通用错误
}

func (x *BroadcastReply) Reset() {
	*x = BroadcastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReply) ProtoMessage() {}

func (x *BroadcastReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReply.ProtoReflect.Descriptor instead.
func (*BroadcastReply) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *BroadcastReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

type DeliverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid        string      `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`               //用户ID
	Clientid      string      `protobuf:"bytes,2,opt,name=clientid,proto3" json:"clientid,omitempty"`           //设备ID
	OmitClientids []string    `protobuf:"bytes,3,rep,name=omitClientids,proto3" json:"omitClientids,omitempty"` //忽略设备ID
	Deliver       *v1.Deliver `protobuf:"bytes,4,opt,name=deliver,proto3" json:"deliver,omitempty"`             //消息
}

func (x *DeliverRequest) Reset() {
	*x = DeliverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverRequest) ProtoMessage() {}

func (x *DeliverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverRequest.ProtoReflect.Descriptor instead.
func (*DeliverRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeliverRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *DeliverRequest) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

func (x *DeliverRequest) GetOmitClientids() []string {
	if x != nil {
		return x.OmitClientids
	}
	return nil
}

func (x *DeliverRequest) GetDeliver() *v1.Deliver {
	if x != nil {
		return x.Deliver
	}
	return nil
}

type DeliverReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"` //结果
	Err *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`  //kratos通用错误
}

func (x *DeliverReply) Reset() {
	*x = DeliverReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverReply) ProtoMessage() {}

func (x *DeliverReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverReply.ProtoReflect.Descriptor instead.
func (*DeliverReply) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeliverReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *DeliverReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ClientId string     `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Action   *v1.Action `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[6]
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
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ActionRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ActionRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
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
		mi := &file_api_gateway_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionReply) ProtoMessage() {}

func (x *ActionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[7]
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
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{7}
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

type DisconnectForceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid     string        `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`                                      //用户ID
	Clientid   string        `protobuf:"bytes,2,opt,name=clientid,proto3" json:"clientid,omitempty"`                                  //设备ID
	DeviceType v1.DeviceType `protobuf:"varint,3,opt,name=deviceType,proto3,enum=protocol.v1.DeviceType" json:"deviceType,omitempty"` //设备类型
}

func (x *DisconnectForceRequest) Reset() {
	*x = DisconnectForceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectForceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectForceRequest) ProtoMessage() {}

func (x *DisconnectForceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectForceRequest.ProtoReflect.Descriptor instead.
func (*DisconnectForceRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *DisconnectForceRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *DisconnectForceRequest) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

func (x *DisconnectForceRequest) GetDeviceType() v1.DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return v1.DeviceType(0)
}

type DisconnectForceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret bool           `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"` //结果
	Err *errors.Status `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`  //kratos通用错误
}

func (x *DisconnectForceReply) Reset() {
	*x = DisconnectForceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectForceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectForceReply) ProtoMessage() {}

func (x *DisconnectForceReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectForceReply.ProtoReflect.Descriptor instead.
func (*DisconnectForceReply) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *DisconnectForceReply) GetRet() bool {
	if x != nil {
		return x.Ret
	}
	return false
}

func (x *DisconnectForceReply) GetErr() *errors.Status {
	if x != nil {
		return x.Err
	}
	return nil
}

var File_api_gateway_v1_service_proto protoreflect.FileDescriptor

var file_api_gateway_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a,
	0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x22, 0x3f, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x6f, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6d, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6d, 0x69, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x9a, 0x01, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x6d, 0x69, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x6a, 0x0a, 0x0d,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0xef, 0x02, 0x0a, 0x0a, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x53, 0x72, 0x76, 0x12, 0x38, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x17,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x12, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x11,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x12, 0x22, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x46, 0x6f, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x78, 0x68, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_gateway_v1_service_proto_rawDescOnce sync.Once
	file_api_gateway_v1_service_proto_rawDescData = file_api_gateway_v1_service_proto_rawDesc
)

func file_api_gateway_v1_service_proto_rawDescGZIP() []byte {
	file_api_gateway_v1_service_proto_rawDescOnce.Do(func() {
		file_api_gateway_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gateway_v1_service_proto_rawDescData)
	})
	return file_api_gateway_v1_service_proto_rawDescData
}

var file_api_gateway_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_gateway_v1_service_proto_goTypes = []interface{}{
	(*SyncRequest)(nil),            // 0: gateway.v1.SyncRequest
	(*SyncReply)(nil),              // 1: gateway.v1.SyncReply
	(*BroadcastRequest)(nil),       // 2: gateway.v1.BroadcastRequest
	(*BroadcastReply)(nil),         // 3: gateway.v1.BroadcastReply
	(*DeliverRequest)(nil),         // 4: gateway.v1.DeliverRequest
	(*DeliverReply)(nil),           // 5: gateway.v1.DeliverReply
	(*ActionRequest)(nil),          // 6: gateway.v1.ActionRequest
	(*ActionReply)(nil),            // 7: gateway.v1.ActionReply
	(*DisconnectForceRequest)(nil), // 8: gateway.v1.DisconnectForceRequest
	(*DisconnectForceReply)(nil),   // 9: gateway.v1.DisconnectForceReply
	(*v1.Sync)(nil),                // 10: protocol.v1.Sync
	(*errors.Status)(nil),          // 11: errors.Status
	(*v1.Deliver)(nil),             // 12: protocol.v1.Deliver
	(*v1.Action)(nil),              // 13: protocol.v1.Action
	(v1.DeviceType)(0),             // 14: protocol.v1.DeviceType
}
var file_api_gateway_v1_service_proto_depIdxs = []int32{
	10, // 0: gateway.v1.SyncRequest.sync:type_name -> protocol.v1.Sync
	11, // 1: gateway.v1.SyncReply.err:type_name -> errors.Status
	12, // 2: gateway.v1.BroadcastRequest.deliver:type_name -> protocol.v1.Deliver
	11, // 3: gateway.v1.BroadcastReply.err:type_name -> errors.Status
	12, // 4: gateway.v1.DeliverRequest.deliver:type_name -> protocol.v1.Deliver
	11, // 5: gateway.v1.DeliverReply.err:type_name -> errors.Status
	13, // 6: gateway.v1.ActionRequest.action:type_name -> protocol.v1.Action
	11, // 7: gateway.v1.ActionReply.err:type_name -> errors.Status
	14, // 8: gateway.v1.DisconnectForceRequest.deviceType:type_name -> protocol.v1.DeviceType
	11, // 9: gateway.v1.DisconnectForceReply.err:type_name -> errors.Status
	0,  // 10: gateway.v1.GatewaySrv.Sync:input_type -> gateway.v1.SyncRequest
	4,  // 11: gateway.v1.GatewaySrv.Deliver:input_type -> gateway.v1.DeliverRequest
	2,  // 12: gateway.v1.GatewaySrv.Broadcast:input_type -> gateway.v1.BroadcastRequest
	6,  // 13: gateway.v1.GatewaySrv.Action:input_type -> gateway.v1.ActionRequest
	8,  // 14: gateway.v1.GatewaySrv.DisconnectedForce:input_type -> gateway.v1.DisconnectForceRequest
	1,  // 15: gateway.v1.GatewaySrv.Sync:output_type -> gateway.v1.SyncReply
	5,  // 16: gateway.v1.GatewaySrv.Deliver:output_type -> gateway.v1.DeliverReply
	3,  // 17: gateway.v1.GatewaySrv.Broadcast:output_type -> gateway.v1.BroadcastReply
	7,  // 18: gateway.v1.GatewaySrv.Action:output_type -> gateway.v1.ActionReply
	9,  // 19: gateway.v1.GatewaySrv.DisconnectedForce:output_type -> gateway.v1.DisconnectForceReply
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_gateway_v1_service_proto_init() }
func file_api_gateway_v1_service_proto_init() {
	if File_api_gateway_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gateway_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_api_gateway_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncReply); i {
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
		file_api_gateway_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
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
		file_api_gateway_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReply); i {
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
		file_api_gateway_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverRequest); i {
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
		file_api_gateway_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverReply); i {
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
		file_api_gateway_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_gateway_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_gateway_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectForceRequest); i {
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
		file_api_gateway_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectForceReply); i {
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
			RawDescriptor: file_api_gateway_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gateway_v1_service_proto_goTypes,
		DependencyIndexes: file_api_gateway_v1_service_proto_depIdxs,
		MessageInfos:      file_api_gateway_v1_service_proto_msgTypes,
	}.Build()
	File_api_gateway_v1_service_proto = out.File
	file_api_gateway_v1_service_proto_rawDesc = nil
	file_api_gateway_v1_service_proto_goTypes = nil
	file_api_gateway_v1_service_proto_depIdxs = nil
}
