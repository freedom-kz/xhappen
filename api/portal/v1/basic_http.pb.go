// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.15.6
// source: api/portal/v1/basic.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationConfigGetCommonConfig = "/portal.v1.Config/GetCommonConfig"

type ConfigHTTPServer interface {
	// GetCommonConfig 获取基础配置
	GetCommonConfig(context.Context, *GetCommonConfigRequest) (*GetCommonConfigReply, error)
}

func RegisterConfigHTTPServer(s *http.Server, srv ConfigHTTPServer) {
	r := s.Route("/")
	r.POST("/basic/getconfig", _Config_GetCommonConfig0_HTTP_Handler(srv))
}

func _Config_GetCommonConfig0_HTTP_Handler(srv ConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommonConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationConfigGetCommonConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommonConfig(ctx, req.(*GetCommonConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCommonConfigReply)
		return ctx.Result(200, reply)
	}
}

type ConfigHTTPClient interface {
	GetCommonConfig(ctx context.Context, req *GetCommonConfigRequest, opts ...http.CallOption) (rsp *GetCommonConfigReply, err error)
}

type ConfigHTTPClientImpl struct {
	cc *http.Client
}

func NewConfigHTTPClient(client *http.Client) ConfigHTTPClient {
	return &ConfigHTTPClientImpl{client}
}

func (c *ConfigHTTPClientImpl) GetCommonConfig(ctx context.Context, in *GetCommonConfigRequest, opts ...http.CallOption) (*GetCommonConfigReply, error) {
	var out GetCommonConfigReply
	pattern := "/basic/getconfig"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationConfigGetCommonConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
