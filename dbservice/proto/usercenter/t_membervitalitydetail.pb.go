// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_membervitalitydetail.proto

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

type TMemberVitalityDetail struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string               `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	ActionType           string               `protobuf:"bytes,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	IncrNum              string               `protobuf:"bytes,4,opt,name=incr_num,json=incrNum,proto3" json:"incr_num,omitempty"`
	Total                string               `protobuf:"bytes,5,opt,name=total,proto3" json:"total,omitempty"`
	Source               string               `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	ExpireDate           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	Status               string               `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	PartitionId          string               `protobuf:"bytes,10,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMemberVitalityDetail) Reset()         { *m = TMemberVitalityDetail{} }
func (m *TMemberVitalityDetail) String() string { return proto.CompactTextString(m) }
func (*TMemberVitalityDetail) ProtoMessage()    {}
func (*TMemberVitalityDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_membervitalitydetail_e60b657bf720ca48, []int{0}
}
func (m *TMemberVitalityDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMemberVitalityDetail.Unmarshal(m, b)
}
func (m *TMemberVitalityDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMemberVitalityDetail.Marshal(b, m, deterministic)
}
func (dst *TMemberVitalityDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMemberVitalityDetail.Merge(dst, src)
}
func (m *TMemberVitalityDetail) XXX_Size() int {
	return xxx_messageInfo_TMemberVitalityDetail.Size(m)
}
func (m *TMemberVitalityDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TMemberVitalityDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TMemberVitalityDetail proto.InternalMessageInfo

func (m *TMemberVitalityDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TMemberVitalityDetail) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMemberVitalityDetail) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *TMemberVitalityDetail) GetIncrNum() string {
	if m != nil {
		return m.IncrNum
	}
	return ""
}

func (m *TMemberVitalityDetail) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *TMemberVitalityDetail) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *TMemberVitalityDetail) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *TMemberVitalityDetail) GetExpireDate() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireDate
	}
	return nil
}

func (m *TMemberVitalityDetail) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TMemberVitalityDetail) GetPartitionId() string {
	if m != nil {
		return m.PartitionId
	}
	return ""
}

func init() {
	proto.RegisterType((*TMemberVitalityDetail)(nil), "usercenter.TMemberVitalityDetail")
}

func init() {
	proto.RegisterFile("usercenter/t_membervitalitydetail.proto", fileDescriptor_t_membervitalitydetail_e60b657bf720ca48)
}

var fileDescriptor_t_membervitalitydetail_e60b657bf720ca48 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xd5, 0x40, 0xd3, 0xd6, 0x41, 0x0c, 0x16, 0x20, 0xd3, 0x85, 0xc2, 0x42, 0xa7, 0x44,
	0x82, 0xb1, 0x1b, 0xea, 0xc2, 0x00, 0x43, 0x55, 0x31, 0xb0, 0x44, 0x4e, 0x7c, 0x04, 0x4b, 0x76,
	0x6c, 0x39, 0x67, 0x44, 0xfe, 0x2c, 0xbf, 0x05, 0xc5, 0x4e, 0xe9, 0xc8, 0xe6, 0xf7, 0xee, 0x3d,
	0x7f, 0x77, 0xe4, 0xde, 0x77, 0xe0, 0x6a, 0x68, 0x11, 0x5c, 0x81, 0xa5, 0x06, 0x5d, 0x81, 0xfb,
	0x92, 0xc8, 0x95, 0xc4, 0x5e, 0x00, 0x72, 0xa9, 0x72, 0xeb, 0x0c, 0x1a, 0x4a, 0x8e, 0xc1, 0xe5,
	0xa6, 0x91, 0xf8, 0xe9, 0xab, 0xbc, 0x36, 0xba, 0x68, 0x8c, 0xe2, 0x6d, 0x53, 0x84, 0x50, 0xe5,
	0x3f, 0x0a, 0x8b, 0xbd, 0x85, 0xae, 0x40, 0xa9, 0xa1, 0x43, 0xae, 0xed, 0xf1, 0x15, 0x3f, 0xba,
	0xfb, 0x49, 0xc8, 0xe5, 0xfe, 0x25, 0x80, 0xde, 0x46, 0xd0, 0x36, 0x80, 0xe8, 0x39, 0x49, 0xa4,
	0x60, 0x93, 0xd5, 0x64, 0xbd, 0xd8, 0x25, 0x52, 0xd0, 0x2b, 0x92, 0x6a, 0x53, 0x49, 0x05, 0x2c,
	0x09, 0xde, 0xa8, 0xe8, 0x0d, 0xc9, 0x78, 0x8d, 0xd2, 0xb4, 0xe5, 0xc0, 0x62, 0x27, 0x61, 0x48,
	0xa2, 0xb5, 0xef, 0x2d, 0xd0, 0x6b, 0x32, 0x97, 0x6d, 0xed, 0xca, 0xd6, 0x6b, 0x76, 0x1a, 0xa6,
	0xb3, 0x41, 0xbf, 0x7a, 0x4d, 0x2f, 0xc8, 0x14, 0x0d, 0x72, 0xc5, 0xa6, 0xc1, 0x8f, 0x62, 0x20,
	0x75, 0xc6, 0xbb, 0x1a, 0x58, 0x1a, 0x49, 0x51, 0xd1, 0x0d, 0xc9, 0x6a, 0x07, 0x1c, 0xa1, 0x14,
	0x1c, 0x81, 0xcd, 0x56, 0x93, 0x75, 0xf6, 0xb0, 0xcc, 0x1b, 0x63, 0x1a, 0x05, 0xf9, 0xe1, 0xe6,
	0x7c, 0x7f, 0x38, 0x71, 0x47, 0x62, 0x7c, 0xcb, 0x31, 0x94, 0xe1, 0xdb, 0x4a, 0x37, 0x96, 0xe7,
	0xff, 0x97, 0x63, 0x3c, 0x94, 0x87, 0x8d, 0x90, 0xa3, 0xef, 0xd8, 0x62, 0xdc, 0x28, 0x28, 0x7a,
	0x4b, 0xce, 0x2c, 0x77, 0x28, 0xc3, 0xf9, 0x52, 0x30, 0x12, 0xa6, 0xd9, 0x9f, 0xf7, 0x2c, 0x9e,
	0xe6, 0xef, 0xa9, 0x36, 0x02, 0x54, 0x57, 0xa5, 0x01, 0xf2, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xbc, 0xaf, 0x35, 0xef, 0xe5, 0x01, 0x00, 0x00,
}
