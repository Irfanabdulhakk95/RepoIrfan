// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supplier.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	supplier.proto

It has these top-level messages:
	AddSupplierReq
	ReadSupplierByKeteranganReq
	ReadSupplierByKeteranganResp
	ReadSupplierResp
	UpdateSupplierReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddSupplierReq struct {
	KodeSupplier string `protobuf:"bytes,1,opt,name=KodeSupplier" json:"KodeSupplier,omitempty"`
	NamaSupplier string `protobuf:"bytes,2,opt,name=NamaSupplier" json:"NamaSupplier,omitempty"`
	Status       int32  `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	CreateBy     string `protobuf:"bytes,4,opt,name=CreateBy" json:"CreateBy,omitempty"`
	Keterangan   string `protobuf:"bytes,5,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *AddSupplierReq) Reset()                    { *m = AddSupplierReq{} }
func (m *AddSupplierReq) String() string            { return proto.CompactTextString(m) }
func (*AddSupplierReq) ProtoMessage()               {}
func (*AddSupplierReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddSupplierReq) GetKodeSupplier() string {
	if m != nil {
		return m.KodeSupplier
	}
	return ""
}

func (m *AddSupplierReq) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

func (m *AddSupplierReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddSupplierReq) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

func (m *AddSupplierReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadSupplierByKeteranganReq struct {
	Keterangan string `protobuf:"bytes,1,opt,name=Keterangan" json:"Keterangan,omitempty"`
}

func (m *ReadSupplierByKeteranganReq) Reset()                    { *m = ReadSupplierByKeteranganReq{} }
func (m *ReadSupplierByKeteranganReq) String() string            { return proto.CompactTextString(m) }
func (*ReadSupplierByKeteranganReq) ProtoMessage()               {}
func (*ReadSupplierByKeteranganReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadSupplierByKeteranganReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

type ReadSupplierByKeteranganResp struct {
	AllKeterangan []*AddSupplierReq `protobuf:"bytes,1,rep,name=allKeterangan" json:"allKeterangan,omitempty"`
}

func (m *ReadSupplierByKeteranganResp) Reset()                    { *m = ReadSupplierByKeteranganResp{} }
func (m *ReadSupplierByKeteranganResp) String() string            { return proto.CompactTextString(m) }
func (*ReadSupplierByKeteranganResp) ProtoMessage()               {}
func (*ReadSupplierByKeteranganResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadSupplierByKeteranganResp) GetAllKeterangan() []*AddSupplierReq {
	if m != nil {
		return m.AllKeterangan
	}
	return nil
}

type ReadSupplierResp struct {
	AllSupplier []*AddSupplierReq `protobuf:"bytes,1,rep,name=allSupplier" json:"allSupplier,omitempty"`
}

func (m *ReadSupplierResp) Reset()                    { *m = ReadSupplierResp{} }
func (m *ReadSupplierResp) String() string            { return proto.CompactTextString(m) }
func (*ReadSupplierResp) ProtoMessage()               {}
func (*ReadSupplierResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadSupplierResp) GetAllSupplier() []*AddSupplierReq {
	if m != nil {
		return m.AllSupplier
	}
	return nil
}

type UpdateSupplierReq struct {
	KodeSupplier string `protobuf:"bytes,1,opt,name=KodeSupplier" json:"KodeSupplier,omitempty"`
	NamaSupplier string `protobuf:"bytes,2,opt,name=NamaSupplier" json:"NamaSupplier,omitempty"`
	Keterangan   string `protobuf:"bytes,3,opt,name=Keterangan" json:"Keterangan,omitempty"`
	Status       int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	UpdateBy     string `protobuf:"bytes,5,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
}

func (m *UpdateSupplierReq) Reset()                    { *m = UpdateSupplierReq{} }
func (m *UpdateSupplierReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateSupplierReq) ProtoMessage()               {}
func (*UpdateSupplierReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateSupplierReq) GetKodeSupplier() string {
	if m != nil {
		return m.KodeSupplier
	}
	return ""
}

func (m *UpdateSupplierReq) GetNamaSupplier() string {
	if m != nil {
		return m.NamaSupplier
	}
	return ""
}

func (m *UpdateSupplierReq) GetKeterangan() string {
	if m != nil {
		return m.Keterangan
	}
	return ""
}

func (m *UpdateSupplierReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateSupplierReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func init() {
	proto.RegisterType((*AddSupplierReq)(nil), "grpc.AddSupplierReq")
	proto.RegisterType((*ReadSupplierByKeteranganReq)(nil), "grpc.ReadSupplierByKeteranganReq")
	proto.RegisterType((*ReadSupplierByKeteranganResp)(nil), "grpc.ReadSupplierByKeteranganResp")
	proto.RegisterType((*ReadSupplierResp)(nil), "grpc.ReadSupplierResp")
	proto.RegisterType((*UpdateSupplierReq)(nil), "grpc.UpdateSupplierReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for SupplierService service

type SupplierServiceClient interface {
	AddSupplier(ctx context.Context, in *AddSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadSupplierByKeterangan(ctx context.Context, in *ReadSupplierByKeteranganReq, opts ...grpc1.CallOption) (*ReadSupplierByKeteranganResp, error)
	ReadSupplier(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadSupplierResp, error)
	UpdateSupplier(ctx context.Context, in *UpdateSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type supplierServiceClient struct {
	cc *grpc1.ClientConn
}

func NewSupplierServiceClient(cc *grpc1.ClientConn) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) AddSupplier(ctx context.Context, in *AddSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/AddSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) ReadSupplierByKeterangan(ctx context.Context, in *ReadSupplierByKeteranganReq, opts ...grpc1.CallOption) (*ReadSupplierByKeteranganResp, error) {
	out := new(ReadSupplierByKeteranganResp)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/ReadSupplierByKeterangan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) ReadSupplier(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadSupplierResp, error) {
	out := new(ReadSupplierResp)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/ReadSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) UpdateSupplier(ctx context.Context, in *UpdateSupplierReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.SupplierService/UpdateSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SupplierService service

type SupplierServiceServer interface {
	AddSupplier(context.Context, *AddSupplierReq) (*google_protobuf.Empty, error)
	ReadSupplierByKeterangan(context.Context, *ReadSupplierByKeteranganReq) (*ReadSupplierByKeteranganResp, error)
	ReadSupplier(context.Context, *google_protobuf.Empty) (*ReadSupplierResp, error)
	UpdateSupplier(context.Context, *UpdateSupplierReq) (*google_protobuf.Empty, error)
}

func RegisterSupplierServiceServer(s *grpc1.Server, srv SupplierServiceServer) {
	s.RegisterService(&_SupplierService_serviceDesc, srv)
}

func _SupplierService_AddSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSupplierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).AddSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/AddSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).AddSupplier(ctx, req.(*AddSupplierReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_ReadSupplierByKeterangan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSupplierByKeteranganReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).ReadSupplierByKeterangan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/ReadSupplierByKeterangan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).ReadSupplierByKeterangan(ctx, req.(*ReadSupplierByKeteranganReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_ReadSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).ReadSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/ReadSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).ReadSupplier(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_UpdateSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupplierReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).UpdateSupplier(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SupplierService/UpdateSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).UpdateSupplier(ctx, req.(*UpdateSupplierReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SupplierService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddSupplier",
			Handler:    _SupplierService_AddSupplier_Handler,
		},
		{
			MethodName: "ReadSupplierByKeterangan",
			Handler:    _SupplierService_ReadSupplierByKeterangan_Handler,
		},
		{
			MethodName: "ReadSupplier",
			Handler:    _SupplierService_ReadSupplier_Handler,
		},
		{
			MethodName: "UpdateSupplier",
			Handler:    _SupplierService_UpdateSupplier_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "supplier.proto",
}

func init() { proto.RegisterFile("supplier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x6d, 0xfa, 0xc5, 0x7b, 0xb7, 0x7d, 0x7d, 0xef, 0x5d, 0xa4, 0x86, 0x54, 0xa4, 0xce, 0xaa,
	0xab, 0x14, 0x2a, 0xb8, 0x10, 0x0a, 0xda, 0xe2, 0xc6, 0x82, 0x8b, 0x14, 0x37, 0xee, 0xa6, 0xcd,
	0x35, 0x14, 0xd2, 0x66, 0x4c, 0xa6, 0x42, 0xfe, 0x93, 0xe0, 0x1f, 0xf0, 0xc7, 0x49, 0x3e, 0x1a,
	0x67, 0x5a, 0xa3, 0x2b, 0x97, 0xf7, 0xcc, 0xb9, 0x67, 0xee, 0x3d, 0xe7, 0x42, 0x27, 0xda, 0x0a,
	0xe1, 0xaf, 0x28, 0xb4, 0x45, 0x18, 0xc8, 0x00, 0xeb, 0x5e, 0x28, 0x96, 0x56, 0xcf, 0x0b, 0x02,
	0xcf, 0xa7, 0x61, 0x8a, 0x2d, 0xb6, 0x8f, 0x43, 0x5a, 0x0b, 0x19, 0x67, 0x14, 0xf6, 0x62, 0x40,
	0xe7, 0xda, 0x75, 0xe7, 0x79, 0xa3, 0x43, 0x4f, 0xc8, 0xa0, 0x3d, 0x0b, 0x5c, 0xda, 0x41, 0xa6,
	0xd1, 0x37, 0x06, 0xbf, 0x1d, 0x0d, 0x4b, 0x38, 0x77, 0x7c, 0xcd, 0x0b, 0x4e, 0x35, 0xe3, 0xa8,
	0x18, 0x76, 0xa1, 0x39, 0x97, 0x5c, 0x6e, 0x23, 0xb3, 0xd6, 0x37, 0x06, 0x0d, 0x27, 0xaf, 0xd0,
	0x82, 0x5f, 0xd3, 0x90, 0xb8, 0xa4, 0x49, 0x6c, 0xd6, 0xd3, 0xbe, 0xa2, 0xc6, 0x53, 0x80, 0x19,
	0x49, 0x0a, 0xf9, 0xc6, 0xe3, 0x1b, 0xb3, 0x91, 0xbe, 0x2a, 0x08, 0x1b, 0x43, 0xcf, 0x21, 0x5e,
	0x8c, 0x3b, 0x89, 0x3f, 0xde, 0x92, 0xd1, 0xf5, 0x76, 0xe3, 0xa0, 0xfd, 0x01, 0x4e, 0xca, 0xdb,
	0x23, 0x81, 0x97, 0xf0, 0x87, 0xfb, 0xbe, 0x26, 0x51, 0x1b, 0xb4, 0x46, 0x47, 0x76, 0x62, 0xa4,
	0xad, 0xfb, 0xe4, 0xe8, 0x54, 0x76, 0x0b, 0xff, 0x54, 0xed, 0x54, 0xef, 0x02, 0x5a, 0xdc, 0xf7,
	0x15, 0x27, 0xcb, 0xd5, 0x54, 0x22, 0x7b, 0x35, 0xe0, 0xff, 0xbd, 0x70, 0xb9, 0xa4, 0x9f, 0x08,
	0x46, 0x77, 0xa9, 0xb6, 0xef, 0x92, 0x12, 0x5c, 0x7d, 0x3f, 0xb8, 0x6c, 0xa8, 0x49, 0x9c, 0x47,
	0x53, 0xd4, 0xa3, 0xb7, 0x2a, 0xfc, 0xdd, 0x7d, 0x30, 0xa7, 0xf0, 0x79, 0xb5, 0x24, 0x1c, 0x43,
	0x4b, 0x59, 0x12, 0x3f, 0xdd, 0xdb, 0xea, 0xda, 0xd9, 0x79, 0xda, 0xbb, 0xf3, 0xb4, 0x6f, 0x92,
	0xf3, 0x64, 0x15, 0x5c, 0x82, 0x59, 0x16, 0x16, 0x9e, 0x65, 0x5a, 0x5f, 0xdc, 0x82, 0xc5, 0xbe,
	0xa3, 0x44, 0x82, 0x55, 0xf0, 0x0a, 0xda, 0x2a, 0x03, 0x4b, 0xc6, 0x49, 0xc6, 0xdc, 0x57, 0xcb,
	0x15, 0xa6, 0xd0, 0xd1, 0xa3, 0xc2, 0xe3, 0x8c, 0x7b, 0x10, 0x60, 0xf9, 0xae, 0x8b, 0x66, 0x8a,
	0x9c, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xaf, 0x28, 0x70, 0xc2, 0x03, 0x00, 0x00,
}
