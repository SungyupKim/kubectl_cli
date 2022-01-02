// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package client

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

// KubectlClientClient is the client API for KubectlClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubectlClientClient interface {
	// Sends a greeting
	GetPods(ctx context.Context, in *GetPodsRequest, opts ...grpc.CallOption) (*GetPodsResponse, error)
}

type kubectlClientClient struct {
	cc grpc.ClientConnInterface
}

func NewKubectlClientClient(cc grpc.ClientConnInterface) KubectlClientClient {
	return &kubectlClientClient{cc}
}

func (c *kubectlClientClient) GetPods(ctx context.Context, in *GetPodsRequest, opts ...grpc.CallOption) (*GetPodsResponse, error) {
	out := new(GetPodsResponse)
	err := c.cc.Invoke(ctx, "/KubectlClient/GetPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubectlClientServer is the server API for KubectlClient service.
// All implementations must embed UnimplementedKubectlClientServer
// for forward compatibility
type KubectlClientServer interface {
	// Sends a greeting
	GetPods(context.Context, *GetPodsRequest) (*GetPodsResponse, error)
	mustEmbedUnimplementedKubectlClientServer()
}

// UnimplementedKubectlClientServer must be embedded to have forward compatible implementations.
type UnimplementedKubectlClientServer struct {
}

func (UnimplementedKubectlClientServer) GetPods(context.Context, *GetPodsRequest) (*GetPodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPods not implemented")
}
func (UnimplementedKubectlClientServer) mustEmbedUnimplementedKubectlClientServer() {}

// UnsafeKubectlClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubectlClientServer will
// result in compilation errors.
type UnsafeKubectlClientServer interface {
	mustEmbedUnimplementedKubectlClientServer()
}

func RegisterKubectlClientServer(s grpc.ServiceRegistrar, srv KubectlClientServer) {
	s.RegisterService(&KubectlClient_ServiceDesc, srv)
}

func _KubectlClient_GetPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubectlClientServer).GetPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KubectlClient/GetPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubectlClientServer).GetPods(ctx, req.(*GetPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KubectlClient_ServiceDesc is the grpc.ServiceDesc for KubectlClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubectlClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KubectlClient",
	HandlerType: (*KubectlClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPods",
			Handler:    _KubectlClient_GetPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}