// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: models/asset_seating.proto

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

type AssetSeating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId   string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	SeatingId string `protobuf:"bytes,2,opt,name=seating_id,json=seatingId,proto3" json:"seating_id,omitempty"`
	Capacity  *int32 `protobuf:"varint,3,opt,name=capacity,proto3,oneof" json:"capacity,omitempty"`
	IsDefault *bool  `protobuf:"varint,4,opt,name=is_default,json=isDefault,proto3,oneof" json:"is_default,omitempty"`
}

func (x *AssetSeating) Reset() {
	*x = AssetSeating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_asset_seating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSeating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSeating) ProtoMessage() {}

func (x *AssetSeating) ProtoReflect() protoreflect.Message {
	mi := &file_models_asset_seating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSeating.ProtoReflect.Descriptor instead.
func (*AssetSeating) Descriptor() ([]byte, []int) {
	return file_models_asset_seating_proto_rawDescGZIP(), []int{0}
}

func (x *AssetSeating) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *AssetSeating) GetSeatingId() string {
	if x != nil {
		return x.SeatingId
	}
	return ""
}

func (x *AssetSeating) GetCapacity() int32 {
	if x != nil && x.Capacity != nil {
		return *x.Capacity
	}
	return 0
}

func (x *AssetSeating) GetIsDefault() bool {
	if x != nil && x.IsDefault != nil {
		return *x.IsDefault
	}
	return false
}

var File_models_asset_seating_proto protoreflect.FileDescriptor

var file_models_asset_seating_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0c, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08,
	0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0,
	0x01, 0x01, 0x52, 0x09, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x01, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x55,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x1e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_asset_seating_proto_rawDescOnce sync.Once
	file_models_asset_seating_proto_rawDescData = file_models_asset_seating_proto_rawDesc
)

func file_models_asset_seating_proto_rawDescGZIP() []byte {
	file_models_asset_seating_proto_rawDescOnce.Do(func() {
		file_models_asset_seating_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_asset_seating_proto_rawDescData)
	})
	return file_models_asset_seating_proto_rawDescData
}

var file_models_asset_seating_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_asset_seating_proto_goTypes = []interface{}{
	(*AssetSeating)(nil), // 0: appcontext.models.AssetSeating
}
var file_models_asset_seating_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_asset_seating_proto_init() }
func file_models_asset_seating_proto_init() {
	if File_models_asset_seating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_asset_seating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSeating); i {
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
	file_models_asset_seating_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_asset_seating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_asset_seating_proto_goTypes,
		DependencyIndexes: file_models_asset_seating_proto_depIdxs,
		MessageInfos:      file_models_asset_seating_proto_msgTypes,
	}.Build()
	File_models_asset_seating_proto = out.File
	file_models_asset_seating_proto_rawDesc = nil
	file_models_asset_seating_proto_goTypes = nil
	file_models_asset_seating_proto_depIdxs = nil
}