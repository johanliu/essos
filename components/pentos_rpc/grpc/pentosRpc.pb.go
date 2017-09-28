// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pentosRpc.proto

/*
Package pentosrpc is a generated protocol buffer package.

It is generated from these files:
	pentosRpc.proto

It has these top-level messages:
	Request
	Reply
*/
package pentosrpc

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
	Arg []string `protobuf:"bytes,1,rep,name=arg" json:"arg,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetArg() []string {
	if m != nil {
		return m.Arg
	}
	return nil
}

type Reply struct {
	Code int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  []string `protobuf:"bytes,2,rep,name=msg" json:"msg,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Reply) GetMsg() []string {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "pentosrpc.Request")
	proto.RegisterType((*Reply)(nil), "pentosrpc.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	// terran
	AddTerran(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	DeleteTerran(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	ListTerrans(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// complete flag
	CheckComplete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	ListIncompleteFlags(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	MarkComplete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AddTerran(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/pentosrpc.Service/AddTerran", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteTerran(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/pentosrpc.Service/DeleteTerran", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListTerrans(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/pentosrpc.Service/ListTerrans", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CheckComplete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/pentosrpc.Service/CheckComplete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListIncompleteFlags(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/pentosrpc.Service/ListIncompleteFlags", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MarkComplete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/pentosrpc.Service/markComplete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	// terran
	AddTerran(context.Context, *Request) (*Reply, error)
	DeleteTerran(context.Context, *Request) (*Reply, error)
	ListTerrans(context.Context, *Request) (*Reply, error)
	// complete flag
	CheckComplete(context.Context, *Request) (*Reply, error)
	ListIncompleteFlags(context.Context, *Request) (*Reply, error)
	MarkComplete(context.Context, *Request) (*Reply, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_AddTerran_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddTerran(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pentosrpc.Service/AddTerran",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddTerran(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteTerran_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteTerran(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pentosrpc.Service/DeleteTerran",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteTerran(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListTerrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListTerrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pentosrpc.Service/ListTerrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListTerrans(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CheckComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CheckComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pentosrpc.Service/CheckComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CheckComplete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListIncompleteFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListIncompleteFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pentosrpc.Service/ListIncompleteFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListIncompleteFlags(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_MarkComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).MarkComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pentosrpc.Service/MarkComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).MarkComplete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pentosrpc.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTerran",
			Handler:    _Service_AddTerran_Handler,
		},
		{
			MethodName: "DeleteTerran",
			Handler:    _Service_DeleteTerran_Handler,
		},
		{
			MethodName: "ListTerrans",
			Handler:    _Service_ListTerrans_Handler,
		},
		{
			MethodName: "CheckComplete",
			Handler:    _Service_CheckComplete_Handler,
		},
		{
			MethodName: "ListIncompleteFlags",
			Handler:    _Service_ListIncompleteFlags_Handler,
		},
		{
			MethodName: "markComplete",
			Handler:    _Service_MarkComplete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pentosRpc.proto",
}

func init() { proto.RegisterFile("pentosRpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x6b, 0x2d, 0x19, 0x15, 0xcb, 0x7a, 0x30, 0xea, 0xa5, 0xe4, 0xd4, 0x8b, 0x7b,
	0xb0, 0xa8, 0x27, 0x0f, 0x5a, 0x11, 0x04, 0x4f, 0xab, 0x2f, 0xb0, 0x6e, 0x86, 0x34, 0x98, 0xcd,
	0x8e, 0xb3, 0xab, 0x50, 0xdf, 0xd5, 0x77, 0x91, 0x6d, 0x82, 0xe0, 0x2d, 0xb9, 0xfd, 0xfc, 0xfc,
	0x1f, 0x1f, 0xc3, 0xc0, 0x11, 0x61, 0x13, 0x9c, 0x57, 0x64, 0x24, 0xb1, 0x0b, 0x4e, 0xa4, 0x6d,
	0xc1, 0x64, 0xf2, 0x73, 0x98, 0x2a, 0xfc, 0xf8, 0x44, 0x1f, 0xc4, 0x0c, 0xc6, 0x9a, 0xcb, 0x2c,
	0x99, 0x8f, 0x17, 0xa9, 0x8a, 0x31, 0xbf, 0x80, 0x89, 0x42, 0xaa, 0x37, 0x42, 0xc0, 0xae, 0x71,
	0x05, 0x66, 0xc9, 0x3c, 0x59, 0x4c, 0xd4, 0x36, 0xc7, 0xb9, 0xf5, 0x65, 0x36, 0x6a, 0xe7, 0xd6,
	0x97, 0x97, 0x3f, 0x23, 0x98, 0xbe, 0x20, 0x7f, 0x55, 0x06, 0xc5, 0x12, 0xd2, 0xbb, 0xa2, 0x78,
	0x45, 0x66, 0xdd, 0x08, 0x21, 0xff, 0x84, 0xb2, 0xb3, 0x9d, 0xcd, 0xfe, 0x75, 0x54, 0x6f, 0xf2,
	0x1d, 0x71, 0x0d, 0x07, 0x0f, 0x58, 0x63, 0xc0, 0x81, 0xdc, 0x15, 0xec, 0x3f, 0x57, 0x3e, 0xb4,
	0x94, 0xef, 0x8d, 0xdd, 0xc0, 0xe1, 0x6a, 0x8d, 0xe6, 0x7d, 0xe5, 0x2c, 0x45, 0x6b, 0x6f, 0xf0,
	0x16, 0x8e, 0xa3, 0xef, 0xa9, 0x31, 0x1d, 0xf9, 0x58, 0xeb, 0xd2, 0x0f, 0x39, 0xd3, 0x6a, 0x1e,
	0xac, 0xbd, 0x3f, 0x85, 0x13, 0xe3, 0xac, 0xb4, 0x95, 0x76, 0xdf, 0x6b, 0x6c, 0xba, 0x85, 0x64,
	0x32, 0x6f, 0x7b, 0xdb, 0xc7, 0x2e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xf4, 0x63, 0x09,
	0xeb, 0x01, 0x00, 0x00,
}
