// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kendaraan.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	kendaraan.proto

It has these top-level messages:
	AddKendaraanReq
	ReadKendaraanByNomorPlatReq
	ReadKendaraanByNomorPlatResp
	ReadKendaraanResp
	UpdateKendaraanReq
	ReadKendaraanByNoPlatReq
	ReadKendaraanByNoPlatResp
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

type AddKendaraanReq struct {
	KodeKendaraan  string `protobuf:"bytes,1,opt,name=KodeKendaraan" json:"KodeKendaraan,omitempty"`
	JenisKendaraan string `protobuf:"bytes,2,opt,name=JenisKendaraan" json:"JenisKendaraan,omitempty"`
	NomorPlat      string `protobuf:"bytes,3,opt,name=NomorPlat" json:"NomorPlat,omitempty"`
	Status         int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreateBy       string `protobuf:"bytes,5,opt,name=CreateBy" json:"CreateBy,omitempty"`
}

func (m *AddKendaraanReq) Reset()                    { *m = AddKendaraanReq{} }
func (m *AddKendaraanReq) String() string            { return proto.CompactTextString(m) }
func (*AddKendaraanReq) ProtoMessage()               {}
func (*AddKendaraanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddKendaraanReq) GetKodeKendaraan() string {
	if m != nil {
		return m.KodeKendaraan
	}
	return ""
}

func (m *AddKendaraanReq) GetJenisKendaraan() string {
	if m != nil {
		return m.JenisKendaraan
	}
	return ""
}

func (m *AddKendaraanReq) GetNomorPlat() string {
	if m != nil {
		return m.NomorPlat
	}
	return ""
}

func (m *AddKendaraanReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddKendaraanReq) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

type ReadKendaraanByNomorPlatReq struct {
	NomorPlat string `protobuf:"bytes,1,opt,name=NomorPlat" json:"NomorPlat,omitempty"`
}

func (m *ReadKendaraanByNomorPlatReq) Reset()                    { *m = ReadKendaraanByNomorPlatReq{} }
func (m *ReadKendaraanByNomorPlatReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKendaraanByNomorPlatReq) ProtoMessage()               {}
func (*ReadKendaraanByNomorPlatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadKendaraanByNomorPlatReq) GetNomorPlat() string {
	if m != nil {
		return m.NomorPlat
	}
	return ""
}

type ReadKendaraanByNomorPlatResp struct {
	KodeKendaraan  string `protobuf:"bytes,1,opt,name=KodeKendaraan" json:"KodeKendaraan,omitempty"`
	JenisKendaraan string `protobuf:"bytes,2,opt,name=JenisKendaraan" json:"JenisKendaraan,omitempty"`
	NomorPlat      string `protobuf:"bytes,3,opt,name=NomorPlat" json:"NomorPlat,omitempty"`
	Status         int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreateBy       string `protobuf:"bytes,5,opt,name=CreateBy" json:"CreateBy,omitempty"`
}

func (m *ReadKendaraanByNomorPlatResp) Reset()                    { *m = ReadKendaraanByNomorPlatResp{} }
func (m *ReadKendaraanByNomorPlatResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKendaraanByNomorPlatResp) ProtoMessage()               {}
func (*ReadKendaraanByNomorPlatResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadKendaraanByNomorPlatResp) GetKodeKendaraan() string {
	if m != nil {
		return m.KodeKendaraan
	}
	return ""
}

func (m *ReadKendaraanByNomorPlatResp) GetJenisKendaraan() string {
	if m != nil {
		return m.JenisKendaraan
	}
	return ""
}

func (m *ReadKendaraanByNomorPlatResp) GetNomorPlat() string {
	if m != nil {
		return m.NomorPlat
	}
	return ""
}

func (m *ReadKendaraanByNomorPlatResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadKendaraanByNomorPlatResp) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

type ReadKendaraanResp struct {
	AllKendaraan []*ReadKendaraanByNomorPlatResp `protobuf:"bytes,1,rep,name=allKendaraan" json:"allKendaraan,omitempty"`
}

func (m *ReadKendaraanResp) Reset()                    { *m = ReadKendaraanResp{} }
func (m *ReadKendaraanResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKendaraanResp) ProtoMessage()               {}
func (*ReadKendaraanResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadKendaraanResp) GetAllKendaraan() []*ReadKendaraanByNomorPlatResp {
	if m != nil {
		return m.AllKendaraan
	}
	return nil
}

type UpdateKendaraanReq struct {
	KodeKendaraan  string `protobuf:"bytes,1,opt,name=KodeKendaraan" json:"KodeKendaraan,omitempty"`
	JenisKendaraan string `protobuf:"bytes,2,opt,name=JenisKendaraan" json:"JenisKendaraan,omitempty"`
	NomorPlat      string `protobuf:"bytes,3,opt,name=NomorPlat" json:"NomorPlat,omitempty"`
	Status         int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	UpdateBy       string `protobuf:"bytes,5,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
}

func (m *UpdateKendaraanReq) Reset()                    { *m = UpdateKendaraanReq{} }
func (m *UpdateKendaraanReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateKendaraanReq) ProtoMessage()               {}
func (*UpdateKendaraanReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateKendaraanReq) GetKodeKendaraan() string {
	if m != nil {
		return m.KodeKendaraan
	}
	return ""
}

func (m *UpdateKendaraanReq) GetJenisKendaraan() string {
	if m != nil {
		return m.JenisKendaraan
	}
	return ""
}

func (m *UpdateKendaraanReq) GetNomorPlat() string {
	if m != nil {
		return m.NomorPlat
	}
	return ""
}

func (m *UpdateKendaraanReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateKendaraanReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

type ReadKendaraanByNoPlatReq struct {
	NomorPlat string `protobuf:"bytes,1,opt,name=NomorPlat" json:"NomorPlat,omitempty"`
}

func (m *ReadKendaraanByNoPlatReq) Reset()                    { *m = ReadKendaraanByNoPlatReq{} }
func (m *ReadKendaraanByNoPlatReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKendaraanByNoPlatReq) ProtoMessage()               {}
func (*ReadKendaraanByNoPlatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadKendaraanByNoPlatReq) GetNomorPlat() string {
	if m != nil {
		return m.NomorPlat
	}
	return ""
}

type ReadKendaraanByNoPlatResp struct {
	AllNoPlat []*ReadKendaraanByNomorPlatResp `protobuf:"bytes,1,rep,name=allNoPlat" json:"allNoPlat,omitempty"`
}

func (m *ReadKendaraanByNoPlatResp) Reset()                    { *m = ReadKendaraanByNoPlatResp{} }
func (m *ReadKendaraanByNoPlatResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKendaraanByNoPlatResp) ProtoMessage()               {}
func (*ReadKendaraanByNoPlatResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadKendaraanByNoPlatResp) GetAllNoPlat() []*ReadKendaraanByNomorPlatResp {
	if m != nil {
		return m.AllNoPlat
	}
	return nil
}

func init() {
	proto.RegisterType((*AddKendaraanReq)(nil), "grpc.AddKendaraanReq")
	proto.RegisterType((*ReadKendaraanByNomorPlatReq)(nil), "grpc.ReadKendaraanByNomorPlatReq")
	proto.RegisterType((*ReadKendaraanByNomorPlatResp)(nil), "grpc.ReadKendaraanByNomorPlatResp")
	proto.RegisterType((*ReadKendaraanResp)(nil), "grpc.ReadKendaraanResp")
	proto.RegisterType((*UpdateKendaraanReq)(nil), "grpc.UpdateKendaraanReq")
	proto.RegisterType((*ReadKendaraanByNoPlatReq)(nil), "grpc.ReadKendaraanByNoPlatReq")
	proto.RegisterType((*ReadKendaraanByNoPlatResp)(nil), "grpc.ReadKendaraanByNoPlatResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for KendaraanService service

type KendaraanServiceClient interface {
	AddKendaraan(ctx context.Context, in *AddKendaraanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadKendaraanByNomorPlat(ctx context.Context, in *ReadKendaraanByNomorPlatReq, opts ...grpc1.CallOption) (*ReadKendaraanByNomorPlatResp, error)
	ReadKendaraan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKendaraanResp, error)
	UpdateKendaraan(ctx context.Context, in *UpdateKendaraanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadKendaraanByNoPlat(ctx context.Context, in *ReadKendaraanByNoPlatReq, opts ...grpc1.CallOption) (*ReadKendaraanByNoPlatResp, error)
}

type kendaraanServiceClient struct {
	cc *grpc1.ClientConn
}

func NewKendaraanServiceClient(cc *grpc1.ClientConn) KendaraanServiceClient {
	return &kendaraanServiceClient{cc}
}

func (c *kendaraanServiceClient) AddKendaraan(ctx context.Context, in *AddKendaraanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KendaraanService/AddKendaraan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kendaraanServiceClient) ReadKendaraanByNomorPlat(ctx context.Context, in *ReadKendaraanByNomorPlatReq, opts ...grpc1.CallOption) (*ReadKendaraanByNomorPlatResp, error) {
	out := new(ReadKendaraanByNomorPlatResp)
	err := grpc1.Invoke(ctx, "/grpc.KendaraanService/ReadKendaraanByNomorPlat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kendaraanServiceClient) ReadKendaraan(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKendaraanResp, error) {
	out := new(ReadKendaraanResp)
	err := grpc1.Invoke(ctx, "/grpc.KendaraanService/ReadKendaraan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kendaraanServiceClient) UpdateKendaraan(ctx context.Context, in *UpdateKendaraanReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KendaraanService/UpdateKendaraan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kendaraanServiceClient) ReadKendaraanByNoPlat(ctx context.Context, in *ReadKendaraanByNoPlatReq, opts ...grpc1.CallOption) (*ReadKendaraanByNoPlatResp, error) {
	out := new(ReadKendaraanByNoPlatResp)
	err := grpc1.Invoke(ctx, "/grpc.KendaraanService/ReadKendaraanByNoPlat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KendaraanService service

type KendaraanServiceServer interface {
	AddKendaraan(context.Context, *AddKendaraanReq) (*google_protobuf.Empty, error)
	ReadKendaraanByNomorPlat(context.Context, *ReadKendaraanByNomorPlatReq) (*ReadKendaraanByNomorPlatResp, error)
	ReadKendaraan(context.Context, *google_protobuf.Empty) (*ReadKendaraanResp, error)
	UpdateKendaraan(context.Context, *UpdateKendaraanReq) (*google_protobuf.Empty, error)
	ReadKendaraanByNoPlat(context.Context, *ReadKendaraanByNoPlatReq) (*ReadKendaraanByNoPlatResp, error)
}

func RegisterKendaraanServiceServer(s *grpc1.Server, srv KendaraanServiceServer) {
	s.RegisterService(&_KendaraanService_serviceDesc, srv)
}

func _KendaraanService_AddKendaraan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKendaraanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KendaraanServiceServer).AddKendaraan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KendaraanService/AddKendaraan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KendaraanServiceServer).AddKendaraan(ctx, req.(*AddKendaraanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KendaraanService_ReadKendaraanByNomorPlat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKendaraanByNomorPlatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KendaraanServiceServer).ReadKendaraanByNomorPlat(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KendaraanService/ReadKendaraanByNomorPlat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KendaraanServiceServer).ReadKendaraanByNomorPlat(ctx, req.(*ReadKendaraanByNomorPlatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KendaraanService_ReadKendaraan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KendaraanServiceServer).ReadKendaraan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KendaraanService/ReadKendaraan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KendaraanServiceServer).ReadKendaraan(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KendaraanService_UpdateKendaraan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKendaraanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KendaraanServiceServer).UpdateKendaraan(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KendaraanService/UpdateKendaraan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KendaraanServiceServer).UpdateKendaraan(ctx, req.(*UpdateKendaraanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KendaraanService_ReadKendaraanByNoPlat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKendaraanByNoPlatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KendaraanServiceServer).ReadKendaraanByNoPlat(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KendaraanService/ReadKendaraanByNoPlat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KendaraanServiceServer).ReadKendaraanByNoPlat(ctx, req.(*ReadKendaraanByNoPlatReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KendaraanService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.KendaraanService",
	HandlerType: (*KendaraanServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddKendaraan",
			Handler:    _KendaraanService_AddKendaraan_Handler,
		},
		{
			MethodName: "ReadKendaraanByNomorPlat",
			Handler:    _KendaraanService_ReadKendaraanByNomorPlat_Handler,
		},
		{
			MethodName: "ReadKendaraan",
			Handler:    _KendaraanService_ReadKendaraan_Handler,
		},
		{
			MethodName: "UpdateKendaraan",
			Handler:    _KendaraanService_UpdateKendaraan_Handler,
		},
		{
			MethodName: "ReadKendaraanByNoPlat",
			Handler:    _KendaraanService_ReadKendaraanByNoPlat_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "kendaraan.proto",
}

func init() { proto.RegisterFile("kendaraan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0xf2, 0x11, 0x19, 0x41, 0x74, 0x12, 0xb0, 0x16, 0xa2, 0xb8, 0x31, 0x86, 0x53, 0x49,
	0xf0, 0x62, 0xe2, 0x41, 0xc5, 0xe0, 0x41, 0x12, 0x63, 0x4a, 0xbc, 0x68, 0x3c, 0x2c, 0x74, 0x25,
	0xc4, 0x42, 0x4b, 0xbb, 0x98, 0xf0, 0x6b, 0xfc, 0x0d, 0x26, 0x9e, 0xfd, 0x6d, 0xa6, 0x2d, 0x6d,
	0x69, 0xa1, 0x88, 0x37, 0x8e, 0xfb, 0xfa, 0xde, 0xcc, 0x7b, 0xd3, 0x19, 0x28, 0xbc, 0xb3, 0x91,
	0x4a, 0x4d, 0x4a, 0x47, 0xb2, 0x61, 0xea, 0x5c, 0xc7, 0x54, 0xdf, 0x34, 0x7a, 0x52, 0xb9, 0xaf,
	0xeb, 0x7d, 0x8d, 0xd5, 0x1d, 0xac, 0x3b, 0x79, 0xab, 0xb3, 0xa1, 0xc1, 0xa7, 0x2e, 0x85, 0x7c,
	0x09, 0x50, 0xb8, 0x51, 0xd5, 0xb6, 0xa7, 0x54, 0xd8, 0x18, 0x4f, 0x21, 0xdf, 0xd6, 0x55, 0xe6,
	0x63, 0xa2, 0x50, 0x15, 0x6a, 0x59, 0x25, 0x0c, 0xe2, 0x19, 0xec, 0xde, 0xb3, 0xd1, 0xc0, 0x0a,
	0x68, 0x5b, 0x0e, 0x2d, 0x82, 0x62, 0x05, 0xb2, 0x0f, 0xfa, 0x50, 0x37, 0x1f, 0x35, 0xca, 0xc5,
	0xa4, 0x43, 0x09, 0x00, 0x2c, 0x41, 0xa6, 0xc3, 0x29, 0x9f, 0x58, 0x62, 0xaa, 0x2a, 0xd4, 0xd2,
	0xca, 0xec, 0x85, 0x12, 0x6c, 0xdf, 0x9a, 0x8c, 0x72, 0xd6, 0x9c, 0x8a, 0x69, 0x47, 0xe4, 0xbf,
	0xc9, 0x25, 0x94, 0x15, 0x46, 0x03, 0xcf, 0xcd, 0xa9, 0x5f, 0xcf, 0xb6, 0x1f, 0x6a, 0x28, 0x44,
	0x1a, 0x92, 0x1f, 0x01, 0x2a, 0xf1, 0x6a, 0xcb, 0xd8, 0xf8, 0xf4, 0x2f, 0xb0, 0x1f, 0xf2, 0xef,
	0x98, 0xbe, 0x83, 0x1c, 0xd5, 0xb4, 0x79, 0xcf, 0xc9, 0xda, 0x4e, 0x83, 0xc8, 0xf6, 0x02, 0xc8,
	0xab, 0xe2, 0x2a, 0x21, 0x1d, 0xf9, 0x16, 0x00, 0x9f, 0x0c, 0x95, 0x72, 0xb6, 0x89, 0x1b, 0xe1,
	0x3a, 0x0b, 0x66, 0xe2, 0xbd, 0xc9, 0x05, 0x88, 0x0b, 0x21, 0xd7, 0x5b, 0x87, 0x57, 0x38, 0x8c,
	0x51, 0x5a, 0x06, 0x5e, 0x43, 0x96, 0x6a, 0x9a, 0x0b, 0xfc, 0x63, 0xa4, 0x81, 0xa8, 0xf1, 0x99,
	0x84, 0x3d, 0x9f, 0xd7, 0x61, 0xe6, 0xc7, 0xa0, 0xc7, 0xf0, 0x0a, 0x72, 0xf3, 0x27, 0x87, 0x45,
	0xb7, 0x66, 0xe4, 0x0c, 0xa5, 0x92, 0xec, 0x1e, 0xae, 0xec, 0x1d, 0xae, 0xdc, 0xb2, 0x0f, 0x97,
	0x24, 0xb0, 0xb7, 0x24, 0xae, 0x37, 0xbe, 0x93, 0xbf, 0x0c, 0x8e, 0xa5, 0x35, 0x32, 0x90, 0x04,
	0x36, 0x21, 0x1f, 0x62, 0x60, 0x8c, 0x1f, 0xe9, 0x60, 0x49, 0xb9, 0x59, 0x8d, 0x16, 0x14, 0x22,
	0xdb, 0x84, 0xa2, 0xcb, 0x5e, 0x5c, 0xb2, 0x15, 0x79, 0x9f, 0xa1, 0xb8, 0xf4, 0x27, 0xe1, 0x51,
	0x4c, 0x12, 0x2f, 0xe9, 0xf1, 0xca, 0xef, 0xb6, 0xc5, 0x6e, 0xc6, 0xe9, 0x76, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x59, 0xed, 0xf0, 0xd1, 0x3d, 0x05, 0x00, 0x00,
}
