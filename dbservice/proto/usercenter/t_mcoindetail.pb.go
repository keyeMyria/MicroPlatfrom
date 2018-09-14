// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_mcoindetail.proto

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

type TMcoinDetail struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Action               string               `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	ActionName           string               `protobuf:"bytes,4,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	McoinNum             string               `protobuf:"bytes,5,opt,name=mcoin_num,json=mcoinNum,proto3" json:"mcoin_num,omitempty"`
	ActionDate           string               `protobuf:"bytes,6,opt,name=action_date,json=actionDate,proto3" json:"action_date,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	Descs                string               `protobuf:"bytes,8,opt,name=descs,proto3" json:"descs,omitempty"`
	PartitionId          string               `protobuf:"bytes,9,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Source               string               `protobuf:"bytes,10,opt,name=source,proto3" json:"source,omitempty"`
	ActivityId           string               `protobuf:"bytes,11,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	Field1               string               `protobuf:"bytes,12,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2               string               `protobuf:"bytes,13,opt,name=field2,proto3" json:"field2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMcoinDetail) Reset()         { *m = TMcoinDetail{} }
func (m *TMcoinDetail) String() string { return proto.CompactTextString(m) }
func (*TMcoinDetail) ProtoMessage()    {}
func (*TMcoinDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_mcoindetail_8deaef0a399fe45c, []int{0}
}
func (m *TMcoinDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMcoinDetail.Unmarshal(m, b)
}
func (m *TMcoinDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMcoinDetail.Marshal(b, m, deterministic)
}
func (dst *TMcoinDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMcoinDetail.Merge(dst, src)
}
func (m *TMcoinDetail) XXX_Size() int {
	return xxx_messageInfo_TMcoinDetail.Size(m)
}
func (m *TMcoinDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TMcoinDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TMcoinDetail proto.InternalMessageInfo

func (m *TMcoinDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMcoinDetail) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMcoinDetail) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *TMcoinDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *TMcoinDetail) GetMcoinNum() string {
	if m != nil {
		return m.McoinNum
	}
	return ""
}

func (m *TMcoinDetail) GetActionDate() string {
	if m != nil {
		return m.ActionDate
	}
	return ""
}

func (m *TMcoinDetail) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMcoinDetail) GetDescs() string {
	if m != nil {
		return m.Descs
	}
	return ""
}

func (m *TMcoinDetail) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func (m *TMcoinDetail) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *TMcoinDetail) GetActivityId() string {
	if m != nil {
		return m.ActivityId
	}
	return ""
}

func (m *TMcoinDetail) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *TMcoinDetail) GetField2() string {
	if m != nil {
		return m.Field2
	}
	return ""
}

func init() {
	proto.RegisterType((*TMcoinDetail)(nil), "usercenter.TMcoinDetail")
}

func init() {
	proto.RegisterFile("usercenter/t_mcoindetail.proto", fileDescriptor_t_mcoindetail_8deaef0a399fe45c)
}

var fileDescriptor_t_mcoindetail_8deaef0a399fe45c = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xd5, 0x96, 0x86, 0xd6, 0x29, 0x0c, 0x11, 0x42, 0x56, 0x91, 0xa0, 0x30, 0x75, 0x4a,
	0x44, 0x19, 0xbb, 0xa1, 0x2e, 0x0c, 0x74, 0xa8, 0x3a, 0xb1, 0x54, 0x8e, 0x7d, 0x0d, 0x96, 0xe2,
	0x38, 0x8a, 0x2f, 0x48, 0xfd, 0x0e, 0x7c, 0x68, 0x94, 0x3b, 0xfa, 0x67, 0xf3, 0xfb, 0xdd, 0xbd,
	0x97, 0x7b, 0x11, 0x8f, 0x6d, 0x80, 0x46, 0x43, 0x85, 0xd0, 0x64, 0xb8, 0x73, 0xda, 0xdb, 0xca,
	0x00, 0x2a, 0x5b, 0xa6, 0x75, 0xe3, 0xd1, 0x27, 0xe2, 0x3c, 0x9f, 0x2e, 0x0b, 0x8b, 0xdf, 0x6d,
	0x9e, 0x6a, 0xef, 0xb2, 0xc2, 0x97, 0xaa, 0x2a, 0x32, 0x5a, 0xca, 0xdb, 0x7d, 0x56, 0xe3, 0xa1,
	0x86, 0x90, 0xa1, 0x75, 0x10, 0x50, 0xb9, 0xfa, 0xfc, 0xe2, 0xa0, 0x97, 0xdf, 0x81, 0x98, 0x6c,
	0x3f, 0xbb, 0xfc, 0x15, 0xe5, 0x27, 0xb7, 0xa2, 0x6f, 0x8d, 0xec, 0xcd, 0x7a, 0xf3, 0xf1, 0xa6,
	0x6f, 0x4d, 0x72, 0x2f, 0x22, 0xe7, 0x73, 0x5b, 0x82, 0xec, 0x13, 0xfb, 0x57, 0x1d, 0x57, 0x1a,
	0xad, 0xaf, 0xe4, 0x80, 0x39, 0xab, 0xe4, 0x49, 0xc4, 0xfc, 0xda, 0x55, 0xca, 0x81, 0xbc, 0xa2,
	0xa1, 0x60, 0xb4, 0x56, 0x0e, 0x92, 0x07, 0x31, 0xa6, 0x3e, 0xbb, 0xaa, 0x75, 0x72, 0x48, 0xe3,
	0x11, 0x81, 0x75, 0xeb, 0x2e, 0xdc, 0x46, 0x21, 0xc8, 0xe8, 0xd2, 0xbd, 0x52, 0x08, 0xc9, 0x52,
	0xc4, 0xba, 0x01, 0x85, 0xc0, 0x0b, 0xd7, 0xb3, 0xde, 0x3c, 0x5e, 0x4c, 0xd3, 0xc2, 0xfb, 0xa2,
	0x84, 0xf4, 0xd8, 0x3b, 0xdd, 0x1e, 0x6b, 0x6e, 0x04, 0xaf, 0x93, 0xf9, 0x4e, 0x0c, 0x0d, 0x04,
	0x1d, 0xe4, 0x88, 0x72, 0x59, 0x24, 0xcf, 0x62, 0x52, 0xab, 0x06, 0x2d, 0x7d, 0xd6, 0x1a, 0x39,
	0xa6, 0x61, 0x7c, 0x62, 0x1f, 0xf4, 0x13, 0x82, 0x6f, 0x1b, 0x0d, 0x52, 0x70, 0x59, 0x56, 0xc7,
	0x73, 0x7f, 0x2c, 0x1e, 0x3a, 0x67, 0x7c, 0x3e, 0xb7, 0x43, 0x6c, 0xdc, 0x5b, 0x28, 0xcd, 0xab,
	0x9c, 0xb0, 0x91, 0xd5, 0x89, 0x2f, 0xe4, 0xcd, 0x05, 0x5f, 0xbc, 0x8f, 0xbe, 0x22, 0xe7, 0x0d,
	0x94, 0x21, 0x8f, 0xa8, 0xcb, 0xdb, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x94, 0xdc, 0x3f,
	0x0a, 0x02, 0x00, 0x00,
}