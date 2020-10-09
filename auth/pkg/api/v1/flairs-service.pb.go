// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: flairs-service.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserAccountType int32

const (
	User_PERSONAL UserAccountType = 0
	User_BUSINESS UserAccountType = 1
)

// Enum value maps for UserAccountType.
var (
	UserAccountType_name = map[int32]string{
		0: "PERSONAL",
		1: "BUSINESS",
	}
	UserAccountType_value = map[string]int32{
		"PERSONAL": 0,
		"BUSINESS": 1,
	}
)

func (x UserAccountType) Enum() *UserAccountType {
	p := new(UserAccountType)
	*p = x
	return p
}

func (x UserAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_flairs_service_proto_enumTypes[0].Descriptor()
}

func (UserAccountType) Type() protoreflect.EnumType {
	return &file_flairs_service_proto_enumTypes[0]
}

func (x UserAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserAccountType.Descriptor instead.
func (UserAccountType) EnumDescriptor() ([]byte, []int) {
	return file_flairs_service_proto_rawDescGZIP(), []int{0, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                   string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string               `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	DOB                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=DOB,proto3" json:"DOB,omitempty"`
	Gender               string               `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	BVN                  string               `protobuf:"bytes,6,opt,name=BVN,proto3" json:"BVN,omitempty"`
	Address              string               `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Street               string               `protobuf:"bytes,8,opt,name=street,proto3" json:"street,omitempty"`
	City                 string               `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	PostalCode           string               `protobuf:"bytes,10,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	State                string               `protobuf:"bytes,11,opt,name=state,proto3" json:"state,omitempty"`
	Country              string               `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	CountryID            uint32               `protobuf:"varint,13,opt,name=countryID,proto3" json:"countryID,omitempty"`
	Photo                string               `protobuf:"bytes,900,opt,name=photo,proto3" json:"photo,omitempty"`
	Passport             string               `protobuf:"bytes,901,opt,name=passport,proto3" json:"passport,omitempty"`
	IDCard               string               `protobuf:"bytes,902,opt,name=IDCard,proto3" json:"IDCard,omitempty"`
	Referrer             string               `protobuf:"bytes,14,opt,name=referrer,proto3" json:"referrer,omitempty"`
	RefCode              string               `protobuf:"bytes,15,opt,name=refCode,proto3" json:"refCode,omitempty"`
	HowDidYouHearAboutUs string               `protobuf:"bytes,16,opt,name=howDidYouHearAboutUs,proto3" json:"howDidYouHearAboutUs,omitempty"`
	UserName             string               `protobuf:"bytes,17,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string               `protobuf:"bytes,18,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerifiedAt      *timestamp.Timestamp `protobuf:"bytes,19,opt,name=emailVerifiedAt,proto3" json:"emailVerifiedAt,omitempty"`
	Password             []byte               `protobuf:"bytes,20,opt,name=password,proto3" json:"password,omitempty"`
	Pin                  []byte               `protobuf:"bytes,21,opt,name=pin,proto3" json:"pin,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,22,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	PhoneVerifiedAt      *timestamp.Timestamp `protobuf:"bytes,23,opt,name=phoneVerifiedAt,proto3" json:"phoneVerifiedAt,omitempty"`
	ACCOUNT_TYPE         UserAccountType      `protobuf:"varint,24,opt,name=ACCOUNT_TYPE,json=ACCOUNTTYPE,proto3,enum=v1.UserAccountType" json:"ACCOUNT_TYPE,omitempty"`
	LastCardRequested    string               `protobuf:"bytes,25,opt,name=lastCardRequested,proto3" json:"lastCardRequested,omitempty"`
	IsProfileCompleted   bool                 `protobuf:"varint,26,opt,name=isProfileCompleted,proto3" json:"isProfileCompleted,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,29,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,30,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_flairs_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetDOB() *timestamp.Timestamp {
	if x != nil {
		return x.DOB
	}
	return nil
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetBVN() string {
	if x != nil {
		return x.BVN
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *User) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *User) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *User) GetCountryID() uint32 {
	if x != nil {
		return x.CountryID
	}
	return 0
}

func (x *User) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *User) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *User) GetIDCard() string {
	if x != nil {
		return x.IDCard
	}
	return ""
}

func (x *User) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *User) GetRefCode() string {
	if x != nil {
		return x.RefCode
	}
	return ""
}

func (x *User) GetHowDidYouHearAboutUs() string {
	if x != nil {
		return x.HowDidYouHearAboutUs
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerifiedAt() *timestamp.Timestamp {
	if x != nil {
		return x.EmailVerifiedAt
	}
	return nil
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *User) GetPin() []byte {
	if x != nil {
		return x.Pin
	}
	return nil
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetPhoneVerifiedAt() *timestamp.Timestamp {
	if x != nil {
		return x.PhoneVerifiedAt
	}
	return nil
}

func (x *User) GetACCOUNT_TYPE() UserAccountType {
	if x != nil {
		return x.ACCOUNT_TYPE
	}
	return User_PERSONAL
}

func (x *User) GetLastCardRequested() string {
	if x != nil {
		return x.LastCardRequested
	}
	return ""
}

func (x *User) GetIsProfileCompleted() bool {
	if x != nil {
		return x.IsProfileCompleted
	}
	return false
}

func (x *User) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AddNewUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Email of user to be added
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AddNewUserRequest) Reset() {
	*x = AddNewUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewUserRequest) ProtoMessage() {}

func (x *AddNewUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewUserRequest.ProtoReflect.Descriptor instead.
func (*AddNewUserRequest) Descriptor() ([]byte, []int) {
	return file_flairs_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddNewUserRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *AddNewUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddNewUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of the created user
	ID string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AddNewUserResponse) Reset() {
	*x = AddNewUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewUserResponse) ProtoMessage() {}

func (x *AddNewUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewUserResponse.ProtoReflect.Descriptor instead.
func (*AddNewUserResponse) Descriptor() ([]byte, []int) {
	return file_flairs_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddNewUserResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *AddNewUserResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_flairs_service_proto protoreflect.FileDescriptor

var file_flairs_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6c, 0x61, 0x69, 0x72, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x44, 0x4f,
	0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x03, 0x44, 0x4f, 0x42, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x42, 0x56, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x42,
	0x56, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x44, 0x12, 0x15, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18,
	0x84, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x85, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x49, 0x44,
	0x43, 0x61, 0x72, 0x64, 0x18, 0x86, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x44, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x66, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x68, 0x6f, 0x77,
	0x44, 0x69, 0x64, 0x59, 0x6f, 0x75, 0x48, 0x65, 0x61, 0x72, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x55,
	0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x68, 0x6f, 0x77, 0x44, 0x69, 0x64, 0x59,
	0x6f, 0x75, 0x48, 0x65, 0x61, 0x72, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x44, 0x0a, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x70, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x54, 0x59, 0x50, 0x45, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x69, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x29, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41,
	0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x22, 0x3b, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x36,
	0x0a, 0x12, 0x61, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0x69, 0x0a, 0x0d, 0x46, 0x6c, 0x61, 0x69, 0x72, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01,
	0x2a, 0x42, 0xe5, 0x01, 0x92, 0x41, 0xe1, 0x01, 0x12, 0x6a, 0x0a, 0x0e, 0x46, 0x6c, 0x61, 0x69,
	0x72, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x53, 0x0a, 0x1c, 0x66, 0x61,
	0x6c, 0x69, 0x72, 0x73, 0x20, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x20, 0x62, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x68, 0x74, 0x74, 0x70,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x70, 0x6c, 0x75, 0x73,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x1a, 0x14, 0x64, 0x61, 0x6d, 0x6d, 0x79,
	0x64, 0x61, 0x72, 0x6d, 0x79, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a,
	0x39, 0x30, 0x30, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03,
	0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_flairs_service_proto_rawDescOnce sync.Once
	file_flairs_service_proto_rawDescData = file_flairs_service_proto_rawDesc
)

func file_flairs_service_proto_rawDescGZIP() []byte {
	file_flairs_service_proto_rawDescOnce.Do(func() {
		file_flairs_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_flairs_service_proto_rawDescData)
	})
	return file_flairs_service_proto_rawDescData
}

var file_flairs_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flairs_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_flairs_service_proto_goTypes = []interface{}{
	(UserAccountType)(0),        // 0: v1.User.accountType
	(*User)(nil),                // 1: v1.User
	(*AddNewUserRequest)(nil),   // 2: v1.addNewUserRequest
	(*AddNewUserResponse)(nil),  // 3: v1.addNewUserResponse
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_flairs_service_proto_depIdxs = []int32{
	4, // 0: v1.User.DOB:type_name -> google.protobuf.Timestamp
	4, // 1: v1.User.emailVerifiedAt:type_name -> google.protobuf.Timestamp
	4, // 2: v1.User.phoneVerifiedAt:type_name -> google.protobuf.Timestamp
	0, // 3: v1.User.ACCOUNT_TYPE:type_name -> v1.User.accountType
	4, // 4: v1.User.createdAt:type_name -> google.protobuf.Timestamp
	4, // 5: v1.User.updatedAt:type_name -> google.protobuf.Timestamp
	2, // 6: v1.FlairsService.AddNewUser:input_type -> v1.addNewUserRequest
	3, // 7: v1.FlairsService.AddNewUser:output_type -> v1.addNewUserResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_flairs_service_proto_init() }
func file_flairs_service_proto_init() {
	if File_flairs_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flairs_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_flairs_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewUserRequest); i {
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
		file_flairs_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewUserResponse); i {
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
			RawDescriptor: file_flairs_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flairs_service_proto_goTypes,
		DependencyIndexes: file_flairs_service_proto_depIdxs,
		EnumInfos:         file_flairs_service_proto_enumTypes,
		MessageInfos:      file_flairs_service_proto_msgTypes,
	}.Build()
	File_flairs_service_proto = out.File
	file_flairs_service_proto_rawDesc = nil
	file_flairs_service_proto_goTypes = nil
	file_flairs_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FlairsServiceClient is the client API for FlairsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlairsServiceClient interface {
	// Create a new user
	AddNewUser(ctx context.Context, in *AddNewUserRequest, opts ...grpc.CallOption) (*AddNewUserResponse, error)
}

type flairsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlairsServiceClient(cc grpc.ClientConnInterface) FlairsServiceClient {
	return &flairsServiceClient{cc}
}

func (c *flairsServiceClient) AddNewUser(ctx context.Context, in *AddNewUserRequest, opts ...grpc.CallOption) (*AddNewUserResponse, error) {
	out := new(AddNewUserResponse)
	err := c.cc.Invoke(ctx, "/v1.FlairsService/AddNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlairsServiceServer is the server API for FlairsService service.
type FlairsServiceServer interface {
	// Create a new user
	AddNewUser(context.Context, *AddNewUserRequest) (*AddNewUserResponse, error)
}

// UnimplementedFlairsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlairsServiceServer struct {
}

func (*UnimplementedFlairsServiceServer) AddNewUser(context.Context, *AddNewUserRequest) (*AddNewUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewUser not implemented")
}

func RegisterFlairsServiceServer(s *grpc.Server, srv FlairsServiceServer) {
	s.RegisterService(&_FlairsService_serviceDesc, srv)
}

func _FlairsService_AddNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlairsServiceServer).AddNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FlairsService/AddNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlairsServiceServer).AddNewUser(ctx, req.(*AddNewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlairsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FlairsService",
	HandlerType: (*FlairsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewUser",
			Handler:    _FlairsService_AddNewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flairs-service.proto",
}
