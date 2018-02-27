// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go.
// source: namespace.proto
// DO NOT EDIT!

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	namespace.proto
	policy.proto
	rule.proto

It has these top-level messages:
	NamespaceSnapshot
	Namespace
	Namespaces
	Resolution
	Retention
	StoragePolicy
	Policy
	MappingRuleSnapshot
	MappingRule
	RollupTarget
	RollupRuleSnapshot
	RollupRule
	RuleSet
*/
package schema

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

type NamespaceSnapshot struct {
	ForRulesetVersion  int32  `protobuf:"varint,1,opt,name=for_ruleset_version,json=forRulesetVersion" json:"for_ruleset_version,omitempty"`
	Tombstoned         bool   `protobuf:"varint,2,opt,name=tombstoned" json:"tombstoned,omitempty"`
	LastUpdatedAtNanos int64  `protobuf:"varint,3,opt,name=last_updated_at_nanos,json=lastUpdatedAtNanos" json:"last_updated_at_nanos,omitempty"`
	LastUpdatedBy      string `protobuf:"bytes,4,opt,name=last_updated_by,json=lastUpdatedBy" json:"last_updated_by,omitempty"`
}

func (m *NamespaceSnapshot) Reset()                    { *m = NamespaceSnapshot{} }
func (m *NamespaceSnapshot) String() string            { return proto.CompactTextString(m) }
func (*NamespaceSnapshot) ProtoMessage()               {}
func (*NamespaceSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Namespace struct {
	Name      string               `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Snapshots []*NamespaceSnapshot `protobuf:"bytes,2,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *Namespace) Reset()                    { *m = Namespace{} }
func (m *Namespace) String() string            { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()               {}
func (*Namespace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Namespace) GetSnapshots() []*NamespaceSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type Namespaces struct {
	Namespaces []*Namespace `protobuf:"bytes,1,rep,name=namespaces" json:"namespaces,omitempty"`
}

func (m *Namespaces) Reset()                    { *m = Namespaces{} }
func (m *Namespaces) String() string            { return proto.CompactTextString(m) }
func (*Namespaces) ProtoMessage()               {}
func (*Namespaces) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Namespaces) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*NamespaceSnapshot)(nil), "schema.NamespaceSnapshot")
	proto.RegisterType((*Namespace)(nil), "schema.Namespace")
	proto.RegisterType((*Namespaces)(nil), "schema.Namespaces")
}

func init() { proto.RegisterFile("namespace.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x3a, 0x87, 0x7d, 0x22, 0xa3, 0x4f, 0x84, 0x78, 0x91, 0xd0, 0x83, 0xe4, 0x54,
	0x98, 0x1e, 0x3c, 0x8a, 0x7e, 0x80, 0x1d, 0x22, 0x8a, 0xb7, 0x90, 0xae, 0x19, 0x13, 0xd6, 0xa4,
	0xe4, 0x65, 0xc2, 0xbe, 0x9c, 0x9f, 0x4d, 0xda, 0x6a, 0xed, 0xe8, 0x2d, 0xe4, 0xf7, 0x7f, 0x8f,
	0xf7, 0xfb, 0xc3, 0xd2, 0x99, 0xda, 0x52, 0x63, 0x36, 0xb6, 0x68, 0x82, 0x8f, 0x1e, 0x17, 0xb4,
	0xd9, 0xd9, 0xda, 0xe4, 0xdf, 0x0c, 0xb2, 0xf5, 0x1f, 0x7b, 0x75, 0xa6, 0xa1, 0x9d, 0x8f, 0x58,
	0xc0, 0xd5, 0xd6, 0x07, 0x1d, 0x0e, 0x7b, 0x4b, 0x36, 0xea, 0x2f, 0x1b, 0xe8, 0xd3, 0x3b, 0xce,
	0x04, 0x93, 0x67, 0x2a, 0xdb, 0xfa, 0xa0, 0x7a, 0xf2, 0xde, 0x03, 0xbc, 0x05, 0x88, 0xbe, 0x2e,
	0x29, 0x7a, 0x67, 0x2b, 0x3e, 0x13, 0x4c, 0x9e, 0xab, 0xd1, 0x0f, 0xae, 0xe0, 0x7a, 0x6f, 0x28,
	0xea, 0x43, 0x53, 0x99, 0x68, 0x2b, 0x6d, 0xa2, 0x76, 0xc6, 0x79, 0xe2, 0x89, 0x60, 0x32, 0x51,
	0xd8, 0xc2, 0xb7, 0x9e, 0x3d, 0xc7, 0x75, 0x4b, 0xf0, 0x0e, 0x96, 0x27, 0x23, 0xe5, 0x91, 0xcf,
	0x05, 0x93, 0xa9, 0xba, 0x1c, 0x85, 0x5f, 0x8e, 0xf9, 0x07, 0xa4, 0xc3, 0xfd, 0x88, 0x30, 0x6f,
	0x45, 0xbb, 0x43, 0x53, 0xd5, 0xbd, 0xf1, 0x11, 0x52, 0xfa, 0xf5, 0x22, 0x3e, 0x13, 0x89, 0xbc,
	0xb8, 0xbf, 0x29, 0x7a, 0xfb, 0x62, 0x62, 0xae, 0xfe, 0xb3, 0xf9, 0x13, 0xc0, 0xc0, 0x09, 0x57,
	0x00, 0x43, 0x87, 0xc4, 0x59, 0xb7, 0x27, 0x9b, 0xec, 0x51, 0xa3, 0x50, 0xb9, 0xe8, 0xaa, 0x7e,
	0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xb4, 0xf1, 0x39, 0x7d, 0x01, 0x00, 0x00,
}