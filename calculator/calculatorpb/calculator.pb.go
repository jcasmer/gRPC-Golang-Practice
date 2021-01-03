// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumRequest struct {
	FirstNum             int64    `protobuf:"varint,1,opt,name=first_num,json=firstNum,proto3" json:"first_num,omitempty"`
	SecondNum            int64    `protobuf:"varint,2,opt,name=second_num,json=secondNum,proto3" json:"second_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstNum() int64 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

func (m *SumRequest) GetSecondNum() int64 {
	if m != nil {
		return m.SecondNum
	}
	return 0
}

type SumResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberManyTimesRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberManyTimesRequest) Reset()         { *m = PrimeNumberManyTimesRequest{} }
func (m *PrimeNumberManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberManyTimesRequest) ProtoMessage()    {}
func (*PrimeNumberManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *PrimeNumberManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberManyTimesRequest.Unmarshal(m, b)
}
func (m *PrimeNumberManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberManyTimesRequest.Merge(m, src)
}
func (m *PrimeNumberManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberManyTimesRequest.Size(m)
}
func (m *PrimeNumberManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberManyTimesRequest proto.InternalMessageInfo

func (m *PrimeNumberManyTimesRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumbertManyTimesResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumbertManyTimesResponse) Reset()         { *m = PrimeNumbertManyTimesResponse{} }
func (m *PrimeNumbertManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumbertManyTimesResponse) ProtoMessage()    {}
func (*PrimeNumbertManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeNumbertManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumbertManyTimesResponse.Unmarshal(m, b)
}
func (m *PrimeNumbertManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumbertManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumbertManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumbertManyTimesResponse.Merge(m, src)
}
func (m *PrimeNumbertManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumbertManyTimesResponse.Size(m)
}
func (m *PrimeNumbertManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumbertManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumbertManyTimesResponse proto.InternalMessageInfo

func (m *PrimeNumbertManyTimesResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAvgRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAvgRequest) Reset()         { *m = ComputeAvgRequest{} }
func (m *ComputeAvgRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAvgRequest) ProtoMessage()    {}
func (*ComputeAvgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *ComputeAvgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAvgRequest.Unmarshal(m, b)
}
func (m *ComputeAvgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAvgRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAvgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAvgRequest.Merge(m, src)
}
func (m *ComputeAvgRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAvgRequest.Size(m)
}
func (m *ComputeAvgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAvgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAvgRequest proto.InternalMessageInfo

func (m *ComputeAvgRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAvgResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAvgResponse) Reset()         { *m = ComputeAvgResponse{} }
func (m *ComputeAvgResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAvgResponse) ProtoMessage()    {}
func (*ComputeAvgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *ComputeAvgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAvgResponse.Unmarshal(m, b)
}
func (m *ComputeAvgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAvgResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAvgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAvgResponse.Merge(m, src)
}
func (m *ComputeAvgResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAvgResponse.Size(m)
}
func (m *ComputeAvgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAvgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAvgResponse proto.InternalMessageInfo

func (m *ComputeAvgResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaximunRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximunRequest) Reset()         { *m = FindMaximunRequest{} }
func (m *FindMaximunRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximunRequest) ProtoMessage()    {}
func (*FindMaximunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaximunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximunRequest.Unmarshal(m, b)
}
func (m *FindMaximunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximunRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximunRequest.Merge(m, src)
}
func (m *FindMaximunRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximunRequest.Size(m)
}
func (m *FindMaximunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximunRequest proto.InternalMessageInfo

func (m *FindMaximunRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximunResponse struct {
	Maximum              int32    `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximunResponse) Reset()         { *m = FindMaximunResponse{} }
func (m *FindMaximunResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximunResponse) ProtoMessage()    {}
func (*FindMaximunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *FindMaximunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximunResponse.Unmarshal(m, b)
}
func (m *FindMaximunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximunResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximunResponse.Merge(m, src)
}
func (m *FindMaximunResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximunResponse.Size(m)
}
func (m *FindMaximunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximunResponse proto.InternalMessageInfo

func (m *FindMaximunResponse) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	DoubleRoot           float64  `protobuf:"fixed64,1,opt,name=double_root,json=doubleRoot,proto3" json:"double_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetDoubleRoot() float64 {
	if m != nil {
		return m.DoubleRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeNumberManyTimesRequest)(nil), "calculator.PrimeNumberManyTimesRequest")
	proto.RegisterType((*PrimeNumbertManyTimesResponse)(nil), "calculator.PrimeNumbertManyTimesResponse")
	proto.RegisterType((*ComputeAvgRequest)(nil), "calculator.ComputeAvgRequest")
	proto.RegisterType((*ComputeAvgResponse)(nil), "calculator.ComputeAvgResponse")
	proto.RegisterType((*FindMaximunRequest)(nil), "calculator.FindMaximunRequest")
	proto.RegisterType((*FindMaximunResponse)(nil), "calculator.FindMaximunResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x51, 0x8b, 0xda, 0x40,
	0x10, 0x36, 0x15, 0x6d, 0x1d, 0x4b, 0xc1, 0x69, 0xb1, 0x12, 0xb1, 0x96, 0x40, 0xa9, 0xa5, 0xa2,
	0xd2, 0x22, 0xed, 0x6b, 0x2b, 0x94, 0xbe, 0x68, 0x4b, 0xe2, 0xd3, 0xbd, 0x48, 0x12, 0xf7, 0x8e,
	0x40, 0x76, 0x37, 0x6e, 0x76, 0xe5, 0xee, 0x77, 0xdd, 0x1f, 0x3c, 0x5c, 0x13, 0xb3, 0x21, 0x7a,
	0xbe, 0xed, 0x7c, 0xf3, 0x7d, 0xdf, 0x0c, 0x7c, 0xb3, 0x30, 0x0a, 0xfd, 0x38, 0x54, 0xb1, 0x2f,
	0xb9, 0x98, 0x16, 0xcf, 0x24, 0x30, 0x8a, 0x49, 0x22, 0xb8, 0xe4, 0x08, 0x05, 0xe2, 0xfc, 0x05,
	0xf0, 0x14, 0x75, 0xc9, 0x4e, 0x91, 0x54, 0x62, 0x1f, 0x5a, 0xb7, 0x91, 0x48, 0xe5, 0x86, 0x29,
	0xda, 0xb3, 0x3e, 0x5a, 0xa3, 0xba, 0xfb, 0x4a, 0x03, 0x2b, 0x45, 0x71, 0x00, 0x90, 0x92, 0x90,
	0xb3, 0xad, 0xee, 0xbe, 0xd0, 0xdd, 0xd6, 0x11, 0x59, 0x29, 0xea, 0x7c, 0x82, 0xb6, 0x76, 0x4a,
	0x13, 0xce, 0x52, 0x82, 0x5d, 0x68, 0x0a, 0x92, 0xaa, 0x58, 0x66, 0x3e, 0x59, 0xe5, 0xcc, 0xa1,
	0xff, 0x5f, 0x44, 0x94, 0xac, 0x14, 0x0d, 0x88, 0x58, 0xfa, 0xec, 0x61, 0x1d, 0x51, 0x92, 0xe6,
	0x1b, 0x74, 0xa1, 0xc9, 0x74, 0x47, 0xcb, 0x1a, 0x6e, 0x56, 0x39, 0x3f, 0x60, 0x60, 0xc8, 0xa4,
	0xa1, 0x3b, 0x3b, 0xaf, 0x71, 0x9a, 0xf7, 0x15, 0x3a, 0x0b, 0x4e, 0x13, 0x25, 0xc9, 0xaf, 0xfd,
	0xdd, 0xb5, 0x29, 0x63, 0x40, 0x93, 0x7c, 0xd6, 0xda, 0x3a, 0x59, 0x8f, 0x01, 0xff, 0x44, 0x6c,
	0xbb, 0xf4, 0xef, 0x23, 0xaa, 0xd8, 0x35, 0xef, 0x29, 0xbc, 0x2d, 0xb1, 0x33, 0xf3, 0x1e, 0xbc,
	0xa4, 0x1a, 0xa2, 0x19, 0x3f, 0x2f, 0x0f, 0x9b, 0x7b, 0x3b, 0xe5, 0x0b, 0xe2, 0x72, 0x2e, 0xaf,
	0xb9, 0xcf, 0x01, 0x4d, 0x72, 0x66, 0x3e, 0x84, 0xf6, 0x96, 0xab, 0x20, 0x26, 0x1b, 0xc1, 0x79,
	0xbe, 0x3e, 0x1c, 0xa1, 0x03, 0xf1, 0xdb, 0x63, 0x1d, 0x3a, 0x8b, 0xd3, 0x35, 0x78, 0x44, 0xec,
	0xa3, 0x90, 0xe0, 0x4f, 0xa8, 0x7b, 0x8a, 0x62, 0x77, 0x62, 0x9c, 0x4e, 0x71, 0x25, 0xf6, 0xfb,
	0x0a, 0x7e, 0x1c, 0xe7, 0xd4, 0x90, 0xc1, 0xbb, 0x73, 0xe9, 0xe2, 0x67, 0x53, 0xf2, 0x4c, 0xfe,
	0xf6, 0x97, 0x0b, 0xc4, 0x6a, 0xe2, 0x4e, 0x6d, 0x66, 0xe1, 0x3f, 0x80, 0x22, 0x30, 0x1c, 0x98,
	0xe2, 0x4a, 0xea, 0xf6, 0x87, 0x4b, 0xed, 0xdc, 0x70, 0x64, 0xe1, 0x1a, 0xda, 0x46, 0x4a, 0x58,
	0x92, 0x54, 0xc3, 0xb6, 0x87, 0x17, 0xfb, 0x85, 0xe7, 0xcc, 0xc2, 0x25, 0x40, 0x91, 0x4e, 0x79,
	0xcd, 0x4a, 0xc4, 0xe5, 0x35, 0xab, 0xa1, 0x3a, 0xb5, 0xdf, 0x6f, 0x6e, 0x5e, 0x9b, 0x3f, 0x3c,
	0x68, 0xea, 0x7f, 0xfd, 0xfd, 0x29, 0x00, 0x00, 0xff, 0xff, 0x09, 0x46, 0xcf, 0x9e, 0x03, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server Streaming
	PrimeNumberManyTimes(ctx context.Context, in *PrimeNumberManyTimesRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberManyTimesClient, error)
	// Client Streaming
	ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAvgClient, error)
	// BiDi Streaming
	FindMaximun(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximunClient, error)
	// error handling
	// this RPC throw an exception if the sent number is negative
	// the error being sent is of type INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberManyTimes(ctx context.Context, in *PrimeNumberManyTimesRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumberManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberManyTimesClient interface {
	Recv() (*PrimeNumbertManyTimesResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberManyTimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberManyTimesClient) Recv() (*PrimeNumbertManyTimesResponse, error) {
	m := new(PrimeNumbertManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAvgClient{stream}
	return x, nil
}

type CalculatorService_ComputeAvgClient interface {
	Send(*ComputeAvgRequest) error
	CloseAndRecv() (*ComputeAvgResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAvgClient) Send(m *ComputeAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAvgClient) CloseAndRecv() (*ComputeAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximun(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximun", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximunClient{stream}
	return x, nil
}

type CalculatorService_FindMaximunClient interface {
	Send(*FindMaximunRequest) error
	Recv() (*FindMaximunResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximunClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximunClient) Send(m *FindMaximunRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximunClient) Recv() (*FindMaximunResponse, error) {
	m := new(FindMaximunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server Streaming
	PrimeNumberManyTimes(*PrimeNumberManyTimesRequest, CalculatorService_PrimeNumberManyTimesServer) error
	// Client Streaming
	ComputeAvg(CalculatorService_ComputeAvgServer) error
	// BiDi Streaming
	FindMaximun(CalculatorService_FindMaximunServer) error
	// error handling
	// this RPC throw an exception if the sent number is negative
	// the error being sent is of type INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberManyTimes(req *PrimeNumberManyTimesRequest, srv CalculatorService_PrimeNumberManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberManyTimes not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAvg(srv CalculatorService_ComputeAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvg not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximun(srv CalculatorService_FindMaximunServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximun not implemented")
}
func (*UnimplementedCalculatorServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberManyTimes(m, &calculatorServicePrimeNumberManyTimesServer{stream})
}

type CalculatorService_PrimeNumberManyTimesServer interface {
	Send(*PrimeNumbertManyTimesResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberManyTimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberManyTimesServer) Send(m *PrimeNumbertManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAvg(&calculatorServiceComputeAvgServer{stream})
}

type CalculatorService_ComputeAvgServer interface {
	SendAndClose(*ComputeAvgResponse) error
	Recv() (*ComputeAvgRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAvgServer) SendAndClose(m *ComputeAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAvgServer) Recv() (*ComputeAvgRequest, error) {
	m := new(ComputeAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximun(&calculatorServiceFindMaximunServer{stream})
}

type CalculatorService_FindMaximunServer interface {
	Send(*FindMaximunResponse) error
	Recv() (*FindMaximunRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximunServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximunServer) Send(m *FindMaximunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximunServer) Recv() (*FindMaximunRequest, error) {
	m := new(FindMaximunRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculatorService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberManyTimes",
			Handler:       _CalculatorService_PrimeNumberManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvg",
			Handler:       _CalculatorService_ComputeAvg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximun",
			Handler:       _CalculatorService_FindMaximun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
