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

// PermissionsServiceClient is the client API for PermissionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionsServiceClient interface {
	PermissionsRequest(ctx context.Context, in *PermissionsRequestRequest, opts ...grpc.CallOption) (*PermissionsRequestResponse, error)
	PermissionsQuery(ctx context.Context, in *PermissionsQueryRequest, opts ...grpc.CallOption) (*PermissionsQueryResponse, error)
	PermissionsGrant(ctx context.Context, in *PermissionsGrantRequest, opts ...grpc.CallOption) (*PermissionsGrantResponse, error)
	PermissionsRevoke(ctx context.Context, in *PermissionsRevokeRequest, opts ...grpc.CallOption) (*PermissionsRevokeResponse, error)
}

type permissionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionsServiceClient(cc grpc.ClientConnInterface) PermissionsServiceClient {
	return &permissionsServiceClient{cc}
}

func (c *permissionsServiceClient) PermissionsRequest(ctx context.Context, in *PermissionsRequestRequest, opts ...grpc.CallOption) (*PermissionsRequestResponse, error) {
	out := new(PermissionsRequestResponse)
	err := c.cc.Invoke(ctx, "/proto.identityhub.v1.PermissionsService/PermissionsRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) PermissionsQuery(ctx context.Context, in *PermissionsQueryRequest, opts ...grpc.CallOption) (*PermissionsQueryResponse, error) {
	out := new(PermissionsQueryResponse)
	err := c.cc.Invoke(ctx, "/proto.identityhub.v1.PermissionsService/PermissionsQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) PermissionsGrant(ctx context.Context, in *PermissionsGrantRequest, opts ...grpc.CallOption) (*PermissionsGrantResponse, error) {
	out := new(PermissionsGrantResponse)
	err := c.cc.Invoke(ctx, "/proto.identityhub.v1.PermissionsService/PermissionsGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) PermissionsRevoke(ctx context.Context, in *PermissionsRevokeRequest, opts ...grpc.CallOption) (*PermissionsRevokeResponse, error) {
	out := new(PermissionsRevokeResponse)
	err := c.cc.Invoke(ctx, "/proto.identityhub.v1.PermissionsService/PermissionsRevoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionsServiceServer is the server API for PermissionsService service.
// All implementations should embed UnimplementedPermissionsServiceServer
// for forward compatibility
type PermissionsServiceServer interface {
	PermissionsRequest(context.Context, *PermissionsRequestRequest) (*PermissionsRequestResponse, error)
	PermissionsQuery(context.Context, *PermissionsQueryRequest) (*PermissionsQueryResponse, error)
	PermissionsGrant(context.Context, *PermissionsGrantRequest) (*PermissionsGrantResponse, error)
	PermissionsRevoke(context.Context, *PermissionsRevokeRequest) (*PermissionsRevokeResponse, error)
}

// UnimplementedPermissionsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPermissionsServiceServer struct {
}

func (UnimplementedPermissionsServiceServer) PermissionsRequest(context.Context, *PermissionsRequestRequest) (*PermissionsRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionsRequest not implemented")
}
func (UnimplementedPermissionsServiceServer) PermissionsQuery(context.Context, *PermissionsQueryRequest) (*PermissionsQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionsQuery not implemented")
}
func (UnimplementedPermissionsServiceServer) PermissionsGrant(context.Context, *PermissionsGrantRequest) (*PermissionsGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionsGrant not implemented")
}
func (UnimplementedPermissionsServiceServer) PermissionsRevoke(context.Context, *PermissionsRevokeRequest) (*PermissionsRevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionsRevoke not implemented")
}

// UnsafePermissionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionsServiceServer will
// result in compilation errors.
type UnsafePermissionsServiceServer interface {
	mustEmbedUnimplementedPermissionsServiceServer()
}

func RegisterPermissionsServiceServer(s grpc.ServiceRegistrar, srv PermissionsServiceServer) {
	s.RegisterService(&_PermissionsService_serviceDesc, srv)
}

func _PermissionsService_PermissionsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).PermissionsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.identityhub.v1.PermissionsService/PermissionsRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).PermissionsRequest(ctx, req.(*PermissionsRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_PermissionsQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).PermissionsQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.identityhub.v1.PermissionsService/PermissionsQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).PermissionsQuery(ctx, req.(*PermissionsQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_PermissionsGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).PermissionsGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.identityhub.v1.PermissionsService/PermissionsGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).PermissionsGrant(ctx, req.(*PermissionsGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_PermissionsRevoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionsRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).PermissionsRevoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.identityhub.v1.PermissionsService/PermissionsRevoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).PermissionsRevoke(ctx, req.(*PermissionsRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PermissionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.identityhub.v1.PermissionsService",
	HandlerType: (*PermissionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PermissionsRequest",
			Handler:    _PermissionsService_PermissionsRequest_Handler,
		},
		{
			MethodName: "PermissionsQuery",
			Handler:    _PermissionsService_PermissionsQuery_Handler,
		},
		{
			MethodName: "PermissionsGrant",
			Handler:    _PermissionsService_PermissionsGrant_Handler,
		},
		{
			MethodName: "PermissionsRevoke",
			Handler:    _PermissionsService_PermissionsRevoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/identityhub/v1/permissions.proto",
}
