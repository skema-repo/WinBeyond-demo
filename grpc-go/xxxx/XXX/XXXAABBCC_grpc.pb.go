// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: XXXAABBCC.proto

package XXX

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

// XXXAABBCCClient is the client API for XXXAABBCC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XXXAABBCCClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type xXXAABBCCClient struct {
	cc grpc.ClientConnInterface
}

func NewXXXAABBCCClient(cc grpc.ClientConnInterface) XXXAABBCCClient {
	return &xXXAABBCCClient{cc}
}

func (c *xXXAABBCCClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxxx.XXX.XXXAABBCC/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XXXAABBCCServer is the server API for XXXAABBCC service.
// All implementations must embed UnimplementedXXXAABBCCServer
// for forward compatibility
type XXXAABBCCServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedXXXAABBCCServer()
}

// UnimplementedXXXAABBCCServer must be embedded to have forward compatible implementations.
type UnimplementedXXXAABBCCServer struct {
}

func (UnimplementedXXXAABBCCServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedXXXAABBCCServer) mustEmbedUnimplementedXXXAABBCCServer() {}

// UnsafeXXXAABBCCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XXXAABBCCServer will
// result in compilation errors.
type UnsafeXXXAABBCCServer interface {
	mustEmbedUnimplementedXXXAABBCCServer()
}

func RegisterXXXAABBCCServer(s grpc.ServiceRegistrar, srv XXXAABBCCServer) {
	s.RegisterService(&XXXAABBCC_ServiceDesc, srv)
}

func _XXXAABBCC_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XXXAABBCCServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxx.XXX.XXXAABBCC/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XXXAABBCCServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// XXXAABBCC_ServiceDesc is the grpc.ServiceDesc for XXXAABBCC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XXXAABBCC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxxx.XXX.XXXAABBCC",
	HandlerType: (*XXXAABBCCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _XXXAABBCC_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "XXXAABBCC.proto",
}
