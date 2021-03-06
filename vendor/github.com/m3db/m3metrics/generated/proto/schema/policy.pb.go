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
// source: policy.proto
// DO NOT EDIT!

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AggregationType int32

const (
	AggregationType_UNKNOWN AggregationType = 0
	AggregationType_LAST    AggregationType = 1
	AggregationType_MIN     AggregationType = 2
	AggregationType_MAX     AggregationType = 3
	AggregationType_MEAN    AggregationType = 4
	AggregationType_MEDIAN  AggregationType = 5
	AggregationType_COUNT   AggregationType = 6
	AggregationType_SUM     AggregationType = 7
	AggregationType_SUMSQ   AggregationType = 8
	AggregationType_STDEV   AggregationType = 9
	AggregationType_P10     AggregationType = 10
	AggregationType_P20     AggregationType = 11
	AggregationType_P30     AggregationType = 12
	AggregationType_P40     AggregationType = 13
	AggregationType_P50     AggregationType = 14
	AggregationType_P60     AggregationType = 15
	AggregationType_P70     AggregationType = 16
	AggregationType_P80     AggregationType = 17
	AggregationType_P90     AggregationType = 18
	AggregationType_P95     AggregationType = 19
	AggregationType_P99     AggregationType = 20
	AggregationType_P999    AggregationType = 21
	AggregationType_P9999   AggregationType = 22
)

var AggregationType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "LAST",
	2:  "MIN",
	3:  "MAX",
	4:  "MEAN",
	5:  "MEDIAN",
	6:  "COUNT",
	7:  "SUM",
	8:  "SUMSQ",
	9:  "STDEV",
	10: "P10",
	11: "P20",
	12: "P30",
	13: "P40",
	14: "P50",
	15: "P60",
	16: "P70",
	17: "P80",
	18: "P90",
	19: "P95",
	20: "P99",
	21: "P999",
	22: "P9999",
}
var AggregationType_value = map[string]int32{
	"UNKNOWN": 0,
	"LAST":    1,
	"MIN":     2,
	"MAX":     3,
	"MEAN":    4,
	"MEDIAN":  5,
	"COUNT":   6,
	"SUM":     7,
	"SUMSQ":   8,
	"STDEV":   9,
	"P10":     10,
	"P20":     11,
	"P30":     12,
	"P40":     13,
	"P50":     14,
	"P60":     15,
	"P70":     16,
	"P80":     17,
	"P90":     18,
	"P95":     19,
	"P99":     20,
	"P999":    21,
	"P9999":   22,
}

func (x AggregationType) String() string {
	return proto.EnumName(AggregationType_name, int32(x))
}
func (AggregationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Resolution struct {
	WindowSize int64 `protobuf:"varint,1,opt,name=window_size,json=windowSize" json:"window_size,omitempty"`
	Precision  int64 `protobuf:"varint,2,opt,name=precision" json:"precision,omitempty"`
}

func (m *Resolution) Reset()                    { *m = Resolution{} }
func (m *Resolution) String() string            { return proto.CompactTextString(m) }
func (*Resolution) ProtoMessage()               {}
func (*Resolution) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Retention struct {
	Period int64 `protobuf:"varint,1,opt,name=period" json:"period,omitempty"`
}

func (m *Retention) Reset()                    { *m = Retention{} }
func (m *Retention) String() string            { return proto.CompactTextString(m) }
func (*Retention) ProtoMessage()               {}
func (*Retention) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type StoragePolicy struct {
	Resolution *Resolution `protobuf:"bytes,1,opt,name=resolution" json:"resolution,omitempty"`
	Retention  *Retention  `protobuf:"bytes,2,opt,name=retention" json:"retention,omitempty"`
}

func (m *StoragePolicy) Reset()                    { *m = StoragePolicy{} }
func (m *StoragePolicy) String() string            { return proto.CompactTextString(m) }
func (*StoragePolicy) ProtoMessage()               {}
func (*StoragePolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *StoragePolicy) GetResolution() *Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *StoragePolicy) GetRetention() *Retention {
	if m != nil {
		return m.Retention
	}
	return nil
}

type Policy struct {
	StoragePolicy    *StoragePolicy    `protobuf:"bytes,1,opt,name=storage_policy,json=storagePolicy" json:"storage_policy,omitempty"`
	AggregationTypes []AggregationType `protobuf:"varint,2,rep,packed,name=aggregation_types,json=aggregationTypes,enum=schema.AggregationType" json:"aggregation_types,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Policy) GetStoragePolicy() *StoragePolicy {
	if m != nil {
		return m.StoragePolicy
	}
	return nil
}

func init() {
	proto.RegisterType((*Resolution)(nil), "schema.Resolution")
	proto.RegisterType((*Retention)(nil), "schema.Retention")
	proto.RegisterType((*StoragePolicy)(nil), "schema.StoragePolicy")
	proto.RegisterType((*Policy)(nil), "schema.Policy")
	proto.RegisterEnum("schema.AggregationType", AggregationType_name, AggregationType_value)
}

func init() { proto.RegisterFile("policy.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xd2, 0x3a, 0xf5, 0xb8, 0x49, 0x27, 0x0b, 0x2d, 0x39, 0x20, 0x51, 0x85, 0x4b,
	0xc5, 0x21, 0x2c, 0x2e, 0x05, 0x2c, 0x71, 0xb1, 0x48, 0x0e, 0x55, 0xb1, 0x1b, 0xec, 0x04, 0xb8,
	0x45, 0x26, 0x5d, 0x99, 0x95, 0x8a, 0xd7, 0xf2, 0x1a, 0x55, 0xe9, 0x33, 0xf0, 0xb4, 0x3c, 0x01,
	0xda, 0xb1, 0x83, 0x49, 0x6f, 0x9f, 0x67, 0x7e, 0xff, 0xf3, 0xc9, 0x32, 0x1c, 0x16, 0xea, 0x56,
	0xae, 0x37, 0x93, 0xa2, 0x54, 0x95, 0x62, 0xb6, 0x5e, 0xff, 0x10, 0x3f, 0xd3, 0xf1, 0x15, 0x40,
	0x2c, 0xb4, 0xba, 0xfd, 0x55, 0x49, 0x95, 0xb3, 0xe7, 0xe0, 0xde, 0xc9, 0xfc, 0x46, 0xdd, 0xad,
	0xb4, 0xbc, 0x17, 0x23, 0xeb, 0xd4, 0x3a, 0xeb, 0xc6, 0x50, 0x8f, 0x12, 0x79, 0x2f, 0xd8, 0x33,
	0x70, 0x8a, 0x52, 0xac, 0xa5, 0x96, 0x2a, 0x1f, 0x75, 0x68, 0xdd, 0x0e, 0xc6, 0x2f, 0xc0, 0x89,
	0x45, 0x25, 0x72, 0xea, 0x3a, 0x01, 0xbb, 0x10, 0xa5, 0x54, 0x37, 0x4d, 0x4d, 0xf3, 0x34, 0xae,
	0xa0, 0x9f, 0x54, 0xaa, 0x4c, 0x33, 0x31, 0x27, 0x21, 0xe6, 0x01, 0x94, 0xff, 0x14, 0x28, 0xec,
	0x7a, 0x6c, 0x52, 0xfb, 0x4d, 0x5a, 0xb9, 0xf8, 0xbf, 0x14, 0x7b, 0x05, 0x4e, 0xb9, 0xbd, 0x44,
	0x1e, 0xae, 0x37, 0x6c, 0x5f, 0x69, 0x16, 0x71, 0x9b, 0x19, 0xff, 0xb6, 0xc0, 0x6e, 0xee, 0x7d,
	0x80, 0x81, 0xae, 0x05, 0x56, 0xf5, 0x27, 0x69, 0x6e, 0x1e, 0x6f, 0x0b, 0x76, 0xf4, 0xe2, 0xbe,
	0xde, 0xb1, 0x9d, 0xc2, 0x30, 0xcd, 0xb2, 0x52, 0x64, 0xa9, 0xe9, 0x5d, 0x55, 0x9b, 0x42, 0xe8,
	0x51, 0xe7, 0xb4, 0x7b, 0x36, 0xf0, 0x9e, 0x6e, 0x0b, 0x82, 0x36, 0xb0, 0xd8, 0x14, 0x22, 0xc6,
	0x74, 0x77, 0xa0, 0x5f, 0xfe, 0xb1, 0xe0, 0xe8, 0x41, 0x8a, 0xb9, 0xd0, 0x5b, 0x46, 0x57, 0xd1,
	0xf5, 0xd7, 0x08, 0x1f, 0xb1, 0x03, 0xd8, 0xfb, 0x14, 0x24, 0x0b, 0xb4, 0x58, 0x0f, 0xba, 0xe1,
	0x65, 0x84, 0x1d, 0x82, 0xe0, 0x1b, 0x76, 0xcd, 0x2e, 0x9c, 0x05, 0x11, 0xee, 0x31, 0x00, 0x3b,
	0x9c, 0x4d, 0x2f, 0x83, 0x08, 0xf7, 0x99, 0x03, 0xfb, 0x1f, 0xaf, 0x97, 0xd1, 0x02, 0x6d, 0x93,
	0x4c, 0x96, 0x21, 0xf6, 0xcc, 0x2c, 0x59, 0x86, 0xc9, 0x67, 0x3c, 0x20, 0x5c, 0x4c, 0x67, 0x5f,
	0xd0, 0x31, 0xeb, 0xf9, 0x6b, 0x8e, 0x40, 0xe0, 0x71, 0x74, 0x09, 0xce, 0x39, 0x1e, 0x12, 0xbc,
	0xe1, 0xd8, 0x27, 0xb8, 0xe0, 0x38, 0x20, 0x78, 0xcb, 0xf1, 0x88, 0xe0, 0x1d, 0x47, 0x24, 0x78,
	0xcf, 0x71, 0x48, 0xe0, 0x73, 0x64, 0x35, 0x5c, 0xe0, 0xe3, 0x1a, 0x7c, 0x7c, 0x62, 0x14, 0xe7,
	0xbe, 0xef, 0xe3, 0xb1, 0xb9, 0x6b, 0xc8, 0xc7, 0x93, 0xef, 0x36, 0xfd, 0x7a, 0xe7, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x02, 0x6c, 0x63, 0x8a, 0x02, 0x00, 0x00,
}
