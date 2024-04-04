// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: rootadminweb.proto

package __

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

// RootAdminWebServiceClient is the client API for RootAdminWebService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RootAdminWebServiceClient interface {
	DoLogin(ctx context.Context, in *DoLoginRequest, opts ...grpc.CallOption) (*DoLoginResponse, error)
	DoLogout(ctx context.Context, in *DoLogoutRequest, opts ...grpc.CallOption) (*DoLogoutResponse, error)
}

type rootAdminWebServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRootAdminWebServiceClient(cc grpc.ClientConnInterface) RootAdminWebServiceClient {
	return &rootAdminWebServiceClient{cc}
}

func (c *rootAdminWebServiceClient) DoLogin(ctx context.Context, in *DoLoginRequest, opts ...grpc.CallOption) (*DoLoginResponse, error) {
	out := new(DoLoginResponse)
	err := c.cc.Invoke(ctx, "/rootadminweb.RootAdminWebService/DoLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootAdminWebServiceClient) DoLogout(ctx context.Context, in *DoLogoutRequest, opts ...grpc.CallOption) (*DoLogoutResponse, error) {
	out := new(DoLogoutResponse)
	err := c.cc.Invoke(ctx, "/rootadminweb.RootAdminWebService/DoLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RootAdminWebServiceServer is the server API for RootAdminWebService service.
// All implementations must embed UnimplementedRootAdminWebServiceServer
// for forward compatibility
type RootAdminWebServiceServer interface {
	DoLogin(context.Context, *DoLoginRequest) (*DoLoginResponse, error)
	DoLogout(context.Context, *DoLogoutRequest) (*DoLogoutResponse, error)
	mustEmbedUnimplementedRootAdminWebServiceServer()
}

// UnimplementedRootAdminWebServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRootAdminWebServiceServer struct {
}

func (UnimplementedRootAdminWebServiceServer) DoLogin(context.Context, *DoLoginRequest) (*DoLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLogin not implemented")
}
func (UnimplementedRootAdminWebServiceServer) DoLogout(context.Context, *DoLogoutRequest) (*DoLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLogout not implemented")
}
func (UnimplementedRootAdminWebServiceServer) mustEmbedUnimplementedRootAdminWebServiceServer() {}

// UnsafeRootAdminWebServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RootAdminWebServiceServer will
// result in compilation errors.
type UnsafeRootAdminWebServiceServer interface {
	mustEmbedUnimplementedRootAdminWebServiceServer()
}

func RegisterRootAdminWebServiceServer(s grpc.ServiceRegistrar, srv RootAdminWebServiceServer) {
	s.RegisterService(&RootAdminWebService_ServiceDesc, srv)
}

func _RootAdminWebService_DoLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootAdminWebServiceServer).DoLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rootadminweb.RootAdminWebService/DoLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootAdminWebServiceServer).DoLogin(ctx, req.(*DoLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootAdminWebService_DoLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootAdminWebServiceServer).DoLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rootadminweb.RootAdminWebService/DoLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootAdminWebServiceServer).DoLogout(ctx, req.(*DoLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RootAdminWebService_ServiceDesc is the grpc.ServiceDesc for RootAdminWebService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RootAdminWebService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rootadminweb.RootAdminWebService",
	HandlerType: (*RootAdminWebServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoLogin",
			Handler:    _RootAdminWebService_DoLogin_Handler,
		},
		{
			MethodName: "DoLogout",
			Handler:    _RootAdminWebService_DoLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rootadminweb.proto",
}