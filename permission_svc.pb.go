// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: permission_svc.proto

package appcontextsvc_client

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	base "gitlab.com/place-me/appcontextsvc-client-go/base"
	models "gitlab.com/place-me/appcontextsvc-client-go/models"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field name should match the noun "Permission" in the method name.
	// There will be a maximum number of items returned based on the page_size field in the request.
	Permissions []*models.Permission `protobuf:"bytes,1,rep,name=Permissions,proto3" json:"Permissions,omitempty"`
	PageIndex   int32                `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32                `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	TotalCount  int32                `protobuf:"varint,4,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListPermissionsResponse) Reset() {
	*x = ListPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsResponse) ProtoMessage() {}

func (x *ListPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsResponse.ProtoReflect.Descriptor instead.
func (*ListPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_permission_svc_proto_rawDescGZIP(), []int{0}
}

func (x *ListPermissionsResponse) GetPermissions() []*models.Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *ListPermissionsResponse) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *ListPermissionsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPermissionsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPermissionRequest) Reset() {
	*x = GetPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionRequest) ProtoMessage() {}

func (x *GetPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_svc_proto_rawDescGZIP(), []int{1}
}

func (x *GetPermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Permission resource to create.
	// The field name should match the Noun in the method name.
	Permission *models.Permission `protobuf:"bytes,1,opt,name=Permission,proto3" json:"Permission,omitempty"`
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionRequest.ProtoReflect.Descriptor instead.
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_svc_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePermissionRequest) GetPermission() *models.Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type UpdatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Permission resource which replaces the resource on the server.
	Permission *models.Permission `protobuf:"bytes,1,opt,name=Permission,proto3" json:"Permission,omitempty"`
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_svc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePermissionRequest) GetPermission() *models.Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type DeletePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePermissionRequest) Reset() {
	*x = DeletePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionRequest) ProtoMessage() {}

func (x *DeletePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionRequest.ProtoReflect.Descriptor instead.
func (*DeletePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permission_svc_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_permission_svc_proto protoreflect.FileDescriptor

var file_permission_svc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x76, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x62, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0a,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xf1, 0x05, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x32, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x52, 0x6f, 0x6c, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x84,
	0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x65, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x70,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2d, 0x67, 0x6f, 0x3b, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73,
	0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x20, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_svc_proto_rawDescOnce sync.Once
	file_permission_svc_proto_rawDescData = file_permission_svc_proto_rawDesc
)

func file_permission_svc_proto_rawDescGZIP() []byte {
	file_permission_svc_proto_rawDescOnce.Do(func() {
		file_permission_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_svc_proto_rawDescData)
	})
	return file_permission_svc_proto_rawDescData
}

var file_permission_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_permission_svc_proto_goTypes = []interface{}{
	(*ListPermissionsResponse)(nil), // 0: appcontextsvc_client.services.ListPermissionsResponse
	(*GetPermissionRequest)(nil),    // 1: appcontextsvc_client.services.GetPermissionRequest
	(*CreatePermissionRequest)(nil), // 2: appcontextsvc_client.services.CreatePermissionRequest
	(*UpdatePermissionRequest)(nil), // 3: appcontextsvc_client.services.UpdatePermissionRequest
	(*DeletePermissionRequest)(nil), // 4: appcontextsvc_client.services.DeletePermissionRequest
	(*models.Permission)(nil),       // 5: appcontextsvc_client.models.Permission
	(*base.ListRequest)(nil),        // 6: appcontextsvc_client.base.ListRequest
	(*emptypb.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_permission_svc_proto_depIdxs = []int32{
	5, // 0: appcontextsvc_client.services.ListPermissionsResponse.Permissions:type_name -> appcontextsvc_client.models.Permission
	5, // 1: appcontextsvc_client.services.CreatePermissionRequest.Permission:type_name -> appcontextsvc_client.models.Permission
	5, // 2: appcontextsvc_client.services.UpdatePermissionRequest.Permission:type_name -> appcontextsvc_client.models.Permission
	6, // 3: appcontextsvc_client.services.PermissionService.ListPermissions:input_type -> appcontextsvc_client.base.ListRequest
	1, // 4: appcontextsvc_client.services.PermissionService.GetPermission:input_type -> appcontextsvc_client.services.GetPermissionRequest
	2, // 5: appcontextsvc_client.services.PermissionService.CreatePermission:input_type -> appcontextsvc_client.services.CreatePermissionRequest
	3, // 6: appcontextsvc_client.services.PermissionService.UpdatePermission:input_type -> appcontextsvc_client.services.UpdatePermissionRequest
	4, // 7: appcontextsvc_client.services.PermissionService.DeletePermission:input_type -> appcontextsvc_client.services.DeletePermissionRequest
	0, // 8: appcontextsvc_client.services.PermissionService.ListPermissions:output_type -> appcontextsvc_client.services.ListPermissionsResponse
	5, // 9: appcontextsvc_client.services.PermissionService.GetPermission:output_type -> appcontextsvc_client.models.Permission
	5, // 10: appcontextsvc_client.services.PermissionService.CreatePermission:output_type -> appcontextsvc_client.models.Permission
	5, // 11: appcontextsvc_client.services.PermissionService.UpdatePermission:output_type -> appcontextsvc_client.models.Permission
	7, // 12: appcontextsvc_client.services.PermissionService.DeletePermission:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_permission_svc_proto_init() }
func file_permission_svc_proto_init() {
	if File_permission_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_permission_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_permission_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePermissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_permission_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_permission_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePermissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_permission_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_permission_svc_proto_goTypes,
		DependencyIndexes: file_permission_svc_proto_depIdxs,
		MessageInfos:      file_permission_svc_proto_msgTypes,
	}.Build()
	File_permission_svc_proto = out.File
	file_permission_svc_proto_rawDesc = nil
	file_permission_svc_proto_goTypes = nil
	file_permission_svc_proto_depIdxs = nil
}
