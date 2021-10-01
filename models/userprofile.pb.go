// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: models/userprofile.proto

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

type ProfileState int32

const (
	ProfileState_pending  ProfileState = 0
	ProfileState_active   ProfileState = 1
	ProfileState_inactive ProfileState = 2
)

// Enum value maps for ProfileState.
var (
	ProfileState_name = map[int32]string{
		0: "pending",
		1: "active",
		2: "inactive",
	}
	ProfileState_value = map[string]int32{
		"pending":  0,
		"active":   1,
		"inactive": 2,
	}
)

func (x ProfileState) Enum() *ProfileState {
	p := new(ProfileState)
	*p = x
	return p
}

func (x ProfileState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfileState) Descriptor() protoreflect.EnumDescriptor {
	return file_models_userprofile_proto_enumTypes[0].Descriptor()
}

func (ProfileState) Type() protoreflect.EnumType {
	return &file_models_userprofile_proto_enumTypes[0]
}

func (x ProfileState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfileState.Descriptor instead.
func (ProfileState) EnumDescriptor() ([]byte, []int) {
	return file_models_userprofile_proto_rawDescGZIP(), []int{0}
}

type Userprofile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *string       `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	TenantId     *string       `protobuf:"bytes,2,opt,name=tenantId,proto3,oneof" json:"tenantId,omitempty"`
	KratosId     *string       `protobuf:"bytes,3,opt,name=kratosId,proto3,oneof" json:"kratosId,omitempty"`
	City         *string       `protobuf:"bytes,4,opt,name=city,proto3,oneof" json:"city,omitempty"`
	Country      *string       `protobuf:"bytes,5,opt,name=country,proto3,oneof" json:"country,omitempty"`
	Department   *string       `protobuf:"bytes,6,opt,name=department,proto3,oneof" json:"department,omitempty"`
	Division     *string       `protobuf:"bytes,7,opt,name=division,proto3,oneof" json:"division,omitempty"`
	Email        *string       `protobuf:"bytes,8,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Firstname    *string       `protobuf:"bytes,9,opt,name=firstname,proto3,oneof" json:"firstname,omitempty"`
	Origin       *string       `protobuf:"bytes,10,opt,name=origin,proto3,oneof" json:"origin,omitempty"`
	OriginId     *string       `protobuf:"bytes,11,opt,name=originId,proto3,oneof" json:"originId,omitempty"`
	Lastname     *string       `protobuf:"bytes,12,opt,name=lastname,proto3,oneof" json:"lastname,omitempty"`
	ProfileState *ProfileState `protobuf:"varint,13,opt,name=profileState,proto3,enum=appcontextsvc_client.models.ProfileState,oneof" json:"profileState,omitempty"`
	PhoneNumber  *string       `protobuf:"bytes,14,opt,name=phoneNumber,proto3,oneof" json:"phoneNumber,omitempty"`
	Salutation   *string       `protobuf:"bytes,15,opt,name=salutation,proto3,oneof" json:"salutation,omitempty"`
	State        *string       `protobuf:"bytes,16,opt,name=state,proto3,oneof" json:"state,omitempty"`
	Street       *string       `protobuf:"bytes,17,opt,name=street,proto3,oneof" json:"street,omitempty"`
	Title        *string       `protobuf:"bytes,18,opt,name=title,proto3,oneof" json:"title,omitempty"`
	ZipCode      *string       `protobuf:"bytes,19,opt,name=zipCode,proto3,oneof" json:"zipCode,omitempty"`
}

func (x *Userprofile) Reset() {
	*x = Userprofile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_userprofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userprofile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userprofile) ProtoMessage() {}

func (x *Userprofile) ProtoReflect() protoreflect.Message {
	mi := &file_models_userprofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userprofile.ProtoReflect.Descriptor instead.
func (*Userprofile) Descriptor() ([]byte, []int) {
	return file_models_userprofile_proto_rawDescGZIP(), []int{0}
}

func (x *Userprofile) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Userprofile) GetTenantId() string {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return ""
}

func (x *Userprofile) GetKratosId() string {
	if x != nil && x.KratosId != nil {
		return *x.KratosId
	}
	return ""
}

func (x *Userprofile) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *Userprofile) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *Userprofile) GetDepartment() string {
	if x != nil && x.Department != nil {
		return *x.Department
	}
	return ""
}

func (x *Userprofile) GetDivision() string {
	if x != nil && x.Division != nil {
		return *x.Division
	}
	return ""
}

func (x *Userprofile) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Userprofile) GetFirstname() string {
	if x != nil && x.Firstname != nil {
		return *x.Firstname
	}
	return ""
}

func (x *Userprofile) GetOrigin() string {
	if x != nil && x.Origin != nil {
		return *x.Origin
	}
	return ""
}

func (x *Userprofile) GetOriginId() string {
	if x != nil && x.OriginId != nil {
		return *x.OriginId
	}
	return ""
}

func (x *Userprofile) GetLastname() string {
	if x != nil && x.Lastname != nil {
		return *x.Lastname
	}
	return ""
}

func (x *Userprofile) GetProfileState() ProfileState {
	if x != nil && x.ProfileState != nil {
		return *x.ProfileState
	}
	return ProfileState_pending
}

func (x *Userprofile) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *Userprofile) GetSalutation() string {
	if x != nil && x.Salutation != nil {
		return *x.Salutation
	}
	return ""
}

func (x *Userprofile) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

func (x *Userprofile) GetStreet() string {
	if x != nil && x.Street != nil {
		return *x.Street
	}
	return ""
}

func (x *Userprofile) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Userprofile) GetZipCode() string {
	if x != nil && x.ZipCode != nil {
		return *x.ZipCode
	}
	return ""
}

var File_models_userprofile_proto protoreflect.FileDescriptor

var file_models_userprofile_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x07, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa,
	0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0,
	0x01, 0x01, 0xd0, 0x01, 0x01, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x08, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x08, 0x72, 0x06,
	0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x08, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x52, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x48, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0d, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0e, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x10,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x48, 0x11, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x48, 0x12, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x61, 0x6c, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x2a,
	0x35, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x69, 0x6e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x10, 0x02, 0x42, 0x55, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x76, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x1e, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2e, 0x56, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_userprofile_proto_rawDescOnce sync.Once
	file_models_userprofile_proto_rawDescData = file_models_userprofile_proto_rawDesc
)

func file_models_userprofile_proto_rawDescGZIP() []byte {
	file_models_userprofile_proto_rawDescOnce.Do(func() {
		file_models_userprofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_userprofile_proto_rawDescData)
	})
	return file_models_userprofile_proto_rawDescData
}

var file_models_userprofile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_userprofile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_userprofile_proto_goTypes = []interface{}{
	(ProfileState)(0),   // 0: appcontextsvc_client.models.ProfileState
	(*Userprofile)(nil), // 1: appcontextsvc_client.models.Userprofile
}
var file_models_userprofile_proto_depIdxs = []int32{
	0, // 0: appcontextsvc_client.models.Userprofile.profileState:type_name -> appcontextsvc_client.models.ProfileState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_userprofile_proto_init() }
func file_models_userprofile_proto_init() {
	if File_models_userprofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_userprofile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userprofile); i {
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
	file_models_userprofile_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_userprofile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_userprofile_proto_goTypes,
		DependencyIndexes: file_models_userprofile_proto_depIdxs,
		EnumInfos:         file_models_userprofile_proto_enumTypes,
		MessageInfos:      file_models_userprofile_proto_msgTypes,
	}.Build()
	File_models_userprofile_proto = out.File
	file_models_userprofile_proto_rawDesc = nil
	file_models_userprofile_proto_goTypes = nil
	file_models_userprofile_proto_depIdxs = nil
}
