// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: models/permission.proto

package models

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PermissionNamespace int32

const (
	PermissionNamespace_roles        PermissionNamespace = 0
	PermissionNamespace_userprofiles PermissionNamespace = 1
	PermissionNamespace_resources    PermissionNamespace = 2
)

// Enum value maps for PermissionNamespace.
var (
	PermissionNamespace_name = map[int32]string{
		0: "roles",
		1: "userprofiles",
		2: "resources",
	}
	PermissionNamespace_value = map[string]int32{
		"roles":        0,
		"userprofiles": 1,
		"resources":    2,
	}
)

func (x PermissionNamespace) Enum() *PermissionNamespace {
	p := new(PermissionNamespace)
	*p = x
	return p
}

func (x PermissionNamespace) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionNamespace) Descriptor() protoreflect.EnumDescriptor {
	return file_models_permission_proto_enumTypes[0].Descriptor()
}

func (PermissionNamespace) Type() protoreflect.EnumType {
	return &file_models_permission_proto_enumTypes[0]
}

func (x PermissionNamespace) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionNamespace.Descriptor instead.
func (PermissionNamespace) EnumDescriptor() ([]byte, []int) {
	return file_models_permission_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *string              `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Namespace   *PermissionNamespace `protobuf:"varint,3,opt,name=namespace,proto3,enum=appcontext.models.PermissionNamespace,oneof" json:"namespace,omitempty"`
	Relation    *string              `protobuf:"bytes,4,opt,name=relation,proto3,oneof" json:"relation,omitempty"`
	Caption     *string              `protobuf:"bytes,5,opt,name=caption,proto3,oneof" json:"caption,omitempty"`
	Description *string              `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ManifestId  string               `protobuf:"bytes,7,opt,name=manifestId,proto3" json:"manifestId,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_models_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_models_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Permission) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Permission) GetNamespace() PermissionNamespace {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return PermissionNamespace_roles
}

func (x *Permission) GetRelation() string {
	if x != nil && x.Relation != nil {
		return *x.Relation
	}
	return ""
}

func (x *Permission) GetCaption() string {
	if x != nil && x.Caption != nil {
		return *x.Caption
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Permission) GetManifestId() string {
	if x != nil {
		return x.ManifestId
	}
	return ""
}

var File_models_permission_proto protoreflect.FileDescriptor

var file_models_permission_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x70, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0xe0, 0x41,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x1e, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x49, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x48, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x1e, 0x48, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x28, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x48, 0x03, 0x52, 0x07,
	0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xf4, 0x03, 0x48, 0x04, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a,
	0x0a, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x41, 0x0a,
	0x13, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x10, 0x02,
	0x42, 0x55, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x1e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x56, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_permission_proto_rawDescOnce sync.Once
	file_models_permission_proto_rawDescData = file_models_permission_proto_rawDesc
)

func file_models_permission_proto_rawDescGZIP() []byte {
	file_models_permission_proto_rawDescOnce.Do(func() {
		file_models_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_permission_proto_rawDescData)
	})
	return file_models_permission_proto_rawDescData
}

var file_models_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_permission_proto_goTypes = []interface{}{
	(PermissionNamespace)(0), // 0: appcontext.models.PermissionNamespace
	(*Permission)(nil),       // 1: appcontext.models.Permission
}
var file_models_permission_proto_depIdxs = []int32{
	0, // 0: appcontext.models.Permission.namespace:type_name -> appcontext.models.PermissionNamespace
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_permission_proto_init() }
func file_models_permission_proto_init() {
	if File_models_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
	file_models_permission_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_permission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_permission_proto_goTypes,
		DependencyIndexes: file_models_permission_proto_depIdxs,
		EnumInfos:         file_models_permission_proto_enumTypes,
		MessageInfos:      file_models_permission_proto_msgTypes,
	}.Build()
	File_models_permission_proto = out.File
	file_models_permission_proto_rawDesc = nil
	file_models_permission_proto_goTypes = nil
	file_models_permission_proto_depIdxs = nil
}
