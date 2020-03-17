// Code generated by protoc-gen-go. DO NOT EDIT.
// source: matcha.proto

package matcha

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User_Gender int32

const (
	User_male   User_Gender = 0
	User_female User_Gender = 1
	User_both   User_Gender = 2
)

var User_Gender_name = map[int32]string{
	0: "male",
	1: "female",
	2: "both",
}

var User_Gender_value = map[string]int32{
	"male":   0,
	"female": 1,
	"both":   2,
}

func (x User_Gender) String() string {
	return proto.EnumName(User_Gender_name, int32(x))
}

func (User_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{7, 0}
}

type ResetPassRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	NewPass              string   `protobuf:"bytes,2,opt,name=newPass,proto3" json:"newPass,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPassRequest) Reset()         { *m = ResetPassRequest{} }
func (m *ResetPassRequest) String() string { return proto.CompactTextString(m) }
func (*ResetPassRequest) ProtoMessage()    {}
func (*ResetPassRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{0}
}

func (m *ResetPassRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPassRequest.Unmarshal(m, b)
}
func (m *ResetPassRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPassRequest.Marshal(b, m, deterministic)
}
func (m *ResetPassRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPassRequest.Merge(m, src)
}
func (m *ResetPassRequest) XXX_Size() int {
	return xxx_messageInfo_ResetPassRequest.Size(m)
}
func (m *ResetPassRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPassRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPassRequest proto.InternalMessageInfo

func (m *ResetPassRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ResetPassRequest) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

func (m *ResetPassRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type SendEmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailRequest) Reset()         { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()    {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{1}
}

func (m *SendEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailRequest.Unmarshal(m, b)
}
func (m *SendEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailRequest.Marshal(b, m, deterministic)
}
func (m *SendEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailRequest.Merge(m, src)
}
func (m *SendEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendEmailRequest.Size(m)
}
func (m *SendEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailRequest proto.InternalMessageInfo

func (m *SendEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type VerifyRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{3}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VerifyRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ImageData struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageData) Reset()         { *m = ImageData{} }
func (m *ImageData) String() string { return proto.CompactTextString(m) }
func (*ImageData) ProtoMessage()    {}
func (*ImageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{4}
}

func (m *ImageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageData.Unmarshal(m, b)
}
func (m *ImageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageData.Marshal(b, m, deterministic)
}
func (m *ImageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageData.Merge(m, src)
}
func (m *ImageData) XXX_Size() int {
	return xxx_messageInfo_ImageData.Size(m)
}
func (m *ImageData) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageData.DiscardUnknown(m)
}

var xxx_messageInfo_ImageData proto.InternalMessageInfo

func (m *ImageData) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{5}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type LoginRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{6}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type User struct {
	UserName             string      `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	FirstName            string      `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string      `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Password             string      `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email                string      `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Gender               User_Gender `protobuf:"varint,6,opt,name=gender,proto3,enum=matcha.User_Gender" json:"gender,omitempty"`
	Preference           User_Gender `protobuf:"varint,7,opt,name=preference,proto3,enum=matcha.User_Gender" json:"preference,omitempty"`
	Bio                  string      `protobuf:"bytes,8,opt,name=bio,proto3" json:"bio,omitempty"`
	Hash                 string      `protobuf:"bytes,9,opt,name=hash,proto3" json:"hash,omitempty"`
	Tags                 []string    `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{7}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetGender() User_Gender {
	if m != nil {
		return m.Gender
	}
	return User_male
}

func (m *User) GetPreference() User_Gender {
	if m != nil {
		return m.Preference
	}
	return User_male
}

func (m *User) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *User) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *User) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Reply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb514d74c502fc4f, []int{8}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("matcha.User_Gender", User_Gender_name, User_Gender_value)
	proto.RegisterType((*ResetPassRequest)(nil), "matcha.resetPassRequest")
	proto.RegisterType((*SendEmailRequest)(nil), "matcha.sendEmailRequest")
	proto.RegisterType((*CreateRequest)(nil), "matcha.createRequest")
	proto.RegisterType((*VerifyRequest)(nil), "matcha.verifyRequest")
	proto.RegisterType((*ImageData)(nil), "matcha.imageData")
	proto.RegisterType((*Empty)(nil), "matcha.empty")
	proto.RegisterType((*LoginRequest)(nil), "matcha.loginRequest")
	proto.RegisterType((*User)(nil), "matcha.User")
	proto.RegisterType((*Reply)(nil), "matcha.Reply")
}

func init() { proto.RegisterFile("matcha.proto", fileDescriptor_bb514d74c502fc4f) }

var fileDescriptor_bb514d74c502fc4f = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcb, 0x8e, 0xd3, 0x3c,
	0x14, 0xfe, 0x93, 0xb6, 0x69, 0x73, 0xd4, 0x8e, 0xfa, 0x9b, 0x41, 0x8a, 0x2a, 0x16, 0x9d, 0x08,
	0x50, 0x11, 0x52, 0x35, 0xd3, 0x61, 0x01, 0x4b, 0x24, 0x2e, 0x3b, 0x84, 0xc2, 0x65, 0xef, 0xa6,
	0x27, 0x69, 0xa4, 0x24, 0x0e, 0xb6, 0xcb, 0xa8, 0x3b, 0xde, 0x82, 0xe7, 0xe1, 0x1d, 0x78, 0x20,
	0x64, 0x3b, 0x4e, 0x93, 0xc0, 0x8c, 0x58, 0xb2, 0x3b, 0x97, 0xef, 0x3b, 0x97, 0xf8, 0x7c, 0x81,
	0x69, 0x41, 0x65, 0xbc, 0xa7, 0xeb, 0x8a, 0x33, 0xc9, 0x88, 0x67, 0xbc, 0xf0, 0x33, 0xcc, 0x39,
	0x0a, 0x94, 0xef, 0xa9, 0x10, 0x11, 0x7e, 0x39, 0xa0, 0x90, 0xe4, 0x1c, 0x46, 0x58, 0xd0, 0x2c,
	0x0f, 0x9c, 0xa5, 0xb3, 0xf2, 0x23, 0xe3, 0x90, 0x00, 0xc6, 0x25, 0xde, 0x28, 0x5c, 0xe0, 0xea,
	0xb8, 0x75, 0x09, 0x81, 0xe1, 0x9e, 0x8a, 0x7d, 0x30, 0xd0, 0x61, 0x6d, 0x87, 0x2b, 0x98, 0x0b,
	0x2c, 0x77, 0xaf, 0x15, 0xf5, 0xce, 0xba, 0xe1, 0x77, 0x07, 0x66, 0x31, 0x47, 0x2a, 0xd1, 0xe2,
	0x16, 0x30, 0x39, 0x08, 0xe4, 0xef, 0x68, 0x81, 0x35, 0xb4, 0xf1, 0xc9, 0x03, 0xf0, 0x93, 0x8c,
	0x0b, 0xa9, 0x93, 0x66, 0x8e, 0x53, 0x40, 0x31, 0x73, 0x5a, 0x27, 0xcd, 0x34, 0x8d, 0xaf, 0x72,
	0x15, 0x15, 0xe2, 0x86, 0xf1, 0x5d, 0x30, 0x34, 0x39, 0xeb, 0x9f, 0x26, 0x1b, 0xb5, 0x27, 0x7b,
	0x01, 0xb3, 0xaf, 0xc8, 0xb3, 0xe4, 0x78, 0xf7, 0x87, 0xb1, 0xeb, 0xbb, 0xad, 0xf5, 0x2f, 0xc0,
	0xcf, 0x0a, 0x9a, 0xe2, 0x2b, 0x2a, 0xa9, 0xa2, 0x69, 0x47, 0xd3, 0xa6, 0x91, 0x71, 0xc2, 0xb1,
	0x2a, 0x56, 0xc9, 0x63, 0xf8, 0x06, 0xa6, 0x39, 0x4b, 0xb3, 0xf2, 0x6f, 0xd6, 0x6f, 0x2f, 0xe1,
	0x76, 0x97, 0x08, 0x7f, 0xba, 0x30, 0xfc, 0x24, 0x90, 0xff, 0x2b, 0xdf, 0x8f, 0x3c, 0x05, 0x2f,
	0xc5, 0x72, 0x87, 0x3c, 0xf0, 0x96, 0xce, 0xea, 0x6c, 0x73, 0x6f, 0x5d, 0x9f, 0xa0, 0x9a, 0x72,
	0xfd, 0x56, 0xa7, 0xa2, 0x1a, 0x42, 0xae, 0x01, 0x2a, 0x8e, 0x09, 0x72, 0x2c, 0x63, 0x0c, 0xc6,
	0xb7, 0x13, 0x5a, 0x30, 0x32, 0x87, 0xc1, 0x36, 0x63, 0xc1, 0x44, 0x77, 0x55, 0x66, 0xf3, 0x18,
	0xfe, 0xe9, 0x31, 0x54, 0x4c, 0xd2, 0x54, 0x04, 0xb0, 0x1c, 0xa8, 0x98, 0xb2, 0xc3, 0x15, 0x78,
	0xa6, 0x1e, 0x99, 0xc0, 0xb0, 0xa0, 0x39, 0xce, 0xff, 0x23, 0x00, 0x5e, 0x82, 0xda, 0x76, 0x54,
	0x74, 0xcb, 0xe4, 0x7e, 0xee, 0x86, 0x17, 0x30, 0x8a, 0xb0, 0xca, 0x8f, 0x4a, 0x00, 0x05, 0x0a,
	0x61, 0x1f, 0xd2, 0x8f, 0xac, 0xbb, 0xf9, 0xe6, 0xc0, 0x59, 0xc2, 0x78, 0xca, 0xb4, 0x8c, 0xf4,
	0x17, 0x79, 0x06, 0x7e, 0x73, 0xff, 0x24, 0xb0, 0x7b, 0xf4, 0x25, 0xb1, 0x98, 0xd9, 0x8c, 0x69,
	0xf1, 0x1c, 0x66, 0x8d, 0x1a, 0x75, 0x99, 0x86, 0xd9, 0x17, 0x69, 0x8f, 0xb9, 0x61, 0x56, 0x44,
	0x2f, 0xe3, 0x98, 0x1d, 0x4a, 0x49, 0xd6, 0xe0, 0x99, 0x00, 0xb9, 0x6f, 0x91, 0x1d, 0x95, 0xf5,
	0x5b, 0xaf, 0xc1, 0x33, 0xc7, 0x7e, 0xc2, 0x77, 0x8e, 0xbf, 0xdf, 0xf0, 0x87, 0x0b, 0x63, 0xdb,
	0xeb, 0x31, 0x8c, 0x53, 0x94, 0x1f, 0x30, 0x4f, 0x48, 0x83, 0xd2, 0xb7, 0xbd, 0x98, 0xb6, 0x5f,
	0x90, 0x3c, 0xd2, 0x38, 0x6d, 0x76, 0x12, 0x3d, 0xd8, 0x15, 0x4c, 0x6a, 0x98, 0xb8, 0x6d, 0xf8,
	0x0e, 0xe1, 0xd2, 0x21, 0x97, 0xe0, 0x6b, 0x0d, 0x69, 0xfe, 0xb9, 0x4d, 0xb6, 0x65, 0xd5, 0xdf,
	0xf7, 0x09, 0xc0, 0xa1, 0xda, 0x51, 0x89, 0x7f, 0x18, 0xa7, 0x07, 0xbd, 0xaa, 0xc5, 0xfc, 0x51,
	0xa9, 0xf3, 0x7f, 0x9b, 0x6b, 0xf4, 0xbd, 0xf8, 0x3d, 0x44, 0x1e, 0xc2, 0x50, 0x2a, 0x74, 0xb7,
	0x52, 0xaf, 0xf0, 0xd6, 0xd3, 0xff, 0xe2, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0xbe,
	0xb6, 0x2a, 0x9b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ForgotPasswordClient is the client API for ForgotPassword service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForgotPasswordClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*Reply, error)
	ResetPassword(ctx context.Context, in *ResetPassRequest, opts ...grpc.CallOption) (*Reply, error)
}

type forgotPasswordClient struct {
	cc *grpc.ClientConn
}

func NewForgotPasswordClient(cc *grpc.ClientConn) ForgotPasswordClient {
	return &forgotPasswordClient{cc}
}

func (c *forgotPasswordClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.forgotPassword/sendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forgotPasswordClient) ResetPassword(ctx context.Context, in *ResetPassRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.forgotPassword/resetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForgotPasswordServer is the server API for ForgotPassword service.
type ForgotPasswordServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*Reply, error)
	ResetPassword(context.Context, *ResetPassRequest) (*Reply, error)
}

// UnimplementedForgotPasswordServer can be embedded to have forward compatible implementations.
type UnimplementedForgotPasswordServer struct {
}

func (*UnimplementedForgotPasswordServer) SendEmail(ctx context.Context, req *SendEmailRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (*UnimplementedForgotPasswordServer) ResetPassword(ctx context.Context, req *ResetPassRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}

func RegisterForgotPasswordServer(s *grpc.Server, srv ForgotPasswordServer) {
	s.RegisterService(&_ForgotPassword_serviceDesc, srv)
}

func _ForgotPassword_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForgotPasswordServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.forgotPassword/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForgotPasswordServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForgotPassword_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForgotPasswordServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.forgotPassword/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForgotPasswordServer).ResetPassword(ctx, req.(*ResetPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ForgotPassword_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matcha.forgotPassword",
	HandlerType: (*ForgotPasswordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendEmail",
			Handler:    _ForgotPassword_SendEmail_Handler,
		},
		{
			MethodName: "resetPassword",
			Handler:    _ForgotPassword_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matcha.proto",
}

// CreateAccountClient is the client API for CreateAccount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateAccountClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Reply, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*Reply, error)
}

type createAccountClient struct {
	cc *grpc.ClientConn
}

func NewCreateAccountClient(cc *grpc.ClientConn) CreateAccountClient {
	return &createAccountClient{cc}
}

func (c *createAccountClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.createAccount/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createAccountClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.createAccount/verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateAccountServer is the server API for CreateAccount service.
type CreateAccountServer interface {
	Create(context.Context, *CreateRequest) (*Reply, error)
	Verify(context.Context, *VerifyRequest) (*Reply, error)
}

// UnimplementedCreateAccountServer can be embedded to have forward compatible implementations.
type UnimplementedCreateAccountServer struct {
}

func (*UnimplementedCreateAccountServer) Create(ctx context.Context, req *CreateRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCreateAccountServer) Verify(ctx context.Context, req *VerifyRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterCreateAccountServer(s *grpc.Server, srv CreateAccountServer) {
	s.RegisterService(&_CreateAccount_serviceDesc, srv)
}

func _CreateAccount_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateAccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.createAccount/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateAccountServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateAccount_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateAccountServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.createAccount/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateAccountServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateAccount_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matcha.createAccount",
	HandlerType: (*CreateAccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CreateAccount_Create_Handler,
		},
		{
			MethodName: "verify",
			Handler:    _CreateAccount_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matcha.proto",
}

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	GetSelf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (Account_GetUsersClient, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Reply, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Reply, error)
	ImageTest(ctx context.Context, in *ImageData, opts ...grpc.CallOption) (*ImageData, error)
	Test(ctx context.Context, in *Reply, opts ...grpc.CallOption) (*Reply, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetSelf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/matcha.Account/getSelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/matcha.Account/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUsers(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (Account_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Account_serviceDesc.Streams[0], "/matcha.Account/getUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Account_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type accountGetUsersClient struct {
	grpc.ClientStream
}

func (x *accountGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.Account/loginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.Account/updateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ImageTest(ctx context.Context, in *ImageData, opts ...grpc.CallOption) (*ImageData, error) {
	out := new(ImageData)
	err := c.cc.Invoke(ctx, "/matcha.Account/imageTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Test(ctx context.Context, in *Reply, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/matcha.Account/test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	GetSelf(context.Context, *Empty) (*User, error)
	GetUser(context.Context, *User) (*User, error)
	GetUsers(*CreateRequest, Account_GetUsersServer) error
	LoginUser(context.Context, *LoginRequest) (*Reply, error)
	UpdateUser(context.Context, *User) (*Reply, error)
	ImageTest(context.Context, *ImageData) (*ImageData, error)
	Test(context.Context, *Reply) (*Reply, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) GetSelf(ctx context.Context, req *Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelf not implemented")
}
func (*UnimplementedAccountServer) GetUser(ctx context.Context, req *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedAccountServer) GetUsers(req *CreateRequest, srv Account_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedAccountServer) LoginUser(ctx context.Context, req *LoginRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedAccountServer) UpdateUser(ctx context.Context, req *User) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedAccountServer) ImageTest(ctx context.Context, req *ImageData) (*ImageData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageTest not implemented")
}
func (*UnimplementedAccountServer) Test(ctx context.Context, req *Reply) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_GetSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.Account/GetSelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetSelf(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.Account/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountServer).GetUsers(m, &accountGetUsersServer{stream})
}

type Account_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type accountGetUsersServer struct {
	grpc.ServerStream
}

func (x *accountGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _Account_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.Account/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.Account/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_ImageTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ImageTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.Account/ImageTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ImageTest(ctx, req.(*ImageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matcha.Account/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Test(ctx, req.(*Reply))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matcha.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSelf",
			Handler:    _Account_GetSelf_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _Account_GetUser_Handler,
		},
		{
			MethodName: "loginUser",
			Handler:    _Account_LoginUser_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _Account_UpdateUser_Handler,
		},
		{
			MethodName: "imageTest",
			Handler:    _Account_ImageTest_Handler,
		},
		{
			MethodName: "test",
			Handler:    _Account_Test_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getUsers",
			Handler:       _Account_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "matcha.proto",
}
