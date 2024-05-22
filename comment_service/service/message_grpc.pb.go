// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: message.proto

package service

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

const (
	MessageService_CreateCommentMessageTwoStream_FullMethodName = "/message.MessageService/CreateCommentMessageTwoStream"
	MessageService_GetCommentTwoStream_FullMethodName           = "/message.MessageService/GetCommentTwoStream"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateCommentMessageTwoStream(ctx context.Context, opts ...grpc.CallOption) (MessageService_CreateCommentMessageTwoStreamClient, error)
	GetCommentTwoStream(ctx context.Context, opts ...grpc.CallOption) (MessageService_GetCommentTwoStreamClient, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateCommentMessageTwoStream(ctx context.Context, opts ...grpc.CallOption) (MessageService_CreateCommentMessageTwoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageService_ServiceDesc.Streams[0], MessageService_CreateCommentMessageTwoStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceCreateCommentMessageTwoStreamClient{stream}
	return x, nil
}

type MessageService_CreateCommentMessageTwoStreamClient interface {
	Send(*CreateMessageRequest) error
	Recv() (*CreateMessageResponse, error)
	grpc.ClientStream
}

type messageServiceCreateCommentMessageTwoStreamClient struct {
	grpc.ClientStream
}

func (x *messageServiceCreateCommentMessageTwoStreamClient) Send(m *CreateMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceCreateCommentMessageTwoStreamClient) Recv() (*CreateMessageResponse, error) {
	m := new(CreateMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) GetCommentTwoStream(ctx context.Context, opts ...grpc.CallOption) (MessageService_GetCommentTwoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageService_ServiceDesc.Streams[1], MessageService_GetCommentTwoStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceGetCommentTwoStreamClient{stream}
	return x, nil
}

type MessageService_GetCommentTwoStreamClient interface {
	Send(*GetCommentRequest) error
	Recv() (*GetCommentResponse, error)
	grpc.ClientStream
}

type messageServiceGetCommentTwoStreamClient struct {
	grpc.ClientStream
}

func (x *messageServiceGetCommentTwoStreamClient) Send(m *GetCommentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceGetCommentTwoStreamClient) Recv() (*GetCommentResponse, error) {
	m := new(GetCommentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	CreateCommentMessageTwoStream(MessageService_CreateCommentMessageTwoStreamServer) error
	GetCommentTwoStream(MessageService_GetCommentTwoStreamServer) error
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) CreateCommentMessageTwoStream(MessageService_CreateCommentMessageTwoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCommentMessageTwoStream not implemented")
}
func (UnimplementedMessageServiceServer) GetCommentTwoStream(MessageService_GetCommentTwoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommentTwoStream not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_CreateCommentMessageTwoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).CreateCommentMessageTwoStream(&messageServiceCreateCommentMessageTwoStreamServer{stream})
}

type MessageService_CreateCommentMessageTwoStreamServer interface {
	Send(*CreateMessageResponse) error
	Recv() (*CreateMessageRequest, error)
	grpc.ServerStream
}

type messageServiceCreateCommentMessageTwoStreamServer struct {
	grpc.ServerStream
}

func (x *messageServiceCreateCommentMessageTwoStreamServer) Send(m *CreateMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceCreateCommentMessageTwoStreamServer) Recv() (*CreateMessageRequest, error) {
	m := new(CreateMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageService_GetCommentTwoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).GetCommentTwoStream(&messageServiceGetCommentTwoStreamServer{stream})
}

type MessageService_GetCommentTwoStreamServer interface {
	Send(*GetCommentResponse) error
	Recv() (*GetCommentRequest, error)
	grpc.ServerStream
}

type messageServiceGetCommentTwoStreamServer struct {
	grpc.ServerStream
}

func (x *messageServiceGetCommentTwoStreamServer) Send(m *GetCommentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceGetCommentTwoStreamServer) Recv() (*GetCommentRequest, error) {
	m := new(GetCommentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCommentMessageTwoStream",
			Handler:       _MessageService_CreateCommentMessageTwoStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetCommentTwoStream",
			Handler:       _MessageService_GetCommentTwoStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
