// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_membermcoinreset.proto

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

type TMemberMcoinReset struct {
	OperateDate          string   `protobuf:"bytes,1,opt,name=operate_date,json=operateDate,proto3" json:"operate_date,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	McoinNum             string   `protobuf:"bytes,3,opt,name=mcoin_num,json=mcoinNum,proto3" json:"mcoin_num,omitempty"`
	Descs                string   `protobuf:"bytes,4,opt,name=descs,proto3" json:"descs,omitempty"`
	PartitionId          string   `protobuf:"bytes,5,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TMemberMcoinReset) Reset()         { *m = TMemberMcoinReset{} }
func (m *TMemberMcoinReset) String() string { return proto.CompactTextString(m) }
func (*TMemberMcoinReset) ProtoMessage()    {}
func (*TMemberMcoinReset) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_membermcoinreset_88cdaeef4cee42e7, []int{0}
}
func (m *TMemberMcoinReset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberMcoinReset.Unmarshal(m, b)
}
func (m *TMemberMcoinReset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberMcoinReset.Marshal(b, m, deterministic)
}
func (dst *TMemberMcoinReset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberMcoinReset.Merge(dst, src)
}
func (m *TMemberMcoinReset) XXX_Size() int {
	return xxx_messageInfo_TMemberMcoinReset.Size(m)
}
func (m *TMemberMcoinReset) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberMcoinReset.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberMcoinReset proto.InternalMessageInfo

func (m *TMemberMcoinReset) GetOperateDate() string {
	if m != nil {
		return m.OperateDate
	}
	return ""
}

func (m *TMemberMcoinReset) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMemberMcoinReset) GetMcoinNum() string {
	if m != nil {
		return m.McoinNum
	}
	return ""
}

func (m *TMemberMcoinReset) GetDescs() string {
	if m != nil {
		return m.Descs
	}
	return ""
}

func (m *TMemberMcoinReset) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberMcoinReset)(nil), "usercenter.TMemberMcoinReset")
}

func init() {
	proto.RegisterFile("usercenter/t_membermcoinreset.proto", fileDescriptor_t_membermcoinreset_88cdaeef4cee42e7)
}

var fileDescriptor_t_membermcoinreset_88cdaeef4cee42e7 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0xa0, 0x51, 0x6b, 0x58, 0xb0, 0x10, 0xb2, 0x60, 0xe1, 0x63, 0x61, 0xaa, 0x07,
	0x46, 0x36, 0xc4, 0xc2, 0x50, 0x86, 0x8a, 0x89, 0x25, 0xb2, 0xe3, 0x23, 0x58, 0xca, 0xf9, 0x2c,
	0xfb, 0x3c, 0xf0, 0x7f, 0xf8, 0xa1, 0x28, 0x4e, 0xd5, 0x6c, 0xf7, 0x7e, 0xe8, 0x9e, 0x57, 0x3c,
	0x96, 0x0c, 0xa9, 0x87, 0xc0, 0x90, 0x34, 0x77, 0x08, 0x68, 0x21, 0x61, 0x4f, 0x3e, 0x24, 0xc8,
	0xc0, 0xdb, 0x98, 0x88, 0x49, 0x8a, 0xa5, 0x74, 0xf3, 0x32, 0x78, 0xfe, 0x29, 0x76, 0xdb, 0x13,
	0xea, 0x81, 0x46, 0x13, 0x06, 0x5d, 0x4b, 0xb6, 0x7c, 0xeb, 0xc8, 0xbf, 0x11, 0xb2, 0x66, 0x8f,
	0x90, 0xd9, 0x60, 0x5c, 0xae, 0xf9, 0xd1, 0xc3, 0x5f, 0x23, 0x2e, 0x3f, 0x77, 0x15, 0xb2, 0x9b,
	0x20, 0xfb, 0x09, 0x22, 0xef, 0xc5, 0x05, 0x45, 0x48, 0x86, 0xa1, 0x73, 0x86, 0x41, 0x35, 0x77,
	0xcd, 0xd3, 0x66, 0x7f, 0x7e, 0xf0, 0xde, 0x0c, 0x83, 0xbc, 0x16, 0x2d, 0x92, 0xf5, 0x23, 0xa8,
	0x93, 0x1a, 0x1e, 0x94, 0xbc, 0x15, 0x9b, 0xba, 0xb6, 0x0b, 0x05, 0xd5, 0x69, 0x8d, 0xd6, 0xd5,
	0xf8, 0x28, 0x28, 0xaf, 0xc4, 0xca, 0x41, 0xee, 0xb3, 0x3a, 0xab, 0xc1, 0x2c, 0x26, 0x5a, 0x34,
	0x89, 0x3d, 0x7b, 0x0a, 0x9d, 0x77, 0x6a, 0x35, 0xd3, 0x8e, 0xde, 0xbb, 0x7b, 0x5d, 0x7f, 0xb5,
	0x48, 0x0e, 0xc6, 0x6c, 0xdb, 0xba, 0xfb, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xf9, 0x95,
	0x9b, 0x27, 0x01, 0x00, 0x00,
}
