// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_updmmclientusers.proto

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TUpdMmclientUsers struct {
	StatTime             string   `protobuf:"bytes,1,opt,name=stat_time,json=statTime,proto3" json:"stat_time,omitempty"`
	Msisdn               string   `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TUpdMmclientUsers) Reset()         { *m = TUpdMmclientUsers{} }
func (m *TUpdMmclientUsers) String() string { return proto.CompactTextString(m) }
func (*TUpdMmclientUsers) ProtoMessage()    {}
func (*TUpdMmclientUsers) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_updmmclientusers_96ad50160ac3f30f, []int{0}
}
func (m *TUpdMmclientUsers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TUpdMmclientUsers.Unmarshal(m, b)
}
func (m *TUpdMmclientUsers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TUpdMmclientUsers.Marshal(b, m, deterministic)
}
func (dst *TUpdMmclientUsers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TUpdMmclientUsers.Merge(dst, src)
}
func (m *TUpdMmclientUsers) XXX_Size() int {
	return xxx_messageInfo_TUpdMmclientUsers.Size(m)
}
func (m *TUpdMmclientUsers) XXX_DiscardUnknown() {
	xxx_messageInfo_TUpdMmclientUsers.DiscardUnknown(m)
}

var xxx_messageInfo_TUpdMmclientUsers proto.InternalMessageInfo

func (m *TUpdMmclientUsers) GetStatTime() string {
	if m != nil {
		return m.StatTime
	}
	return ""
}

func (m *TUpdMmclientUsers) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TUpdMmclientUsers) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*TUpdMmclientUsers)(nil), "usercenter.TUpdMmclientUsers")
}

func init() {
	proto.RegisterFile("usercenter/t_updmmclientusers.proto", fileDescriptor_t_updmmclientusers_96ad50160ac3f30f)
}

var fileDescriptor_t_updmmclientusers_96ad50160ac3f30f = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8d, 0xb1, 0xcf, 0x82, 0x30,
	0x10, 0x47, 0xc3, 0xf7, 0x25, 0x04, 0xba, 0xc9, 0x60, 0x88, 0x2e, 0x46, 0x17, 0x27, 0x3a, 0x38,
	0xba, 0xb9, 0xbb, 0x18, 0x58, 0x5c, 0x10, 0xe8, 0x89, 0x4d, 0xb8, 0xb6, 0xe1, 0xae, 0x83, 0xff,
	0xbd, 0x29, 0x90, 0xb0, 0xf5, 0xf7, 0xfa, 0x72, 0x4f, 0x9c, 0x3c, 0xc1, 0xd8, 0x81, 0x61, 0x18,
	0x25, 0xd7, 0xde, 0x29, 0xc4, 0x6e, 0xd0, 0x60, 0x38, 0x70, 0x2a, 0xdc, 0x68, 0xd9, 0x66, 0x62,
	0x95, 0x76, 0xd7, 0x5e, 0xf3, 0xc7, 0xb7, 0x45, 0x67, 0x51, 0xf6, 0x76, 0x68, 0x4c, 0x2f, 0x27,
	0xa9, 0xf5, 0x6f, 0xe9, 0xf8, 0xeb, 0x80, 0x24, 0x6b, 0x04, 0xe2, 0x06, 0xdd, 0xfa, 0x9a, 0x0f,
	0x1d, 0x5f, 0x62, 0x53, 0x56, 0x4e, 0xdd, 0x97, 0x46, 0x15, 0x1a, 0xd9, 0x5e, 0xa4, 0xc4, 0x0d,
	0xd7, 0x41, 0xce, 0xa3, 0x43, 0x74, 0x4e, 0x1f, 0x49, 0x00, 0xa5, 0x46, 0xc8, 0xb6, 0x22, 0x46,
	0xd2, 0xa4, 0x4c, 0xfe, 0x37, 0xfd, 0x2c, 0x2b, 0xf0, 0xe0, 0x78, 0xca, 0xff, 0x67, 0x3e, 0xaf,
	0x5b, 0xf2, 0x8c, 0xd1, 0x2a, 0x18, 0xa8, 0x8d, 0xa7, 0xe4, 0xe5, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x98, 0x96, 0x28, 0x7a, 0xe2, 0x00, 0x00, 0x00,
}
