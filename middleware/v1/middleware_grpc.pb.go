// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: middleware/v1/middleware.proto

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
	OneMiddleware_ParseIP_FullMethodName = "/oneMiddleware.v1.oneMiddleware/ParseIP"
)

// OneMiddlewareClient is the client API for OneMiddleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OneMiddlewareClient interface {
	ParseIP(ctx context.Context, in *ParseIPReq, opts ...grpc.CallOption) (*ParseIPReply, error)
}

type oneMiddlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewOneMiddlewareClient(cc grpc.ClientConnInterface) OneMiddlewareClient {
	return &oneMiddlewareClient{cc}
}

func (c *oneMiddlewareClient) ParseIP(ctx context.Context, in *ParseIPReq, opts ...grpc.CallOption) (*ParseIPReply, error) {
	out := new(ParseIPReply)
	err := c.cc.Invoke(ctx, OneMiddleware_ParseIP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OneMiddlewareServer is the server API for OneMiddleware service.
// All implementations must embed UnimplementedOneMiddlewareServer
// for forward compatibility
type OneMiddlewareServer interface {
	ParseIP(context.Context, *ParseIPReq) (*ParseIPReply, error)
	mustEmbedUnimplementedOneMiddlewareServer()
}

// UnimplementedOneMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedOneMiddlewareServer struct {
}

func (UnimplementedOneMiddlewareServer) ParseIP(context.Context, *ParseIPReq) (*ParseIPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseIP not implemented")
}
func (UnimplementedOneMiddlewareServer) mustEmbedUnimplementedOneMiddlewareServer() {}

// UnsafeOneMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OneMiddlewareServer will
// result in compilation errors.
type UnsafeOneMiddlewareServer interface {
	mustEmbedUnimplementedOneMiddlewareServer()
}

func RegisterOneMiddlewareServer(s grpc.ServiceRegistrar, srv OneMiddlewareServer) {
	s.RegisterService(&OneMiddleware_ServiceDesc, srv)
}

func _OneMiddleware_ParseIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseIPReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OneMiddlewareServer).ParseIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OneMiddleware_ParseIP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OneMiddlewareServer).ParseIP(ctx, req.(*ParseIPReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OneMiddleware_ServiceDesc is the grpc.ServiceDesc for OneMiddleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OneMiddleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oneMiddleware.v1.oneMiddleware",
	HandlerType: (*OneMiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseIP",
			Handler:    _OneMiddleware_ParseIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middleware/v1/middleware.proto",
}
