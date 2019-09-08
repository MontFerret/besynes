// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package dto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Query struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string              `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Params               map[string]*any.Any `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cdp                  string              `protobuf:"bytes,4,opt,name=cdp,proto3" json:"cdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Query) GetParams() map[string]*any.Any {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Query) GetCdp() string {
	if m != nil {
		return m.Cdp
	}
	return ""
}

func init() {
	proto.RegisterType((*Query)(nil), "Query")
	proto.RegisterMapType((map[string]*any.Any)(nil), "Query.ParamsEntry")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2c, 0x4d, 0x2d,
	0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0xa0, 0x52, 0x4a, 0xbb, 0x18, 0xb9, 0x58, 0x03,
	0x41, 0x4a, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98,
	0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x98, 0xc0, 0x22, 0x60, 0xb6,
	0x90, 0x16, 0x17, 0x5b, 0x41, 0x62, 0x51, 0x62, 0x6e, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7,
	0x91, 0x90, 0x1e, 0x58, 0xaf, 0x5e, 0x00, 0x58, 0xd0, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0xaa,
	0x42, 0x48, 0x80, 0x8b, 0x39, 0x39, 0xa5, 0x40, 0x82, 0x05, 0xac, 0x1d, 0xc4, 0x94, 0xf2, 0xe7,
	0xe2, 0x46, 0x52, 0x08, 0x52, 0x90, 0x9d, 0x5a, 0x09, 0xb5, 0x11, 0xc4, 0x14, 0xd2, 0xe2, 0x62,
	0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x05, 0xdb, 0xc9, 0x6d, 0x24, 0xa2, 0x07, 0x71, 0xb7, 0x1e, 0xcc,
	0xdd, 0x7a, 0x8e, 0x79, 0x95, 0x41, 0x10, 0x25, 0x56, 0x4c, 0x16, 0x8c, 0x4e, 0xac, 0x51, 0xcc,
	0x29, 0x25, 0xf9, 0x49, 0x6c, 0x60, 0x79, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x0a,
	0xdd, 0xf4, 0xf4, 0x00, 0x00, 0x00,
}
