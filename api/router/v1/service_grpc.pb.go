// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.6
// source: api/router/v1/service.proto

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
	Router_DeviceBind_FullMethodName              = "/sequence.v1.Router/DeviceBind"
	Router_DeviceAuth_FullMethodName              = "/sequence.v1.Router/DeviceAuth"
	Router_DeviceUnBind_FullMethodName            = "/sequence.v1.Router/DeviceUnBind"
	Router_GetServerByUserIds_FullMethodName      = "/sequence.v1.Router/GetServerByUserIds"
	Router_GetLocalServerByUserIds_FullMethodName = "/sequence.v1.Router/GetLocalServerByUserIds"
	Router_SaveRoomServer_FullMethodName          = "/sequence.v1.Router/SaveRoomServer"
	Router_SaveLocalRoomServer_FullMethodName     = "/sequence.v1.Router/SaveLocalRoomServer"
	Router_GetRoomServerByID_FullMethodName       = "/sequence.v1.Router/GetRoomServerByID"
	Router_GetLocalRoomServerByID_FullMethodName  = "/sequence.v1.Router/GetLocalRoomServerByID"
)

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterClient interface {
	// 设备相关操作
	DeviceBind(ctx context.Context, in *DeviceBindRequest, opts ...grpc.CallOption) (*DeviceBindReply, error)
	DeviceAuth(ctx context.Context, in *DeviceAuthRequest, opts ...grpc.CallOption) (*DeviceAuthReply, error)
	DeviceUnBind(ctx context.Context, in *DeviceUnBindRequest, opts ...grpc.CallOption) (*DeviceUnBindReply, error)
	// 查询用户路由
	GetServerByUserIds(ctx context.Context, in *GetServerByUserIdsRequest, opts ...grpc.CallOption) (*GetServerByUserIdsReply, error)
	GetLocalServerByUserIds(ctx context.Context, in *GetLocalServerByUserIdsRequest, opts ...grpc.CallOption) (*GetLocalServerByUserIdsReply, error)
	// 查询房间路由
	SaveRoomServer(ctx context.Context, in *SaveRoomServerRequest, opts ...grpc.CallOption) (*SaveRoomServerReply, error)
	SaveLocalRoomServer(ctx context.Context, in *SaveLocalRoomServerRequest, opts ...grpc.CallOption) (*SaveLocalRoomServerReply, error)
	GetRoomServerByID(ctx context.Context, in *GetRoomServerByIDRequest, opts ...grpc.CallOption) (*GetRoomServerByIDReply, error)
	GetLocalRoomServerByID(ctx context.Context, in *GetLocalRoomServerByIDRequest, opts ...grpc.CallOption) (*GetLocalRoomServerByIDReply, error)
}

type routerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterClient(cc grpc.ClientConnInterface) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) DeviceBind(ctx context.Context, in *DeviceBindRequest, opts ...grpc.CallOption) (*DeviceBindReply, error) {
	out := new(DeviceBindReply)
	err := c.cc.Invoke(ctx, Router_DeviceBind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) DeviceAuth(ctx context.Context, in *DeviceAuthRequest, opts ...grpc.CallOption) (*DeviceAuthReply, error) {
	out := new(DeviceAuthReply)
	err := c.cc.Invoke(ctx, Router_DeviceAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) DeviceUnBind(ctx context.Context, in *DeviceUnBindRequest, opts ...grpc.CallOption) (*DeviceUnBindReply, error) {
	out := new(DeviceUnBindReply)
	err := c.cc.Invoke(ctx, Router_DeviceUnBind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) GetServerByUserIds(ctx context.Context, in *GetServerByUserIdsRequest, opts ...grpc.CallOption) (*GetServerByUserIdsReply, error) {
	out := new(GetServerByUserIdsReply)
	err := c.cc.Invoke(ctx, Router_GetServerByUserIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) GetLocalServerByUserIds(ctx context.Context, in *GetLocalServerByUserIdsRequest, opts ...grpc.CallOption) (*GetLocalServerByUserIdsReply, error) {
	out := new(GetLocalServerByUserIdsReply)
	err := c.cc.Invoke(ctx, Router_GetLocalServerByUserIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) SaveRoomServer(ctx context.Context, in *SaveRoomServerRequest, opts ...grpc.CallOption) (*SaveRoomServerReply, error) {
	out := new(SaveRoomServerReply)
	err := c.cc.Invoke(ctx, Router_SaveRoomServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) SaveLocalRoomServer(ctx context.Context, in *SaveLocalRoomServerRequest, opts ...grpc.CallOption) (*SaveLocalRoomServerReply, error) {
	out := new(SaveLocalRoomServerReply)
	err := c.cc.Invoke(ctx, Router_SaveLocalRoomServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) GetRoomServerByID(ctx context.Context, in *GetRoomServerByIDRequest, opts ...grpc.CallOption) (*GetRoomServerByIDReply, error) {
	out := new(GetRoomServerByIDReply)
	err := c.cc.Invoke(ctx, Router_GetRoomServerByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) GetLocalRoomServerByID(ctx context.Context, in *GetLocalRoomServerByIDRequest, opts ...grpc.CallOption) (*GetLocalRoomServerByIDReply, error) {
	out := new(GetLocalRoomServerByIDReply)
	err := c.cc.Invoke(ctx, Router_GetLocalRoomServerByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
// All implementations must embed UnimplementedRouterServer
// for forward compatibility
type RouterServer interface {
	// 设备相关操作
	DeviceBind(context.Context, *DeviceBindRequest) (*DeviceBindReply, error)
	DeviceAuth(context.Context, *DeviceAuthRequest) (*DeviceAuthReply, error)
	DeviceUnBind(context.Context, *DeviceUnBindRequest) (*DeviceUnBindReply, error)
	// 查询用户路由
	GetServerByUserIds(context.Context, *GetServerByUserIdsRequest) (*GetServerByUserIdsReply, error)
	GetLocalServerByUserIds(context.Context, *GetLocalServerByUserIdsRequest) (*GetLocalServerByUserIdsReply, error)
	// 查询房间路由
	SaveRoomServer(context.Context, *SaveRoomServerRequest) (*SaveRoomServerReply, error)
	SaveLocalRoomServer(context.Context, *SaveLocalRoomServerRequest) (*SaveLocalRoomServerReply, error)
	GetRoomServerByID(context.Context, *GetRoomServerByIDRequest) (*GetRoomServerByIDReply, error)
	GetLocalRoomServerByID(context.Context, *GetLocalRoomServerByIDRequest) (*GetLocalRoomServerByIDReply, error)
	mustEmbedUnimplementedRouterServer()
}

// UnimplementedRouterServer must be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (UnimplementedRouterServer) DeviceBind(context.Context, *DeviceBindRequest) (*DeviceBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceBind not implemented")
}
func (UnimplementedRouterServer) DeviceAuth(context.Context, *DeviceAuthRequest) (*DeviceAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceAuth not implemented")
}
func (UnimplementedRouterServer) DeviceUnBind(context.Context, *DeviceUnBindRequest) (*DeviceUnBindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceUnBind not implemented")
}
func (UnimplementedRouterServer) GetServerByUserIds(context.Context, *GetServerByUserIdsRequest) (*GetServerByUserIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerByUserIds not implemented")
}
func (UnimplementedRouterServer) GetLocalServerByUserIds(context.Context, *GetLocalServerByUserIdsRequest) (*GetLocalServerByUserIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalServerByUserIds not implemented")
}
func (UnimplementedRouterServer) SaveRoomServer(context.Context, *SaveRoomServerRequest) (*SaveRoomServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRoomServer not implemented")
}
func (UnimplementedRouterServer) SaveLocalRoomServer(context.Context, *SaveLocalRoomServerRequest) (*SaveLocalRoomServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLocalRoomServer not implemented")
}
func (UnimplementedRouterServer) GetRoomServerByID(context.Context, *GetRoomServerByIDRequest) (*GetRoomServerByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomServerByID not implemented")
}
func (UnimplementedRouterServer) GetLocalRoomServerByID(context.Context, *GetLocalRoomServerByIDRequest) (*GetLocalRoomServerByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalRoomServerByID not implemented")
}
func (UnimplementedRouterServer) mustEmbedUnimplementedRouterServer() {}

// UnsafeRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServer will
// result in compilation errors.
type UnsafeRouterServer interface {
	mustEmbedUnimplementedRouterServer()
}

func RegisterRouterServer(s grpc.ServiceRegistrar, srv RouterServer) {
	s.RegisterService(&Router_ServiceDesc, srv)
}

func _Router_DeviceBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).DeviceBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_DeviceBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).DeviceBind(ctx, req.(*DeviceBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_DeviceAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).DeviceAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_DeviceAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).DeviceAuth(ctx, req.(*DeviceAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_DeviceUnBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceUnBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).DeviceUnBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_DeviceUnBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).DeviceUnBind(ctx, req.(*DeviceUnBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_GetServerByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerByUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).GetServerByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_GetServerByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).GetServerByUserIds(ctx, req.(*GetServerByUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_GetLocalServerByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalServerByUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).GetLocalServerByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_GetLocalServerByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).GetLocalServerByUserIds(ctx, req.(*GetLocalServerByUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SaveRoomServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRoomServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SaveRoomServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_SaveRoomServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SaveRoomServer(ctx, req.(*SaveRoomServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SaveLocalRoomServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLocalRoomServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SaveLocalRoomServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_SaveLocalRoomServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SaveLocalRoomServer(ctx, req.(*SaveLocalRoomServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_GetRoomServerByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomServerByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).GetRoomServerByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_GetRoomServerByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).GetRoomServerByID(ctx, req.(*GetRoomServerByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_GetLocalRoomServerByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalRoomServerByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).GetLocalRoomServerByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_GetLocalRoomServerByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).GetLocalRoomServerByID(ctx, req.(*GetLocalRoomServerByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Router_ServiceDesc is the grpc.ServiceDesc for Router service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Router_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sequence.v1.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceBind",
			Handler:    _Router_DeviceBind_Handler,
		},
		{
			MethodName: "DeviceAuth",
			Handler:    _Router_DeviceAuth_Handler,
		},
		{
			MethodName: "DeviceUnBind",
			Handler:    _Router_DeviceUnBind_Handler,
		},
		{
			MethodName: "GetServerByUserIds",
			Handler:    _Router_GetServerByUserIds_Handler,
		},
		{
			MethodName: "GetLocalServerByUserIds",
			Handler:    _Router_GetLocalServerByUserIds_Handler,
		},
		{
			MethodName: "SaveRoomServer",
			Handler:    _Router_SaveRoomServer_Handler,
		},
		{
			MethodName: "SaveLocalRoomServer",
			Handler:    _Router_SaveLocalRoomServer_Handler,
		},
		{
			MethodName: "GetRoomServerByID",
			Handler:    _Router_GetRoomServerByID_Handler,
		},
		{
			MethodName: "GetLocalRoomServerByID",
			Handler:    _Router_GetLocalRoomServerByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/router/v1/service.proto",
}
