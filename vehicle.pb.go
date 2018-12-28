// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vehicle.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Vehicle_Type int32

const (
	Vehicle_NEW    Vehicle_Type = 0
	Vehicle_USED   Vehicle_Type = 1
	Vehicle_LOANER Vehicle_Type = 3
)

var Vehicle_Type_name = map[int32]string{
	0: "NEW",
	1: "USED",
	3: "LOANER",
}

var Vehicle_Type_value = map[string]int32{
	"NEW":    0,
	"USED":   1,
	"LOANER": 3,
}

func (x Vehicle_Type) String() string {
	return proto.EnumName(Vehicle_Type_name, int32(x))
}

func (Vehicle_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{1, 0}
}

type VehicleFinancial struct {
	Price                int32    `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Cost                 int32    `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehicleFinancial) Reset()         { *m = VehicleFinancial{} }
func (m *VehicleFinancial) String() string { return proto.CompactTextString(m) }
func (*VehicleFinancial) ProtoMessage()    {}
func (*VehicleFinancial) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{0}
}

func (m *VehicleFinancial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleFinancial.Unmarshal(m, b)
}
func (m *VehicleFinancial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleFinancial.Marshal(b, m, deterministic)
}
func (m *VehicleFinancial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleFinancial.Merge(m, src)
}
func (m *VehicleFinancial) XXX_Size() int {
	return xxx_messageInfo_VehicleFinancial.Size(m)
}
func (m *VehicleFinancial) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleFinancial.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleFinancial proto.InternalMessageInfo

func (m *VehicleFinancial) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *VehicleFinancial) GetCost() int32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

type Vehicle struct {
	Vin                  string            `protobuf:"bytes,1,opt,name=vin,proto3" json:"vin,omitempty"`
	Year                 int32             `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Make                 string            `protobuf:"bytes,3,opt,name=make,proto3" json:"make,omitempty"`
	Model                string            `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Color                string            `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Type                 Vehicle_Type      `protobuf:"varint,6,opt,name=type,proto3,enum=main.Vehicle_Type" json:"type,omitempty"`
	VehicleFinancial     *VehicleFinancial `protobuf:"bytes,7,opt,name=vehicleFinancial,proto3" json:"vehicleFinancial,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Vehicle) Reset()         { *m = Vehicle{} }
func (m *Vehicle) String() string { return proto.CompactTextString(m) }
func (*Vehicle) ProtoMessage()    {}
func (*Vehicle) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{1}
}

func (m *Vehicle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vehicle.Unmarshal(m, b)
}
func (m *Vehicle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vehicle.Marshal(b, m, deterministic)
}
func (m *Vehicle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vehicle.Merge(m, src)
}
func (m *Vehicle) XXX_Size() int {
	return xxx_messageInfo_Vehicle.Size(m)
}
func (m *Vehicle) XXX_DiscardUnknown() {
	xxx_messageInfo_Vehicle.DiscardUnknown(m)
}

var xxx_messageInfo_Vehicle proto.InternalMessageInfo

func (m *Vehicle) GetVin() string {
	if m != nil {
		return m.Vin
	}
	return ""
}

func (m *Vehicle) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Vehicle) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

func (m *Vehicle) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Vehicle) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Vehicle) GetType() Vehicle_Type {
	if m != nil {
		return m.Type
	}
	return Vehicle_NEW
}

func (m *Vehicle) GetVehicleFinancial() *VehicleFinancial {
	if m != nil {
		return m.VehicleFinancial
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.Vehicle_Type", Vehicle_Type_name, Vehicle_Type_value)
	proto.RegisterType((*VehicleFinancial)(nil), "main.VehicleFinancial")
	proto.RegisterType((*Vehicle)(nil), "main.Vehicle")
}

func init() { proto.RegisterFile("vehicle.proto", fileDescriptor_416ab71f8212867c) }

var fileDescriptor_416ab71f8212867c = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0x47, 0xcd, 0x36, 0x6d, 0x75, 0x44, 0x09, 0x83, 0x48, 0x8e, 0xa5, 0xa0, 0xf4, 0xd4, 0xc3,
	0x7a, 0xf5, 0xa2, 0x58, 0x4f, 0xb2, 0x42, 0xfc, 0x77, 0x8e, 0x31, 0x60, 0xb0, 0x6d, 0x4a, 0x2d,
	0x85, 0x5e, 0xfd, 0xe4, 0x92, 0xc9, 0x2a, 0xe8, 0xde, 0x7e, 0xf3, 0x66, 0x5e, 0x92, 0x09, 0x1c,
	0xcd, 0xf6, 0xdd, 0x99, 0xd6, 0xd6, 0xc3, 0xe8, 0x27, 0x8f, 0xbc, 0xd3, 0xae, 0x2f, 0x2f, 0x41,
	0x3c, 0x47, 0x7c, 0xeb, 0x7a, 0xdd, 0x1b, 0xa7, 0x5b, 0x3c, 0x81, 0x74, 0x18, 0x9d, 0xb1, 0x92,
	0x15, 0xac, 0x4a, 0x55, 0x2c, 0x10, 0x81, 0x1b, 0xff, 0x39, 0xc9, 0x15, 0x41, 0xca, 0xe5, 0xd7,
	0x0a, 0xf2, 0xad, 0x8e, 0x02, 0x92, 0xd9, 0xf5, 0xe4, 0x1c, 0xa8, 0x10, 0x83, 0xb1, 0x58, 0x3d,
	0xfe, 0x18, 0x21, 0x07, 0xd6, 0xe9, 0x0f, 0x2b, 0x13, 0x1a, 0xa3, 0x1c, 0xee, 0xeb, 0xfc, 0x9b,
	0x6d, 0x25, 0x27, 0x18, 0x8b, 0x40, 0x8d, 0x6f, 0xfd, 0x28, 0xd3, 0x48, 0xa9, 0xc0, 0x73, 0xe0,
	0xd3, 0x32, 0x58, 0x99, 0x15, 0xac, 0x3a, 0x5e, 0x63, 0x1d, 0x96, 0xa8, 0xb7, 0x4f, 0xa8, 0x1f,
	0x97, 0xc1, 0x2a, 0xea, 0xe3, 0x35, 0x88, 0xf9, 0xdf, 0x5e, 0x32, 0x2f, 0x58, 0x75, 0xb8, 0x3e,
	0xfd, 0xe3, 0xfc, 0x76, 0xd5, 0xce, 0x7c, 0x79, 0x06, 0x3c, 0x9c, 0x88, 0x39, 0x24, 0x9b, 0xe6,
	0x45, 0xec, 0xe1, 0x3e, 0xf0, 0xa7, 0x87, 0xe6, 0x46, 0x30, 0x04, 0xc8, 0xee, 0xee, 0xaf, 0x36,
	0x8d, 0x12, 0xc9, 0x6b, 0x46, 0xff, 0x79, 0xf1, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x52, 0x2b, 0xa9,
	0x75, 0x60, 0x01, 0x00, 0x00,
}