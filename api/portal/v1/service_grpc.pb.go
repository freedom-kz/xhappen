// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.6
// source: api/portal/v1/service.proto

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
	User_SendSMSCode_FullMethodName    = "/portal.v1.User/SendSMSCode"
	User_LoginByMobile_FullMethodName  = "/portal.v1.User/LoginByMobile"
	User_Logout_FullMethodName         = "/portal.v1.User/Logout"
	User_DeRegister_FullMethodName     = "/portal.v1.User/DeRegister"
	User_GetUserProfile_FullMethodName = "/portal.v1.User/GetUserProfile"
	User_GetSelfProfile_FullMethodName = "/portal.v1.User/GetSelfProfile"
	User_UpdateProfile_FullMethodName  = "/portal.v1.User/UpdateProfile"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// Sends SMS code
	SendSMSCode(ctx context.Context, in *SMSCodeRequest, opts ...grpc.CallOption) (*SMSCodeReply, error)
	// mobile login
	LoginByMobile(ctx context.Context, in *LoginByMobileRequest, opts ...grpc.CallOption) (*LoginByMobileReply, error)
	// mobile login
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	// mobile login
	DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DeRegisterReply, error)
	// get user profile
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileReply, error)
	// get self profile
	GetSelfProfile(ctx context.Context, in *GetSelfProfileRequest, opts ...grpc.CallOption) (*GetSelfProfileReply, error)
	// update profile
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) SendSMSCode(ctx context.Context, in *SMSCodeRequest, opts ...grpc.CallOption) (*SMSCodeReply, error) {
	out := new(SMSCodeReply)
	err := c.cc.Invoke(ctx, User_SendSMSCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginByMobile(ctx context.Context, in *LoginByMobileRequest, opts ...grpc.CallOption) (*LoginByMobileReply, error) {
	out := new(LoginByMobileReply)
	err := c.cc.Invoke(ctx, User_LoginByMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, User_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DeRegisterReply, error) {
	out := new(DeRegisterReply)
	err := c.cc.Invoke(ctx, User_DeRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileReply, error) {
	out := new(GetUserProfileReply)
	err := c.cc.Invoke(ctx, User_GetUserProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetSelfProfile(ctx context.Context, in *GetSelfProfileRequest, opts ...grpc.CallOption) (*GetSelfProfileReply, error) {
	out := new(GetSelfProfileReply)
	err := c.cc.Invoke(ctx, User_GetSelfProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileReply, error) {
	out := new(UpdateProfileReply)
	err := c.cc.Invoke(ctx, User_UpdateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// Sends SMS code
	SendSMSCode(context.Context, *SMSCodeRequest) (*SMSCodeReply, error)
	// mobile login
	LoginByMobile(context.Context, *LoginByMobileRequest) (*LoginByMobileReply, error)
	// mobile login
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// mobile login
	DeRegister(context.Context, *DeRegisterRequest) (*DeRegisterReply, error)
	// get user profile
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileReply, error)
	// get self profile
	GetSelfProfile(context.Context, *GetSelfProfileRequest) (*GetSelfProfileReply, error)
	// update profile
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) SendSMSCode(context.Context, *SMSCodeRequest) (*SMSCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMSCode not implemented")
}
func (UnimplementedUserServer) LoginByMobile(context.Context, *LoginByMobileRequest) (*LoginByMobileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByMobile not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *LogoutRequest) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) DeRegister(context.Context, *DeRegisterRequest) (*DeRegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeRegister not implemented")
}
func (UnimplementedUserServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserServer) GetSelfProfile(context.Context, *GetSelfProfileRequest) (*GetSelfProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelfProfile not implemented")
}
func (UnimplementedUserServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_SendSMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMSCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendSMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SendSMSCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendSMSCode(ctx, req.(*SMSCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_LoginByMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginByMobile(ctx, req.(*LoginByMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeRegister(ctx, req.(*DeRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetSelfProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSelfProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetSelfProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetSelfProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetSelfProfile(ctx, req.(*GetSelfProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "portal.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSMSCode",
			Handler:    _User_SendSMSCode_Handler,
		},
		{
			MethodName: "LoginByMobile",
			Handler:    _User_LoginByMobile_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _User_Logout_Handler,
		},
		{
			MethodName: "DeRegister",
			Handler:    _User_DeRegister_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _User_GetUserProfile_Handler,
		},
		{
			MethodName: "GetSelfProfile",
			Handler:    _User_GetSelfProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _User_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/portal/v1/service.proto",
}
