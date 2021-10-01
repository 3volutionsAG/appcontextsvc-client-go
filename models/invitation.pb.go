// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: models/invitation.proto

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

type State int32

const (
	State_unknown   State = 0
	State_confirmed State = 1
	State_expired   State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "unknown",
		1: "confirmed",
		2: "expired",
	}
	State_value = map[string]int32{
		"unknown":   0,
		"confirmed": 1,
		"expired":   2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_models_invitation_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_models_invitation_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_models_invitation_proto_rawDescGZIP(), []int{0}
}

type InvitationState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State State  `protobuf:"varint,2,opt,name=state,proto3,enum=appcontextsvc_client.models.State" json:"state,omitempty"`
}

func (x *InvitationState) Reset() {
	*x = InvitationState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_invitation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationState) ProtoMessage() {}

func (x *InvitationState) ProtoReflect() protoreflect.Message {
	mi := &file_models_invitation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationState.ProtoReflect.Descriptor instead.
func (*InvitationState) Descriptor() ([]byte, []int) {
	return file_models_invitation_proto_rawDescGZIP(), []int{0}
}

func (x *InvitationState) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvitationState) GetState() State {
	if x != nil {
		return x.State
	}
	return State_unknown
}

type Invitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserprofileId string `protobuf:"bytes,2,opt,name=userprofileId,proto3" json:"userprofileId,omitempty"`
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_invitation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_models_invitation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_models_invitation_proto_rawDescGZIP(), []int{1}
}

func (x *Invitation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invitation) GetUserprofileId() string {
	if x != nil {
		return x.UserprofileId
	}
	return ""
}

var File_models_invitation_proto protoreflect.FileDescriptor

var file_models_invitation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x70, 0x70, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5c, 0x0a, 0x0a, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08,
	0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x2a, 0x30, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x02, 0x42, 0x55, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa,
	0x02, 0x1e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_invitation_proto_rawDescOnce sync.Once
	file_models_invitation_proto_rawDescData = file_models_invitation_proto_rawDesc
)

func file_models_invitation_proto_rawDescGZIP() []byte {
	file_models_invitation_proto_rawDescOnce.Do(func() {
		file_models_invitation_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_invitation_proto_rawDescData)
	})
	return file_models_invitation_proto_rawDescData
}

var file_models_invitation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_invitation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_invitation_proto_goTypes = []interface{}{
	(State)(0),              // 0: appcontextsvc_client.models.State
	(*InvitationState)(nil), // 1: appcontextsvc_client.models.InvitationState
	(*Invitation)(nil),      // 2: appcontextsvc_client.models.Invitation
}
var file_models_invitation_proto_depIdxs = []int32{
	0, // 0: appcontextsvc_client.models.InvitationState.state:type_name -> appcontextsvc_client.models.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_invitation_proto_init() }
func file_models_invitation_proto_init() {
	if File_models_invitation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_invitation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationState); i {
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
		file_models_invitation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invitation); i {
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
			RawDescriptor: file_models_invitation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_invitation_proto_goTypes,
		DependencyIndexes: file_models_invitation_proto_depIdxs,
		EnumInfos:         file_models_invitation_proto_enumTypes,
		MessageInfos:      file_models_invitation_proto_msgTypes,
	}.Build()
	File_models_invitation_proto = out.File
	file_models_invitation_proto_rawDesc = nil
	file_models_invitation_proto_goTypes = nil
	file_models_invitation_proto_depIdxs = nil
}
