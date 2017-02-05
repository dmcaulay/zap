// Code generated by protoc-gen-go.
// source: benchmarks.proto
// DO NOT EDIT!

/*
Package benchmarks is a generated protocol buffer package.

It is generated from these files:
	benchmarks.proto

It has these top-level messages:
	BenchmarkContext
*/
package benchmarks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BenchmarkContext struct {
	Int32Field           int32                      `protobuf:"varint,1,opt,name=int32_field,json=int32Field" json:"int32_field,omitempty"`
	Int64Field           int64                      `protobuf:"varint,2,opt,name=int64_field,json=int64Field" json:"int64_field,omitempty"`
	FloatField           float32                    `protobuf:"fixed32,3,opt,name=float_field,json=floatField" json:"float_field,omitempty"`
	StringField          string                     `protobuf:"bytes,4,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	BoolField            bool                       `protobuf:"varint,5,opt,name=bool_field,json=boolField" json:"bool_field,omitempty"`
	TimestampField       *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=timestamp_field,json=timestampField" json:"timestamp_field,omitempty"`
	ErrorField           string                     `protobuf:"bytes,7,opt,name=error_field,json=errorField" json:"error_field,omitempty"`
	DurationField        int64                      `protobuf:"varint,8,opt,name=duration_field,json=durationField" json:"duration_field,omitempty"`
	UserDefinedTypeField []byte                     `protobuf:"bytes,9,opt,name=user_defined_type_field,json=userDefinedTypeField,proto3" json:"user_defined_type_field,omitempty"`
	AnotherStringField   string                     `protobuf:"bytes,10,opt,name=another_string_field,json=anotherStringField" json:"another_string_field,omitempty"`
}

func (m *BenchmarkContext) Reset()                    { *m = BenchmarkContext{} }
func (m *BenchmarkContext) String() string            { return proto.CompactTextString(m) }
func (*BenchmarkContext) ProtoMessage()               {}
func (*BenchmarkContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BenchmarkContext) GetInt32Field() int32 {
	if m != nil {
		return m.Int32Field
	}
	return 0
}

func (m *BenchmarkContext) GetInt64Field() int64 {
	if m != nil {
		return m.Int64Field
	}
	return 0
}

func (m *BenchmarkContext) GetFloatField() float32 {
	if m != nil {
		return m.FloatField
	}
	return 0
}

func (m *BenchmarkContext) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

func (m *BenchmarkContext) GetBoolField() bool {
	if m != nil {
		return m.BoolField
	}
	return false
}

func (m *BenchmarkContext) GetTimestampField() *google_protobuf.Timestamp {
	if m != nil {
		return m.TimestampField
	}
	return nil
}

func (m *BenchmarkContext) GetErrorField() string {
	if m != nil {
		return m.ErrorField
	}
	return ""
}

func (m *BenchmarkContext) GetDurationField() int64 {
	if m != nil {
		return m.DurationField
	}
	return 0
}

func (m *BenchmarkContext) GetUserDefinedTypeField() []byte {
	if m != nil {
		return m.UserDefinedTypeField
	}
	return nil
}

func (m *BenchmarkContext) GetAnotherStringField() string {
	if m != nil {
		return m.AnotherStringField
	}
	return ""
}

func init() {
	proto.RegisterType((*BenchmarkContext)(nil), "uber.zap.benchmarks.BenchmarkContext")
}

func init() { proto.RegisterFile("benchmarks.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0xd1, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x06, 0xe0, 0x14, 0x04, 0xe1, 0x80, 0x48, 0x90, 0x44, 0x43, 0x62, 0xa8, 0x26, 0x26, 0xbd,
	0x1a, 0x06, 0x90, 0x07, 0x00, 0xe3, 0x03, 0x4c, 0xae, 0xbc, 0x59, 0x3a, 0x77, 0x06, 0x8d, 0xa3,
	0x5d, 0xba, 0x2e, 0x11, 0xdf, 0xc9, 0x77, 0x34, 0xb4, 0xdd, 0x88, 0xb7, 0xff, 0xf9, 0xce, 0x72,
	0xfe, 0x0e, 0x86, 0x31, 0xca, 0xcf, 0xfd, 0x81, 0xeb, 0xaf, 0x22, 0xc8, 0xb5, 0x32, 0x6a, 0x74,
	0x53, 0xc6, 0xa8, 0x83, 0x1f, 0x9e, 0x07, 0xe7, 0xd1, 0x64, 0xba, 0x53, 0x6a, 0x97, 0xe1, 0xcc,
	0x92, 0xb8, 0x4c, 0x67, 0x46, 0x1c, 0xb0, 0x30, 0xfc, 0x90, 0xbb, 0xad, 0xc7, 0xdf, 0x26, 0x0c,
	0xd7, 0x95, 0xdf, 0x28, 0x69, 0xf0, 0xdb, 0x8c, 0xa6, 0xd0, 0x13, 0xd2, 0x2c, 0xe6, 0x51, 0x2a,
	0x30, 0x4b, 0xee, 0x08, 0x25, 0xac, 0x15, 0x82, 0x8d, 0xde, 0x4e, 0x89, 0x07, 0xab, 0xa5, 0x07,
	0x0d, 0x4a, 0x58, 0xd3, 0x82, 0xd5, 0xb2, 0x06, 0x69, 0xa6, 0xb8, 0xf1, 0xa0, 0x49, 0x09, 0x6b,
	0x84, 0x60, 0x23, 0x07, 0x1e, 0xa0, 0x5f, 0x18, 0x2d, 0xe4, 0xce, 0x8b, 0x0b, 0x4a, 0x58, 0x37,
	0xec, 0xb9, 0xcc, 0x91, 0x7b, 0x80, 0x58, 0xa9, 0xcc, 0x83, 0x16, 0x25, 0xac, 0x13, 0x76, 0x4f,
	0x89, 0x1b, 0x6f, 0xe0, 0xba, 0x2e, 0xe3, 0x4d, 0x9b, 0x12, 0xd6, 0x9b, 0x4f, 0x02, 0x57, 0x3a,
	0xa8, 0x4a, 0x07, 0xdb, 0xca, 0x85, 0x83, 0x7a, 0xa5, 0xbe, 0x13, 0xb5, 0x56, 0xda, 0x7f, 0xe0,
	0xd2, 0x5e, 0x01, 0x36, 0x72, 0xe0, 0x09, 0x06, 0x49, 0xa9, 0xb9, 0x11, 0x4a, 0x7a, 0xd3, 0xb1,
	0x65, 0xaf, 0xaa, 0xd4, 0xb1, 0x17, 0xb8, 0x2d, 0x0b, 0xd4, 0x51, 0x82, 0xa9, 0x90, 0x98, 0x44,
	0xe6, 0x98, 0xa3, 0xf7, 0x5d, 0x4a, 0x58, 0x3f, 0x1c, 0x9f, 0xc6, 0xaf, 0x6e, 0xba, 0x3d, 0xe6,
	0xe8, 0xd6, 0x9e, 0x61, 0xcc, 0xa5, 0x32, 0x7b, 0xd4, 0xd1, 0xbf, 0xd7, 0x00, 0x7b, 0xc7, 0xc8,
	0xcf, 0xde, 0xcf, 0x8f, 0xb2, 0xee, 0x7f, 0xc0, 0xf9, 0xf7, 0xc6, 0x6d, 0x5b, 0x71, 0xf1, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0x03, 0xcd, 0x26, 0x0e, 0x02, 0x00, 0x00,
}
