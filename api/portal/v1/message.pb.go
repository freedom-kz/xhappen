// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.15.6
// source: api/portal/v1/message.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{0}
}

type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{1}
}

type EditMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditMessageRequest) Reset() {
	*x = EditMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMessageRequest) ProtoMessage() {}

func (x *EditMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMessageRequest.ProtoReflect.Descriptor instead.
func (*EditMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{2}
}

type EditMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditMessageReply) Reset() {
	*x = EditMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMessageReply) ProtoMessage() {}

func (x *EditMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMessageReply.ProtoReflect.Descriptor instead.
func (*EditMessageReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{3}
}

type RevokeMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevokeMessageRequest) Reset() {
	*x = RevokeMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeMessageRequest) ProtoMessage() {}

func (x *RevokeMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeMessageRequest.ProtoReflect.Descriptor instead.
func (*RevokeMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{4}
}

type RevokeMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevokeMessageReply) Reset() {
	*x = RevokeMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeMessageReply) ProtoMessage() {}

func (x *RevokeMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeMessageReply.ProtoReflect.Descriptor instead.
func (*RevokeMessageReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{5}
}

type ListMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMessageRequest) Reset() {
	*x = ListMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessageRequest) ProtoMessage() {}

func (x *ListMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessageRequest.ProtoReflect.Descriptor instead.
func (*ListMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{6}
}

type ListMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMessageReply) Reset() {
	*x = ListMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessageReply) ProtoMessage() {}

func (x *ListMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessageReply.ProtoReflect.Descriptor instead.
func (*ListMessageReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{7}
}

type MarkMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkMessageRequest) Reset() {
	*x = MarkMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMessageRequest) ProtoMessage() {}

func (x *MarkMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMessageRequest.ProtoReflect.Descriptor instead.
func (*MarkMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{8}
}

type MarkMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkMessageReply) Reset() {
	*x = MarkMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_portal_v1_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMessageReply) ProtoMessage() {}

func (x *MarkMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_portal_v1_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMessageReply.ProtoReflect.Descriptor instead.
func (*MarkMessageReply) Descriptor() ([]byte, []int) {
	return file_api_portal_v1_message_proto_rawDescGZIP(), []int{9}
}

var File_api_portal_v1_message_proto protoreflect.FileDescriptor

var file_api_portal_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x14, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x64, 0x69,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x4d,
	0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa1, 0x04, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x68, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x68, 0x0a, 0x0b, 0x45,
	0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01,
	0x2a, 0x22, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x6e, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01,
	0x2a, 0x22, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x68, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12,
	0x68, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x34, 0x0a, 0x0a, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x78, 0x68, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0xa2, 0x02, 0x09, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_portal_v1_message_proto_rawDescOnce sync.Once
	file_api_portal_v1_message_proto_rawDescData = file_api_portal_v1_message_proto_rawDesc
)

func file_api_portal_v1_message_proto_rawDescGZIP() []byte {
	file_api_portal_v1_message_proto_rawDescOnce.Do(func() {
		file_api_portal_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_portal_v1_message_proto_rawDescData)
	})
	return file_api_portal_v1_message_proto_rawDescData
}

var file_api_portal_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_portal_v1_message_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),   // 0: portal.v1.SendMessageRequest
	(*SendMessageReply)(nil),     // 1: portal.v1.SendMessageReply
	(*EditMessageRequest)(nil),   // 2: portal.v1.EditMessageRequest
	(*EditMessageReply)(nil),     // 3: portal.v1.EditMessageReply
	(*RevokeMessageRequest)(nil), // 4: portal.v1.RevokeMessageRequest
	(*RevokeMessageReply)(nil),   // 5: portal.v1.RevokeMessageReply
	(*ListMessageRequest)(nil),   // 6: portal.v1.ListMessageRequest
	(*ListMessageReply)(nil),     // 7: portal.v1.ListMessageReply
	(*MarkMessageRequest)(nil),   // 8: portal.v1.MarkMessageRequest
	(*MarkMessageReply)(nil),     // 9: portal.v1.MarkMessageReply
}
var file_api_portal_v1_message_proto_depIdxs = []int32{
	0, // 0: portal.v1.Message.SendMessage:input_type -> portal.v1.SendMessageRequest
	2, // 1: portal.v1.Message.EditMessage:input_type -> portal.v1.EditMessageRequest
	4, // 2: portal.v1.Message.RevokeMessage:input_type -> portal.v1.RevokeMessageRequest
	6, // 3: portal.v1.Message.ListMessage:input_type -> portal.v1.ListMessageRequest
	8, // 4: portal.v1.Message.MarkMessage:input_type -> portal.v1.MarkMessageRequest
	1, // 5: portal.v1.Message.SendMessage:output_type -> portal.v1.SendMessageReply
	3, // 6: portal.v1.Message.EditMessage:output_type -> portal.v1.EditMessageReply
	5, // 7: portal.v1.Message.RevokeMessage:output_type -> portal.v1.RevokeMessageReply
	7, // 8: portal.v1.Message.ListMessage:output_type -> portal.v1.ListMessageReply
	9, // 9: portal.v1.Message.MarkMessage:output_type -> portal.v1.MarkMessageReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_portal_v1_message_proto_init() }
func file_api_portal_v1_message_proto_init() {
	if File_api_portal_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_portal_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_api_portal_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReply); i {
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
		file_api_portal_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditMessageRequest); i {
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
		file_api_portal_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditMessageReply); i {
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
		file_api_portal_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeMessageRequest); i {
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
		file_api_portal_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeMessageReply); i {
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
		file_api_portal_v1_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessageRequest); i {
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
		file_api_portal_v1_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessageReply); i {
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
		file_api_portal_v1_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkMessageRequest); i {
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
		file_api_portal_v1_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkMessageReply); i {
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
			RawDescriptor: file_api_portal_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_portal_v1_message_proto_goTypes,
		DependencyIndexes: file_api_portal_v1_message_proto_depIdxs,
		MessageInfos:      file_api_portal_v1_message_proto_msgTypes,
	}.Build()
	File_api_portal_v1_message_proto = out.File
	file_api_portal_v1_message_proto_rawDesc = nil
	file_api_portal_v1_message_proto_goTypes = nil
	file_api_portal_v1_message_proto_depIdxs = nil
}
