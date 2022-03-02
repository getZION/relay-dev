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

// CommunitiesServiceClient is the client API for CommunitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunitiesServiceClient interface {
	// Create a community
	CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityResponse, error)
	// Delete a community you own
	DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityResponse, error)
	// Edit a community you own
	EditCommunity(ctx context.Context, in *EditCommunityRequest, opts ...grpc.CallOption) (*EditCommunityResponse, error)
	// Join a specific community
	JoinCommunity(ctx context.Context, in *JoinCommunityRequest, opts ...grpc.CallOption) (*JoinCommunityResponse, error)
	// List communities
	ListCommunity(ctx context.Context, in *ListCommunityRequest, opts ...grpc.CallOption) (*ListCommunityResponse, error)
}

type communitiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunitiesServiceClient(cc grpc.ClientConnInterface) CommunitiesServiceClient {
	return &communitiesServiceClient{cc}
}

func (c *communitiesServiceClient) CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityResponse, error) {
	out := new(CreateCommunityResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.CommunitiesService/CreateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communitiesServiceClient) DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityResponse, error) {
	out := new(DeleteCommunityResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.CommunitiesService/DeleteCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communitiesServiceClient) EditCommunity(ctx context.Context, in *EditCommunityRequest, opts ...grpc.CallOption) (*EditCommunityResponse, error) {
	out := new(EditCommunityResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.CommunitiesService/EditCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communitiesServiceClient) JoinCommunity(ctx context.Context, in *JoinCommunityRequest, opts ...grpc.CallOption) (*JoinCommunityResponse, error) {
	out := new(JoinCommunityResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.CommunitiesService/JoinCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communitiesServiceClient) ListCommunity(ctx context.Context, in *ListCommunityRequest, opts ...grpc.CallOption) (*ListCommunityResponse, error) {
	out := new(ListCommunityResponse)
	err := c.cc.Invoke(ctx, "/proto.zion.v1.CommunitiesService/ListCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunitiesServiceServer is the server API for CommunitiesService service.
// All implementations should embed UnimplementedCommunitiesServiceServer
// for forward compatibility
type CommunitiesServiceServer interface {
	// Create a community
	CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityResponse, error)
	// Delete a community you own
	DeleteCommunity(context.Context, *DeleteCommunityRequest) (*DeleteCommunityResponse, error)
	// Edit a community you own
	EditCommunity(context.Context, *EditCommunityRequest) (*EditCommunityResponse, error)
	// Join a specific community
	JoinCommunity(context.Context, *JoinCommunityRequest) (*JoinCommunityResponse, error)
	// List communities
	ListCommunity(context.Context, *ListCommunityRequest) (*ListCommunityResponse, error)
}

// UnimplementedCommunitiesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCommunitiesServiceServer struct {
}

func (UnimplementedCommunitiesServiceServer) CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunity not implemented")
}
func (UnimplementedCommunitiesServiceServer) DeleteCommunity(context.Context, *DeleteCommunityRequest) (*DeleteCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunity not implemented")
}
func (UnimplementedCommunitiesServiceServer) EditCommunity(context.Context, *EditCommunityRequest) (*EditCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCommunity not implemented")
}
func (UnimplementedCommunitiesServiceServer) JoinCommunity(context.Context, *JoinCommunityRequest) (*JoinCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCommunity not implemented")
}
func (UnimplementedCommunitiesServiceServer) ListCommunity(context.Context, *ListCommunityRequest) (*ListCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommunity not implemented")
}

// UnsafeCommunitiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunitiesServiceServer will
// result in compilation errors.
type UnsafeCommunitiesServiceServer interface {
	mustEmbedUnimplementedCommunitiesServiceServer()
}

func RegisterCommunitiesServiceServer(s grpc.ServiceRegistrar, srv CommunitiesServiceServer) {
	s.RegisterService(&_CommunitiesService_serviceDesc, srv)
}

func _CommunitiesService_CreateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunitiesServiceServer).CreateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.CommunitiesService/CreateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunitiesServiceServer).CreateCommunity(ctx, req.(*CreateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunitiesService_DeleteCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunitiesServiceServer).DeleteCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.CommunitiesService/DeleteCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunitiesServiceServer).DeleteCommunity(ctx, req.(*DeleteCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunitiesService_EditCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunitiesServiceServer).EditCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.CommunitiesService/EditCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunitiesServiceServer).EditCommunity(ctx, req.(*EditCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunitiesService_JoinCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunitiesServiceServer).JoinCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.CommunitiesService/JoinCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunitiesServiceServer).JoinCommunity(ctx, req.(*JoinCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunitiesService_ListCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunitiesServiceServer).ListCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.zion.v1.CommunitiesService/ListCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunitiesServiceServer).ListCommunity(ctx, req.(*ListCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunitiesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.zion.v1.CommunitiesService",
	HandlerType: (*CommunitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunity",
			Handler:    _CommunitiesService_CreateCommunity_Handler,
		},
		{
			MethodName: "DeleteCommunity",
			Handler:    _CommunitiesService_DeleteCommunity_Handler,
		},
		{
			MethodName: "EditCommunity",
			Handler:    _CommunitiesService_EditCommunity_Handler,
		},
		{
			MethodName: "JoinCommunity",
			Handler:    _CommunitiesService_JoinCommunity_Handler,
		},
		{
			MethodName: "ListCommunity",
			Handler:    _CommunitiesService_ListCommunity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/zion/v1/communities.proto",
}
