// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: seating_svc.proto

package appcontextsvc_client

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	models "gitlab.com/place-me/appcontextsvc-client-go/models"
	common "gitlab.com/place-me/place-to-go/grpc/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListSeatingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seatings   []*models.Seating `protobuf:"bytes,1,rep,name=seatings,proto3" json:"seatings,omitempty"`
	PageNumber int32             `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32             `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalCount int32             `protobuf:"varint,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListSeatingsResponse) Reset() {
	*x = ListSeatingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seating_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSeatingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSeatingsResponse) ProtoMessage() {}

func (x *ListSeatingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seating_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSeatingsResponse.ProtoReflect.Descriptor instead.
func (*ListSeatingsResponse) Descriptor() ([]byte, []int) {
	return file_seating_svc_proto_rawDescGZIP(), []int{0}
}

func (x *ListSeatingsResponse) GetSeatings() []*models.Seating {
	if x != nil {
		return x.Seatings
	}
	return nil
}

func (x *ListSeatingsResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListSeatingsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSeatingsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetSeatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSeatingRequest) Reset() {
	*x = GetSeatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seating_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeatingRequest) ProtoMessage() {}

func (x *GetSeatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seating_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeatingRequest.ProtoReflect.Descriptor instead.
func (*GetSeatingRequest) Descriptor() ([]byte, []int) {
	return file_seating_svc_proto_rawDescGZIP(), []int{1}
}

func (x *GetSeatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateSeatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seating *models.Seating `protobuf:"bytes,1,opt,name=seating,proto3" json:"seating,omitempty"`
}

func (x *CreateSeatingRequest) Reset() {
	*x = CreateSeatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seating_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeatingRequest) ProtoMessage() {}

func (x *CreateSeatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seating_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeatingRequest.ProtoReflect.Descriptor instead.
func (*CreateSeatingRequest) Descriptor() ([]byte, []int) {
	return file_seating_svc_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeatingRequest) GetSeating() *models.Seating {
	if x != nil {
		return x.Seating
	}
	return nil
}

type UpdateSeatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seating    *models.Seating        `protobuf:"bytes,1,opt,name=seating,proto3" json:"seating,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSeatingRequest) Reset() {
	*x = UpdateSeatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seating_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSeatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeatingRequest) ProtoMessage() {}

func (x *UpdateSeatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seating_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeatingRequest.ProtoReflect.Descriptor instead.
func (*UpdateSeatingRequest) Descriptor() ([]byte, []int) {
	return file_seating_svc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSeatingRequest) GetSeating() *models.Seating {
	if x != nil {
		return x.Seating
	}
	return nil
}

func (x *UpdateSeatingRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteSeatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Seating to be deleted.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSeatingRequest) Reset() {
	*x = DeleteSeatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seating_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSeatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSeatingRequest) ProtoMessage() {}

func (x *DeleteSeatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seating_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSeatingRequest.ProtoReflect.Descriptor instead.
func (*DeleteSeatingRequest) Descriptor() ([]byte, []int) {
	return file_seating_svc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSeatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_seating_svc_proto protoreflect.FileDescriptor

var file_seating_svc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73,
	0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbe, 0x03, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x52, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x65, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x3b, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x20, 0x41, 0x70, 0x70,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seating_svc_proto_rawDescOnce sync.Once
	file_seating_svc_proto_rawDescData = file_seating_svc_proto_rawDesc
)

func file_seating_svc_proto_rawDescGZIP() []byte {
	file_seating_svc_proto_rawDescOnce.Do(func() {
		file_seating_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_seating_svc_proto_rawDescData)
	})
	return file_seating_svc_proto_rawDescData
}

var file_seating_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_seating_svc_proto_goTypes = []interface{}{
	(*ListSeatingsResponse)(nil),  // 0: appcontext.services.ListSeatingsResponse
	(*GetSeatingRequest)(nil),     // 1: appcontext.services.GetSeatingRequest
	(*CreateSeatingRequest)(nil),  // 2: appcontext.services.CreateSeatingRequest
	(*UpdateSeatingRequest)(nil),  // 3: appcontext.services.UpdateSeatingRequest
	(*DeleteSeatingRequest)(nil),  // 4: appcontext.services.DeleteSeatingRequest
	(*models.Seating)(nil),        // 5: appcontext.models.Seating
	(*fieldmaskpb.FieldMask)(nil), // 6: google.protobuf.FieldMask
	(*common.ListRequest)(nil),    // 7: placeme.common.ListRequest
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_seating_svc_proto_depIdxs = []int32{
	5, // 0: appcontext.services.ListSeatingsResponse.seatings:type_name -> appcontext.models.Seating
	5, // 1: appcontext.services.CreateSeatingRequest.seating:type_name -> appcontext.models.Seating
	5, // 2: appcontext.services.UpdateSeatingRequest.seating:type_name -> appcontext.models.Seating
	6, // 3: appcontext.services.UpdateSeatingRequest.update_mask:type_name -> google.protobuf.FieldMask
	7, // 4: appcontext.services.SeatingService.ListSeatings:input_type -> placeme.common.ListRequest
	1, // 5: appcontext.services.SeatingService.GetSeating:input_type -> appcontext.services.GetSeatingRequest
	2, // 6: appcontext.services.SeatingService.CreateSeating:input_type -> appcontext.services.CreateSeatingRequest
	3, // 7: appcontext.services.SeatingService.UpdateSeating:input_type -> appcontext.services.UpdateSeatingRequest
	4, // 8: appcontext.services.SeatingService.DeleteSeating:input_type -> appcontext.services.DeleteSeatingRequest
	0, // 9: appcontext.services.SeatingService.ListSeatings:output_type -> appcontext.services.ListSeatingsResponse
	5, // 10: appcontext.services.SeatingService.GetSeating:output_type -> appcontext.models.Seating
	5, // 11: appcontext.services.SeatingService.CreateSeating:output_type -> appcontext.models.Seating
	5, // 12: appcontext.services.SeatingService.UpdateSeating:output_type -> appcontext.models.Seating
	8, // 13: appcontext.services.SeatingService.DeleteSeating:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_seating_svc_proto_init() }
func file_seating_svc_proto_init() {
	if File_seating_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seating_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSeatingsResponse); i {
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
		file_seating_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeatingRequest); i {
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
		file_seating_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeatingRequest); i {
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
		file_seating_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSeatingRequest); i {
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
		file_seating_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSeatingRequest); i {
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
			RawDescriptor: file_seating_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seating_svc_proto_goTypes,
		DependencyIndexes: file_seating_svc_proto_depIdxs,
		MessageInfos:      file_seating_svc_proto_msgTypes,
	}.Build()
	File_seating_svc_proto = out.File
	file_seating_svc_proto_rawDesc = nil
	file_seating_svc_proto_goTypes = nil
	file_seating_svc_proto_depIdxs = nil
}
