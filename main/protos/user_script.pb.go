// Code generated by protoc-gen-go.
// source: user_script.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	user_script.proto

It has these top-level messages:
	UserSql
	Script
	ObjScript
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type UserSql struct {
	User   string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Script string `protobuf:"bytes,2,opt,name=script" json:"script,omitempty"`
}

func (m *UserSql) Reset()                    { *m = UserSql{} }
func (m *UserSql) String() string            { return proto1.CompactTextString(m) }
func (*UserSql) ProtoMessage()               {}
func (*UserSql) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserSql) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserSql) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

type Script struct {
	Script string `protobuf:"bytes,1,opt,name=script" json:"script,omitempty"`
}

func (m *Script) Reset()                    { *m = Script{} }
func (m *Script) String() string            { return proto1.CompactTextString(m) }
func (*Script) ProtoMessage()               {}
func (*Script) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Script) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

type ObjScript struct {
	ObjScript string `protobuf:"bytes,1,opt,name=objScript" json:"objScript,omitempty"`
}

func (m *ObjScript) Reset()                    { *m = ObjScript{} }
func (m *ObjScript) String() string            { return proto1.CompactTextString(m) }
func (*ObjScript) ProtoMessage()               {}
func (*ObjScript) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ObjScript) GetObjScript() string {
	if m != nil {
		return m.ObjScript
	}
	return ""
}

func init() {
	proto1.RegisterType((*UserSql)(nil), "proto.UserSql")
	proto1.RegisterType((*Script)(nil), "proto.Script")
	proto1.RegisterType((*ObjScript)(nil), "proto.ObjScript")
}

func init() { proto1.RegisterFile("user_script.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x4a, 0xa6, 0x5c, 0xec, 0xa1, 0xc5, 0xa9, 0x45, 0xc1, 0x85, 0x39, 0x42, 0x42, 0x5c, 0x2c,
	0x20, 0x65, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x18, 0x17, 0x1b, 0x44,
	0x97, 0x04, 0x13, 0x58, 0x14, 0xca, 0x53, 0x52, 0xe0, 0x62, 0x0b, 0x06, 0xb3, 0x90, 0x54, 0x30,
	0xa2, 0xa8, 0xd0, 0xe4, 0xe2, 0xf4, 0x4f, 0xca, 0x82, 0x2a, 0x92, 0xe1, 0xe2, 0xcc, 0x87, 0x71,
	0xa0, 0xea, 0x10, 0x02, 0x46, 0x39, 0x5c, 0x5c, 0x8e, 0xfe, 0x4e, 0xc1, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x42, 0xfa, 0x5c, 0xdc, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0x8e, 0xc9, 0xc9, 0xa9, 0xc5,
	0xc5, 0x42, 0x7c, 0x10, 0xf7, 0xea, 0x41, 0x5d, 0x29, 0x85, 0xc6, 0x57, 0x62, 0x10, 0xd2, 0xe5,
	0xe2, 0x08, 0x48, 0x2c, 0x2a, 0x4e, 0x05, 0xf9, 0x81, 0x17, 0x2a, 0x0b, 0x31, 0x59, 0x4a, 0x00,
	0xca, 0x85, 0xbb, 0x44, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0x50, 0x99, 0xf4, 0x14, 0x01, 0x00, 0x00,
}
