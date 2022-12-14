// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/email_user/email_user.proto

package email_user

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

// SendEmailClient is the client API for SendEmail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendEmailClient interface {
	SendEmail(ctx context.Context, in *EmailMessage, opts ...grpc.CallOption) (*EmailResponse, error)
}

type sendEmailClient struct {
	cc grpc.ClientConnInterface
}

func NewSendEmailClient(cc grpc.ClientConnInterface) SendEmailClient {
	return &sendEmailClient{cc}
}

func (c *sendEmailClient) SendEmail(ctx context.Context, in *EmailMessage, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/SendEmail/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendEmailServer is the server API for SendEmail service.
// All implementations must embed UnimplementedSendEmailServer
// for forward compatibility
type SendEmailServer interface {
	SendEmail(context.Context, *EmailMessage) (*EmailResponse, error)
	mustEmbedUnimplementedSendEmailServer()
}

// UnimplementedSendEmailServer must be embedded to have forward compatible implementations.
type UnimplementedSendEmailServer struct {
}

func (UnimplementedSendEmailServer) SendEmail(context.Context, *EmailMessage) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedSendEmailServer) mustEmbedUnimplementedSendEmailServer() {}

// UnsafeSendEmailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendEmailServer will
// result in compilation errors.
type UnsafeSendEmailServer interface {
	mustEmbedUnimplementedSendEmailServer()
}

func RegisterSendEmailServer(s grpc.ServiceRegistrar, srv SendEmailServer) {
	s.RegisterService(&SendEmail_ServiceDesc, srv)
}

func _SendEmail_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendEmailServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendEmail/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendEmailServer).SendEmail(ctx, req.(*EmailMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SendEmail_ServiceDesc is the grpc.ServiceDesc for SendEmail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendEmail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SendEmail",
	HandlerType: (*SendEmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _SendEmail_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/email_user/email_user.proto",
}

// GetAuthenticatedUserClient is the client API for GetAuthenticatedUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetAuthenticatedUserClient interface {
	GetAuthenticatedUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EmailUser, error)
}

type getAuthenticatedUserClient struct {
	cc grpc.ClientConnInterface
}

func NewGetAuthenticatedUserClient(cc grpc.ClientConnInterface) GetAuthenticatedUserClient {
	return &getAuthenticatedUserClient{cc}
}

func (c *getAuthenticatedUserClient) GetAuthenticatedUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EmailUser, error) {
	out := new(EmailUser)
	err := c.cc.Invoke(ctx, "/GetAuthenticatedUser/GetAuthenticatedUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetAuthenticatedUserServer is the server API for GetAuthenticatedUser service.
// All implementations must embed UnimplementedGetAuthenticatedUserServer
// for forward compatibility
type GetAuthenticatedUserServer interface {
	GetAuthenticatedUser(context.Context, *emptypb.Empty) (*EmailUser, error)
	mustEmbedUnimplementedGetAuthenticatedUserServer()
}

// UnimplementedGetAuthenticatedUserServer must be embedded to have forward compatible implementations.
type UnimplementedGetAuthenticatedUserServer struct {
}

func (UnimplementedGetAuthenticatedUserServer) GetAuthenticatedUser(context.Context, *emptypb.Empty) (*EmailUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthenticatedUser not implemented")
}
func (UnimplementedGetAuthenticatedUserServer) mustEmbedUnimplementedGetAuthenticatedUserServer() {}

// UnsafeGetAuthenticatedUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetAuthenticatedUserServer will
// result in compilation errors.
type UnsafeGetAuthenticatedUserServer interface {
	mustEmbedUnimplementedGetAuthenticatedUserServer()
}

func RegisterGetAuthenticatedUserServer(s grpc.ServiceRegistrar, srv GetAuthenticatedUserServer) {
	s.RegisterService(&GetAuthenticatedUser_ServiceDesc, srv)
}

func _GetAuthenticatedUser_GetAuthenticatedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetAuthenticatedUserServer).GetAuthenticatedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetAuthenticatedUser/GetAuthenticatedUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetAuthenticatedUserServer).GetAuthenticatedUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GetAuthenticatedUser_ServiceDesc is the grpc.ServiceDesc for GetAuthenticatedUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetAuthenticatedUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetAuthenticatedUser",
	HandlerType: (*GetAuthenticatedUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthenticatedUser",
			Handler:    _GetAuthenticatedUser_GetAuthenticatedUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/email_user/email_user.proto",
}
