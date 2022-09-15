// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: remote_connection_svc.proto

package appcontextsvc_client

import (
	context "context"
	models "gitlab.com/place-me/appcontextsvc-client-go/models"
	common "gitlab.com/place-me/place-to-go/grpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RemoteConnectionServiceClient is the client API for RemoteConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteConnectionServiceClient interface {
	ListRemoteConnections(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListRemoteConnectionsResponse, error)
	GetRemoteConnection(ctx context.Context, in *GetRemoteConnectionRequest, opts ...grpc.CallOption) (*models.RemoteConnection, error)
	CreateRemoteConnection(ctx context.Context, in *CreateRemoteConnectionRequest, opts ...grpc.CallOption) (*models.RemoteConnection, error)
	UpdateRemoteConnection(ctx context.Context, in *UpdateRemoteConnectionRequest, opts ...grpc.CallOption) (*models.RemoteConnection, error)
	DeleteRemoteConnection(ctx context.Context, in *DeleteRemoteConnectionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type remoteConnectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteConnectionServiceClient(cc grpc.ClientConnInterface) RemoteConnectionServiceClient {
	return &remoteConnectionServiceClient{cc}
}

func (c *remoteConnectionServiceClient) ListRemoteConnections(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListRemoteConnectionsResponse, error) {
	out := new(ListRemoteConnectionsResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteConnectionService/ListRemoteConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteConnectionServiceClient) GetRemoteConnection(ctx context.Context, in *GetRemoteConnectionRequest, opts ...grpc.CallOption) (*models.RemoteConnection, error) {
	out := new(models.RemoteConnection)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteConnectionService/GetRemoteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteConnectionServiceClient) CreateRemoteConnection(ctx context.Context, in *CreateRemoteConnectionRequest, opts ...grpc.CallOption) (*models.RemoteConnection, error) {
	out := new(models.RemoteConnection)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteConnectionService/CreateRemoteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteConnectionServiceClient) UpdateRemoteConnection(ctx context.Context, in *UpdateRemoteConnectionRequest, opts ...grpc.CallOption) (*models.RemoteConnection, error) {
	out := new(models.RemoteConnection)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteConnectionService/UpdateRemoteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteConnectionServiceClient) DeleteRemoteConnection(ctx context.Context, in *DeleteRemoteConnectionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteConnectionService/DeleteRemoteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteConnectionServiceServer is the server API for RemoteConnectionService service.
// All implementations must embed UnimplementedRemoteConnectionServiceServer
// for forward compatibility
type RemoteConnectionServiceServer interface {
	ListRemoteConnections(context.Context, *common.ListRequest) (*ListRemoteConnectionsResponse, error)
	GetRemoteConnection(context.Context, *GetRemoteConnectionRequest) (*models.RemoteConnection, error)
	CreateRemoteConnection(context.Context, *CreateRemoteConnectionRequest) (*models.RemoteConnection, error)
	UpdateRemoteConnection(context.Context, *UpdateRemoteConnectionRequest) (*models.RemoteConnection, error)
	DeleteRemoteConnection(context.Context, *DeleteRemoteConnectionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRemoteConnectionServiceServer()
}

// UnimplementedRemoteConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteConnectionServiceServer struct {
}

func (UnimplementedRemoteConnectionServiceServer) ListRemoteConnections(context.Context, *common.ListRequest) (*ListRemoteConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRemoteConnections not implemented")
}
func (UnimplementedRemoteConnectionServiceServer) GetRemoteConnection(context.Context, *GetRemoteConnectionRequest) (*models.RemoteConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemoteConnection not implemented")
}
func (UnimplementedRemoteConnectionServiceServer) CreateRemoteConnection(context.Context, *CreateRemoteConnectionRequest) (*models.RemoteConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRemoteConnection not implemented")
}
func (UnimplementedRemoteConnectionServiceServer) UpdateRemoteConnection(context.Context, *UpdateRemoteConnectionRequest) (*models.RemoteConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRemoteConnection not implemented")
}
func (UnimplementedRemoteConnectionServiceServer) DeleteRemoteConnection(context.Context, *DeleteRemoteConnectionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRemoteConnection not implemented")
}
func (UnimplementedRemoteConnectionServiceServer) mustEmbedUnimplementedRemoteConnectionServiceServer() {
}

// UnsafeRemoteConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteConnectionServiceServer will
// result in compilation errors.
type UnsafeRemoteConnectionServiceServer interface {
	mustEmbedUnimplementedRemoteConnectionServiceServer()
}

func RegisterRemoteConnectionServiceServer(s grpc.ServiceRegistrar, srv RemoteConnectionServiceServer) {
	s.RegisterService(&RemoteConnectionService_ServiceDesc, srv)
}

func _RemoteConnectionService_ListRemoteConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteConnectionServiceServer).ListRemoteConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteConnectionService/ListRemoteConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteConnectionServiceServer).ListRemoteConnections(ctx, req.(*common.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteConnectionService_GetRemoteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteConnectionServiceServer).GetRemoteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteConnectionService/GetRemoteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteConnectionServiceServer).GetRemoteConnection(ctx, req.(*GetRemoteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteConnectionService_CreateRemoteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRemoteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteConnectionServiceServer).CreateRemoteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteConnectionService/CreateRemoteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteConnectionServiceServer).CreateRemoteConnection(ctx, req.(*CreateRemoteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteConnectionService_UpdateRemoteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRemoteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteConnectionServiceServer).UpdateRemoteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteConnectionService/UpdateRemoteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteConnectionServiceServer).UpdateRemoteConnection(ctx, req.(*UpdateRemoteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteConnectionService_DeleteRemoteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRemoteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteConnectionServiceServer).DeleteRemoteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteConnectionService/DeleteRemoteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteConnectionServiceServer).DeleteRemoteConnection(ctx, req.(*DeleteRemoteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteConnectionService_ServiceDesc is the grpc.ServiceDesc for RemoteConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appcontext.services.RemoteConnectionService",
	HandlerType: (*RemoteConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRemoteConnections",
			Handler:    _RemoteConnectionService_ListRemoteConnections_Handler,
		},
		{
			MethodName: "GetRemoteConnection",
			Handler:    _RemoteConnectionService_GetRemoteConnection_Handler,
		},
		{
			MethodName: "CreateRemoteConnection",
			Handler:    _RemoteConnectionService_CreateRemoteConnection_Handler,
		},
		{
			MethodName: "UpdateRemoteConnection",
			Handler:    _RemoteConnectionService_UpdateRemoteConnection_Handler,
		},
		{
			MethodName: "DeleteRemoteConnection",
			Handler:    _RemoteConnectionService_DeleteRemoteConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_connection_svc.proto",
}
