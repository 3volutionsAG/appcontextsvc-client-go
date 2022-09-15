// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: ory/keto/acl/v1alpha1/write_service.proto

package acl

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

// WriteServiceClient is the client API for WriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriteServiceClient interface {
	// Writes one or more relation tuples in a single transaction.
	TransactRelationTuples(ctx context.Context, in *TransactRelationTuplesRequest, opts ...grpc.CallOption) (*TransactRelationTuplesResponse, error)
}

type writeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteServiceClient(cc grpc.ClientConnInterface) WriteServiceClient {
	return &writeServiceClient{cc}
}

func (c *writeServiceClient) TransactRelationTuples(ctx context.Context, in *TransactRelationTuplesRequest, opts ...grpc.CallOption) (*TransactRelationTuplesResponse, error) {
	out := new(TransactRelationTuplesResponse)
	err := c.cc.Invoke(ctx, "/ory.keto.acl.v1alpha1.WriteService/TransactRelationTuples", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteServiceServer is the server API for WriteService service.
// All implementations must embed UnimplementedWriteServiceServer
// for forward compatibility
type WriteServiceServer interface {
	// Writes one or more relation tuples in a single transaction.
	TransactRelationTuples(context.Context, *TransactRelationTuplesRequest) (*TransactRelationTuplesResponse, error)
	mustEmbedUnimplementedWriteServiceServer()
}

// UnimplementedWriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWriteServiceServer struct {
}

func (UnimplementedWriteServiceServer) TransactRelationTuples(context.Context, *TransactRelationTuplesRequest) (*TransactRelationTuplesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactRelationTuples not implemented")
}
func (UnimplementedWriteServiceServer) mustEmbedUnimplementedWriteServiceServer() {}

// UnsafeWriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriteServiceServer will
// result in compilation errors.
type UnsafeWriteServiceServer interface {
	mustEmbedUnimplementedWriteServiceServer()
}

func RegisterWriteServiceServer(s grpc.ServiceRegistrar, srv WriteServiceServer) {
	s.RegisterService(&WriteService_ServiceDesc, srv)
}

func _WriteService_TransactRelationTuples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactRelationTuplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteServiceServer).TransactRelationTuples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ory.keto.acl.v1alpha1.WriteService/TransactRelationTuples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteServiceServer).TransactRelationTuples(ctx, req.(*TransactRelationTuplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriteService_ServiceDesc is the grpc.ServiceDesc for WriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ory.keto.acl.v1alpha1.WriteService",
	HandlerType: (*WriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransactRelationTuples",
			Handler:    _WriteService_TransactRelationTuples_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ory/keto/acl/v1alpha1/write_service.proto",
}
