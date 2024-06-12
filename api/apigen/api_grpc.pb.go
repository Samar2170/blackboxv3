// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: api.proto

package blackboxv3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BlackBoxService_Signup_FullMethodName           = "/blackboxv3.BlackBoxService/Signup"
	BlackBoxService_Login_FullMethodName            = "/blackboxv3.BlackBoxService/Login"
	BlackBoxService_CreateChannel_FullMethodName    = "/blackboxv3.BlackBoxService/CreateChannel"
	BlackBoxService_GetChannels_FullMethodName      = "/blackboxv3.BlackBoxService/GetChannels"
	BlackBoxService_GetChannel_FullMethodName       = "/blackboxv3.BlackBoxService/GetChannel"
	BlackBoxService_DeleteChannel_FullMethodName    = "/blackboxv3.BlackBoxService/DeleteChannel"
	BlackBoxService_SendMessage_FullMethodName      = "/blackboxv3.BlackBoxService/SendMessage"
	BlackBoxService_GetMessages_FullMethodName      = "/blackboxv3.BlackBoxService/GetMessages"
	BlackBoxService_SendMediaMessage_FullMethodName = "/blackboxv3.BlackBoxService/SendMediaMessage"
)

// BlackBoxServiceClient is the client API for BlackBoxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlackBoxServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	GetChannels(ctx context.Context, in *GetChannelsRequest, opts ...grpc.CallOption) (*GetChannelsResponse, error)
	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	SendMediaMessage(ctx context.Context, opts ...grpc.CallOption) (BlackBoxService_SendMediaMessageClient, error)
}

type blackBoxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlackBoxServiceClient(cc grpc.ClientConnInterface) BlackBoxServiceClient {
	return &blackBoxServiceClient{cc}
}

func (c *blackBoxServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_Signup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_CreateChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) GetChannels(ctx context.Context, in *GetChannelsRequest, opts ...grpc.CallOption) (*GetChannelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChannelsResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_GetChannels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChannelResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_GetChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_DeleteChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, BlackBoxService_GetMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackBoxServiceClient) SendMediaMessage(ctx context.Context, opts ...grpc.CallOption) (BlackBoxService_SendMediaMessageClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlackBoxService_ServiceDesc.Streams[0], BlackBoxService_SendMediaMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &blackBoxServiceSendMediaMessageClient{ClientStream: stream}
	return x, nil
}

type BlackBoxService_SendMediaMessageClient interface {
	Send(*SendMediaMessageRequest) error
	CloseAndRecv() (*SendMediaMessageResponse, error)
	grpc.ClientStream
}

type blackBoxServiceSendMediaMessageClient struct {
	grpc.ClientStream
}

func (x *blackBoxServiceSendMediaMessageClient) Send(m *SendMediaMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blackBoxServiceSendMediaMessageClient) CloseAndRecv() (*SendMediaMessageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendMediaMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlackBoxServiceServer is the server API for BlackBoxService service.
// All implementations must embed UnimplementedBlackBoxServiceServer
// for forward compatibility
type BlackBoxServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	GetChannels(context.Context, *GetChannelsRequest) (*GetChannelsResponse, error)
	GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	SendMediaMessage(BlackBoxService_SendMediaMessageServer) error
	mustEmbedUnimplementedBlackBoxServiceServer()
}

// UnimplementedBlackBoxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlackBoxServiceServer struct {
}

func (UnimplementedBlackBoxServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedBlackBoxServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBlackBoxServiceServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedBlackBoxServiceServer) GetChannels(context.Context, *GetChannelsRequest) (*GetChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannels not implemented")
}
func (UnimplementedBlackBoxServiceServer) GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}
func (UnimplementedBlackBoxServiceServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedBlackBoxServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedBlackBoxServiceServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedBlackBoxServiceServer) SendMediaMessage(BlackBoxService_SendMediaMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMediaMessage not implemented")
}
func (UnimplementedBlackBoxServiceServer) mustEmbedUnimplementedBlackBoxServiceServer() {}

// UnsafeBlackBoxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlackBoxServiceServer will
// result in compilation errors.
type UnsafeBlackBoxServiceServer interface {
	mustEmbedUnimplementedBlackBoxServiceServer()
}

func RegisterBlackBoxServiceServer(s grpc.ServiceRegistrar, srv BlackBoxServiceServer) {
	s.RegisterService(&BlackBoxService_ServiceDesc, srv)
}

func _BlackBoxService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_CreateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_GetChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).GetChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_GetChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).GetChannels(ctx, req.(*GetChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_GetChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).GetChannel(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_DeleteChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackBoxServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlackBoxService_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackBoxServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackBoxService_SendMediaMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlackBoxServiceServer).SendMediaMessage(&blackBoxServiceSendMediaMessageServer{ServerStream: stream})
}

type BlackBoxService_SendMediaMessageServer interface {
	SendAndClose(*SendMediaMessageResponse) error
	Recv() (*SendMediaMessageRequest, error)
	grpc.ServerStream
}

type blackBoxServiceSendMediaMessageServer struct {
	grpc.ServerStream
}

func (x *blackBoxServiceSendMediaMessageServer) SendAndClose(m *SendMediaMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blackBoxServiceSendMediaMessageServer) Recv() (*SendMediaMessageRequest, error) {
	m := new(SendMediaMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlackBoxService_ServiceDesc is the grpc.ServiceDesc for BlackBoxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlackBoxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blackboxv3.BlackBoxService",
	HandlerType: (*BlackBoxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _BlackBoxService_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _BlackBoxService_Login_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _BlackBoxService_CreateChannel_Handler,
		},
		{
			MethodName: "GetChannels",
			Handler:    _BlackBoxService_GetChannels_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _BlackBoxService_GetChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _BlackBoxService_DeleteChannel_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _BlackBoxService_SendMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _BlackBoxService_GetMessages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMediaMessage",
			Handler:       _BlackBoxService_SendMediaMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
