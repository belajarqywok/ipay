// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipay.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	ipay.proto

It has these top-level messages:
	CreateCustomerRequest
	CreateCustomerResponse
	ServiceError
	ServiceGenericReply
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CreateCustomerRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Phonenumber string `protobuf:"bytes,3,opt,name=phonenumber" json:"phonenumber,omitempty"`
}

func (m *CreateCustomerRequest) Reset()                    { *m = CreateCustomerRequest{} }
func (m *CreateCustomerRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCustomerRequest) ProtoMessage()               {}
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateCustomerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateCustomerRequest) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

type CreateCustomerResponse struct {
	Reply       *ServiceGenericReply `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
	CustomerId  string               `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Name        string               `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Email       string               `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phonenumber string               `protobuf:"bytes,5,opt,name=phonenumber" json:"phonenumber,omitempty"`
}

func (m *CreateCustomerResponse) Reset()                    { *m = CreateCustomerResponse{} }
func (m *CreateCustomerResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCustomerResponse) ProtoMessage()               {}
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateCustomerResponse) GetReply() *ServiceGenericReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *CreateCustomerResponse) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *CreateCustomerResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCustomerResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateCustomerResponse) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

type ServiceError struct {
	Code    string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ServiceError) Reset()                    { *m = ServiceError{} }
func (m *ServiceError) String() string            { return proto.CompactTextString(m) }
func (*ServiceError) ProtoMessage()               {}
func (*ServiceError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ServiceError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ServiceGenericReply struct {
	Success bool            `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Errors  []*ServiceError `protobuf:"bytes,2,rep,name=errors" json:"errors,omitempty"`
}

func (m *ServiceGenericReply) Reset()                    { *m = ServiceGenericReply{} }
func (m *ServiceGenericReply) String() string            { return proto.CompactTextString(m) }
func (*ServiceGenericReply) ProtoMessage()               {}
func (*ServiceGenericReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServiceGenericReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ServiceGenericReply) GetErrors() []*ServiceError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCustomerRequest)(nil), "CreateCustomerRequest")
	proto.RegisterType((*CreateCustomerResponse)(nil), "CreateCustomerResponse")
	proto.RegisterType((*ServiceError)(nil), "ServiceError")
	proto.RegisterType((*ServiceGenericReply)(nil), "ServiceGenericReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IpayService service

type IpayServiceClient interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error)
}

type ipayServiceClient struct {
	cc *grpc.ClientConn
}

func NewIpayServiceClient(cc *grpc.ClientConn) IpayServiceClient {
	return &ipayServiceClient{cc}
}

func (c *ipayServiceClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error) {
	out := new(CreateCustomerResponse)
	err := grpc.Invoke(ctx, "/IpayService/CreateCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IpayService service

type IpayServiceServer interface {
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error)
}

func RegisterIpayServiceServer(s *grpc.Server, srv IpayServiceServer) {
	s.RegisterService(&_IpayService_serviceDesc, srv)
}

func _IpayService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpayServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IpayService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpayServiceServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IpayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "IpayService",
	HandlerType: (*IpayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _IpayService_CreateCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipay.proto",
}

func init() { proto.RegisterFile("ipay.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x9f, 0xbe, 0x3e, 0x3a, 0xa9, 0x1e, 0xd6, 0x5a, 0x43, 0x2f, 0x96, 0x80, 0x50, 0x3c,
	0xac, 0x50, 0xaf, 0x9e, 0x2c, 0x22, 0xbd, 0xae, 0xe0, 0xc1, 0x8b, 0x6c, 0xb7, 0x43, 0x5d, 0xc8,
	0xbe, 0xb8, 0x9b, 0x08, 0xf9, 0x58, 0x7e, 0x43, 0xc9, 0x76, 0x53, 0xaa, 0x06, 0x6f, 0x3b, 0xf3,
	0xcf, 0xcc, 0xff, 0x97, 0x99, 0x01, 0x90, 0x96, 0x57, 0xd4, 0x3a, 0x53, 0x98, 0x4c, 0xc0, 0xf9,
	0xd2, 0x21, 0x2f, 0x70, 0x59, 0xfa, 0xc2, 0x28, 0x74, 0x0c, 0xdf, 0x4b, 0xf4, 0x05, 0x21, 0xd0,
	0xd7, 0x5c, 0x61, 0xda, 0x99, 0x75, 0xe6, 0xc7, 0x2c, 0xbc, 0xc9, 0x18, 0x06, 0xa8, 0xb8, 0xcc,
	0xd3, 0x6e, 0x48, 0xee, 0x02, 0x32, 0x83, 0xc4, 0xbe, 0x19, 0x8d, 0xba, 0x54, 0x6b, 0x74, 0x69,
	0x2f, 0x68, 0x87, 0xa9, 0xec, 0xb3, 0x03, 0x93, 0x9f, 0x2e, 0xde, 0x1a, 0xed, 0x91, 0x5c, 0xc3,
	0xc0, 0xa1, 0xcd, 0xab, 0xe0, 0x93, 0x2c, 0xc6, 0xf4, 0x09, 0xdd, 0x87, 0x14, 0xf8, 0x88, 0x1a,
	0x9d, 0x14, 0xac, 0xd6, 0xd8, 0xee, 0x13, 0x72, 0x09, 0x89, 0x88, 0xf5, 0xaf, 0x72, 0x13, 0x21,
	0xa0, 0x49, 0xad, 0x36, 0x7b, 0xe6, 0x5e, 0x1b, 0x73, 0xff, 0x0f, 0xe6, 0xc1, 0x6f, 0xe6, 0x3b,
	0x18, 0x45, 0x94, 0x07, 0xe7, 0x8c, 0xab, 0x7b, 0x0b, 0xb3, 0xd9, 0xcf, 0xa3, 0x7e, 0x93, 0x14,
	0xfe, 0x2b, 0xf4, 0x9e, 0x6f, 0x31, 0xc2, 0x34, 0x61, 0xf6, 0x0c, 0x67, 0x2d, 0x3f, 0x52, 0x17,
	0xf8, 0x52, 0x08, 0xf4, 0x3e, 0xf4, 0x39, 0x62, 0x4d, 0x48, 0xae, 0x60, 0x88, 0xb5, 0x8f, 0x4f,
	0xbb, 0xb3, 0xde, 0x3c, 0x59, 0x9c, 0xd0, 0x43, 0x77, 0x16, 0xc5, 0x05, 0x83, 0x64, 0x65, 0x79,
	0x15, 0x35, 0xb2, 0x84, 0xd3, 0xef, 0x73, 0x25, 0x13, 0xda, 0xba, 0xce, 0xe9, 0x05, 0x6d, 0x5f,
	0x40, 0xf6, 0xef, 0x7e, 0x0a, 0x23, 0x61, 0x14, 0x0d, 0x47, 0xc1, 0xad, 0x7c, 0x81, 0xad, 0xc9,
	0xb9, 0xde, 0xde, 0x70, 0x2b, 0xd7, 0xc3, 0x70, 0x25, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x1c, 0xe5, 0xa8, 0x71, 0x33, 0x02, 0x00, 0x00,
}
