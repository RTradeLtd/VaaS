// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgs.proto

package main

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

type GenerationRequest struct {
	SearchPrefix         string   `protobuf:"bytes,1,opt,name=searchPrefix,proto3" json:"searchPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerationRequest) Reset()         { *m = GenerationRequest{} }
func (m *GenerationRequest) String() string { return proto.CompactTextString(m) }
func (*GenerationRequest) ProtoMessage()    {}
func (*GenerationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgs_64be00d33e26adbc, []int{0}
}
func (m *GenerationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerationRequest.Unmarshal(m, b)
}
func (m *GenerationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerationRequest.Marshal(b, m, deterministic)
}
func (dst *GenerationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerationRequest.Merge(dst, src)
}
func (m *GenerationRequest) XXX_Size() int {
	return xxx_messageInfo_GenerationRequest.Size(m)
}
func (m *GenerationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerationRequest proto.InternalMessageInfo

func (m *GenerationRequest) GetSearchPrefix() string {
	if m != nil {
		return m.SearchPrefix
	}
	return ""
}

type GenerationResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerationResponse) Reset()         { *m = GenerationResponse{} }
func (m *GenerationResponse) String() string { return proto.CompactTextString(m) }
func (*GenerationResponse) ProtoMessage()    {}
func (*GenerationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgs_64be00d33e26adbc, []int{1}
}
func (m *GenerationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerationResponse.Unmarshal(m, b)
}
func (m *GenerationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerationResponse.Marshal(b, m, deterministic)
}
func (dst *GenerationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerationResponse.Merge(dst, src)
}
func (m *GenerationResponse) XXX_Size() int {
	return xxx_messageInfo_GenerationResponse.Size(m)
}
func (m *GenerationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerationResponse proto.InternalMessageInfo

func (m *GenerationResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GenerationResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerationRequest)(nil), "main.GenerationRequest")
	proto.RegisterType((*GenerationResponse)(nil), "main.GenerationResponse")
}

func init() { proto.RegisterFile("msgs.proto", fileDescriptor_msgs_64be00d33e26adbc) }

var fileDescriptor_msgs_64be00d33e26adbc = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2d, 0x4e, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x32, 0xe7, 0x12,
	0x74, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x52, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0x08, 0x28, 0x4a, 0x4d, 0xcb,
	0xac, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x11, 0x53, 0x72, 0xe0, 0x12, 0x42, 0xd6,
	0x58, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xd5, 0x00,
	0x62, 0x0a, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x81, 0x45,
	0x61, 0xdc, 0x24, 0x36, 0xb0, 0x3b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xfb, 0x16,
	0x1c, 0x95, 0x00, 0x00, 0x00,
}