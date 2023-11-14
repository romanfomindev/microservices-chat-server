// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: access.proto

package access_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccessServiceClient is the client API for AccessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessServiceClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type accessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessServiceClient(cc grpc.ClientConnInterface) AccessServiceClient {
	return &accessServiceClient{cc}
}

func (c *accessServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/access_v1.AccessService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessServiceServer is the server API for AccessService service.
// All implementations must embed UnimplementedAccessServiceServer
// for forward compatibility
type AccessServiceServer interface {
	Check(context.Context, *CheckRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAccessServiceServer()
}

// UnimplementedAccessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccessServiceServer struct {
}

func (UnimplementedAccessServiceServer) Check(context.Context, *CheckRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedAccessServiceServer) mustEmbedUnimplementedAccessServiceServer() {}

// UnsafeAccessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessServiceServer will
// result in compilation errors.
type UnsafeAccessServiceServer interface {
	mustEmbedUnimplementedAccessServiceServer()
}

func RegisterAccessServiceServer(s grpc.ServiceRegistrar, srv AccessServiceServer) {
	s.RegisterService(&AccessService_ServiceDesc, srv)
}

func _AccessService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access_v1.AccessService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessService_ServiceDesc is the grpc.ServiceDesc for AccessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "access_v1.AccessService",
	HandlerType: (*AccessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _AccessService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access.proto",
}
