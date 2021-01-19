// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol // import "library/protocol"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	RequestJson          []byte   `protobuf:"bytes,1,opt,name=requestJson,proto3" json:"requestJson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_b6b4d65f4dc81bc1, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetRequestJson() []byte {
	if m != nil {
		return m.RequestJson
	}
	return nil
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ResultJson           []byte   `protobuf:"bytes,3,opt,name=resultJson,proto3" json:"resultJson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_b6b4d65f4dc81bc1, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetResultJson() []byte {
	if m != nil {
		return m.ResultJson
	}
	return nil
}

type Args struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Args                 []byte   `protobuf:"bytes,4,opt,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}
func (*Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_b6b4d65f4dc81bc1, []int{2}
}
func (m *Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args.Unmarshal(m, b)
}
func (m *Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args.Marshal(b, m, deterministic)
}
func (dst *Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args.Merge(dst, src)
}
func (m *Args) XXX_Size() int {
	return xxx_messageInfo_Args.Size(m)
}
func (m *Args) XXX_DiscardUnknown() {
	xxx_messageInfo_Args.DiscardUnknown(m)
}

var xxx_messageInfo_Args proto.InternalMessageInfo

func (m *Args) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Args) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Args) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Args) GetArgs() []byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterType((*Response)(nil), "protocol.Response")
	proto.RegisterType((*Args)(nil), "protocol.Args")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcServiceClient is the client API for RpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcServiceClient interface {
	Invoke(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Response, error)
}

type rpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcServiceClient(cc *grpc.ClientConn) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) Invoke(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.RpcService/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServiceServer is the server API for RpcService service.
type RpcServiceServer interface {
	Invoke(context.Context, *Args) (*Response, error)
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.RpcService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).Invoke(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.RpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _RpcService_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_protocol_b6b4d65f4dc81bc1) }

var fileDescriptor_protocol_b6b4d65f4dc81bc1 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x40,
	0x0c, 0x46, 0x09, 0x0d, 0x69, 0x63, 0x50, 0x55, 0x79, 0x40, 0x11, 0x03, 0x8a, 0x32, 0x55, 0x42,
	0x0a, 0x12, 0xac, 0x2c, 0xb0, 0xc1, 0x84, 0x8e, 0x8d, 0xad, 0x4d, 0x4d, 0x88, 0x48, 0xe3, 0x60,
	0xa7, 0x91, 0xf8, 0xf7, 0x28, 0x6e, 0x0e, 0xb2, 0xbd, 0x67, 0xdf, 0x7d, 0xf6, 0x1d, 0x2c, 0x5b,
	0xe1, 0x8e, 0x0b, 0xae, 0x73, 0x03, 0x5c, 0x78, 0xcf, 0x6e, 0x60, 0xee, 0xe8, 0xfb, 0x40, 0xda,
	0x61, 0x0a, 0xe7, 0x72, 0xc4, 0x17, 0xe5, 0x26, 0x09, 0xd2, 0x60, 0x7d, 0xe1, 0xa6, 0xa5, 0xec,
	0x15, 0x16, 0x8e, 0xb4, 0xe5, 0x46, 0x09, 0x11, 0xc2, 0x82, 0x77, 0x64, 0xc7, 0xce, 0x9c, 0x31,
	0xae, 0x60, 0xb6, 0xd7, 0x32, 0x39, 0x4d, 0x83, 0x75, 0xec, 0x06, 0xc4, 0x6b, 0x00, 0x21, 0x3d,
	0xd4, 0xc7, 0xc8, 0x99, 0x45, 0x4e, 0x2a, 0xd9, 0x07, 0x84, 0x8f, 0x52, 0x2a, 0x26, 0x30, 0xef,
	0x49, 0xb4, 0x1a, 0xe7, 0xc6, 0xce, 0xeb, 0xd0, 0x51, 0x92, 0xbe, 0x2a, 0x68, 0xcc, 0xf5, 0x8a,
	0x97, 0x10, 0xed, 0xa9, 0xfb, 0xe4, 0x9d, 0xe5, 0xc6, 0x6e, 0xb4, 0x61, 0xb3, 0x8d, 0x94, 0x9a,
	0x84, 0x36, 0xcd, 0xf8, 0xee, 0x01, 0xc0, 0xb5, 0xc5, 0xdb, 0x78, 0x33, 0x87, 0xe8, 0xb9, 0xe9,
	0xf9, 0x8b, 0x70, 0x99, 0xff, 0xfd, 0xcc, 0xb0, 0xc7, 0x15, 0xfe, 0xbb, 0x7f, 0x69, 0x76, 0xf2,
	0x84, 0xef, 0xab, 0xba, 0xda, 0xca, 0x46, 0x7e, 0x6e, 0x7d, 0x7b, 0x1b, 0x19, 0xdd, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x68, 0xa7, 0x9d, 0x52, 0x5b, 0x01, 0x00, 0x00,
}