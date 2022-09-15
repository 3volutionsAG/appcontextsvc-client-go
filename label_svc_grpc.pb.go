// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: label_svc.proto

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

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LabelServiceClient interface {
	ListLabels(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListLabelResponse, error)
	CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*models.Label, error)
	DeleteLabel(ctx context.Context, in *DeleteLabelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*models.Label, error)
	UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*models.Label, error)
}

type labelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelServiceClient(cc grpc.ClientConnInterface) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) ListLabels(ctx context.Context, in *common.ListRequest, opts ...grpc.CallOption) (*ListLabelResponse, error) {
	out := new(ListLabelResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.LabelService/ListLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*models.Label, error) {
	out := new(models.Label)
	err := c.cc.Invoke(ctx, "/appcontext.services.LabelService/CreateLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) DeleteLabel(ctx context.Context, in *DeleteLabelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/appcontext.services.LabelService/DeleteLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*models.Label, error) {
	out := new(models.Label)
	err := c.cc.Invoke(ctx, "/appcontext.services.LabelService/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*models.Label, error) {
	out := new(models.Label)
	err := c.cc.Invoke(ctx, "/appcontext.services.LabelService/UpdateLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
// All implementations must embed UnimplementedLabelServiceServer
// for forward compatibility
type LabelServiceServer interface {
	ListLabels(context.Context, *common.ListRequest) (*ListLabelResponse, error)
	CreateLabel(context.Context, *CreateLabelRequest) (*models.Label, error)
	DeleteLabel(context.Context, *DeleteLabelRequest) (*emptypb.Empty, error)
	GetLabel(context.Context, *GetLabelRequest) (*models.Label, error)
	UpdateLabel(context.Context, *UpdateLabelRequest) (*models.Label, error)
	mustEmbedUnimplementedLabelServiceServer()
}

// UnimplementedLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLabelServiceServer struct {
}

func (UnimplementedLabelServiceServer) ListLabels(context.Context, *common.ListRequest) (*ListLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabels not implemented")
}
func (UnimplementedLabelServiceServer) CreateLabel(context.Context, *CreateLabelRequest) (*models.Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabel not implemented")
}
func (UnimplementedLabelServiceServer) DeleteLabel(context.Context, *DeleteLabelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLabel not implemented")
}
func (UnimplementedLabelServiceServer) GetLabel(context.Context, *GetLabelRequest) (*models.Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}
func (UnimplementedLabelServiceServer) UpdateLabel(context.Context, *UpdateLabelRequest) (*models.Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLabel not implemented")
}
func (UnimplementedLabelServiceServer) mustEmbedUnimplementedLabelServiceServer() {}

// UnsafeLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabelServiceServer will
// result in compilation errors.
type UnsafeLabelServiceServer interface {
	mustEmbedUnimplementedLabelServiceServer()
}

func RegisterLabelServiceServer(s grpc.ServiceRegistrar, srv LabelServiceServer) {
	s.RegisterService(&LabelService_ServiceDesc, srv)
}

func _LabelService_ListLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).ListLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.LabelService/ListLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).ListLabels(ctx, req.(*common.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_CreateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).CreateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.LabelService/CreateLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).CreateLabel(ctx, req.(*CreateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_DeleteLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).DeleteLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.LabelService/DeleteLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).DeleteLabel(ctx, req.(*DeleteLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.LabelService/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_UpdateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).UpdateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.LabelService/UpdateLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).UpdateLabel(ctx, req.(*UpdateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LabelService_ServiceDesc is the grpc.ServiceDesc for LabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appcontext.services.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLabels",
			Handler:    _LabelService_ListLabels_Handler,
		},
		{
			MethodName: "CreateLabel",
			Handler:    _LabelService_CreateLabel_Handler,
		},
		{
			MethodName: "DeleteLabel",
			Handler:    _LabelService_DeleteLabel_Handler,
		},
		{
			MethodName: "GetLabel",
			Handler:    _LabelService_GetLabel_Handler,
		},
		{
			MethodName: "UpdateLabel",
			Handler:    _LabelService_UpdateLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "label_svc.proto",
}