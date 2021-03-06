// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_usemcoinlist.proto

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TUsemcoinlist struct {
	Seq                  string               `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Msisdn               string               `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Usedate              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=usedate,proto3" json:"usedate,omitempty"`
	Usetype              string               `protobuf:"bytes,4,opt,name=usetype,proto3" json:"usetype,omitempty"`
	Value                string               `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Descr                string               `protobuf:"bytes,6,opt,name=descr,proto3" json:"descr,omitempty"`
	Mcoin                string               `protobuf:"bytes,7,opt,name=mcoin,proto3" json:"mcoin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TUsemcoinlist) Reset()         { *m = TUsemcoinlist{} }
func (m *TUsemcoinlist) String() string { return proto.CompactTextString(m) }
func (*TUsemcoinlist) ProtoMessage()    {}
func (*TUsemcoinlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_usemcoinlist_9342ee56b3cbb2bf, []int{0}
}
func (m *TUsemcoinlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TUsemcoinlist.Unmarshal(m, b)
}
func (m *TUsemcoinlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TUsemcoinlist.Marshal(b, m, deterministic)
}
func (dst *TUsemcoinlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TUsemcoinlist.Merge(dst, src)
}
func (m *TUsemcoinlist) XXX_Size() int {
	return xxx_messageInfo_TUsemcoinlist.Size(m)
}
func (m *TUsemcoinlist) XXX_DiscardUnknown() {
	xxx_messageInfo_TUsemcoinlist.DiscardUnknown(m)
}

var xxx_messageInfo_TUsemcoinlist proto.InternalMessageInfo

func (m *TUsemcoinlist) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *TUsemcoinlist) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *TUsemcoinlist) GetUsedate() *timestamp.Timestamp {
	if m != nil {
		return m.Usedate
	}
	return nil
}

func (m *TUsemcoinlist) GetUsetype() string {
	if m != nil {
		return m.Usetype
	}
	return ""
}

func (m *TUsemcoinlist) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TUsemcoinlist) GetDescr() string {
	if m != nil {
		return m.Descr
	}
	return ""
}

func (m *TUsemcoinlist) GetMcoin() string {
	if m != nil {
		return m.Mcoin
	}
	return ""
}

func init() {
	proto.RegisterType((*TUsemcoinlist)(nil), "usercenter.TUsemcoinlist")
}

func init() {
	proto.RegisterFile("usercenter/t_usemcoinlist.proto", fileDescriptor_t_usemcoinlist_9342ee56b3cbb2bf)
}

var fileDescriptor_t_usemcoinlist_9342ee56b3cbb2bf = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xa9, 0xeb, 0x76, 0x35, 0x22, 0x48, 0x10, 0x09, 0x7b, 0x71, 0xf1, 0xb4, 0xa7, 0x04,
	0xd4, 0x9b, 0x37, 0x1f, 0x61, 0x59, 0x2f, 0x5e, 0xa4, 0x7f, 0xc6, 0x18, 0x48, 0x3a, 0xb5, 0x33,
	0x11, 0x7c, 0x46, 0x5f, 0x4a, 0x3a, 0xb5, 0xd4, 0x5b, 0x7e, 0x3f, 0xbe, 0xcc, 0xf7, 0xa9, 0xdb,
	0x4c, 0x30, 0x34, 0xd0, 0x31, 0x0c, 0x8e, 0xdf, 0x32, 0x41, 0x6a, 0x30, 0x74, 0x31, 0x10, 0xdb,
	0x7e, 0x40, 0x46, 0xad, 0x96, 0xc0, 0xf6, 0xc9, 0x07, 0xfe, 0xc8, 0xb5, 0x6d, 0x30, 0x39, 0x8f,
	0xb1, 0xea, 0xbc, 0x93, 0x50, 0x9d, 0xdf, 0x5d, 0xcf, 0xdf, 0x3d, 0x90, 0xe3, 0x90, 0x80, 0xb8,
	0x4a, 0xfd, 0xf2, 0x9a, 0x0e, 0xdd, 0xfd, 0x14, 0xea, 0xf2, 0xf8, 0xf2, 0xaf, 0x40, 0x5f, 0xa9,
	0x15, 0xc1, 0xa7, 0x29, 0x76, 0xc5, 0xfe, 0xfc, 0x30, 0x3e, 0xf5, 0x8d, 0x2a, 0x13, 0x05, 0x6a,
	0x3b, 0x73, 0x22, 0xf2, 0x8f, 0xf4, 0xa3, 0xda, 0x64, 0x82, 0xb6, 0x62, 0x30, 0xab, 0x5d, 0xb1,
	0xbf, 0xb8, 0xdf, 0x5a, 0x8f, 0xe8, 0x23, 0xd8, 0xb9, 0xdf, 0x1e, 0xe7, 0xba, 0xc3, 0x1c, 0xd5,
	0x46, 0x7e, 0x8d, 0xbb, 0xcc, 0xa9, 0x9c, 0x9b, 0x51, 0x5f, 0xab, 0xf5, 0x57, 0x15, 0x33, 0x98,
	0xb5, 0xf8, 0x09, 0x46, 0xdb, 0x02, 0x35, 0x83, 0x29, 0x27, 0x2b, 0x30, 0x5a, 0x99, 0x6c, 0x36,
	0x93, 0x15, 0x78, 0x3e, 0x7b, 0x2d, 0x13, 0xb6, 0x10, 0xa9, 0x2e, 0x65, 0xc2, 0xc3, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0x02, 0x00, 0x60, 0x4a, 0x01, 0x00, 0x00,
}
