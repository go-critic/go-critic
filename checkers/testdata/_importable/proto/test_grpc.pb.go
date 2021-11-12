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

// TestingClient is the client API for Testing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestingClient interface {
	Call(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
}

type testingClient struct {
	cc grpc.ClientConnInterface
}

func NewTestingClient(cc grpc.ClientConnInterface) TestingClient {
	return &testingClient{cc}
}

func (c *testingClient) Call(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/Testing/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestingServer is the server API for Testing service.
// All implementations must embed UnimplementedTestingServer
// for forward compatibility
type TestingServer interface {
	Call(context.Context, *Test) (*Test, error)
	mustEmbedUnimplementedTestingServer()
}

// UnimplementedTestingServer must be embedded to have forward compatible implementations.
type UnimplementedTestingServer struct {
}

func (UnimplementedTestingServer) Call(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedTestingServer) mustEmbedUnimplementedTestingServer() {}

// UnsafeTestingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestingServer will
// result in compilation errors.
type UnsafeTestingServer interface {
	mustEmbedUnimplementedTestingServer()
}

func RegisterTestingServer(s grpc.ServiceRegistrar, srv TestingServer) {
	s.RegisterService(&Testing_ServiceDesc, srv)
}

func _Testing_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Testing/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServer).Call(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

// Testing_ServiceDesc is the grpc.ServiceDesc for Testing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Testing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Testing",
	HandlerType: (*TestingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Testing_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
