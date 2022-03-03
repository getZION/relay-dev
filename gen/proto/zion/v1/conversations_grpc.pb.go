// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConversationServiceClient is the client API for ConversationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversationServiceClient interface {
	// Create a conversation
	CreateConversation(ctx context.Context, in *CreateConversationRequest, opts ...grpc.CallOption) (*CreateConversationResponse, error)
	// Delete a comment of mine, or in a conversation I started
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	// Delete a conversation I started
	DeleteConversation(ctx context.Context, in *DeleteConversationRequest, opts ...grpc.CallOption) (*DeleteConversationResponse, error)
	// Get list of my activity (notifications)
	GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityResponse, error)
	// Get a conversation and its comments
	GetConversation(ctx context.Context, in *GetConversationRequest, opts ...grpc.CallOption) (*GetConversationResponse, error)
	// Get my feed of conversations
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error)
}

type conversationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversationServiceClient(cc grpc.ClientConnInterface) ConversationServiceClient {
	return &conversationServiceClient{cc}
}

func (c *conversationServiceClient) CreateConversation(ctx context.Context, in *CreateConversationRequest, opts ...grpc.CallOption) (*CreateConversationResponse, error) {
	out := new(CreateConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.ConversationService/CreateConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.ConversationService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) DeleteConversation(ctx context.Context, in *DeleteConversationRequest, opts ...grpc.CallOption) (*DeleteConversationResponse, error) {
	out := new(DeleteConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.ConversationService/DeleteConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityResponse, error) {
	out := new(GetActivityResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.ConversationService/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) GetConversation(ctx context.Context, in *GetConversationRequest, opts ...grpc.CallOption) (*GetConversationResponse, error) {
	out := new(GetConversationResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.ConversationService/GetConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error) {
	out := new(GetFeedResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.ConversationService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServiceServer is the server API for ConversationService service.
// All implementations should embed UnimplementedConversationServiceServer
// for forward compatibility
type ConversationServiceServer interface {
	// Create a conversation
	CreateConversation(context.Context, *CreateConversationRequest) (*CreateConversationResponse, error)
	// Delete a comment of mine, or in a conversation I started
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	// Delete a conversation I started
	DeleteConversation(context.Context, *DeleteConversationRequest) (*DeleteConversationResponse, error)
	// Get list of my activity (notifications)
	GetActivity(context.Context, *GetActivityRequest) (*GetActivityResponse, error)
	// Get a conversation and its comments
	GetConversation(context.Context, *GetConversationRequest) (*GetConversationResponse, error)
	// Get my feed of conversations
	GetFeed(context.Context, *GetFeedRequest) (*GetFeedResponse, error)
}

// UnimplementedConversationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConversationServiceServer struct {
}

func (UnimplementedConversationServiceServer) CreateConversation(context.Context, *CreateConversationRequest) (*CreateConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConversation not implemented")
}
func (UnimplementedConversationServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedConversationServiceServer) DeleteConversation(context.Context, *DeleteConversationRequest) (*DeleteConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConversation not implemented")
}
func (UnimplementedConversationServiceServer) GetActivity(context.Context, *GetActivityRequest) (*GetActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedConversationServiceServer) GetConversation(context.Context, *GetConversationRequest) (*GetConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversation not implemented")
}
func (UnimplementedConversationServiceServer) GetFeed(context.Context, *GetFeedRequest) (*GetFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}

// UnsafeConversationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversationServiceServer will
// result in compilation errors.
type UnsafeConversationServiceServer interface {
	mustEmbedUnimplementedConversationServiceServer()
}

func RegisterConversationServiceServer(s grpc.ServiceRegistrar, srv ConversationServiceServer) {
	s.RegisterService(&_ConversationService_serviceDesc, srv)
}

func _ConversationService_CreateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).CreateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.ConversationService/CreateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).CreateConversation(ctx, req.(*CreateConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.ConversationService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_DeleteConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).DeleteConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.ConversationService/DeleteConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).DeleteConversation(ctx, req.(*DeleteConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.ConversationService/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).GetActivity(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_GetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).GetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.ConversationService/GetConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).GetConversation(ctx, req.(*GetConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.ConversationService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.zion.v1.ConversationService",
	HandlerType: (*ConversationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConversation",
			Handler:    _ConversationService_CreateConversation_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _ConversationService_DeleteComment_Handler,
		},
		{
			MethodName: "DeleteConversation",
			Handler:    _ConversationService_DeleteConversation_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _ConversationService_GetActivity_Handler,
		},
		{
			MethodName: "GetConversation",
			Handler:    _ConversationService_GetConversation_Handler,
		},
		{
			MethodName: "GetFeed",
			Handler:    _ConversationService_GetFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/zion/v1/conversations.proto",
}
