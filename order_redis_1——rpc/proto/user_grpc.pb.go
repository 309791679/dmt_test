// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: user.proto

package grpc_

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

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	//创建
	CreateUser(ctx context.Context, in *UniversalUserRequest, opts ...grpc.CallOption) (*UniversalResponse, error)
	//修改密码
	ChangeUser(ctx context.Context, in *UniversalUserRequest, opts ...grpc.CallOption) (*UniversalResponse, error)
	//查询
	QueryUser(ctx context.Context, in *UniversalUserRequest, opts ...grpc.CallOption) (*UniversalResponse, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) CreateUser(ctx context.Context, in *UniversalUserRequest, opts ...grpc.CallOption) (*UniversalResponse, error) {
	out := new(UniversalResponse)
	err := c.cc.Invoke(ctx, "/usercenter/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) ChangeUser(ctx context.Context, in *UniversalUserRequest, opts ...grpc.CallOption) (*UniversalResponse, error) {
	out := new(UniversalResponse)
	err := c.cc.Invoke(ctx, "/usercenter/ChangeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) QueryUser(ctx context.Context, in *UniversalUserRequest, opts ...grpc.CallOption) (*UniversalResponse, error) {
	out := new(UniversalResponse)
	err := c.cc.Invoke(ctx, "/usercenter/QueryUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	//创建
	CreateUser(context.Context, *UniversalUserRequest) (*UniversalResponse, error)
	//修改密码
	ChangeUser(context.Context, *UniversalUserRequest) (*UniversalResponse, error)
	//查询
	QueryUser(context.Context, *UniversalUserRequest) (*UniversalResponse, error)

}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) CreateUser(context.Context, *UniversalUserRequest) (*UniversalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsercenterServer) ChangeUser(context.Context, *UniversalUserRequest) (*UniversalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUser not implemented")
}
func (UnimplementedUsercenterServer) QueryUser(context.Context, *UniversalUserRequest) (*UniversalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}


func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UniversalUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).CreateUser(ctx, req.(*UniversalUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_ChangeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UniversalUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).ChangeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter/ChangeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).ChangeUser(ctx, req.(*UniversalUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UniversalUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter/QueryUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).QueryUser(ctx, req.(*UniversalUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Usercenter_CreateUser_Handler,
		},
		{
			MethodName: "ChangeUser",
			Handler:    _Usercenter_ChangeUser_Handler,
		},
		{
			MethodName: "QueryUser",
			Handler:    _Usercenter_QueryUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
