// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: app/portal/internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Trace  *Trace  `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	Auth   *Auth   `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetTrace() *Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt *Auth_Jwt `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Auth) GetJwt() *Auth_Jwt {
	if x != nil {
		return x.Jwt
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Data_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	Etcd     *Data_Etcd     `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`
	Kafka    *Data_Kafka    `protobuf:"bytes,4,opt,name=kafka,proto3" json:"kafka,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data) GetEtcd() *Data_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Data) GetKafka() *Data_Kafka {
	if x != nil {
		return x.Kafka
	}
	return nil
}

type Auth_Jwt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret     string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Expiration int32  `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *Auth_Jwt) Reset() {
	*x = Auth_Jwt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth_Jwt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_Jwt) ProtoMessage() {}

func (x *Auth_Jwt) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_Jwt.ProtoReflect.Descriptor instead.
func (*Auth_Jwt) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Auth_Jwt) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Auth_Jwt) GetExpiration() int32 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Password     string               `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Db           int32                `protobuf:"varint,3,opt,name=db,proto3" json:"db,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,4,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=readTimeout,proto3" json:"readTimeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,6,opt,name=writeTimeout,proto3" json:"writeTimeout,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{4, 1}
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

type Data_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,2,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,3,opt,name=readTimeout,proto3" json:"readTimeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=writeTimeout,proto3" json:"writeTimeout,omitempty"`
}

func (x *Data_Etcd) Reset() {
	*x = Data_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Etcd) ProtoMessage() {}

func (x *Data_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Etcd.ProtoReflect.Descriptor instead.
func (*Data_Etcd) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{4, 2}
}

func (x *Data_Etcd) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Etcd) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Etcd) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Etcd) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

type Data_Kafka struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	SmsCodeTopic string `protobuf:"bytes,2,opt,name=smsCodeTopic,proto3" json:"smsCodeTopic,omitempty"`
}

func (x *Data_Kafka) Reset() {
	*x = Data_Kafka{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_portal_internal_conf_conf_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Kafka) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Kafka) ProtoMessage() {}

func (x *Data_Kafka) ProtoReflect() protoreflect.Message {
	mi := &file_app_portal_internal_conf_conf_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Kafka.ProtoReflect.Descriptor instead.
func (*Data_Kafka) Descriptor() ([]byte, []int) {
	return file_app_portal_internal_conf_conf_proto_rawDescGZIP(), []int{4, 3}
}

func (x *Data_Kafka) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Kafka) GetSmsCodeTopic() string {
	if x != nil {
		return x.SmsCodeTopic
	}
	return ""
}

var File_app_portal_internal_conf_conf_proto protoreflect.FileDescriptor

var file_app_portal_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xac, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12,
	0x2a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x23, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x6d, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x26, 0x0a,
	0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x77, 0x74,
	0x52, 0x03, 0x6a, 0x77, 0x74, 0x1a, 0x3d, 0x0a, 0x03, 0x4a, 0x77, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b, 0x0a, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x4f, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x4f, 0x0a, 0x04, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x9a, 0x06, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64,
	0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x65, 0x74, 0x63,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04,
	0x65, 0x74, 0x63, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x05, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x80,
	0x02, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x64, 0x62, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x1a, 0xd3, 0x01, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x3b,
	0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x72,
	0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x3f, 0x0a, 0x05, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6d, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x42, 0x27, 0x5a, 0x25, 0x78, 0x68, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_portal_internal_conf_conf_proto_rawDescOnce sync.Once
	file_app_portal_internal_conf_conf_proto_rawDescData = file_app_portal_internal_conf_conf_proto_rawDesc
)

func file_app_portal_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_app_portal_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_app_portal_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_portal_internal_conf_conf_proto_rawDescData)
	})
	return file_app_portal_internal_conf_conf_proto_rawDescData
}

var file_app_portal_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_app_portal_internal_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Trace)(nil),               // 1: kratos.api.Trace
	(*Auth)(nil),                // 2: kratos.api.Auth
	(*Server)(nil),              // 3: kratos.api.Server
	(*Data)(nil),                // 4: kratos.api.Data
	(*Auth_Jwt)(nil),            // 5: kratos.api.Auth.Jwt
	(*Server_HTTP)(nil),         // 6: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 7: kratos.api.Server.GRPC
	(*Data_Database)(nil),       // 8: kratos.api.Data.Database
	(*Data_Redis)(nil),          // 9: kratos.api.Data.Redis
	(*Data_Etcd)(nil),           // 10: kratos.api.Data.Etcd
	(*Data_Kafka)(nil),          // 11: kratos.api.Data.Kafka
	(*durationpb.Duration)(nil), // 12: google.protobuf.Duration
}
var file_app_portal_internal_conf_conf_proto_depIdxs = []int32{
	3,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	4,  // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	1,  // 2: kratos.api.Bootstrap.trace:type_name -> kratos.api.Trace
	2,  // 3: kratos.api.Bootstrap.auth:type_name -> kratos.api.Auth
	5,  // 4: kratos.api.Auth.jwt:type_name -> kratos.api.Auth.Jwt
	6,  // 5: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	7,  // 6: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	8,  // 7: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	9,  // 8: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	10, // 9: kratos.api.Data.etcd:type_name -> kratos.api.Data.Etcd
	11, // 10: kratos.api.Data.kafka:type_name -> kratos.api.Data.Kafka
	12, // 11: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	12, // 12: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	12, // 13: kratos.api.Data.Redis.dialTimeout:type_name -> google.protobuf.Duration
	12, // 14: kratos.api.Data.Redis.readTimeout:type_name -> google.protobuf.Duration
	12, // 15: kratos.api.Data.Redis.writeTimeout:type_name -> google.protobuf.Duration
	12, // 16: kratos.api.Data.Etcd.dialTimeout:type_name -> google.protobuf.Duration
	12, // 17: kratos.api.Data.Etcd.readTimeout:type_name -> google.protobuf.Duration
	12, // 18: kratos.api.Data.Etcd.writeTimeout:type_name -> google.protobuf.Duration
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_app_portal_internal_conf_conf_proto_init() }
func file_app_portal_internal_conf_conf_proto_init() {
	if File_app_portal_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_portal_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth_Jwt); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Etcd); i {
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
		file_app_portal_internal_conf_conf_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Kafka); i {
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
			RawDescriptor: file_app_portal_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_portal_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_app_portal_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_app_portal_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_app_portal_internal_conf_conf_proto = out.File
	file_app_portal_internal_conf_conf_proto_rawDesc = nil
	file_app_portal_internal_conf_conf_proto_goTypes = nil
	file_app_portal_internal_conf_conf_proto_depIdxs = nil
}
