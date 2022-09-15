// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: remote_asset_svc.proto

package appcontextsvc_client

import (
	context "context"
	models "gitlab.com/place-me/appcontextsvc-client-go/models"
	common "gitlab.com/place-me/place-to-go/grpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RemoteAssetServiceClient is the client API for RemoteAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteAssetServiceClient interface {
	GetRemoteAsset(ctx context.Context, in *GetRemoteAssetRequest, opts ...grpc.CallOption) (*models.RemoteAsset, error)
	ListRemoteAssets(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListRemoteAssetsResponse, error)
	ImportRemoteAssets(ctx context.Context, in *ImportRemoteAssetsRequest, opts ...grpc.CallOption) (*ImportRemoteAssetsResponse, error)
}

type remoteAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteAssetServiceClient(cc grpc.ClientConnInterface) RemoteAssetServiceClient {
	return &remoteAssetServiceClient{cc}
}

func (c *remoteAssetServiceClient) GetRemoteAsset(ctx context.Context, in *GetRemoteAssetRequest, opts ...grpc.CallOption) (*models.RemoteAsset, error) {
	out := new(models.RemoteAsset)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteAssetService/GetRemoteAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteAssetServiceClient) ListRemoteAssets(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListRemoteAssetsResponse, error) {
	out := new(ListRemoteAssetsResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteAssetService/ListRemoteAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteAssetServiceClient) ImportRemoteAssets(ctx context.Context, in *ImportRemoteAssetsRequest, opts ...grpc.CallOption) (*ImportRemoteAssetsResponse, error) {
	out := new(ImportRemoteAssetsResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RemoteAssetService/ImportRemoteAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteAssetServiceServer is the server API for RemoteAssetService service.
// All implementations must embed UnimplementedRemoteAssetServiceServer
// for forward compatibility
type RemoteAssetServiceServer interface {
	GetRemoteAsset(context.Context, *GetRemoteAssetRequest) (*models.RemoteAsset, error)
	ListRemoteAssets(context.Context, *common.ListRequest) (*ListRemoteAssetsResponse, error)
	ImportRemoteAssets(context.Context, *ImportRemoteAssetsRequest) (*ImportRemoteAssetsResponse, error)
	mustEmbedUnimplementedRemoteAssetServiceServer()
}

// UnimplementedRemoteAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteAssetServiceServer struct {
}

func (UnimplementedRemoteAssetServiceServer) GetRemoteAsset(context.Context, *GetRemoteAssetRequest) (*models.RemoteAsset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemoteAsset not implemented")
}
func (UnimplementedRemoteAssetServiceServer) ListRemoteAssets(context.Context, *common.ListRequest) (*ListRemoteAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRemoteAssets not implemented")
}
func (UnimplementedRemoteAssetServiceServer) ImportRemoteAssets(context.Context, *ImportRemoteAssetsRequest) (*ImportRemoteAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportRemoteAssets not implemented")
}
func (UnimplementedRemoteAssetServiceServer) mustEmbedUnimplementedRemoteAssetServiceServer() {}

// UnsafeRemoteAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteAssetServiceServer will
// result in compilation errors.
type UnsafeRemoteAssetServiceServer interface {
	mustEmbedUnimplementedRemoteAssetServiceServer()
}

func RegisterRemoteAssetServiceServer(s grpc.ServiceRegistrar, srv RemoteAssetServiceServer) {
	s.RegisterService(&RemoteAssetService_ServiceDesc, srv)
}

func _RemoteAssetService_GetRemoteAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteAssetServiceServer).GetRemoteAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteAssetService/GetRemoteAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteAssetServiceServer).GetRemoteAsset(ctx, req.(*GetRemoteAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteAssetService_ListRemoteAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteAssetServiceServer).ListRemoteAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteAssetService/ListRemoteAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteAssetServiceServer).ListRemoteAssets(ctx, req.(*common.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteAssetService_ImportRemoteAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportRemoteAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteAssetServiceServer).ImportRemoteAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RemoteAssetService/ImportRemoteAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteAssetServiceServer).ImportRemoteAssets(ctx, req.(*ImportRemoteAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteAssetService_ServiceDesc is the grpc.ServiceDesc for RemoteAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appcontext.services.RemoteAssetService",
	HandlerType: (*RemoteAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRemoteAsset",
			Handler:    _RemoteAssetService_GetRemoteAsset_Handler,
		},
		{
			MethodName: "ListRemoteAssets",
			Handler:    _RemoteAssetService_ListRemoteAssets_Handler,
		},
		{
			MethodName: "ImportRemoteAssets",
			Handler:    _RemoteAssetService_ImportRemoteAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_asset_svc.proto",
}
