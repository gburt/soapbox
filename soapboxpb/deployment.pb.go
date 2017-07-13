// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployment.proto

package soapboxpb

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

type ListDeploymentResponse struct {
	Deployments []*Deployment `protobuf:"bytes,1,rep,name=deployments" json:"deployments,omitempty"`
}

func (m *ListDeploymentResponse) Reset()                    { *m = ListDeploymentResponse{} }
func (m *ListDeploymentResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDeploymentResponse) ProtoMessage()               {}
func (*ListDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ListDeploymentResponse) GetDeployments() []*Deployment {
	if m != nil {
		return m.Deployments
	}
	return nil
}

type GetDeploymentRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetDeploymentRequest) Reset()                    { *m = GetDeploymentRequest{} }
func (m *GetDeploymentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeploymentRequest) ProtoMessage()               {}
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GetDeploymentRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Deployment struct {
	Id          int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Application *Application `protobuf:"bytes,2,opt,name=application" json:"application,omitempty"`
	Sha1OrTag   string       `protobuf:"bytes,3,opt,name=sha1_or_tag,json=sha1OrTag" json:"sha1_or_tag,omitempty"`
	State       string       `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Env         *Environment `protobuf:"bytes,5,opt,name=env" json:"env,omitempty"`
	CreatedAt   string       `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Deployment) Reset()                    { *m = Deployment{} }
func (m *Deployment) String() string            { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()               {}
func (*Deployment) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Deployment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Deployment) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Deployment) GetSha1OrTag() string {
	if m != nil {
		return m.Sha1OrTag
	}
	return ""
}

func (m *Deployment) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Deployment) GetEnv() *Environment {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Deployment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type StartDeploymentResponse struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *StartDeploymentResponse) Reset()                    { *m = StartDeploymentResponse{} }
func (m *StartDeploymentResponse) String() string            { return proto.CompactTextString(m) }
func (*StartDeploymentResponse) ProtoMessage()               {}
func (*StartDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *StartDeploymentResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetDeploymentStatusRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetDeploymentStatusRequest) Reset()                    { *m = GetDeploymentStatusRequest{} }
func (m *GetDeploymentStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeploymentStatusRequest) ProtoMessage()               {}
func (*GetDeploymentStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *GetDeploymentStatusRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetDeploymentStatusResponse struct {
	State string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
}

func (m *GetDeploymentStatusResponse) Reset()                    { *m = GetDeploymentStatusResponse{} }
func (m *GetDeploymentStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDeploymentStatusResponse) ProtoMessage()               {}
func (*GetDeploymentStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *GetDeploymentStatusResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type TeardownDeploymentRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *TeardownDeploymentRequest) Reset()                    { *m = TeardownDeploymentRequest{} }
func (m *TeardownDeploymentRequest) String() string            { return proto.CompactTextString(m) }
func (*TeardownDeploymentRequest) ProtoMessage()               {}
func (*TeardownDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *TeardownDeploymentRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ListDeploymentResponse)(nil), "soapbox.ListDeploymentResponse")
	proto.RegisterType((*GetDeploymentRequest)(nil), "soapbox.GetDeploymentRequest")
	proto.RegisterType((*Deployment)(nil), "soapbox.Deployment")
	proto.RegisterType((*StartDeploymentResponse)(nil), "soapbox.StartDeploymentResponse")
	proto.RegisterType((*GetDeploymentStatusRequest)(nil), "soapbox.GetDeploymentStatusRequest")
	proto.RegisterType((*GetDeploymentStatusResponse)(nil), "soapbox.GetDeploymentStatusResponse")
	proto.RegisterType((*TeardownDeploymentRequest)(nil), "soapbox.TeardownDeploymentRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deployments service

type DeploymentsClient interface {
	ListDeployments(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListDeploymentResponse, error)
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*Deployment, error)
	StartDeployment(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (*StartDeploymentResponse, error)
	GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusRequest, opts ...grpc.CallOption) (*GetDeploymentStatusResponse, error)
	TeardownDeployment(ctx context.Context, in *TeardownDeploymentRequest, opts ...grpc.CallOption) (*Empty, error)
}

type deploymentsClient struct {
	cc *grpc.ClientConn
}

func NewDeploymentsClient(cc *grpc.ClientConn) DeploymentsClient {
	return &deploymentsClient{cc}
}

func (c *deploymentsClient) ListDeployments(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListDeploymentResponse, error) {
	out := new(ListDeploymentResponse)
	err := grpc.Invoke(ctx, "/soapbox.Deployments/ListDeployments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentsClient) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*Deployment, error) {
	out := new(Deployment)
	err := grpc.Invoke(ctx, "/soapbox.Deployments/GetDeployment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentsClient) StartDeployment(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (*StartDeploymentResponse, error) {
	out := new(StartDeploymentResponse)
	err := grpc.Invoke(ctx, "/soapbox.Deployments/StartDeployment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentsClient) GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusRequest, opts ...grpc.CallOption) (*GetDeploymentStatusResponse, error) {
	out := new(GetDeploymentStatusResponse)
	err := grpc.Invoke(ctx, "/soapbox.Deployments/GetDeploymentStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentsClient) TeardownDeployment(ctx context.Context, in *TeardownDeploymentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/soapbox.Deployments/TeardownDeployment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Deployments service

type DeploymentsServer interface {
	ListDeployments(context.Context, *Empty) (*ListDeploymentResponse, error)
	GetDeployment(context.Context, *GetDeploymentRequest) (*Deployment, error)
	StartDeployment(context.Context, *Deployment) (*StartDeploymentResponse, error)
	GetDeploymentStatus(context.Context, *GetDeploymentStatusRequest) (*GetDeploymentStatusResponse, error)
	TeardownDeployment(context.Context, *TeardownDeploymentRequest) (*Empty, error)
}

func RegisterDeploymentsServer(s *grpc.Server, srv DeploymentsServer) {
	s.RegisterService(&_Deployments_serviceDesc, srv)
}

func _Deployments_ListDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).ListDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Deployments/ListDeployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).ListDeployments(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployments_GetDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).GetDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Deployments/GetDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).GetDeployment(ctx, req.(*GetDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployments_StartDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Deployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).StartDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Deployments/StartDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).StartDeployment(ctx, req.(*Deployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployments_GetDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).GetDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Deployments/GetDeploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).GetDeploymentStatus(ctx, req.(*GetDeploymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployments_TeardownDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).TeardownDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Deployments/TeardownDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).TeardownDeployment(ctx, req.(*TeardownDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "soapbox.Deployments",
	HandlerType: (*DeploymentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDeployments",
			Handler:    _Deployments_ListDeployments_Handler,
		},
		{
			MethodName: "GetDeployment",
			Handler:    _Deployments_GetDeployment_Handler,
		},
		{
			MethodName: "StartDeployment",
			Handler:    _Deployments_StartDeployment_Handler,
		},
		{
			MethodName: "GetDeploymentStatus",
			Handler:    _Deployments_GetDeploymentStatus_Handler,
		},
		{
			MethodName: "TeardownDeployment",
			Handler:    _Deployments_TeardownDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployment.proto",
}

func init() { proto.RegisterFile("deployment.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0x4e, 0x1a, 0x7b, 0xd2, 0x09, 0x77, 0xa7, 0x73, 0x45, 0xe3, 0xca, 0x69, 0x58, 0xe5, 0x88,
	0x28, 0x07, 0xf6, 0xd0, 0xf7, 0x93, 0x3b, 0x84, 0x22, 0x14, 0xd2, 0x3e, 0xf9, 0x52, 0x36, 0xcd,
	0x52, 0x03, 0x6d, 0x76, 0xcd, 0x6e, 0xab, 0xfd, 0x47, 0xfc, 0xbf, 0xfc, 0x8f, 0xa4, 0x6d, 0xba,
	0xf9, 0xe1, 0x56, 0xdf, 0xb2, 0xdf, 0x7c, 0x33, 0xdf, 0xcc, 0x37, 0x13, 0x78, 0x94, 0x72, 0xb9,
	0x10, 0x9b, 0x25, 0xcf, 0xf5, 0xb5, 0x2c, 0x84, 0x16, 0xf8, 0x50, 0x09, 0x26, 0x13, 0xf1, 0x93,
	0x9c, 0x96, 0x1f, 0x7b, 0x9c, 0x3c, 0x66, 0x52, 0x2e, 0xb2, 0x19, 0xd3, 0x99, 0xc8, 0x0f, 0x10,
	0xcf, 0xd7, 0x59, 0x21, 0xf2, 0x2a, 0x9b, 0x8e, 0xe0, 0xc9, 0x97, 0x4c, 0xe9, 0x3b, 0x53, 0x35,
	0xe6, 0x4a, 0x8a, 0x5c, 0x71, 0xfc, 0x00, 0x7e, 0xa5, 0xa5, 0x02, 0x37, 0xf4, 0x22, 0x7f, 0x70,
	0x71, 0x7d, 0x10, 0xa9, 0x65, 0xd4, 0x79, 0xf4, 0x0a, 0xfa, 0x9f, 0x79, 0xa3, 0xde, 0xf7, 0x15,
	0x57, 0x1a, 0xcf, 0xa0, 0x93, 0xa5, 0x81, 0x1b, 0xba, 0x51, 0x37, 0xee, 0x64, 0x29, 0xfd, 0xed,
	0x02, 0x54, 0xac, 0x76, 0x18, 0x3f, 0x82, 0x5f, 0xeb, 0x3f, 0xe8, 0x84, 0x6e, 0xe4, 0x0f, 0xfa,
	0x46, 0xfd, 0xb6, 0x8a, 0xc5, 0x75, 0x22, 0xbe, 0x00, 0x5f, 0x7d, 0x63, 0xef, 0xa7, 0xa2, 0x98,
	0x6a, 0x36, 0x0f, 0xbc, 0xd0, 0x8d, 0x7a, 0x71, 0x6f, 0x0b, 0x8d, 0x8a, 0x09, 0x9b, 0x63, 0x1f,
	0xba, 0x4a, 0x33, 0xcd, 0x83, 0x07, 0xbb, 0xc8, 0xfe, 0x81, 0x57, 0xe0, 0xf1, 0x7c, 0x1d, 0x74,
	0x5b, 0x2a, 0xf7, 0x95, 0x5d, 0xf1, 0x96, 0x80, 0x97, 0x00, 0xb3, 0x82, 0x33, 0xcd, 0xd3, 0x29,
	0xd3, 0xc1, 0xc9, 0xbe, 0x78, 0x89, 0xdc, 0x6a, 0xfa, 0x06, 0x9e, 0x8e, 0x35, 0x2b, 0x6c, 0x6e,
	0xb6, 0xc7, 0x7f, 0x07, 0xa4, 0x61, 0xd3, 0x58, 0x33, 0xbd, 0x52, 0xc7, 0xcc, 0xba, 0x81, 0xe7,
	0x56, 0x76, 0x59, 0xdc, 0x0c, 0xe5, 0xd6, 0x86, 0xa2, 0x6f, 0xe1, 0xd9, 0x84, 0xb3, 0x22, 0x15,
	0x3f, 0xf2, 0xff, 0xae, 0x63, 0xf0, 0xcb, 0x03, 0xbf, 0x62, 0x29, 0xbc, 0x83, 0xf3, 0xe6, 0x5d,
	0x28, 0x3c, 0xab, 0x7c, 0x59, 0x4a, 0xbd, 0x21, 0x2f, 0xcd, 0xdb, 0x7e, 0x41, 0xd4, 0xc1, 0x7b,
	0x38, 0x6d, 0xf4, 0x8d, 0x97, 0x26, 0xc7, 0x76, 0x24, 0xc4, 0x76, 0x5e, 0xd4, 0xc1, 0x21, 0x9c,
	0xb7, 0x7c, 0x45, 0x1b, 0x93, 0x84, 0x06, 0x3c, 0xb2, 0x06, 0xea, 0x60, 0x02, 0x17, 0x16, 0x2b,
	0xf1, 0x95, 0xbd, 0xb1, 0xc6, 0x5a, 0xc8, 0xeb, 0x7f, 0x93, 0x8c, 0xc6, 0x10, 0xf0, 0x6f, 0xe7,
	0x91, 0x9a, 0xec, 0xa3, 0x6b, 0x21, 0x2d, 0x8f, 0xa9, 0xf3, 0xc9, 0xff, 0xda, 0x2b, 0x21, 0x99,
	0x24, 0x27, 0xbb, 0x9f, 0xf6, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xab, 0x22, 0xed,
	0x06, 0x04, 0x00, 0x00,
}