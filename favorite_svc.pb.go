// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: favorite_svc.proto

package appcontextsvc_client

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	models "gitlab.com/place-me/appcontextsvc-client-go/models"
	common "gitlab.com/place-me/place-to-go/grpc/common"
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

type ListFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Favorites  []*models.Favorite `protobuf:"bytes,1,rep,name=favorites,proto3" json:"favorites,omitempty"`
	PageNumber int32              `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32              `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalCount int32              `protobuf:"varint,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListFavoriteResponse) Reset() {
	*x = ListFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFavoriteResponse) ProtoMessage() {}

func (x *ListFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFavoriteResponse.ProtoReflect.Descriptor instead.
func (*ListFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_favorite_svc_proto_rawDescGZIP(), []int{0}
}

func (x *ListFavoriteResponse) GetFavorites() []*models.Favorite {
	if x != nil {
		return x.Favorites
	}
	return nil
}

func (x *ListFavoriteResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListFavoriteResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFavoriteResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CreateFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Favorite *models.Favorite `protobuf:"bytes,1,opt,name=favorite,proto3" json:"favorite,omitempty"`
}

func (x *CreateFavoriteRequest) Reset() {
	*x = CreateFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFavoriteRequest) ProtoMessage() {}

func (x *CreateFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFavoriteRequest.ProtoReflect.Descriptor instead.
func (*CreateFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_svc_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFavoriteRequest) GetFavorite() *models.Favorite {
	if x != nil {
		return x.Favorite
	}
	return nil
}

type DeleteFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFavoriteRequest) Reset() {
	*x = DeleteFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFavoriteRequest) ProtoMessage() {}

func (x *DeleteFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFavoriteRequest.ProtoReflect.Descriptor instead.
func (*DeleteFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_svc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFavoriteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFavoriteRequest) Reset() {
	*x = GetFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavoriteRequest) ProtoMessage() {}

func (x *GetFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavoriteRequest.ProtoReflect.Descriptor instead.
func (*GetFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_favorite_svc_proto_rawDescGZIP(), []int{3}
}

func (x *GetFavoriteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_favorite_svc_proto protoreflect.FileDescriptor

var file_favorite_svc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22,
	0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xf0, 0x02, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x53, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x27, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x42, 0x65, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x70,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2d, 0x67, 0x6f, 0x3b, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73,
	0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x20, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favorite_svc_proto_rawDescOnce sync.Once
	file_favorite_svc_proto_rawDescData = file_favorite_svc_proto_rawDesc
)

func file_favorite_svc_proto_rawDescGZIP() []byte {
	file_favorite_svc_proto_rawDescOnce.Do(func() {
		file_favorite_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorite_svc_proto_rawDescData)
	})
	return file_favorite_svc_proto_rawDescData
}

var file_favorite_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_favorite_svc_proto_goTypes = []interface{}{
	(*ListFavoriteResponse)(nil),  // 0: appcontext.services.ListFavoriteResponse
	(*CreateFavoriteRequest)(nil), // 1: appcontext.services.CreateFavoriteRequest
	(*DeleteFavoriteRequest)(nil), // 2: appcontext.services.DeleteFavoriteRequest
	(*GetFavoriteRequest)(nil),    // 3: appcontext.services.GetFavoriteRequest
	(*models.Favorite)(nil),       // 4: appcontext.models.Favorite
	(*common.ListRequest)(nil),    // 5: placeme.common.ListRequest
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_favorite_svc_proto_depIdxs = []int32{
	4, // 0: appcontext.services.ListFavoriteResponse.favorites:type_name -> appcontext.models.Favorite
	4, // 1: appcontext.services.CreateFavoriteRequest.favorite:type_name -> appcontext.models.Favorite
	5, // 2: appcontext.services.FavoriteService.ListFavorites:input_type -> placeme.common.ListRequest
	1, // 3: appcontext.services.FavoriteService.CreateFavorite:input_type -> appcontext.services.CreateFavoriteRequest
	2, // 4: appcontext.services.FavoriteService.DeleteFavorite:input_type -> appcontext.services.DeleteFavoriteRequest
	3, // 5: appcontext.services.FavoriteService.GetFavorite:input_type -> appcontext.services.GetFavoriteRequest
	0, // 6: appcontext.services.FavoriteService.ListFavorites:output_type -> appcontext.services.ListFavoriteResponse
	4, // 7: appcontext.services.FavoriteService.CreateFavorite:output_type -> appcontext.models.Favorite
	6, // 8: appcontext.services.FavoriteService.DeleteFavorite:output_type -> google.protobuf.Empty
	4, // 9: appcontext.services.FavoriteService.GetFavorite:output_type -> appcontext.models.Favorite
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_favorite_svc_proto_init() }
func file_favorite_svc_proto_init() {
	if File_favorite_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorite_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFavoriteResponse); i {
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
		file_favorite_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFavoriteRequest); i {
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
		file_favorite_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFavoriteRequest); i {
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
		file_favorite_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavoriteRequest); i {
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
			RawDescriptor: file_favorite_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorite_svc_proto_goTypes,
		DependencyIndexes: file_favorite_svc_proto_depIdxs,
		MessageInfos:      file_favorite_svc_proto_msgTypes,
	}.Build()
	File_favorite_svc_proto = out.File
	file_favorite_svc_proto_rawDesc = nil
	file_favorite_svc_proto_goTypes = nil
	file_favorite_svc_proto_depIdxs = nil
}
