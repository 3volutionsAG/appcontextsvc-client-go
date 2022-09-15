// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: models/favorite.proto

package models

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FavoriteEntityType int32

const (
	FavoriteEntityType_FAVORITE_ENTITY_TYPE_UNSPECIFIED FavoriteEntityType = 0
	FavoriteEntityType_FAVORITE_ENTITY_TYPE_ASSET       FavoriteEntityType = 1
)

// Enum value maps for FavoriteEntityType.
var (
	FavoriteEntityType_name = map[int32]string{
		0: "FAVORITE_ENTITY_TYPE_UNSPECIFIED",
		1: "FAVORITE_ENTITY_TYPE_ASSET",
	}
	FavoriteEntityType_value = map[string]int32{
		"FAVORITE_ENTITY_TYPE_UNSPECIFIED": 0,
		"FAVORITE_ENTITY_TYPE_ASSET":       1,
	}
)

func (x FavoriteEntityType) Enum() *FavoriteEntityType {
	p := new(FavoriteEntityType)
	*p = x
	return p
}

func (x FavoriteEntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FavoriteEntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_models_favorite_proto_enumTypes[0].Descriptor()
}

func (FavoriteEntityType) Type() protoreflect.EnumType {
	return &file_models_favorite_proto_enumTypes[0]
}

func (x FavoriteEntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FavoriteEntityType.Descriptor instead.
func (FavoriteEntityType) EnumDescriptor() ([]byte, []int) {
	return file_models_favorite_proto_rawDescGZIP(), []int{0}
}

type Favorite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *string            `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	TenantId      *string            `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3,oneof" json:"tenant_id,omitempty"`
	UserprofileId string             `protobuf:"bytes,3,opt,name=userprofile_id,json=userprofileId,proto3" json:"userprofile_id,omitempty"`
	EntityId      string             `protobuf:"bytes,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	EntityType    FavoriteEntityType `protobuf:"varint,5,opt,name=entity_type,json=entityType,proto3,enum=appcontext.models.FavoriteEntityType" json:"entity_type,omitempty"`
}

func (x *Favorite) Reset() {
	*x = Favorite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Favorite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Favorite) ProtoMessage() {}

func (x *Favorite) ProtoReflect() protoreflect.Message {
	mi := &file_models_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Favorite.ProtoReflect.Descriptor instead.
func (*Favorite) Descriptor() ([]byte, []int) {
	return file_models_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *Favorite) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Favorite) GetTenantId() string {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return ""
}

func (x *Favorite) GetUserprofileId() string {
	if x != nil {
		return x.UserprofileId
	}
	return ""
}

func (x *Favorite) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *Favorite) GetEntityType() FavoriteEntityType {
	if x != nil {
		return x.EntityType
	}
	return FavoriteEntityType_FAVORITE_ENTITY_TYPE_UNSPECIFIED
}

var File_models_favorite_proto protoreflect.FileDescriptor

var file_models_favorite_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42,
	0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0,
	0x01, 0x01, 0x48, 0x01, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x32, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06,
	0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0,
	0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x46, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x2a, 0x5a, 0x0a, 0x12,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x46, 0x41, 0x56, 0x4f, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x46, 0x41, 0x56, 0x4f,
	0x52, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x01, 0x42, 0x55, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f,
	0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02,
	0x1e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_favorite_proto_rawDescOnce sync.Once
	file_models_favorite_proto_rawDescData = file_models_favorite_proto_rawDesc
)

func file_models_favorite_proto_rawDescGZIP() []byte {
	file_models_favorite_proto_rawDescOnce.Do(func() {
		file_models_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_favorite_proto_rawDescData)
	})
	return file_models_favorite_proto_rawDescData
}

var file_models_favorite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_favorite_proto_goTypes = []interface{}{
	(FavoriteEntityType)(0), // 0: appcontext.models.FavoriteEntityType
	(*Favorite)(nil),        // 1: appcontext.models.Favorite
}
var file_models_favorite_proto_depIdxs = []int32{
	0, // 0: appcontext.models.Favorite.entity_type:type_name -> appcontext.models.FavoriteEntityType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_favorite_proto_init() }
func file_models_favorite_proto_init() {
	if File_models_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Favorite); i {
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
	file_models_favorite_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_favorite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_favorite_proto_goTypes,
		DependencyIndexes: file_models_favorite_proto_depIdxs,
		EnumInfos:         file_models_favorite_proto_enumTypes,
		MessageInfos:      file_models_favorite_proto_msgTypes,
	}.Build()
	File_models_favorite_proto = out.File
	file_models_favorite_proto_rawDesc = nil
	file_models_favorite_proto_goTypes = nil
	file_models_favorite_proto_depIdxs = nil
}