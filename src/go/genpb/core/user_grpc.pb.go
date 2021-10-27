// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package core

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindBy(ctx context.Context, in *FindParam, opts ...grpc.CallOption) (*UserEither, error)
	Create(ctx context.Context, in *CreateParam, opts ...grpc.CallOption) (*UserEither, error)
	Update(ctx context.Context, in *UpdateParam, opts ...grpc.CallOption) (*UserEither, error)
	ComparePassword(ctx context.Context, in *AuthenticateParam, opts ...grpc.CallOption) (*UserEither, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindBy(ctx context.Context, in *FindParam, opts ...grpc.CallOption) (*UserEither, error) {
	out := new(UserEither)
	err := c.cc.Invoke(ctx, "/core.userService/find_by", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateParam, opts ...grpc.CallOption) (*UserEither, error) {
	out := new(UserEither)
	err := c.cc.Invoke(ctx, "/core.userService/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateParam, opts ...grpc.CallOption) (*UserEither, error) {
	out := new(UserEither)
	err := c.cc.Invoke(ctx, "/core.userService/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ComparePassword(ctx context.Context, in *AuthenticateParam, opts ...grpc.CallOption) (*UserEither, error) {
	out := new(UserEither)
	err := c.cc.Invoke(ctx, "/core.userService/compare_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FindBy(context.Context, *FindParam) (*UserEither, error)
	Create(context.Context, *CreateParam) (*UserEither, error)
	Update(context.Context, *UpdateParam) (*UserEither, error)
	ComparePassword(context.Context, *AuthenticateParam) (*UserEither, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindBy(context.Context, *FindParam) (*UserEither, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBy not implemented")
}
func (UnimplementedUserServiceServer) Create(context.Context, *CreateParam) (*UserEither, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateParam) (*UserEither, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) ComparePassword(context.Context, *AuthenticateParam) (*UserEither, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComparePassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.userService/find_by",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindBy(ctx, req.(*FindParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.userService/create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.userService/update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ComparePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ComparePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.userService/compare_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ComparePassword(ctx, req.(*AuthenticateParam))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.userService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "find_by",
			Handler:    _UserService_FindBy_Handler,
		},
		{
			MethodName: "create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "compare_password",
			Handler:    _UserService_ComparePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/user.proto",
}
