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
	Sequence_GetSequenceByIds_FullMethodName             = "/sequence.v1.Sequence/GetSequenceByIds"
	Sequence_GetLocalSequenceByIds_FullMethodName        = "/sequence.v1.Sequence/GetLocalSequenceByIds"
	Sequence_GetCurrentSequenceByIds_FullMethodName      = "/sequence.v1.Sequence/GetCurrentSequenceByIds"
	Sequence_GetLocalCurrentSequenceByIds_FullMethodName = "/sequence.v1.Sequence/GetLocalCurrentSequenceByIds"
)

// SequenceClient is the client API for Sequence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SequenceClient interface {
	GetSequenceByIds(ctx context.Context, in *GetSequenceByIdsRequest, opts ...grpc.CallOption) (*GetSequenceByIdsReply, error)
	GetLocalSequenceByIds(ctx context.Context, in *GetLocalSequenceByIdsRequest, opts ...grpc.CallOption) (*GetLocalSequenceByIdsReply, error)
	GetCurrentSequenceByIds(ctx context.Context, in *GetCurrentSequenceByIdsRequest, opts ...grpc.CallOption) (*GetCurrentSequenceByIdsReply, error)
	GetLocalCurrentSequenceByIds(ctx context.Context, in *GetLocalCurrentSequenceByIdsRequest, opts ...grpc.CallOption) (*GetLocalCurrentSequenceByIdsReply, error)
}

type sequenceClient struct {
	cc grpc.ClientConnInterface
}

func NewSequenceClient(cc grpc.ClientConnInterface) SequenceClient {
	return &sequenceClient{cc}
}

func (c *sequenceClient) GetSequenceByIds(ctx context.Context, in *GetSequenceByIdsRequest, opts ...grpc.CallOption) (*GetSequenceByIdsReply, error) {
	out := new(GetSequenceByIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GetSequenceByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceClient) GetLocalSequenceByIds(ctx context.Context, in *GetLocalSequenceByIdsRequest, opts ...grpc.CallOption) (*GetLocalSequenceByIdsReply, error) {
	out := new(GetLocalSequenceByIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GetLocalSequenceByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceClient) GetCurrentSequenceByIds(ctx context.Context, in *GetCurrentSequenceByIdsRequest, opts ...grpc.CallOption) (*GetCurrentSequenceByIdsReply, error) {
	out := new(GetCurrentSequenceByIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GetCurrentSequenceByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceClient) GetLocalCurrentSequenceByIds(ctx context.Context, in *GetLocalCurrentSequenceByIdsRequest, opts ...grpc.CallOption) (*GetLocalCurrentSequenceByIdsReply, error) {
	out := new(GetLocalCurrentSequenceByIdsReply)
	err := c.cc.Invoke(ctx, Sequence_GetLocalCurrentSequenceByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequenceServer is the server API for Sequence service.
// All implementations must embed UnimplementedSequenceServer
// for forward compatibility
type SequenceServer interface {
	GetSequenceByIds(context.Context, *GetSequenceByIdsRequest) (*GetSequenceByIdsReply, error)
	GetLocalSequenceByIds(context.Context, *GetLocalSequenceByIdsRequest) (*GetLocalSequenceByIdsReply, error)
	GetCurrentSequenceByIds(context.Context, *GetCurrentSequenceByIdsRequest) (*GetCurrentSequenceByIdsReply, error)
	GetLocalCurrentSequenceByIds(context.Context, *GetLocalCurrentSequenceByIdsRequest) (*GetLocalCurrentSequenceByIdsReply, error)
	mustEmbedUnimplementedSequenceServer()
}

// UnimplementedSequenceServer must be embedded to have forward compatible implementations.
type UnimplementedSequenceServer struct {
}

func (UnimplementedSequenceServer) GetSequenceByIds(context.Context, *GetSequenceByIdsRequest) (*GetSequenceByIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSequenceByIds not implemented")
}
func (UnimplementedSequenceServer) GetLocalSequenceByIds(context.Context, *GetLocalSequenceByIdsRequest) (*GetLocalSequenceByIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalSequenceByIds not implemented")
}
func (UnimplementedSequenceServer) GetCurrentSequenceByIds(context.Context, *GetCurrentSequenceByIdsRequest) (*GetCurrentSequenceByIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentSequenceByIds not implemented")
}
func (UnimplementedSequenceServer) GetLocalCurrentSequenceByIds(context.Context, *GetLocalCurrentSequenceByIdsRequest) (*GetLocalCurrentSequenceByIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalCurrentSequenceByIds not implemented")
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

func _Sequence_GetSequenceByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSequenceByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GetSequenceByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GetSequenceByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GetSequenceByIds(ctx, req.(*GetSequenceByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequence_GetLocalSequenceByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalSequenceByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GetLocalSequenceByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GetLocalSequenceByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GetLocalSequenceByIds(ctx, req.(*GetLocalSequenceByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequence_GetCurrentSequenceByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentSequenceByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GetCurrentSequenceByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GetCurrentSequenceByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GetCurrentSequenceByIds(ctx, req.(*GetCurrentSequenceByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sequence_GetLocalCurrentSequenceByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalCurrentSequenceByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServer).GetLocalCurrentSequenceByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sequence_GetLocalCurrentSequenceByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServer).GetLocalCurrentSequenceByIds(ctx, req.(*GetLocalCurrentSequenceByIdsRequest))
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
			MethodName: "GetSequenceByIds",
			Handler:    _Sequence_GetSequenceByIds_Handler,
		},
		{
			MethodName: "GetLocalSequenceByIds",
			Handler:    _Sequence_GetLocalSequenceByIds_Handler,
		},
		{
			MethodName: "GetCurrentSequenceByIds",
			Handler:    _Sequence_GetCurrentSequenceByIds_Handler,
		},
		{
			MethodName: "GetLocalCurrentSequenceByIds",
			Handler:    _Sequence_GetLocalCurrentSequenceByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sequence/v1/service.proto",
}