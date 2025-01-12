// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.6
// source: api/portal/v1/room.proto

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
	Room_CreateRoom_FullMethodName = "/portal.v1.Room/CreateRoom"
)

// RoomClient is the client API for Room service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomClient interface {
	// 获取基础配置
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomReply, error)
}

type roomClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomClient(cc grpc.ClientConnInterface) RoomClient {
	return &roomClient{cc}
}

func (c *roomClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomReply, error) {
	out := new(CreateRoomReply)
	err := c.cc.Invoke(ctx, Room_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServer is the server API for Room service.
// All implementations must embed UnimplementedRoomServer
// for forward compatibility
type RoomServer interface {
	// 获取基础配置
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomReply, error)
	mustEmbedUnimplementedRoomServer()
}

// UnimplementedRoomServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServer struct {
}

func (UnimplementedRoomServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServer) mustEmbedUnimplementedRoomServer() {}

// UnsafeRoomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServer will
// result in compilation errors.
type UnsafeRoomServer interface {
	mustEmbedUnimplementedRoomServer()
}

func RegisterRoomServer(s grpc.ServiceRegistrar, srv RoomServer) {
	s.RegisterService(&Room_ServiceDesc, srv)
}

func _Room_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Room_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Room_ServiceDesc is the grpc.ServiceDesc for Room service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Room_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "portal.v1.Room",
	HandlerType: (*RoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _Room_CreateRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/portal/v1/room.proto",
}
