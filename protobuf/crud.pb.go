// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crud.proto

package crud

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type HostnameResponse struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameResponse) Reset()         { *m = HostnameResponse{} }
func (m *HostnameResponse) String() string { return proto.CompactTextString(m) }
func (*HostnameResponse) ProtoMessage()    {}
func (*HostnameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{0}
}
func (m *HostnameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameResponse.Unmarshal(m, b)
}
func (m *HostnameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameResponse.Marshal(b, m, deterministic)
}
func (dst *HostnameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameResponse.Merge(dst, src)
}
func (m *HostnameResponse) XXX_Size() int {
	return xxx_messageInfo_HostnameResponse.Size(m)
}
func (m *HostnameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameResponse proto.InternalMessageInfo

func (m *HostnameResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type Request struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReadResponse struct {
	Data                 []*Request `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_478bbe1b22b2e995, []int{2}
}
func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (dst *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(dst, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetData() []*Request {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HostnameResponse)(nil), "crud.HostnameResponse")
	proto.RegisterType((*Request)(nil), "crud.Request")
	proto.RegisterType((*ReadResponse)(nil), "crud.ReadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CRUDServiceClient is the client API for CRUDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CRUDServiceClient interface {
	Read(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadResponse, error)
	Hostname(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HostnameResponse, error)
}

type cRUDServiceClient struct {
	cc *grpc.ClientConn
}

func NewCRUDServiceClient(cc *grpc.ClientConn) CRUDServiceClient {
	return &cRUDServiceClient{cc}
}

func (c *cRUDServiceClient) Read(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/crud.CRUDService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) Hostname(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HostnameResponse, error) {
	out := new(HostnameResponse)
	err := c.cc.Invoke(ctx, "/crud.CRUDService/Hostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServiceServer is the server API for CRUDService service.
type CRUDServiceServer interface {
	Read(context.Context, *empty.Empty) (*ReadResponse, error)
	Hostname(context.Context, *empty.Empty) (*HostnameResponse, error)
}

func RegisterCRUDServiceServer(s *grpc.Server, srv CRUDServiceServer) {
	s.RegisterService(&_CRUDService_serviceDesc, srv)
}

func _CRUDService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.CRUDService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).Read(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_Hostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).Hostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.CRUDService/Hostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).Hostname(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CRUDService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crud.CRUDService",
	HandlerType: (*CRUDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _CRUDService_Read_Handler,
		},
		{
			MethodName: "Hostname",
			Handler:    _CRUDService_Hostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}

func init() { proto.RegisterFile("crud.proto", fileDescriptor_478bbe1b22b2e995) }

var fileDescriptor_478bbe1b22b2e995 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2e, 0x2a, 0x4d,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0xc1, 0x62, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x25,
	0x4a, 0x7a, 0x5c, 0x02, 0x1e, 0xf9, 0xc5, 0x25, 0x79, 0x89, 0xb9, 0xa9, 0x41, 0xa9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x19, 0x50, 0x31, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x38, 0x5f, 0x49, 0x99, 0x8b, 0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48,
	0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x0a, 0xc6, 0x55, 0x32, 0xe4, 0xe2,
	0x09, 0x4a, 0x4d, 0x4c, 0x81, 0x1b, 0xa8, 0xc8, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xab, 0x07, 0x76, 0x22, 0xd4, 0x98, 0x20, 0xb0, 0x94, 0x51, 0x23,
	0x23, 0x17, 0xb7, 0x73, 0x50, 0xa8, 0x4b, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x09,
	0x17, 0x0b, 0xc8, 0x08, 0x21, 0x31, 0x3d, 0x88, 0xeb, 0xf5, 0x60, 0xae, 0xd7, 0x73, 0x05, 0xb9,
	0x5e, 0x4a, 0x08, 0x66, 0x08, 0xc2, 0x1a, 0x25, 0x06, 0x21, 0x1b, 0x2e, 0x0e, 0x98, 0x6f, 0x70,
	0xea, 0x14, 0x83, 0xe8, 0x44, 0xf7, 0xb5, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xa5, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x05, 0x05, 0x6d, 0xb5, 0x43, 0x01, 0x00, 0x00,
}
