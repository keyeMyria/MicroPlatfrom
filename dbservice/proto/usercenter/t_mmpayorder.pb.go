// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercenter/t_mmpayorder.proto

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

type TMmPayOrder struct {
	OrderId              string               `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TradeId              string               `protobuf:"bytes,2,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	Mobile               string               `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status               string               `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	DelFlag              string               `protobuf:"bytes,5,opt,name=del_flag,json=delFlag,proto3" json:"del_flag,omitempty"`
	GoodsId              string               `protobuf:"bytes,6,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	OrderDate            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=order_date,json=orderDate,proto3" json:"order_date,omitempty"`
	PayDate              *timestamp.Timestamp `protobuf:"bytes,8,opt,name=pay_date,json=payDate,proto3" json:"pay_date,omitempty"`
	McoinNum             string               `protobuf:"bytes,9,opt,name=mcoin_num,json=mcoinNum,proto3" json:"mcoin_num,omitempty"`
	ScoreNum             string               `protobuf:"bytes,10,opt,name=score_num,json=scoreNum,proto3" json:"score_num,omitempty"`
	PayMoney             string               `protobuf:"bytes,11,opt,name=pay_money,json=payMoney,proto3" json:"pay_money,omitempty"`
	NotifyDate           *timestamp.Timestamp `protobuf:"bytes,12,opt,name=notify_date,json=notifyDate,proto3" json:"notify_date,omitempty"`
	NotifyResult         string               `protobuf:"bytes,13,opt,name=notify_result,json=notifyResult,proto3" json:"notify_result,omitempty"`
	NotifyMsg            string               `protobuf:"bytes,14,opt,name=notify_msg,json=notifyMsg,proto3" json:"notify_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TMmPayOrder) Reset()         { *m = TMmPayOrder{} }
func (m *TMmPayOrder) String() string { return proto.CompactTextString(m) }
func (*TMmPayOrder) ProtoMessage()    {}
func (*TMmPayOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_t_mmpayorder_8770ceb74f1bf701, []int{0}
}
func (m *TMmPayOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TMmPayOrder.Unmarshal(m, b)
}
func (m *TMmPayOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TMmPayOrder.Marshal(b, m, deterministic)
}
func (dst *TMmPayOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TMmPayOrder.Merge(dst, src)
}
func (m *TMmPayOrder) XXX_Size() int {
	return xxx_messageInfo_TMmPayOrder.Size(m)
}
func (m *TMmPayOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_TMmPayOrder.DiscardUnknown(m)
}

var xxx_messageInfo_TMmPayOrder proto.InternalMessageInfo

func (m *TMmPayOrder) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *TMmPayOrder) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

func (m *TMmPayOrder) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *TMmPayOrder) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TMmPayOrder) GetDelFlag() string {
	if m != nil {
		return m.DelFlag
	}
	return ""
}

func (m *TMmPayOrder) GetGoodsId() string {
	if m != nil {
		return m.GoodsId
	}
	return ""
}

func (m *TMmPayOrder) GetOrderDate() *timestamp.Timestamp {
	if m != nil {
		return m.OrderDate
	}
	return nil
}

func (m *TMmPayOrder) GetPayDate() *timestamp.Timestamp {
	if m != nil {
		return m.PayDate
	}
	return nil
}

func (m *TMmPayOrder) GetMcoinNum() string {
	if m != nil {
		return m.McoinNum
	}
	return ""
}

func (m *TMmPayOrder) GetScoreNum() string {
	if m != nil {
		return m.ScoreNum
	}
	return ""
}

func (m *TMmPayOrder) GetPayMoney() string {
	if m != nil {
		return m.PayMoney
	}
	return ""
}

func (m *TMmPayOrder) GetNotifyDate() *timestamp.Timestamp {
	if m != nil {
		return m.NotifyDate
	}
	return nil
}

func (m *TMmPayOrder) GetNotifyResult() string {
	if m != nil {
		return m.NotifyResult
	}
	return ""
}

func (m *TMmPayOrder) GetNotifyMsg() string {
	if m != nil {
		return m.NotifyMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*TMmPayOrder)(nil), "usercenter.TMmPayOrder")
}

func init() {
	proto.RegisterFile("usercenter/t_mmpayorder.proto", fileDescriptor_t_mmpayorder_8770ceb74f1bf701)
}

var fileDescriptor_t_mmpayorder_8770ceb74f1bf701 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x8b, 0xdb, 0x30,
	0x14, 0xc6, 0x49, 0xd3, 0x3a, 0xb6, 0x9c, 0x74, 0xf0, 0x50, 0xd4, 0x94, 0x40, 0x68, 0x97, 0x4c,
	0x36, 0xb4, 0x74, 0x28, 0xd9, 0xca, 0x71, 0x90, 0x21, 0x77, 0x47, 0xc8, 0x74, 0x8b, 0x91, 0xad,
	0x17, 0x9d, 0x41, 0xb2, 0x84, 0x24, 0x0f, 0xfe, 0xbb, 0xef, 0x1f, 0x38, 0xf4, 0xe4, 0x90, 0x31,
	0x9b, 0xbe, 0xef, 0xf7, 0xde, 0xf7, 0x3e, 0x91, 0xcd, 0xe0, 0xc0, 0xb6, 0xd0, 0x7b, 0xb0, 0x95,
	0xaf, 0x95, 0x32, 0x6c, 0xd4, 0x96, 0x83, 0x2d, 0x8d, 0xd5, 0x5e, 0x17, 0xe4, 0x86, 0xd7, 0x7b,
	0xd1, 0xf9, 0xb7, 0xa1, 0x29, 0x5b, 0xad, 0x2a, 0xa1, 0x25, 0xeb, 0x45, 0x85, 0x43, 0xcd, 0x70,
	0xa9, 0x8c, 0x1f, 0x0d, 0xb8, 0xca, 0x77, 0x0a, 0x9c, 0x67, 0xca, 0xdc, 0x5e, 0x31, 0xe8, 0xe7,
	0xfb, 0x9c, 0xe4, 0xe7, 0xa3, 0x7a, 0x61, 0xe3, 0x73, 0x88, 0x2f, 0xbe, 0x93, 0x14, 0xef, 0xd4,
	0x1d, 0xa7, 0xb3, 0xed, 0x6c, 0x97, 0x9d, 0x16, 0xa8, 0x0f, 0x3c, 0x20, 0x6f, 0x19, 0x87, 0x80,
	0x3e, 0x45, 0x84, 0xfa, 0xc0, 0x8b, 0x6f, 0x24, 0x51, 0xba, 0xe9, 0x24, 0xd0, 0x39, 0x82, 0x49,
	0x05, 0xdf, 0x79, 0xe6, 0x07, 0x47, 0x3f, 0x47, 0x3f, 0xaa, 0x10, 0xc5, 0x41, 0xd6, 0x17, 0xc9,
	0x04, 0xfd, 0x12, 0xa3, 0x38, 0xc8, 0x47, 0xc9, 0x44, 0x40, 0x42, 0x6b, 0xee, 0xc2, 0x95, 0x24,
	0x22, 0xd4, 0x07, 0x5e, 0xfc, 0x23, 0x24, 0x76, 0xe3, 0xcc, 0x03, 0x5d, 0x6c, 0x67, 0xbb, 0xfc,
	0xf7, 0xba, 0x14, 0x5a, 0x0b, 0x09, 0xe5, 0xf5, 0xcb, 0xe5, 0xf9, 0xfa, 0xc3, 0x53, 0x86, 0xd3,
	0x0f, 0xcc, 0x43, 0xf1, 0x97, 0xa4, 0x86, 0x8d, 0x71, 0x31, 0xbd, 0xbb, 0xb8, 0x30, 0x6c, 0xc4,
	0xb5, 0x1f, 0x24, 0x53, 0xad, 0xee, 0xfa, 0xba, 0x1f, 0x14, 0xcd, 0xb0, 0x4d, 0x8a, 0xc6, 0xd3,
	0xa0, 0x02, 0x74, 0xad, 0xb6, 0x80, 0x90, 0x44, 0x88, 0xc6, 0x04, 0xc3, 0x41, 0xa5, 0x7b, 0x18,
	0x69, 0x1e, 0xa1, 0x61, 0xe3, 0x31, 0xe8, 0x62, 0x4f, 0xf2, 0x5e, 0xfb, 0xee, 0x32, 0x15, 0x5a,
	0xde, 0x2d, 0x44, 0xe2, 0x38, 0x76, 0xfa, 0x45, 0x56, 0xd3, 0xb2, 0x05, 0x37, 0x48, 0x4f, 0x57,
	0x98, 0xbe, 0x8c, 0xe6, 0x09, 0xbd, 0x62, 0x43, 0xa6, 0x95, 0x5a, 0x39, 0x41, 0xbf, 0xe2, 0x44,
	0x16, 0x9d, 0xa3, 0x13, 0xff, 0xd3, 0xd7, 0x44, 0x69, 0x0e, 0xd2, 0x35, 0x09, 0x5e, 0xfb, 0xf3,
	0x11, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xb8, 0xa9, 0x3c, 0x70, 0x02, 0x00, 0x00,
}
