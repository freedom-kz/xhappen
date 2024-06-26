// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.6
// source: api/sequence/v1/service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sequence_GenSequenceByUserIds_FullMethodName            = "/sequence.v1.Sequence/GenSequenceByUserIds"
	Sequence_GetCurrentSequenceByUserIds_FullMethodName     = "/sequence.v1.Sequence/GetCurrentSequenceByUserIds"
	Sequence_GenRoomSequenceByRoomIds_FullMethodName        = "/sequence.v1.Sequence/GenRoomSequenceByRoomIds"
	Sequence_GetCurrentRoomSequenceByRoomIds_FullMethodName = "/sequence.v1.Sequence/GetCurrentRoomSequenceByRoomIds"
)

// SequenceClient is the client API for Sequence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SequenceClient interface {
	// 批量获取用户消息序列号（生产消费序列号）
	GenSequenceByUserIds(ctx context.Context, in *GenSequenceByUserIdsRequest, opts ...grpc.CallOption) (*GenSequenceByUserIdsReply, error)
	// 批量获取用户当前序列号（获取当前序列号）
	GetCurrentSequenceByUserIds(ctx context.Context, in *GetCurrentSequenceByUserIdsRequest, opts ...grpc.CallOption) (*GetCurrentSequenceByUserIdsReply, error)
	// 批量获取房间消息序列号（生产房间消息序列号）
	GenRoomSequenceByRoomIds(ctx context.Context, in *GenRoomSequenceByRoomIdsRequest, opts ...grpc.CallOption) (*GenRoomSequenceByRoomIdsReply, error)
	// 批量获取房间当前消息序列号（获取房间当前消息序列号）
	GetCurrentRoomSequenceByRoomIds(ctx context.Context, in *GetCurrentSequenceByRoomIdsRequest, opts ...grpc.CallOption) (*GetCurrentSequenceByRoomIdsReply, error)
}

type sequenceClient struct {
	cc grpc.ClientConnInterface
}

func NewSequenceClient(cc grpc.ClientConnInterface) SequenceClient {
	return &sequenceClient{cc}
}

func (c *sequenceClient) GenSequenceByUserIds(ctx context.Context, in *GenSequenceByUserIdsRequest, opts ...grpc.CallOption) (*GenSequenceByUserIdsReply, error) {
	out := new(GenSequenceByUserIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GenSequenceByUserIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceClient) GetCurrentSequenceByUserIds(ctx context.Context, in *GetCurrentSequenceByUserIdsRequest, opts ...grpc.CallOption) (*GetCurrentSequenceByUserIdsReply, error) {
	out := new(GetCurrentSequenceByUserIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GetCurrentSequenceByUserIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceClient) GenRoomSequenceByRoomIds(ctx context.Context, in *GenRoomSequenceByRoomIdsRequest, opts ...grpc.CallOption) (*GenRoomSequenceByRoomIdsReply, error) {
	out := new(GenRoomSequenceByRoomIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GenRoomSequenceByRoomIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceClient) GetCurrentRoomSequenceByRoomIds(ctx context.Context, in *GetCurrentSequenceByRoomIdsRequest, opts ...grpc.CallOption) (*GetCurrentSequenceByRoomIdsReply, error) {
	out := new(GetCurrentSequenceByRoomIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GetCurrentRoomSequenceByRoomIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequenceServer is the server API for Sequence service.
// All implementations must embed UnimplementedSequenceServer
// for forward compatibility
type SequenceServer interface {
	// 批量获取用户消息序列号（生产消费序列号）
	GenSequenceByUserIds(context.Context, *GenSequenceByUserIdsRequest) (*GenSequenceByUserIdsReply, error)
	// 批量获取用户当前序列号（获取当前序列号）
	GetCurrentSequenceByUserIds(context.Context, *GetCurrentSequenceByUserIdsRequest) (*GetCurrentSequenceByUserIdsReply, error)
	// 批量获取房间消息序列号（生产房间消息序列号）
	GenRoomSequenceByRoomIds(context.Context, *GenRoomSequenceByRoomIdsRequest) (*GenRoomSequenceByRoomIdsReply, error)
	// 批量获取房间当前消息序列号（获取房间当前消息序列号）
	GetCurrentRoomSequenceByRoomIds(context.Context, *GetCurrentSequenceByRoomIdsRequest) (*GetCurrentSequenceByRoomIdsReply, error)
	mustEmbedUnimplementedSequenceServer()
}

// UnimplementedSequenceServer must be embedded to have forward compatible implementations.
type UnimplementedSequenceServer struct {
}

func (UnimplementedSequenceServer) GenSequenceByUserIds(context.Context, *GenSequenceByUserIdsRequest) (*GenSequenceByUserIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenSequenceByUserIds not implemented")
}
func (UnimplementedSequenceServer) GetCurrentSequenceByUserIds(context.Context, *GetCurrentSequenceByUserIdsRequest) (*GetCurrentSequenceByUserIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentSequenceByUserIds not implemented")
}
func (UnimplementedSequenceServer) GenRoomSequenceByRoomIds(context.Context, *GenRoomSequenceByRoomIdsRequest) (*GenRoomSequenceByRoomIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenRoomSequenceByRoomIds not implemented")
}
func (UnimplementedSequenceServer) GetCurrentRoomSequenceByRoomIds(context.Context, *GetCurrentSequenceByRoomIdsRequest) (*GetCurrentSequenceByRoomIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentRoomSequenceByRoomIds not implemented")
}
func (UnimplementedSequenceServer) mustEmbedUnimplementedSequenceServer() {}

// UnsafeSequenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SequenceServer will
// result in compilation errors.
type UnsafeSequenceServer interface {
	mustEmbedUnimplementedSequenceServer()
}

func RegisterSequenceServer(s grpc.ServiceRegistrar, srv SequenceServer) {
	s.RegisterService(&Sequence_ServiceDesc, srv)
}

func _Sequence_GenSequenceByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenSequenceByUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GenSequenceByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GenSequenceByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GenSequenceByUserIds(ctx, req.(*GenSequenceByUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequence_GetCurrentSequenceByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentSequenceByUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GetCurrentSequenceByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GetCurrentSequenceByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GetCurrentSequenceByUserIds(ctx, req.(*GetCurrentSequenceByUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequence_GenRoomSequenceByRoomIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenRoomSequenceByRoomIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GenRoomSequenceByRoomIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GenRoomSequenceByRoomIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GenRoomSequenceByRoomIds(ctx, req.(*GenRoomSequenceByRoomIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequence_GetCurrentRoomSequenceByRoomIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentSequenceByRoomIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GetCurrentRoomSequenceByRoomIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GetCurrentRoomSequenceByRoomIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GetCurrentRoomSequenceByRoomIds(ctx, req.(*GetCurrentSequenceByRoomIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sequence_ServiceDesc is the grpc.ServiceDesc for Sequence service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sequence_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sequence.v1.Sequence",
	HandlerType: (*SequenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenSequenceByUserIds",
			Handler:    _Sequence_GenSequenceByUserIds_Handler,
		},
		{
			MethodName: "GetCurrentSequenceByUserIds",
			Handler:    _Sequence_GetCurrentSequenceByUserIds_Handler,
		},
		{
			MethodName: "GenRoomSequenceByRoomIds",
			Handler:    _Sequence_GenRoomSequenceByRoomIds_Handler,
		},
		{
			MethodName: "GetCurrentRoomSequenceByRoomIds",
			Handler:    _Sequence_GetCurrentRoomSequenceByRoomIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sequence/v1/service.proto",
}
