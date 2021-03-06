// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterClient interface {
	// Sends a greeting
	RpcCallBU(ctx context.Context, in *RpcRequest, opts ...grpc.CallOption) (*RpcResponse, error)
}

type routerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterClient(cc grpc.ClientConnInterface) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) RpcCallBU(ctx context.Context, in *RpcRequest, opts ...grpc.CallOption) (*RpcResponse, error) {
	out := new(RpcResponse)
	err := c.cc.Invoke(ctx, "/Router.Router/rpcCallBU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
// All implementations must embed UnimplementedRouterServer
// for forward compatibility
type RouterServer interface {
	// Sends a greeting
	RpcCallBU(context.Context, *RpcRequest) (*RpcResponse, error)
	mustEmbedUnimplementedRouterServer()
}

// UnimplementedRouterServer must be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (UnimplementedRouterServer) RpcCallBU(context.Context, *RpcRequest) (*RpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcCallBU not implemented")
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

func _Router_RpcCallBU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).RpcCallBU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router.Router/rpcCallBU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).RpcCallBU(ctx, req.(*RpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Router_ServiceDesc is the grpc.ServiceDesc for Router service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Router_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Router.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "rpcCallBU",
			Handler:    _Router_RpcCallBU_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-router.proto",
}
