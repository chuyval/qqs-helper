// Code generated by protoc-gen-go. DO NOT EDIT.
// source: q2.proto

package models // import "github.com/chuyval/qqs-helper/q2/pkg/models"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Stuff struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Stuff:
	//	*Stuff_ValueList
	Stuff                isStuff_Stuff `protobuf_oneof:"stuff"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Stuff) Reset()         { *m = Stuff{} }
func (m *Stuff) String() string { return proto.CompactTextString(m) }
func (*Stuff) ProtoMessage()    {}
func (*Stuff) Descriptor() ([]byte, []int) {
	return fileDescriptor_q2_0cf3d1b9d1b3ee61, []int{0}
}
func (m *Stuff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stuff.Unmarshal(m, b)
}
func (m *Stuff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stuff.Marshal(b, m, deterministic)
}
func (dst *Stuff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stuff.Merge(dst, src)
}
func (m *Stuff) XXX_Size() int {
	return xxx_messageInfo_Stuff.Size(m)
}
func (m *Stuff) XXX_DiscardUnknown() {
	xxx_messageInfo_Stuff.DiscardUnknown(m)
}

var xxx_messageInfo_Stuff proto.InternalMessageInfo

type isStuff_Stuff interface {
	isStuff_Stuff()
}

type Stuff_ValueList struct {
	ValueList *SomeValueList `protobuf:"bytes,11,opt,name=value_list,json=valueList,proto3,oneof"`
}

func (*Stuff_ValueList) isStuff_Stuff() {}

func (m *Stuff) GetStuff() isStuff_Stuff {
	if m != nil {
		return m.Stuff
	}
	return nil
}

func (m *Stuff) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Stuff) GetValueList() *SomeValueList {
	if x, ok := m.GetStuff().(*Stuff_ValueList); ok {
		return x.ValueList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Stuff) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Stuff_OneofMarshaler, _Stuff_OneofUnmarshaler, _Stuff_OneofSizer, []interface{}{
		(*Stuff_ValueList)(nil),
	}
}

func _Stuff_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Stuff)
	// stuff
	switch x := m.Stuff.(type) {
	case *Stuff_ValueList:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ValueList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Stuff.Stuff has unexpected type %T", x)
	}
	return nil
}

func _Stuff_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Stuff)
	switch tag {
	case 11: // stuff.value_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SomeValueList)
		err := b.DecodeMessage(msg)
		m.Stuff = &Stuff_ValueList{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Stuff_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Stuff)
	// stuff
	switch x := m.Stuff.(type) {
	case *Stuff_ValueList:
		s := proto.Size(x.ValueList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SomeValueList struct {
	SomeValue            []*SomeValue `protobuf:"bytes,1,rep,name=some_value,json=someValue,proto3" json:"some_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SomeValueList) Reset()         { *m = SomeValueList{} }
func (m *SomeValueList) String() string { return proto.CompactTextString(m) }
func (*SomeValueList) ProtoMessage()    {}
func (*SomeValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_q2_0cf3d1b9d1b3ee61, []int{1}
}
func (m *SomeValueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SomeValueList.Unmarshal(m, b)
}
func (m *SomeValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SomeValueList.Marshal(b, m, deterministic)
}
func (dst *SomeValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeValueList.Merge(dst, src)
}
func (m *SomeValueList) XXX_Size() int {
	return xxx_messageInfo_SomeValueList.Size(m)
}
func (m *SomeValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeValueList.DiscardUnknown(m)
}

var xxx_messageInfo_SomeValueList proto.InternalMessageInfo

func (m *SomeValueList) GetSomeValue() []*SomeValue {
	if m != nil {
		return m.SomeValue
	}
	return nil
}

type SomeValue struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SomeValue) Reset()         { *m = SomeValue{} }
func (m *SomeValue) String() string { return proto.CompactTextString(m) }
func (*SomeValue) ProtoMessage()    {}
func (*SomeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_q2_0cf3d1b9d1b3ee61, []int{2}
}
func (m *SomeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SomeValue.Unmarshal(m, b)
}
func (m *SomeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SomeValue.Marshal(b, m, deterministic)
}
func (dst *SomeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeValue.Merge(dst, src)
}
func (m *SomeValue) XXX_Size() int {
	return xxx_messageInfo_SomeValue.Size(m)
}
func (m *SomeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeValue.DiscardUnknown(m)
}

var xxx_messageInfo_SomeValue proto.InternalMessageInfo

func (m *SomeValue) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Stuff)(nil), "proto.Stuff")
	proto.RegisterType((*SomeValueList)(nil), "proto.SomeValueList")
	proto.RegisterType((*SomeValue)(nil), "proto.SomeValue")
}

func init() { proto.RegisterFile("q2.proto", fileDescriptor_q2_0cf3d1b9d1b3ee61) }

var fileDescriptor_q2_0cf3d1b9d1b3ee61 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x34, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xe1, 0x5c, 0xac, 0xc1, 0x25, 0xa5, 0x69,
	0x69, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99,
	0x29, 0x42, 0xa6, 0x5c, 0x5c, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0xf1, 0x39, 0x99, 0xc5, 0x25, 0x12,
	0xdc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x10, 0xbd, 0x7a, 0xc1, 0xf9, 0xb9, 0xa9, 0x61, 0x20,
	0x49, 0x9f, 0xcc, 0xe2, 0x12, 0x0f, 0x86, 0x20, 0xce, 0x32, 0x18, 0xc7, 0x89, 0x9d, 0x8b, 0xb5,
	0x18, 0x64, 0x9e, 0x92, 0x03, 0x17, 0x2f, 0x8a, 0x32, 0x21, 0x7d, 0x2e, 0xae, 0xe2, 0xfc, 0xdc,
	0xd4, 0x78, 0xb0, 0x5a, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x01, 0x74, 0x03, 0x83, 0x38,
	0x8b, 0x61, 0x4c, 0x25, 0x69, 0x2e, 0x4e, 0xb8, 0x38, 0xba, 0xf3, 0x9c, 0x74, 0xa3, 0xb4, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0x33, 0x4a, 0x2b, 0xcb, 0x12,
	0x73, 0xf4, 0x0b, 0x0b, 0x8b, 0x75, 0x33, 0x52, 0x73, 0x0a, 0x52, 0x8b, 0xf4, 0x0b, 0x8d, 0xf4,
	0x0b, 0xb2, 0xd3, 0xf5, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x8a, 0x93, 0xd8, 0xc0, 0xf6, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x86, 0x7b, 0x47, 0x00, 0x01, 0x00, 0x00,
}
