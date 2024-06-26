// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.6
// source: api/gateway/v1/service.proto

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
	GatewaySrv_Sync_FullMethodName              = "/gateway.v1.GatewaySrv/Sync"
	GatewaySrv_Deliver_FullMethodName           = "/gateway.v1.GatewaySrv/Deliver"
	GatewaySrv_Broadcast_FullMethodName         = "/gateway.v1.GatewaySrv/Broadcast"
	GatewaySrv_Action_FullMethodName            = "/gateway.v1.GatewaySrv/Action"
	GatewaySrv_DisconnectedForce_FullMethodName = "/gateway.v1.GatewaySrv/DisconnectedForce"
)

// GatewaySrvClient is the client API for GatewaySrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewaySrvClient interface {
	// 接收同步消息
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncReply, error)
	// 接收下行消息
	Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverReply, error)
	// 广播
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastReply, error)
	// 指令
	Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionReply, error)
	// 强制断开连接
	DisconnectedForce(ctx context.Context, in *DisconnectForceRequest, opts ...grpc.CallOption) (*DisconnectForceReply, error)
}

type gatewaySrvClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewaySrvClient(cc grpc.ClientConnInterface) GatewaySrvClient {
	return &gatewaySrvClient{cc}
}

func (c *gatewaySrvClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncReply, error) {
	out := new(SyncReply)
	err := c.cc.Invoke(ctx, GatewaySrv_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewaySrvClient) Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverReply, error) {
	out := new(DeliverReply)
	err := c.cc.Invoke(ctx, GatewaySrv_Deliver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewaySrvClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastReply, error) {
	out := new(BroadcastReply)
	err := c.cc.Invoke(ctx, GatewaySrv_Broadcast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewaySrvClient) Action(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionReply, error) {
	out := new(ActionReply)
	err := c.cc.Invoke(ctx, GatewaySrv_Action_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewaySrvClient) DisconnectedForce(ctx context.Context, in *DisconnectForceRequest, opts ...grpc.CallOption) (*DisconnectForceReply, error) {
	out := new(DisconnectForceReply)
	err := c.cc.Invoke(ctx, GatewaySrv_DisconnectedForce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewaySrvServer is the server API for GatewaySrv service.
// All implementations must embed UnimplementedGatewaySrvServer
// for forward compatibility
type GatewaySrvServer interface {
	// 接收同步消息
	Sync(context.Context, *SyncRequest) (*SyncReply, error)
	// 接收下行消息
	Deliver(context.Context, *DeliverRequest) (*DeliverReply, error)
	// 广播
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastReply, error)
	// 指令
	Action(context.Context, *ActionRequest) (*ActionReply, error)
	// 强制断开连接
	DisconnectedForce(context.Context, *DisconnectForceRequest) (*DisconnectForceReply, error)
	mustEmbedUnimplementedGatewaySrvServer()
}

// UnimplementedGatewaySrvServer must be embedded to have forward compatible implementations.
type UnimplementedGatewaySrvServer struct {
}

func (UnimplementedGatewaySrvServer) Sync(context.Context, *SyncRequest) (*SyncReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedGatewaySrvServer) Deliver(context.Context, *DeliverRequest) (*DeliverReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}
func (UnimplementedGatewaySrvServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedGatewaySrvServer) Action(context.Context, *ActionRequest) (*ActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedGatewaySrvServer) DisconnectedForce(context.Context, *DisconnectForceRequest) (*DisconnectForceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectedForce not implemented")
}
func (UnimplementedGatewaySrvServer) mustEmbedUnimplementedGatewaySrvServer() {}

// UnsafeGatewaySrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewaySrvServer will
// result in compilation errors.
type UnsafeGatewaySrvServer interface {
	mustEmbedUnimplementedGatewaySrvServer()
}

func RegisterGatewaySrvServer(s grpc.ServiceRegistrar, srv GatewaySrvServer) {
	s.RegisterService(&GatewaySrv_ServiceDesc, srv)
}

func _GatewaySrv_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySrvServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySrv_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySrvServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewaySrv_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySrvServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySrv_Deliver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySrvServer).Deliver(ctx, req.(*DeliverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewaySrv_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySrvServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySrv_Broadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySrvServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewaySrv_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySrvServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySrv_Action_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySrvServer).Action(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewaySrv_DisconnectedForce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectForceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySrvServer).DisconnectedForce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySrv_DisconnectedForce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySrvServer).DisconnectedForce(ctx, req.(*DisconnectForceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewaySrv_ServiceDesc is the grpc.ServiceDesc for GatewaySrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewaySrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.v1.GatewaySrv",
	HandlerType: (*GatewaySrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _GatewaySrv_Sync_Handler,
		},
		{
			MethodName: "Deliver",
			Handler:    _GatewaySrv_Deliver_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _GatewaySrv_Broadcast_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _GatewaySrv_Action_Handler,
		},
		{
			MethodName: "DisconnectedForce",
			Handler:    _GatewaySrv_DisconnectedForce_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/gateway/v1/service.proto",
}
