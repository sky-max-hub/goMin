// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ServerTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerTimeRequest) Reset() {
	*x = ServerTimeRequest{}
	mi := &file_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerTimeRequest) ProtoMessage() {}

func (x *ServerTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerTimeRequest.ProtoReflect.Descriptor instead.
func (*ServerTimeRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type ServerTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *ServerTimeResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ServerTimeResponse) Reset() {
	*x = ServerTimeResponse{}
	mi := &file_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerTimeResponse) ProtoMessage() {}

func (x *ServerTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerTimeResponse.ProtoReflect.Descriptor instead.
func (*ServerTimeResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServerTimeResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ServerTimeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ServerTimeResponse) GetData() *ServerTimeResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ServerTimeResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime uint64 `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
}

func (x *ServerTimeResponseData) Reset() {
	*x = ServerTimeResponseData{}
	mi := &file_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerTimeResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerTimeResponseData) ProtoMessage() {}

func (x *ServerTimeResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerTimeResponseData.ProtoReflect.Descriptor instead.
func (*ServerTimeResponseData) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *ServerTimeResponseData) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x16, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x50, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_proto_goTypes = []any{
	(*ServerTimeRequest)(nil),      // 0: server.ServerTimeRequest
	(*ServerTimeResponse)(nil),     // 1: server.ServerTimeResponse
	(*ServerTimeResponseData)(nil), // 2: server.ServerTimeResponseData
}
var file_server_proto_depIdxs = []int32{
	2, // 0: server.ServerTimeResponse.data:type_name -> server.ServerTimeResponseData
	0, // 1: server.Server.GetServerTime:input_type -> server.ServerTimeRequest
	1, // 2: server.Server.GetServerTime:output_type -> server.ServerTimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	GetServerTime(ctx context.Context, in *ServerTimeRequest, opts ...grpc.CallOption) (*ServerTimeResponse, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetServerTime(ctx context.Context, in *ServerTimeRequest, opts ...grpc.CallOption) (*ServerTimeResponse, error) {
	out := new(ServerTimeResponse)
	err := c.cc.Invoke(ctx, "/server.Server/GetServerTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	GetServerTime(context.Context, *ServerTimeRequest) (*ServerTimeResponse, error)
}

// UnimplementedServerServer can be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (*UnimplementedServerServer) GetServerTime(context.Context, *ServerTimeRequest) (*ServerTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTime not implemented")
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_GetServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Server/GetServerTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetServerTime(ctx, req.(*ServerTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerTime",
			Handler:    _Server_GetServerTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
