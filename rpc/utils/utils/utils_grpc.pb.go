// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: utils.proto

package utils

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

// UtilerClient is the client API for Utiler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UtilerClient interface {
	MD5(ctx context.Context, in *MD5Reqeust, opts ...grpc.CallOption) (*MD5Reponse, error)
}

type utilerClient struct {
	cc grpc.ClientConnInterface
}

func NewUtilerClient(cc grpc.ClientConnInterface) UtilerClient {
	return &utilerClient{cc}
}

func (c *utilerClient) MD5(ctx context.Context, in *MD5Reqeust, opts ...grpc.CallOption) (*MD5Reponse, error) {
	out := new(MD5Reponse)
	err := c.cc.Invoke(ctx, "/utils.Utiler/MD5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilerServer is the server API for Utiler service.
// All implementations must embed UnimplementedUtilerServer
// for forward compatibility
type UtilerServer interface {
	MD5(context.Context, *MD5Reqeust) (*MD5Reponse, error)
	mustEmbedUnimplementedUtilerServer()
}

// UnimplementedUtilerServer must be embedded to have forward compatible implementations.
type UnimplementedUtilerServer struct {
}

func (UnimplementedUtilerServer) MD5(context.Context, *MD5Reqeust) (*MD5Reponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MD5 not implemented")
}
func (UnimplementedUtilerServer) mustEmbedUnimplementedUtilerServer() {}

// UnsafeUtilerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UtilerServer will
// result in compilation errors.
type UnsafeUtilerServer interface {
	mustEmbedUnimplementedUtilerServer()
}

func RegisterUtilerServer(s grpc.ServiceRegistrar, srv UtilerServer) {
	s.RegisterService(&Utiler_ServiceDesc, srv)
}

func _Utiler_MD5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MD5Reqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilerServer).MD5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utils.Utiler/MD5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilerServer).MD5(ctx, req.(*MD5Reqeust))
	}
	return interceptor(ctx, in, info, handler)
}

// Utiler_ServiceDesc is the grpc.ServiceDesc for Utiler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Utiler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "utils.Utiler",
	HandlerType: (*UtilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MD5",
			Handler:    _Utiler_MD5_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "utils.proto",
}