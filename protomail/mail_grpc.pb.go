// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protomail

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

// MailClient is the client API for Mail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailClient interface {
	ConnectClient(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponce, error)
	LoginClient(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Empty, error)
	LogoutClient(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteClient(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
	ListMailboxes(ctx context.Context, in *ListMailboxesRequest, opts ...grpc.CallOption) (Mail_ListMailboxesClient, error)
	SelectMailbox(ctx context.Context, in *SelectMailboxRequest, opts ...grpc.CallOption) (*Empty, error)
}

type mailClient struct {
	cc grpc.ClientConnInterface
}

func NewMailClient(cc grpc.ClientConnInterface) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) ConnectClient(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponce, error) {
	out := new(ConnectResponce)
	err := c.cc.Invoke(ctx, "/protomail.Mail/ConnectClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailClient) LoginClient(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protomail.Mail/LoginClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailClient) LogoutClient(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protomail.Mail/LogoutClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailClient) DeleteClient(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protomail.Mail/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailClient) ListMailboxes(ctx context.Context, in *ListMailboxesRequest, opts ...grpc.CallOption) (Mail_ListMailboxesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Mail_ServiceDesc.Streams[0], "/protomail.Mail/ListMailboxes", opts...)
	if err != nil {
		return nil, err
	}
	x := &mailListMailboxesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mail_ListMailboxesClient interface {
	Recv() (*ListMailboxesResponce, error)
	grpc.ClientStream
}

type mailListMailboxesClient struct {
	grpc.ClientStream
}

func (x *mailListMailboxesClient) Recv() (*ListMailboxesResponce, error) {
	m := new(ListMailboxesResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mailClient) SelectMailbox(ctx context.Context, in *SelectMailboxRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protomail.Mail/SelectMailbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServer is the server API for Mail service.
// All implementations must embed UnimplementedMailServer
// for forward compatibility
type MailServer interface {
	ConnectClient(context.Context, *ConnectRequest) (*ConnectResponce, error)
	LoginClient(context.Context, *LoginRequest) (*Empty, error)
	LogoutClient(context.Context, *LogoutRequest) (*Empty, error)
	DeleteClient(context.Context, *DeleteRequest) (*Empty, error)
	ListMailboxes(*ListMailboxesRequest, Mail_ListMailboxesServer) error
	SelectMailbox(context.Context, *SelectMailboxRequest) (*Empty, error)
	mustEmbedUnimplementedMailServer()
}

// UnimplementedMailServer must be embedded to have forward compatible implementations.
type UnimplementedMailServer struct {
}

func (UnimplementedMailServer) ConnectClient(context.Context, *ConnectRequest) (*ConnectResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectClient not implemented")
}
func (UnimplementedMailServer) LoginClient(context.Context, *LoginRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginClient not implemented")
}
func (UnimplementedMailServer) LogoutClient(context.Context, *LogoutRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutClient not implemented")
}
func (UnimplementedMailServer) DeleteClient(context.Context, *DeleteRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedMailServer) ListMailboxes(*ListMailboxesRequest, Mail_ListMailboxesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMailboxes not implemented")
}
func (UnimplementedMailServer) SelectMailbox(context.Context, *SelectMailboxRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectMailbox not implemented")
}
func (UnimplementedMailServer) mustEmbedUnimplementedMailServer() {}

// UnsafeMailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailServer will
// result in compilation errors.
type UnsafeMailServer interface {
	mustEmbedUnimplementedMailServer()
}

func RegisterMailServer(s grpc.ServiceRegistrar, srv MailServer) {
	s.RegisterService(&Mail_ServiceDesc, srv)
}

func _Mail_ConnectClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).ConnectClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomail.Mail/ConnectClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).ConnectClient(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mail_LoginClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).LoginClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomail.Mail/LoginClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).LoginClient(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mail_LogoutClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).LogoutClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomail.Mail/LogoutClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).LogoutClient(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mail_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomail.Mail/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).DeleteClient(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mail_ListMailboxes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMailboxesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MailServer).ListMailboxes(m, &mailListMailboxesServer{stream})
}

type Mail_ListMailboxesServer interface {
	Send(*ListMailboxesResponce) error
	grpc.ServerStream
}

type mailListMailboxesServer struct {
	grpc.ServerStream
}

func (x *mailListMailboxesServer) Send(m *ListMailboxesResponce) error {
	return x.ServerStream.SendMsg(m)
}

func _Mail_SelectMailbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectMailboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).SelectMailbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protomail.Mail/SelectMailbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).SelectMailbox(ctx, req.(*SelectMailboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mail_ServiceDesc is the grpc.ServiceDesc for Mail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protomail.Mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectClient",
			Handler:    _Mail_ConnectClient_Handler,
		},
		{
			MethodName: "LoginClient",
			Handler:    _Mail_LoginClient_Handler,
		},
		{
			MethodName: "LogoutClient",
			Handler:    _Mail_LogoutClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Mail_DeleteClient_Handler,
		},
		{
			MethodName: "SelectMailbox",
			Handler:    _Mail_SelectMailbox_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMailboxes",
			Handler:       _Mail_ListMailboxes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mail.proto",
}
