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
// source: placement.proto
// DO NOT EDIT!

/*
Package placementpb is a generated protocol buffer package.

It is generated from these files:
	placement.proto

It has these top-level messages:
	Placement
	Instance
	Shard
	PlacementSnapshots
*/
package placementpb

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

type ShardState int32

const (
	ShardState_INITIALIZING ShardState = 0
	ShardState_AVAILABLE    ShardState = 1
	ShardState_LEAVING      ShardState = 2
)

var ShardState_name = map[int32]string{
	0: "INITIALIZING",
	1: "AVAILABLE",
	2: "LEAVING",
}
var ShardState_value = map[string]int32{
	"INITIALIZING": 0,
	"AVAILABLE":    1,
	"LEAVING":      2,
}

func (x ShardState) String() string {
	return proto.EnumName(ShardState_name, int32(x))
}
func (ShardState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Placement struct {
	Instances     map[string]*Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReplicaFactor uint32               `protobuf:"varint,2,opt,name=replica_factor,json=replicaFactor" json:"replica_factor,omitempty"`
	NumShards     uint32               `protobuf:"varint,3,opt,name=num_shards,json=numShards" json:"num_shards,omitempty"`
	IsSharded     bool                 `protobuf:"varint,4,opt,name=is_sharded,json=isSharded" json:"is_sharded,omitempty"`
	// cutover_time is the placement-level cutover time and determines when the clients
	// watching the placement deems the placement as "in effect" and can use it to determine
	// shard placement.
	CutoverTime int64 `protobuf:"varint,5,opt,name=cutover_time,json=cutoverTime" json:"cutover_time,omitempty"`
	IsMirrored  bool  `protobuf:"varint,6,opt,name=is_mirrored,json=isMirrored" json:"is_mirrored,omitempty"`
	// max_shard_set_id stores the maximum shard set id used to guarantee unique
	// shard set id generations across placement changes.
	MaxShardSetId uint32 `protobuf:"varint,7,opt,name=max_shard_set_id,json=maxShardSetId" json:"max_shard_set_id,omitempty"`
}

func (m *Placement) Reset()                    { *m = Placement{} }
func (m *Placement) String() string            { return proto.CompactTextString(m) }
func (*Placement) ProtoMessage()               {}
func (*Placement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Placement) GetInstances() map[string]*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type Instance struct {
	Id         string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Rack       string   `protobuf:"bytes,2,opt,name=rack" json:"rack,omitempty"`
	Zone       string   `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	Weight     uint32   `protobuf:"varint,4,opt,name=weight" json:"weight,omitempty"`
	Endpoint   string   `protobuf:"bytes,5,opt,name=endpoint" json:"endpoint,omitempty"`
	Shards     []*Shard `protobuf:"bytes,6,rep,name=shards" json:"shards,omitempty"`
	ShardSetId uint32   `protobuf:"varint,7,opt,name=shard_set_id,json=shardSetId" json:"shard_set_id,omitempty"`
	Hostname   string   `protobuf:"bytes,8,opt,name=hostname" json:"hostname,omitempty"`
	Port       uint32   `protobuf:"varint,9,opt,name=port" json:"port,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Instance) GetShards() []*Shard {
	if m != nil {
		return m.Shards
	}
	return nil
}

type Shard struct {
	Id       uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	State    ShardState `protobuf:"varint,2,opt,name=state,enum=placementpb.ShardState" json:"state,omitempty"`
	SourceId string     `protobuf:"bytes,3,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	// Shard-level cutover and cutoff times determine when the shards have been cut over or
	// cut off from the source instance to the destination instance. The placement-level
	// cutover times are usually (but not required to be) earlier than shard-level cutover
	// times if the clients watching the placement need to send traffic to the shards before
	// they are ready to cut over or after they are ready to cut off (e.g., for warmup purposes).
	CutoverNanos int64 `protobuf:"varint,4,opt,name=cutover_nanos,json=cutoverNanos" json:"cutover_nanos,omitempty"`
	CutoffNanos  int64 `protobuf:"varint,5,opt,name=cutoff_nanos,json=cutoffNanos" json:"cutoff_nanos,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PlacementSnapshots struct {
	Snapshots []*Placement `protobuf:"bytes,1,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *PlacementSnapshots) Reset()                    { *m = PlacementSnapshots{} }
func (m *PlacementSnapshots) String() string            { return proto.CompactTextString(m) }
func (*PlacementSnapshots) ProtoMessage()               {}
func (*PlacementSnapshots) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlacementSnapshots) GetSnapshots() []*Placement {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func init() {
	proto.RegisterType((*Placement)(nil), "placementpb.Placement")
	proto.RegisterType((*Instance)(nil), "placementpb.Instance")
	proto.RegisterType((*Shard)(nil), "placementpb.Shard")
	proto.RegisterType((*PlacementSnapshots)(nil), "placementpb.PlacementSnapshots")
	proto.RegisterEnum("placementpb.ShardState", ShardState_name, ShardState_value)
}

func init() { proto.RegisterFile("placement.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x53, 0xd1, 0x6f, 0xd3, 0x3e,
	0x10, 0xfe, 0x25, 0x5d, 0xbb, 0xfa, 0xb2, 0xf4, 0x57, 0x9d, 0xc4, 0x88, 0x86, 0x10, 0xa1, 0x68,
	0xa2, 0x1a, 0xa2, 0x0f, 0x83, 0x07, 0xb4, 0xb7, 0x82, 0x06, 0x0a, 0x2a, 0x13, 0x72, 0xa6, 0x3d,
	0xf0, 0x12, 0x79, 0x89, 0x4b, 0xad, 0x2d, 0x76, 0x64, 0xbb, 0x63, 0xe3, 0x5f, 0xe2, 0xdf, 0x43,
	0xe2, 0x15, 0xc5, 0x49, 0xba, 0x4d, 0xec, 0xed, 0xee, 0xbb, 0xcf, 0xbe, 0xfb, 0x3e, 0x9f, 0xe1,
	0xff, 0xea, 0x92, 0xe5, 0xbc, 0xe4, 0xd2, 0xce, 0x2a, 0xad, 0xac, 0xc2, 0x60, 0x03, 0x54, 0xe7,
	0x93, 0x3f, 0x3e, 0x90, 0xaf, 0x5d, 0x8e, 0x1f, 0x80, 0x08, 0x69, 0x2c, 0x93, 0x39, 0x37, 0x91,
	0x17, 0xf7, 0xa6, 0xc1, 0xe1, 0xfe, 0xec, 0x0e, 0x7d, 0xb6, 0xa1, 0xce, 0x92, 0x8e, 0x77, 0x2c,
	0xad, 0xbe, 0xa1, 0xb7, 0xe7, 0x70, 0x1f, 0x46, 0x9a, 0x57, 0x97, 0x22, 0x67, 0xd9, 0x92, 0xe5,
	0x56, 0xe9, 0xc8, 0x8f, 0xbd, 0x69, 0x48, 0xc3, 0x16, 0xfd, 0xe8, 0x40, 0x7c, 0x0a, 0x20, 0xd7,
	0x65, 0x66, 0x56, 0x4c, 0x17, 0x26, 0xea, 0x39, 0x0a, 0x91, 0xeb, 0x32, 0x75, 0x40, 0x5d, 0x16,
	0xa6, 0xa9, 0xf2, 0x22, 0xda, 0x8a, 0xbd, 0xe9, 0x90, 0x12, 0x61, 0xd2, 0x06, 0xc0, 0xe7, 0xb0,
	0x93, 0xaf, 0xad, 0xba, 0xe2, 0x3a, 0xb3, 0xa2, 0xe4, 0x51, 0x3f, 0xf6, 0xa6, 0x3d, 0x1a, 0xb4,
	0xd8, 0xa9, 0x28, 0x39, 0x3e, 0x83, 0x40, 0x98, 0xac, 0x14, 0x5a, 0x2b, 0xcd, 0x8b, 0x68, 0xe0,
	0xae, 0x00, 0x61, 0xbe, 0xb4, 0x08, 0xbe, 0x84, 0x71, 0xc9, 0xae, 0x9b, 0x1e, 0x99, 0xe1, 0x36,
	0x13, 0x45, 0xb4, 0xdd, 0x8c, 0x5a, 0xb2, 0x6b, 0xd7, 0x29, 0xe5, 0x36, 0x29, 0xf6, 0x52, 0x18,
	0xdd, 0x97, 0x8b, 0x63, 0xe8, 0x5d, 0xf0, 0x9b, 0xc8, 0x8b, 0xbd, 0x29, 0xa1, 0x75, 0x88, 0xaf,
	0xa0, 0x7f, 0xc5, 0x2e, 0xd7, 0xdc, 0x89, 0x0d, 0x0e, 0x1f, 0xdd, 0xb3, 0xad, 0x3b, 0x4d, 0x1b,
	0xce, 0x91, 0xff, 0xce, 0x9b, 0xfc, 0xf6, 0x60, 0xd8, 0xe1, 0x38, 0x02, 0x5f, 0x14, 0xed, 0x75,
	0xbe, 0x28, 0x10, 0x61, 0x4b, 0xb3, 0xfc, 0xc2, 0x5d, 0x46, 0xa8, 0x8b, 0x6b, 0xec, 0xa7, 0x92,
	0xdc, 0x59, 0x45, 0xa8, 0x8b, 0x71, 0x17, 0x06, 0x3f, 0xb8, 0xf8, 0xbe, 0xb2, 0xce, 0xa1, 0x90,
	0xb6, 0x19, 0xee, 0xc1, 0x90, 0xcb, 0xa2, 0x52, 0x42, 0x5a, 0x67, 0x0d, 0xa1, 0x9b, 0x1c, 0x0f,
	0x60, 0xd0, 0x9a, 0x3e, 0x70, 0x2f, 0x8c, 0xf7, 0x46, 0x75, 0xb2, 0x69, 0xcb, 0xc0, 0x18, 0x76,
	0x1e, 0xb0, 0x07, 0xcc, 0xc6, 0x9b, 0xba, 0xd3, 0x4a, 0x19, 0x2b, 0x59, 0xc9, 0xa3, 0x61, 0xd3,
	0xa9, 0xcb, 0xeb, 0x89, 0x2b, 0xa5, 0x6d, 0x44, 0xdc, 0x29, 0x17, 0x4f, 0x7e, 0x79, 0xd0, 0x77,
	0x3d, 0xee, 0x68, 0x0e, 0x9d, 0xe6, 0xd7, 0xd0, 0x37, 0x96, 0xd9, 0xc6, 0xc1, 0xd1, 0xe1, 0xe3,
	0x7f, 0xc7, 0x4a, 0xeb, 0x32, 0x6d, 0x58, 0xf8, 0x04, 0x88, 0x51, 0x6b, 0x9d, 0xf3, 0x7a, 0xae,
	0xc6, 0x93, 0x61, 0x03, 0x24, 0x05, 0xbe, 0x80, 0xb0, 0x5b, 0x0f, 0xc9, 0xa4, 0x32, 0xce, 0x9e,
	0x1e, 0xed, 0x76, 0xe6, 0xa4, 0xc6, 0xba, 0x1d, 0x5a, 0x2e, 0x5b, 0xce, 0x9d, 0x1d, 0x5a, 0x2e,
	0x1d, 0x65, 0xf2, 0x19, 0x70, 0xb3, 0xf2, 0xa9, 0x64, 0x95, 0x59, 0x29, 0x6b, 0xf0, 0x2d, 0x10,
	0xd3, 0x25, 0xed, 0x37, 0xd9, 0x7d, 0xf8, 0x9b, 0xd0, 0x5b, 0xe2, 0xc1, 0x11, 0xc0, 0xad, 0x0a,
	0x1c, 0xc3, 0x4e, 0x72, 0x92, 0x9c, 0x26, 0xf3, 0x45, 0xf2, 0x2d, 0x39, 0xf9, 0x34, 0xfe, 0x0f,
	0x43, 0x20, 0xf3, 0xb3, 0x79, 0xb2, 0x98, 0xbf, 0x5f, 0x1c, 0x8f, 0x3d, 0x0c, 0x60, 0x7b, 0x71,
	0x3c, 0x3f, 0xab, 0x6b, 0xfe, 0xf9, 0xc0, 0x7d, 0xdd, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xef, 0xda, 0xb7, 0x14, 0xcd, 0x03, 0x00, 0x00,
}
