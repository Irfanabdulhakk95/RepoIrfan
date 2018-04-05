// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventori.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	inventori.proto

It has these top-level messages:
	AddInventoriReq
	ReadInventoriByNamaInventoriReq
	ReadInventoriByNamaInventoriResp
	ReadInventoriResp
	UpdateInventoriReq
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

type AddInventoriReq struct {
	KodeInventori string `protobuf:"bytes,1,opt,name=KodeInventori" json:"KodeInventori,omitempty"`
	NamaInventori string `protobuf:"bytes,2,opt,name=NamaInventori" json:"NamaInventori,omitempty"`
	Jumlah        int32  `protobuf:"varint,3,opt,name=Jumlah" json:"Jumlah,omitempty"`
	Status        int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreateBy      string `protobuf:"bytes,5,opt,name=CreateBy" json:"CreateBy,omitempty"`
}

func (m *AddInventoriReq) Reset()                    { *m = AddInventoriReq{} }
func (m *AddInventoriReq) String() string            { return proto.CompactTextString(m) }
func (*AddInventoriReq) ProtoMessage()               {}
func (*AddInventoriReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddInventoriReq) GetKodeInventori() string {
	if m != nil {
		return m.KodeInventori
	}
	return ""
}

func (m *AddInventoriReq) GetNamaInventori() string {
	if m != nil {
		return m.NamaInventori
	}
	return ""
}

func (m *AddInventoriReq) GetJumlah() int32 {
	if m != nil {
		return m.Jumlah
	}
	return 0
}

func (m *AddInventoriReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AddInventoriReq) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

type ReadInventoriByNamaInventoriReq struct {
	NamaInventori string `protobuf:"bytes,1,opt,name=NamaInventori" json:"NamaInventori,omitempty"`
}

func (m *ReadInventoriByNamaInventoriReq) Reset()                    { *m = ReadInventoriByNamaInventoriReq{} }
func (m *ReadInventoriByNamaInventoriReq) String() string            { return proto.CompactTextString(m) }
func (*ReadInventoriByNamaInventoriReq) ProtoMessage()               {}
func (*ReadInventoriByNamaInventoriReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadInventoriByNamaInventoriReq) GetNamaInventori() string {
	if m != nil {
		return m.NamaInventori
	}
	return ""
}

type ReadInventoriByNamaInventoriResp struct {
	KodeInventori string `protobuf:"bytes,1,opt,name=KodeInventori" json:"KodeInventori,omitempty"`
	NamaInventori string `protobuf:"bytes,2,opt,name=NamaInventori" json:"NamaInventori,omitempty"`
	Jumlah        int32  `protobuf:"varint,3,opt,name=Jumlah" json:"Jumlah,omitempty"`
	Status        int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	CreateBy      string `protobuf:"bytes,5,opt,name=CreateBy" json:"CreateBy,omitempty"`
}

func (m *ReadInventoriByNamaInventoriResp) Reset()         { *m = ReadInventoriByNamaInventoriResp{} }
func (m *ReadInventoriByNamaInventoriResp) String() string { return proto.CompactTextString(m) }
func (*ReadInventoriByNamaInventoriResp) ProtoMessage()    {}
func (*ReadInventoriByNamaInventoriResp) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

func (m *ReadInventoriByNamaInventoriResp) GetKodeInventori() string {
	if m != nil {
		return m.KodeInventori
	}
	return ""
}

func (m *ReadInventoriByNamaInventoriResp) GetNamaInventori() string {
	if m != nil {
		return m.NamaInventori
	}
	return ""
}

func (m *ReadInventoriByNamaInventoriResp) GetJumlah() int32 {
	if m != nil {
		return m.Jumlah
	}
	return 0
}

func (m *ReadInventoriByNamaInventoriResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReadInventoriByNamaInventoriResp) GetCreateBy() string {
	if m != nil {
		return m.CreateBy
	}
	return ""
}

type ReadInventoriResp struct {
	AllInventori []*ReadInventoriByNamaInventoriResp `protobuf:"bytes,1,rep,name=allInventori" json:"allInventori,omitempty"`
}

func (m *ReadInventoriResp) Reset()                    { *m = ReadInventoriResp{} }
func (m *ReadInventoriResp) String() string            { return proto.CompactTextString(m) }
func (*ReadInventoriResp) ProtoMessage()               {}
func (*ReadInventoriResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadInventoriResp) GetAllInventori() []*ReadInventoriByNamaInventoriResp {
	if m != nil {
		return m.AllInventori
	}
	return nil
}

type UpdateInventoriReq struct {
	KodeInventori string `protobuf:"bytes,1,opt,name=KodeInventori" json:"KodeInventori,omitempty"`
	NamaInventori string `protobuf:"bytes,2,opt,name=NamaInventori" json:"NamaInventori,omitempty"`
	Jumlah        int32  `protobuf:"varint,3,opt,name=Jumlah" json:"Jumlah,omitempty"`
	Status        int32  `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	UpdateBy      string `protobuf:"bytes,5,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
}

func (m *UpdateInventoriReq) Reset()                    { *m = UpdateInventoriReq{} }
func (m *UpdateInventoriReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateInventoriReq) ProtoMessage()               {}
func (*UpdateInventoriReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateInventoriReq) GetKodeInventori() string {
	if m != nil {
		return m.KodeInventori
	}
	return ""
}

func (m *UpdateInventoriReq) GetNamaInventori() string {
	if m != nil {
		return m.NamaInventori
	}
	return ""
}

func (m *UpdateInventoriReq) GetJumlah() int32 {
	if m != nil {
		return m.Jumlah
	}
	return 0
}

func (m *UpdateInventoriReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdateInventoriReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func init() {
	proto.RegisterType((*AddInventoriReq)(nil), "grpc.AddInventoriReq")
	proto.RegisterType((*ReadInventoriByNamaInventoriReq)(nil), "grpc.ReadInventoriByNamaInventoriReq")
	proto.RegisterType((*ReadInventoriByNamaInventoriResp)(nil), "grpc.ReadInventoriByNamaInventoriResp")
	proto.RegisterType((*ReadInventoriResp)(nil), "grpc.ReadInventoriResp")
	proto.RegisterType((*UpdateInventoriReq)(nil), "grpc.UpdateInventoriReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for InventoriService service

type InventoriServiceClient interface {
	AddInventori(ctx context.Context, in *AddInventoriReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadInventoriByNamaInventori(ctx context.Context, in *ReadInventoriByNamaInventoriReq, opts ...grpc1.CallOption) (*ReadInventoriByNamaInventoriResp, error)
	ReadInventori(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadInventoriResp, error)
	UpdateInventori(ctx context.Context, in *UpdateInventoriReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type inventoriServiceClient struct {
	cc *grpc1.ClientConn
}

func NewInventoriServiceClient(cc *grpc1.ClientConn) InventoriServiceClient {
	return &inventoriServiceClient{cc}
}

func (c *inventoriServiceClient) AddInventori(ctx context.Context, in *AddInventoriReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.InventoriService/AddInventori", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoriServiceClient) ReadInventoriByNamaInventori(ctx context.Context, in *ReadInventoriByNamaInventoriReq, opts ...grpc1.CallOption) (*ReadInventoriByNamaInventoriResp, error) {
	out := new(ReadInventoriByNamaInventoriResp)
	err := grpc1.Invoke(ctx, "/grpc.InventoriService/ReadInventoriByNamaInventori", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoriServiceClient) ReadInventori(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadInventoriResp, error) {
	out := new(ReadInventoriResp)
	err := grpc1.Invoke(ctx, "/grpc.InventoriService/ReadInventori", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoriServiceClient) UpdateInventori(ctx context.Context, in *UpdateInventoriReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.InventoriService/UpdateInventori", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InventoriService service

type InventoriServiceServer interface {
	AddInventori(context.Context, *AddInventoriReq) (*google_protobuf.Empty, error)
	ReadInventoriByNamaInventori(context.Context, *ReadInventoriByNamaInventoriReq) (*ReadInventoriByNamaInventoriResp, error)
	ReadInventori(context.Context, *google_protobuf.Empty) (*ReadInventoriResp, error)
	UpdateInventori(context.Context, *UpdateInventoriReq) (*google_protobuf.Empty, error)
}

func RegisterInventoriServiceServer(s *grpc1.Server, srv InventoriServiceServer) {
	s.RegisterService(&_InventoriService_serviceDesc, srv)
}

func _InventoriService_AddInventori_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInventoriReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoriServiceServer).AddInventori(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.InventoriService/AddInventori",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoriServiceServer).AddInventori(ctx, req.(*AddInventoriReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoriService_ReadInventoriByNamaInventori_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadInventoriByNamaInventoriReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoriServiceServer).ReadInventoriByNamaInventori(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.InventoriService/ReadInventoriByNamaInventori",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoriServiceServer).ReadInventoriByNamaInventori(ctx, req.(*ReadInventoriByNamaInventoriReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoriService_ReadInventori_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoriServiceServer).ReadInventori(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.InventoriService/ReadInventori",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoriServiceServer).ReadInventori(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoriService_UpdateInventori_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInventoriReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoriServiceServer).UpdateInventori(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.InventoriService/UpdateInventori",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoriServiceServer).UpdateInventori(ctx, req.(*UpdateInventoriReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _InventoriService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.InventoriService",
	HandlerType: (*InventoriServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddInventori",
			Handler:    _InventoriService_AddInventori_Handler,
		},
		{
			MethodName: "ReadInventoriByNamaInventori",
			Handler:    _InventoriService_ReadInventoriByNamaInventori_Handler,
		},
		{
			MethodName: "ReadInventori",
			Handler:    _InventoriService_ReadInventori_Handler,
		},
		{
			MethodName: "UpdateInventori",
			Handler:    _InventoriService_UpdateInventori_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "inventori.proto",
}

func init() { proto.RegisterFile("inventori.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xdd, 0x4a, 0x02, 0x41,
	0x14, 0x76, 0xfc, 0xa3, 0x4e, 0x8a, 0x75, 0x20, 0x5b, 0xb6, 0x20, 0x19, 0x2a, 0xbc, 0x1a, 0xc1,
	0x1e, 0x20, 0x32, 0x24, 0x32, 0xe8, 0x62, 0xa5, 0xeb, 0x18, 0xdd, 0xc9, 0x84, 0xd5, 0x59, 0xd7,
	0x51, 0xf0, 0x91, 0x82, 0x9e, 0xa0, 0x97, 0xe8, 0x95, 0x62, 0x77, 0x74, 0xd7, 0x31, 0x35, 0x2f,
	0xbd, 0x3c, 0x1f, 0xdf, 0x7c, 0xf3, 0x9d, 0xef, 0x9c, 0x03, 0xa5, 0xfe, 0x70, 0x2a, 0x86, 0x4a,
	0x06, 0x7d, 0xe6, 0x07, 0x52, 0x49, 0xcc, 0xf6, 0x02, 0xbf, 0x6b, 0x9f, 0xf7, 0xa4, 0xec, 0x79,
	0xa2, 0x16, 0x61, 0x9d, 0xc9, 0x7b, 0x4d, 0x0c, 0x7c, 0x35, 0xd3, 0x14, 0xfa, 0x49, 0xa0, 0x74,
	0xef, 0xba, 0x4f, 0x8b, 0x97, 0x8e, 0x18, 0xe1, 0x15, 0x14, 0x9f, 0xa5, 0x2b, 0x62, 0xcc, 0x22,
	0x15, 0x52, 0x3d, 0x74, 0x4c, 0x30, 0x64, 0xbd, 0xf0, 0x01, 0x4f, 0x58, 0x69, 0xcd, 0x32, 0x40,
	0x2c, 0x43, 0xbe, 0x35, 0x19, 0x78, 0xfc, 0xc3, 0xca, 0x54, 0x48, 0x35, 0xe7, 0xcc, 0xab, 0x10,
	0x6f, 0x2b, 0xae, 0x26, 0x63, 0x2b, 0xab, 0x71, 0x5d, 0xa1, 0x0d, 0x07, 0x0f, 0x81, 0xe0, 0x4a,
	0x34, 0x66, 0x56, 0x2e, 0x12, 0x8c, 0x6b, 0xfa, 0x08, 0x97, 0x8e, 0xe0, 0x89, 0xd7, 0xc6, 0xcc,
	0xf8, 0x6b, 0x6e, 0xdd, 0x34, 0x45, 0xd6, 0x98, 0xa2, 0xdf, 0x04, 0x2a, 0xdb, 0x95, 0xc6, 0xfe,
	0xde, 0xa6, 0xf0, 0x06, 0x27, 0x86, 0xf7, 0xc8, 0x6c, 0x0b, 0x0a, 0xdc, 0xf3, 0x96, 0xbd, 0x66,
	0xaa, 0x47, 0xf5, 0x1b, 0x16, 0x2e, 0x00, 0xfb, 0xaf, 0x55, 0xc7, 0x78, 0x4b, 0xbf, 0x08, 0xe0,
	0xab, 0xef, 0x72, 0x25, 0xf6, 0x69, 0x2b, 0xb4, 0xa3, 0x24, 0x8f, 0x45, 0x5d, 0xff, 0x49, 0xc3,
	0x71, 0xac, 0xdc, 0x16, 0xc1, 0xb4, 0xdf, 0x15, 0x78, 0x07, 0x85, 0xe5, 0xad, 0xc6, 0x53, 0x9d,
	0xc4, 0xca, 0xa6, 0xdb, 0x65, 0xa6, 0x6f, 0x83, 0x2d, 0x6e, 0x83, 0x35, 0xc3, 0xdb, 0xa0, 0x29,
	0x94, 0x70, 0xb1, 0x2d, 0x36, 0xbc, 0xde, 0x25, 0xda, 0x91, 0xbd, 0xe3, 0x04, 0x68, 0x0a, 0x1b,
	0x50, 0x34, 0x58, 0xb8, 0xc1, 0x9b, 0x7d, 0xb6, 0x46, 0x72, 0xae, 0xd1, 0x84, 0xd2, 0xca, 0xe0,
	0xd0, 0xd2, 0xec, 0xbf, 0xf3, 0xdc, 0xdc, 0x7b, 0x27, 0x1f, 0x21, 0xb7, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x38, 0x30, 0xab, 0xa2, 0x50, 0x04, 0x00, 0x00,
}