// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	BaudRate
	Frame
	ToDevice
	Response
	Status
	FromDevice
*/
package messages

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

type BaudRate_Rate int32

const (
	BaudRate_K1000 BaudRate_Rate = 0
	BaudRate_K800  BaudRate_Rate = 1
	BaudRate_K500  BaudRate_Rate = 2
	BaudRate_K250  BaudRate_Rate = 3
	BaudRate_K125  BaudRate_Rate = 4
	BaudRate_K100  BaudRate_Rate = 5
	BaudRate_K83   BaudRate_Rate = 6
	BaudRate_K50   BaudRate_Rate = 7
	BaudRate_K20   BaudRate_Rate = 8
	BaudRate_AUTO  BaudRate_Rate = 9
)

var BaudRate_Rate_name = map[int32]string{
	0: "K1000",
	1: "K800",
	2: "K500",
	3: "K250",
	4: "K125",
	5: "K100",
	6: "K83",
	7: "K50",
	8: "K20",
	9: "AUTO",
}
var BaudRate_Rate_value = map[string]int32{
	"K1000": 0,
	"K800":  1,
	"K500":  2,
	"K250":  3,
	"K125":  4,
	"K100":  5,
	"K83":   6,
	"K50":   7,
	"K20":   8,
	"AUTO":  9,
}

func (x BaudRate_Rate) String() string {
	return proto.EnumName(BaudRate_Rate_name, int32(x))
}
func (BaudRate_Rate) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Frame_IDE int32

const (
	Frame_STD Frame_IDE = 0
	Frame_EXT Frame_IDE = 4
)

var Frame_IDE_name = map[int32]string{
	0: "STD",
	4: "EXT",
}
var Frame_IDE_value = map[string]int32{
	"STD": 0,
	"EXT": 4,
}

func (x Frame_IDE) String() string {
	return proto.EnumName(Frame_IDE_name, int32(x))
}
func (Frame_IDE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Frame_RTR int32

const (
	Frame_DATA   Frame_RTR = 0
	Frame_REMOTE Frame_RTR = 2
)

var Frame_RTR_name = map[int32]string{
	0: "DATA",
	2: "REMOTE",
}
var Frame_RTR_value = map[string]int32{
	"DATA":   0,
	"REMOTE": 2,
}

func (x Frame_RTR) String() string {
	return proto.EnumName(Frame_RTR_name, int32(x))
}
func (Frame_RTR) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type ToDevice_MessageType int32

const (
	ToDevice_SEND_FRAME   ToDevice_MessageType = 0
	ToDevice_INIT         ToDevice_MessageType = 1
	ToDevice_DEINIT       ToDevice_MessageType = 2
	ToDevice_SET_FILTER   ToDevice_MessageType = 3
	ToDevice_SET_BAUDRATE ToDevice_MessageType = 4
	ToDevice_GET_STATUS   ToDevice_MessageType = 5
)

var ToDevice_MessageType_name = map[int32]string{
	0: "SEND_FRAME",
	1: "INIT",
	2: "DEINIT",
	3: "SET_FILTER",
	4: "SET_BAUDRATE",
	5: "GET_STATUS",
}
var ToDevice_MessageType_value = map[string]int32{
	"SEND_FRAME":   0,
	"INIT":         1,
	"DEINIT":       2,
	"SET_FILTER":   3,
	"SET_BAUDRATE": 4,
	"GET_STATUS":   5,
}

func (x ToDevice_MessageType) String() string {
	return proto.EnumName(ToDevice_MessageType_name, int32(x))
}
func (ToDevice_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type FromDevice_MessageType int32

const (
	FromDevice_GET_FRAME FromDevice_MessageType = 0
	FromDevice_STATUS    FromDevice_MessageType = 1
	FromDevice_RESPONSE  FromDevice_MessageType = 2
)

var FromDevice_MessageType_name = map[int32]string{
	0: "GET_FRAME",
	1: "STATUS",
	2: "RESPONSE",
}
var FromDevice_MessageType_value = map[string]int32{
	"GET_FRAME": 0,
	"STATUS":    1,
	"RESPONSE":  2,
}

func (x FromDevice_MessageType) String() string {
	return proto.EnumName(FromDevice_MessageType_name, int32(x))
}
func (FromDevice_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type BaudRate struct {
	Rate BaudRate_Rate `protobuf:"varint,1,opt,name=rate,enum=BaudRate_Rate" json:"rate,omitempty"`
}

func (m *BaudRate) Reset()                    { *m = BaudRate{} }
func (m *BaudRate) String() string            { return proto.CompactTextString(m) }
func (*BaudRate) ProtoMessage()               {}
func (*BaudRate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BaudRate) GetRate() BaudRate_Rate {
	if m != nil {
		return m.Rate
	}
	return BaudRate_K1000
}

type Frame struct {
	Id   int32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Dlc  int32     `protobuf:"varint,2,opt,name=dlc" json:"dlc,omitempty"`
	Ide  Frame_IDE `protobuf:"varint,3,opt,name=ide,enum=Frame_IDE" json:"ide,omitempty"`
	Rtr  Frame_RTR `protobuf:"varint,4,opt,name=rtr,enum=Frame_RTR" json:"rtr,omitempty"`
	Data [][]byte  `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Frame) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Frame) GetDlc() int32 {
	if m != nil {
		return m.Dlc
	}
	return 0
}

func (m *Frame) GetIde() Frame_IDE {
	if m != nil {
		return m.Ide
	}
	return Frame_STD
}

func (m *Frame) GetRtr() Frame_RTR {
	if m != nil {
		return m.Rtr
	}
	return Frame_DATA
}

func (m *Frame) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToDevice struct {
	Type     ToDevice_MessageType `protobuf:"varint,1,opt,name=type,enum=ToDevice_MessageType" json:"type,omitempty"`
	Id       int32                `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Frame    *Frame               `protobuf:"bytes,3,opt,name=frame" json:"frame,omitempty"`
	BaudRate *BaudRate            `protobuf:"bytes,4,opt,name=baudRate" json:"baudRate,omitempty"`
}

func (m *ToDevice) Reset()                    { *m = ToDevice{} }
func (m *ToDevice) String() string            { return proto.CompactTextString(m) }
func (*ToDevice) ProtoMessage()               {}
func (*ToDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ToDevice) GetType() ToDevice_MessageType {
	if m != nil {
		return m.Type
	}
	return ToDevice_SEND_FRAME
}

func (m *ToDevice) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDevice) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *ToDevice) GetBaudRate() *BaudRate {
	if m != nil {
		return m.BaudRate
	}
	return nil
}

type Response struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Status struct {
	Rate BaudRate_Rate `protobuf:"varint,1,opt,name=rate,enum=BaudRate_Rate" json:"rate,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Status) GetRate() BaudRate_Rate {
	if m != nil {
		return m.Rate
	}
	return BaudRate_K1000
}

type FromDevice struct {
	Type     FromDevice_MessageType `protobuf:"varint,1,opt,name=type,enum=FromDevice_MessageType" json:"type,omitempty"`
	Id       int32                  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Frame    *Frame                 `protobuf:"bytes,3,opt,name=frame" json:"frame,omitempty"`
	Response *Response              `protobuf:"bytes,4,opt,name=response" json:"response,omitempty"`
	Status   *Status                `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *FromDevice) Reset()                    { *m = FromDevice{} }
func (m *FromDevice) String() string            { return proto.CompactTextString(m) }
func (*FromDevice) ProtoMessage()               {}
func (*FromDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FromDevice) GetType() FromDevice_MessageType {
	if m != nil {
		return m.Type
	}
	return FromDevice_GET_FRAME
}

func (m *FromDevice) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FromDevice) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *FromDevice) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *FromDevice) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*BaudRate)(nil), "BaudRate")
	proto.RegisterType((*Frame)(nil), "Frame")
	proto.RegisterType((*ToDevice)(nil), "ToDevice")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*FromDevice)(nil), "FromDevice")
	proto.RegisterEnum("BaudRate_Rate", BaudRate_Rate_name, BaudRate_Rate_value)
	proto.RegisterEnum("Frame_IDE", Frame_IDE_name, Frame_IDE_value)
	proto.RegisterEnum("Frame_RTR", Frame_RTR_name, Frame_RTR_value)
	proto.RegisterEnum("ToDevice_MessageType", ToDevice_MessageType_name, ToDevice_MessageType_value)
	proto.RegisterEnum("FromDevice_MessageType", FromDevice_MessageType_name, FromDevice_MessageType_value)
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdf, 0x8a, 0xd3, 0x4e,
	0x18, 0xdd, 0xfc, 0x6d, 0xfa, 0xb5, 0xbf, 0x32, 0x0c, 0x3f, 0xd9, 0x80, 0x0b, 0x2e, 0x03, 0xc2,
	0x8a, 0x12, 0x66, 0xb3, 0x54, 0xf6, 0x36, 0x4b, 0xa6, 0x52, 0x6a, 0x5b, 0xf9, 0x32, 0x05, 0xef,
	0x4a, 0x76, 0x33, 0x6a, 0x59, 0x6b, 0x4a, 0x92, 0x15, 0xf6, 0x4d, 0x7c, 0x0b, 0x1f, 0xcb, 0x6b,
	0xdf, 0x40, 0x66, 0x92, 0x56, 0x7b, 0xe7, 0x85, 0x37, 0xe1, 0xcc, 0x77, 0x4e, 0xc8, 0x39, 0xe7,
	0xcb, 0xc0, 0x68, 0xab, 0xea, 0x3a, 0xff, 0xa8, 0xea, 0x68, 0x57, 0x95, 0x4d, 0xc9, 0xbe, 0x59,
	0x10, 0xdc, 0xe4, 0x0f, 0x05, 0xe6, 0x8d, 0xa2, 0x0c, 0xdc, 0x2a, 0x6f, 0x54, 0x68, 0x9d, 0x5b,
	0x17, 0xa3, 0x78, 0x14, 0xed, 0x89, 0x48, 0x3f, 0xd0, 0x70, 0xec, 0x13, 0xb8, 0x46, 0xdb, 0x07,
	0x6f, 0x76, 0xc9, 0x39, 0x27, 0x27, 0x34, 0x00, 0x77, 0x76, 0xcd, 0x39, 0xb1, 0x0c, 0x1a, 0x73,
	0x4e, 0x6c, 0x83, 0xe2, 0x31, 0x27, 0x8e, 0x41, 0x97, 0xf1, 0x98, 0xb8, 0x2d, 0xe2, 0x9c, 0x78,
	0xb4, 0x07, 0xce, 0xec, 0xfa, 0x8a, 0xf8, 0x06, 0x8c, 0x39, 0xe9, 0x19, 0x10, 0x73, 0x12, 0x68,
	0x51, 0xb2, 0x92, 0x4b, 0xd2, 0x67, 0xdf, 0x2d, 0xf0, 0x26, 0x55, 0xbe, 0x55, 0x74, 0x04, 0xf6,
	0xa6, 0x30, 0xae, 0x3c, 0xb4, 0x37, 0x05, 0x25, 0xe0, 0x14, 0x9f, 0xef, 0x42, 0xdb, 0x0c, 0x34,
	0xa4, 0x67, 0xe0, 0x6c, 0x0a, 0x15, 0x3a, 0xc6, 0x38, 0x44, 0xe6, 0xb5, 0x68, 0x9a, 0x0a, 0xd4,
	0x63, 0xcd, 0x56, 0x4d, 0x15, 0xba, 0x47, 0x2c, 0x4a, 0x44, 0x3d, 0xa6, 0x14, 0xdc, 0x22, 0x6f,
	0xf2, 0xd0, 0x3b, 0x77, 0x2e, 0x86, 0x68, 0x30, 0x3b, 0x05, 0x67, 0x9a, 0x0a, 0xed, 0x2a, 0x93,
	0x29, 0x39, 0xd1, 0x40, 0xbc, 0x97, 0xc4, 0x65, 0x4f, 0xc1, 0x41, 0x89, 0xda, 0x65, 0x9a, 0xc8,
	0x84, 0x9c, 0x50, 0x00, 0x1f, 0xc5, 0x7c, 0x29, 0x05, 0xb1, 0xd9, 0x4f, 0x0b, 0x02, 0x59, 0xa6,
	0xea, 0xeb, 0xe6, 0x4e, 0xd1, 0x17, 0xe0, 0x36, 0x8f, 0xbb, 0x7d, 0x99, 0x4f, 0xa2, 0x3d, 0x11,
	0xcd, 0xdb, 0x0d, 0xc8, 0xc7, 0x9d, 0x42, 0x23, 0xe9, 0xf2, 0xd9, 0x87, 0x7c, 0x67, 0xe0, 0x7d,
	0xd0, 0x1e, 0x4d, 0x9e, 0x41, 0xec, 0xb7, 0x8e, 0xb1, 0x1d, 0xd2, 0xe7, 0x10, 0xdc, 0x76, 0x8b,
	0x31, 0x91, 0x06, 0x71, 0xff, 0xb0, 0x29, 0x3c, 0x50, 0x4c, 0xc1, 0xe0, 0x8f, 0x2f, 0xd1, 0x11,
	0x40, 0x26, 0x16, 0xe9, 0x7a, 0x82, 0xc9, 0x5c, 0xb4, 0x4b, 0x9b, 0x2e, 0xa6, 0x92, 0x58, 0x3a,
	0x41, 0x2a, 0x0c, 0xb6, 0x5b, 0x95, 0x5c, 0x4f, 0xa6, 0x6f, 0xa5, 0x40, 0xe2, 0x50, 0x02, 0x43,
	0x7d, 0xbe, 0x49, 0x56, 0x29, 0x26, 0x52, 0x10, 0x57, 0x2b, 0xde, 0x08, 0xb9, 0xce, 0x64, 0x22,
	0x57, 0x19, 0xf1, 0x18, 0x87, 0x00, 0x55, 0xbd, 0x2b, 0xbf, 0xd4, 0x26, 0x47, 0x79, 0x6f, 0x02,
	0x07, 0x68, 0x97, 0xf7, 0xf4, 0x7f, 0xf0, 0x54, 0x55, 0x95, 0x95, 0x89, 0xd6, 0xc7, 0xf6, 0xc0,
	0x5e, 0x81, 0x9f, 0x35, 0x79, 0xf3, 0x50, 0xff, 0xd5, 0xff, 0xf6, 0xc3, 0x02, 0x98, 0x54, 0xe5,
	0xb6, 0x6b, 0xf5, 0xe5, 0x51, 0xab, 0xa7, 0xd1, 0x6f, 0xea, 0x5f, 0xf4, 0x5a, 0x75, 0x49, 0x0e,
	0xbd, 0xee, 0xa3, 0xe1, 0x81, 0xa2, 0xcf, 0xc0, 0xaf, 0x8d, 0xfd, 0xd0, 0x33, 0xa2, 0x5e, 0xd4,
	0xa6, 0xc1, 0x6e, 0xcc, 0x5e, 0x1f, 0x17, 0xff, 0x1f, 0xf4, 0x75, 0x61, 0xfb, 0xde, 0x01, 0xfc,
	0xae, 0x3b, 0x8b, 0x0e, 0x21, 0x40, 0x91, 0xbd, 0x5b, 0x2e, 0x32, 0x41, 0xec, 0x5b, 0xdf, 0xdc,
	0xc8, 0xab, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xe5, 0x10, 0x82, 0xa3, 0x03, 0x00, 0x00,
}
