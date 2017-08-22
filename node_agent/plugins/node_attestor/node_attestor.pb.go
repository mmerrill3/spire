// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_attestor.proto

/*
Package nodeattestor is a generated protocol buffer package.

It is generated from these files:
	node_attestor.proto

It has these top-level messages:
	AttestedData
	FetchAttestationDataRequest
	FetchAttestationDataResponse
*/
package nodeattestor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sriplugin "github.com/spiffe/sri/common/plugin"

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

// ConfigureRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type ConfigureRequest sriplugin.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*sriplugin.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*sriplugin.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*sriplugin.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/common/plugin/plugin.proto
type ConfigureResponse sriplugin.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*sriplugin.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*sriplugin.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*sriplugin.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type GetPluginInfoRequest sriplugin.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*sriplugin.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*sriplugin.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/common/plugin/plugin.proto
type GetPluginInfoResponse sriplugin.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()         { (*sriplugin.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string { return (*sriplugin.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetCompany()
}

// PluginInfoRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type PluginInfoRequest sriplugin.PluginInfoRequest

func (m *PluginInfoRequest) Reset()         { (*sriplugin.PluginInfoRequest)(m).Reset() }
func (m *PluginInfoRequest) String() string { return (*sriplugin.PluginInfoRequest)(m).String() }
func (*PluginInfoRequest) ProtoMessage()    {}

// PluginInfoReply from public import github.com/spiffe/sri/common/plugin/plugin.proto
type PluginInfoReply sriplugin.PluginInfoReply

func (m *PluginInfoReply) Reset()         { (*sriplugin.PluginInfoReply)(m).Reset() }
func (m *PluginInfoReply) String() string { return (*sriplugin.PluginInfoReply)(m).String() }
func (*PluginInfoReply) ProtoMessage()    {}
func (m *PluginInfoReply) GetPluginInfo() []*GetPluginInfoResponse {
	o := (*sriplugin.PluginInfoReply)(m).GetPluginInfo()
	if o == nil {
		return nil
	}
	s := make([]*GetPluginInfoResponse, len(o))
	for i, x := range o {
		s[i] = (*GetPluginInfoResponse)(x)
	}
	return s
}

// StopRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type StopRequest sriplugin.StopRequest

func (m *StopRequest) Reset()         { (*sriplugin.StopRequest)(m).Reset() }
func (m *StopRequest) String() string { return (*sriplugin.StopRequest)(m).String() }
func (*StopRequest) ProtoMessage()    {}

// StopReply from public import github.com/spiffe/sri/common/plugin/plugin.proto
type StopReply sriplugin.StopReply

func (m *StopReply) Reset()         { (*sriplugin.StopReply)(m).Reset() }
func (m *StopReply) String() string { return (*sriplugin.StopReply)(m).String() }
func (*StopReply) ProtoMessage()    {}

// *A type which contains attestation data for specific platform.
type AttestedData struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AttestedData) Reset()                    { *m = AttestedData{} }
func (m *AttestedData) String() string            { return proto.CompactTextString(m) }
func (*AttestedData) ProtoMessage()               {}
func (*AttestedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AttestedData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestedData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// * Represents an empty request.
type FetchAttestationDataRequest struct {
}

func (m *FetchAttestationDataRequest) Reset()                    { *m = FetchAttestationDataRequest{} }
func (m *FetchAttestationDataRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchAttestationDataRequest) ProtoMessage()               {}
func (*FetchAttestationDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// * Represents the attested data and base SPIFFE ID.
type FetchAttestationDataResponse struct {
	AttestedData *AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	SpiffeId     string        `protobuf:"bytes,2,opt,name=spiffeId" json:"spiffeId,omitempty"`
}

func (m *FetchAttestationDataResponse) Reset()                    { *m = FetchAttestationDataResponse{} }
func (m *FetchAttestationDataResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchAttestationDataResponse) ProtoMessage()               {}
func (*FetchAttestationDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchAttestationDataResponse) GetAttestedData() *AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *FetchAttestationDataResponse) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func init() {
	proto.RegisterType((*AttestedData)(nil), "nodeattestor.AttestedData")
	proto.RegisterType((*FetchAttestationDataRequest)(nil), "nodeattestor.FetchAttestationDataRequest")
	proto.RegisterType((*FetchAttestationDataResponse)(nil), "nodeattestor.FetchAttestationDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	// / Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation.
	FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error)
	// / Applies the plugin configuration and returns configuration errors.
	Configure(ctx context.Context, in *sriplugin.ConfigureRequest, opts ...grpc.CallOption) (*sriplugin.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin.
	GetPluginInfo(ctx context.Context, in *sriplugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*sriplugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error) {
	out := new(FetchAttestationDataResponse)
	err := grpc.Invoke(ctx, "/nodeattestor.NodeAttestor/FetchAttestationData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *sriplugin.ConfigureRequest, opts ...grpc.CallOption) (*sriplugin.ConfigureResponse, error) {
	out := new(sriplugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/nodeattestor.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *sriplugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*sriplugin.GetPluginInfoResponse, error) {
	out := new(sriplugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/nodeattestor.NodeAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	// / Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation.
	FetchAttestationData(context.Context, *FetchAttestationDataRequest) (*FetchAttestationDataResponse, error)
	// / Applies the plugin configuration and returns configuration errors.
	Configure(context.Context, *sriplugin.ConfigureRequest) (*sriplugin.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin.
	GetPluginInfo(context.Context, *sriplugin.GetPluginInfoRequest) (*sriplugin.GetPluginInfoResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_FetchAttestationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAttestationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeattestor.NodeAttestor/FetchAttestationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, req.(*FetchAttestationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sriplugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*sriplugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sriplugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*sriplugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAttestationData",
			Handler:    _NodeAttestor_FetchAttestationData_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_attestor.proto",
}

func init() { proto.RegisterFile("node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0xb6, 0x8d, 0x31, 0x76, 0xc5, 0xcb, 0xea, 0xa1, 0xa1, 0x35, 0x12, 0x4e, 0xd5, 0x03, 0x98,
	0x9a, 0x78, 0x34, 0x69, 0x34, 0x35, 0xbd, 0x98, 0x66, 0x5f, 0xc0, 0x6c, 0x61, 0xa0, 0x9b, 0xc8,
	0x0e, 0xb2, 0xc3, 0x41, 0xdf, 0xc4, 0xb7, 0x35, 0x2c, 0xd0, 0x40, 0x82, 0xc6, 0x13, 0xc3, 0xf7,
	0x37, 0xf3, 0x01, 0xbb, 0xd0, 0x18, 0xc3, 0x9b, 0x24, 0x02, 0x43, 0x58, 0x04, 0x79, 0x81, 0x84,
	0xdc, 0xa9, 0xc0, 0x16, 0x73, 0xef, 0x52, 0x45, 0xfb, 0x72, 0x17, 0x44, 0x98, 0x85, 0x26, 0x57,
	0x49, 0x02, 0xa1, 0x29, 0x54, 0x18, 0x61, 0x96, 0xa1, 0x0e, 0xf3, 0xf7, 0x32, 0x55, 0xed, 0xa3,
	0xf6, 0xfb, 0x0f, 0xcc, 0x59, 0x59, 0x37, 0xc4, 0xcf, 0x92, 0x24, 0xe7, 0xec, 0x98, 0x3e, 0x73,
	0x98, 0x8e, 0xbc, 0xd1, 0x62, 0x22, 0xec, 0x5c, 0x61, 0xb1, 0x24, 0x39, 0x1d, 0x7b, 0xa3, 0x85,
	0x23, 0xec, 0xec, 0x5f, 0xb1, 0xd9, 0x1a, 0x28, 0xda, 0xd7, 0x66, 0x49, 0x0a, 0x75, 0xe5, 0x17,
	0xf0, 0x51, 0x82, 0x21, 0xff, 0x8b, 0xcd, 0x87, 0x69, 0x93, 0xa3, 0x36, 0xc0, 0x1f, 0x99, 0x23,
	0x3b, 0x6b, 0xed, 0xba, 0xb3, 0xa5, 0x1b, 0x74, 0xdb, 0x04, 0xdd, 0xc3, 0x44, 0x4f, 0xcf, 0x5d,
	0x76, 0x5a, 0xf7, 0xdb, 0xc4, 0xf6, 0xac, 0x89, 0x38, 0xbc, 0x2f, 0xbf, 0xc7, 0xcc, 0x79, 0xc5,
	0x18, 0x56, 0x4d, 0x0e, 0xcf, 0xd8, 0xe5, 0xd0, 0x31, 0xfc, 0xa6, 0xbf, 0xee, 0x8f, 0x3e, 0xee,
	0xed, 0x7f, 0xa4, 0x4d, 0xb7, 0x35, 0x9b, 0x3c, 0xa1, 0x4e, 0x54, 0x5a, 0x16, 0xc0, 0x67, 0x81,
	0x29, 0x54, 0xf3, 0xc5, 0x0f, 0x68, 0x9b, 0x3a, 0x1f, 0x26, 0x9b, 0x1c, 0xc1, 0xce, 0x5f, 0x80,
	0xb6, 0x96, 0xde, 0xe8, 0x04, 0xf9, 0x75, 0x47, 0xde, 0x63, 0xda, 0x3c, 0xef, 0x77, 0x41, 0x9d,
	0xb9, 0x3d, 0xda, 0x9d, 0xd8, 0xff, 0x7e, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xdb, 0xe1,
	0xe9, 0x4e, 0x02, 0x00, 0x00,
}