// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appcontextsvc_client

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RbacServiceClient is the client API for RbacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RbacServiceClient interface {
	AssignSubjectToRole(ctx context.Context, in *AssignSubjectToRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveSubjectFromRole(ctx context.Context, in *RemoveSubjectFromRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AssignPermissionToRole(ctx context.Context, in *AssignPermissionToRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AssignedRolesFromSubject(ctx context.Context, in *AssignedRolesFromSubjectRequest, opts ...grpc.CallOption) (*AssignedRolesResponse, error)
	AssignedPermissionsFromSubject(ctx context.Context, in *AssignedPermissionsFromSubjectRequest, opts ...grpc.CallOption) (*AssignedPermissionsResponse, error)
	AssignedSubjectsFromRole(ctx context.Context, in *AssignedSubjectsFromRoleRequest, opts ...grpc.CallOption) (*AssignedSubjectsFromRoleResponse, error)
	AssignedRolesFromPermission(ctx context.Context, in *AssignedRolesFromPermissionRequest, opts ...grpc.CallOption) (*AssignedRolesResponse, error)
	AssignedPermissionsFromRole(ctx context.Context, in *AssignedPermissionsFromRoleRequest, opts ...grpc.CallOption) (*AssignedPermissionsResponse, error)
}

type rbacServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRbacServiceClient(cc grpc.ClientConnInterface) RbacServiceClient {
	return &rbacServiceClient{cc}
}

func (c *rbacServiceClient) AssignSubjectToRole(ctx context.Context, in *AssignSubjectToRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignSubjectToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) RemoveSubjectFromRole(ctx context.Context, in *RemoveSubjectFromRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/RemoveSubjectFromRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignPermissionToRole(ctx context.Context, in *AssignPermissionToRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignPermissionToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/RemovePermissionFromRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignedRolesFromSubject(ctx context.Context, in *AssignedRolesFromSubjectRequest, opts ...grpc.CallOption) (*AssignedRolesResponse, error) {
	out := new(AssignedRolesResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignedRolesFromSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignedPermissionsFromSubject(ctx context.Context, in *AssignedPermissionsFromSubjectRequest, opts ...grpc.CallOption) (*AssignedPermissionsResponse, error) {
	out := new(AssignedPermissionsResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignedPermissionsFromSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignedSubjectsFromRole(ctx context.Context, in *AssignedSubjectsFromRoleRequest, opts ...grpc.CallOption) (*AssignedSubjectsFromRoleResponse, error) {
	out := new(AssignedSubjectsFromRoleResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignedSubjectsFromRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignedRolesFromPermission(ctx context.Context, in *AssignedRolesFromPermissionRequest, opts ...grpc.CallOption) (*AssignedRolesResponse, error) {
	out := new(AssignedRolesResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignedRolesFromPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignedPermissionsFromRole(ctx context.Context, in *AssignedPermissionsFromRoleRequest, opts ...grpc.CallOption) (*AssignedPermissionsResponse, error) {
	out := new(AssignedPermissionsResponse)
	err := c.cc.Invoke(ctx, "/appcontext.services.RbacService/AssignedPermissionsFromRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RbacServiceServer is the server API for RbacService service.
// All implementations must embed UnimplementedRbacServiceServer
// for forward compatibility
type RbacServiceServer interface {
	AssignSubjectToRole(context.Context, *AssignSubjectToRoleRequest) (*empty.Empty, error)
	RemoveSubjectFromRole(context.Context, *RemoveSubjectFromRoleRequest) (*empty.Empty, error)
	AssignPermissionToRole(context.Context, *AssignPermissionToRoleRequest) (*empty.Empty, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*empty.Empty, error)
	AssignedRolesFromSubject(context.Context, *AssignedRolesFromSubjectRequest) (*AssignedRolesResponse, error)
	AssignedPermissionsFromSubject(context.Context, *AssignedPermissionsFromSubjectRequest) (*AssignedPermissionsResponse, error)
	AssignedSubjectsFromRole(context.Context, *AssignedSubjectsFromRoleRequest) (*AssignedSubjectsFromRoleResponse, error)
	AssignedRolesFromPermission(context.Context, *AssignedRolesFromPermissionRequest) (*AssignedRolesResponse, error)
	AssignedPermissionsFromRole(context.Context, *AssignedPermissionsFromRoleRequest) (*AssignedPermissionsResponse, error)
	mustEmbedUnimplementedRbacServiceServer()
}

// UnimplementedRbacServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRbacServiceServer struct {
}

func (UnimplementedRbacServiceServer) AssignSubjectToRole(context.Context, *AssignSubjectToRoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignSubjectToRole not implemented")
}
func (UnimplementedRbacServiceServer) RemoveSubjectFromRole(context.Context, *RemoveSubjectFromRoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubjectFromRole not implemented")
}
func (UnimplementedRbacServiceServer) AssignPermissionToRole(context.Context, *AssignPermissionToRoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignPermissionToRole not implemented")
}
func (UnimplementedRbacServiceServer) RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermissionFromRole not implemented")
}
func (UnimplementedRbacServiceServer) AssignedRolesFromSubject(context.Context, *AssignedRolesFromSubjectRequest) (*AssignedRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignedRolesFromSubject not implemented")
}
func (UnimplementedRbacServiceServer) AssignedPermissionsFromSubject(context.Context, *AssignedPermissionsFromSubjectRequest) (*AssignedPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignedPermissionsFromSubject not implemented")
}
func (UnimplementedRbacServiceServer) AssignedSubjectsFromRole(context.Context, *AssignedSubjectsFromRoleRequest) (*AssignedSubjectsFromRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignedSubjectsFromRole not implemented")
}
func (UnimplementedRbacServiceServer) AssignedRolesFromPermission(context.Context, *AssignedRolesFromPermissionRequest) (*AssignedRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignedRolesFromPermission not implemented")
}
func (UnimplementedRbacServiceServer) AssignedPermissionsFromRole(context.Context, *AssignedPermissionsFromRoleRequest) (*AssignedPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignedPermissionsFromRole not implemented")
}
func (UnimplementedRbacServiceServer) mustEmbedUnimplementedRbacServiceServer() {}

// UnsafeRbacServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RbacServiceServer will
// result in compilation errors.
type UnsafeRbacServiceServer interface {
	mustEmbedUnimplementedRbacServiceServer()
}

func RegisterRbacServiceServer(s grpc.ServiceRegistrar, srv RbacServiceServer) {
	s.RegisterService(&RbacService_ServiceDesc, srv)
}

func _RbacService_AssignSubjectToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignSubjectToRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignSubjectToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignSubjectToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignSubjectToRole(ctx, req.(*AssignSubjectToRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_RemoveSubjectFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubjectFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).RemoveSubjectFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/RemoveSubjectFromRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).RemoveSubjectFromRole(ctx, req.(*RemoveSubjectFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignPermissionToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignPermissionToRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignPermissionToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignPermissionToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignPermissionToRole(ctx, req.(*AssignPermissionToRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_RemovePermissionFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).RemovePermissionFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/RemovePermissionFromRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).RemovePermissionFromRole(ctx, req.(*RemovePermissionFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignedRolesFromSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignedRolesFromSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignedRolesFromSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignedRolesFromSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignedRolesFromSubject(ctx, req.(*AssignedRolesFromSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignedPermissionsFromSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignedPermissionsFromSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignedPermissionsFromSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignedPermissionsFromSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignedPermissionsFromSubject(ctx, req.(*AssignedPermissionsFromSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignedSubjectsFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignedSubjectsFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignedSubjectsFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignedSubjectsFromRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignedSubjectsFromRole(ctx, req.(*AssignedSubjectsFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignedRolesFromPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignedRolesFromPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignedRolesFromPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignedRolesFromPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignedRolesFromPermission(ctx, req.(*AssignedRolesFromPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignedPermissionsFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignedPermissionsFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignedPermissionsFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appcontext.services.RbacService/AssignedPermissionsFromRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignedPermissionsFromRole(ctx, req.(*AssignedPermissionsFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RbacService_ServiceDesc is the grpc.ServiceDesc for RbacService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RbacService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appcontext.services.RbacService",
	HandlerType: (*RbacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignSubjectToRole",
			Handler:    _RbacService_AssignSubjectToRole_Handler,
		},
		{
			MethodName: "RemoveSubjectFromRole",
			Handler:    _RbacService_RemoveSubjectFromRole_Handler,
		},
		{
			MethodName: "AssignPermissionToRole",
			Handler:    _RbacService_AssignPermissionToRole_Handler,
		},
		{
			MethodName: "RemovePermissionFromRole",
			Handler:    _RbacService_RemovePermissionFromRole_Handler,
		},
		{
			MethodName: "AssignedRolesFromSubject",
			Handler:    _RbacService_AssignedRolesFromSubject_Handler,
		},
		{
			MethodName: "AssignedPermissionsFromSubject",
			Handler:    _RbacService_AssignedPermissionsFromSubject_Handler,
		},
		{
			MethodName: "AssignedSubjectsFromRole",
			Handler:    _RbacService_AssignedSubjectsFromRole_Handler,
		},
		{
			MethodName: "AssignedRolesFromPermission",
			Handler:    _RbacService_AssignedRolesFromPermission_Handler,
		},
		{
			MethodName: "AssignedPermissionsFromRole",
			Handler:    _RbacService_AssignedPermissionsFromRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rbac_svc.proto",
}
