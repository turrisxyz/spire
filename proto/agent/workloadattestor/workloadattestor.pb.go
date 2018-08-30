// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workloadattestor.proto

package workloadattestor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/spiffe/spire/proto/common"
import plugin "github.com/spiffe/spire/proto/common/plugin"

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

// Empty from public import github.com/spiffe/spire/proto/common/common.proto
type Empty = common.Empty

// AttestationData from public import github.com/spiffe/spire/proto/common/common.proto
type AttestationData = common.AttestationData

// Selector from public import github.com/spiffe/spire/proto/common/common.proto
type Selector = common.Selector

// Selectors from public import github.com/spiffe/spire/proto/common/common.proto
type Selectors = common.Selectors

// RegistrationEntry from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntry = common.RegistrationEntry

// RegistrationEntries from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntries = common.RegistrationEntries

// ConfigureRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest = plugin.ConfigureRequest

// GlobalConfig from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest_GlobalConfig = plugin.ConfigureRequest_GlobalConfig

// ConfigureResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureResponse = plugin.ConfigureResponse

// GetPluginInfoRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoRequest = plugin.GetPluginInfoRequest

// GetPluginInfoResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoResponse = plugin.GetPluginInfoResponse

// * Represents the workload PID.
type AttestRequest struct {
	// * Workload PID
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestRequest) Reset()         { *m = AttestRequest{} }
func (m *AttestRequest) String() string { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()    {}
func (*AttestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workloadattestor_4b37543ab2590fb4, []int{0}
}
func (m *AttestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestRequest.Unmarshal(m, b)
}
func (m *AttestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestRequest.Marshal(b, m, deterministic)
}
func (dst *AttestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestRequest.Merge(dst, src)
}
func (m *AttestRequest) XXX_Size() int {
	return xxx_messageInfo_AttestRequest.Size(m)
}
func (m *AttestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttestRequest proto.InternalMessageInfo

func (m *AttestRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

// * Represents a list of selectors resolved for a given PID.
type AttestResponse struct {
	// * List of selectors
	Selectors            []*common.Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AttestResponse) Reset()         { *m = AttestResponse{} }
func (m *AttestResponse) String() string { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()    {}
func (*AttestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_workloadattestor_4b37543ab2590fb4, []int{1}
}
func (m *AttestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestResponse.Unmarshal(m, b)
}
func (m *AttestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestResponse.Marshal(b, m, deterministic)
}
func (dst *AttestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestResponse.Merge(dst, src)
}
func (m *AttestResponse) XXX_Size() int {
	return xxx_messageInfo_AttestResponse.Size(m)
}
func (m *AttestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttestResponse proto.InternalMessageInfo

func (m *AttestResponse) GetSelectors() []*common.Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "spire.agent.workloadattestor.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "spire.agent.workloadattestor.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkloadAttestorClient is the client API for WorkloadAttestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkloadAttestorClient interface {
	// * Returns a list of selectors resolved for a given PID
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
	// * Applies the plugin configuration and returns configuration errors
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	// * Returns the version and related metadata of the plugin
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type workloadAttestorClient struct {
	cc *grpc.ClientConn
}

func NewWorkloadAttestorClient(cc *grpc.ClientConn) WorkloadAttestorClient {
	return &workloadAttestorClient{cc}
}

func (c *workloadAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.workloadattestor.WorkloadAttestor/Attest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.workloadattestor.WorkloadAttestor/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.workloadattestor.WorkloadAttestor/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkloadAttestorServer is the server API for WorkloadAttestor service.
type WorkloadAttestorServer interface {
	// * Returns a list of selectors resolved for a given PID
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
	// * Applies the plugin configuration and returns configuration errors
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	// * Returns the version and related metadata of the plugin
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

func RegisterWorkloadAttestorServer(s *grpc.Server, srv WorkloadAttestorServer) {
	s.RegisterService(&_WorkloadAttestor_serviceDesc, srv)
}

func _WorkloadAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.workloadattestor.WorkloadAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.workloadattestor.WorkloadAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.workloadattestor.WorkloadAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkloadAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.workloadattestor.WorkloadAttestor",
	HandlerType: (*WorkloadAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Attest",
			Handler:    _WorkloadAttestor_Attest_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _WorkloadAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _WorkloadAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workloadattestor.proto",
}

func init() {
	proto.RegisterFile("workloadattestor.proto", fileDescriptor_workloadattestor_4b37543ab2590fb4)
}

var fileDescriptor_workloadattestor_4b37543ab2590fb4 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xad, 0xc3, 0xc1, 0x22, 0x93, 0x91, 0xc3, 0x18, 0xc5, 0xc3, 0x1c, 0x28, 0xf3, 0x0f,
	0x29, 0x4e, 0x0f, 0x5e, 0xa7, 0xa0, 0x78, 0x2b, 0xf5, 0x20, 0xec, 0xd6, 0x75, 0x6f, 0x6b, 0xb0,
	0xcd, 0x1b, 0x93, 0x14, 0x3f, 0x9c, 0x5f, 0x4e, 0x4c, 0x52, 0xa5, 0x43, 0xb4, 0xa7, 0x04, 0xde,
	0xdf, 0x93, 0xe7, 0x79, 0xde, 0x90, 0xf1, 0x3b, 0xaa, 0xd7, 0x12, 0xd3, 0x4d, 0x6a, 0x0c, 0x68,
	0x83, 0x8a, 0x49, 0x85, 0x06, 0xe9, 0xa1, 0x96, 0x5c, 0x01, 0x4b, 0x0b, 0x10, 0x86, 0x6d, 0x33,
	0xe1, 0x65, 0xc1, 0xcd, 0x4b, 0xbd, 0x66, 0x19, 0x56, 0x91, 0x96, 0x3c, 0xcf, 0x21, 0xb2, 0x7c,
	0x64, 0xc5, 0x51, 0x86, 0x55, 0x85, 0xc2, 0x1f, 0xee, 0xc1, 0xf0, 0xa6, 0x93, 0x44, 0x96, 0x75,
	0xc1, 0x9b, 0xc3, 0x29, 0x67, 0x47, 0x64, 0xb8, 0xb4, 0xc6, 0x09, 0xbc, 0xd5, 0xa0, 0x0d, 0x1d,
	0x91, 0x9e, 0xe4, 0x9b, 0x49, 0x30, 0x0d, 0xe6, 0x7b, 0xc9, 0xd7, 0x75, 0x76, 0x4f, 0x0e, 0x1a,
	0x44, 0x4b, 0x14, 0x1a, 0xe8, 0x35, 0x19, 0x68, 0x28, 0x21, 0x33, 0xa8, 0xf4, 0x24, 0x98, 0xf6,
	0xe6, 0xfb, 0x8b, 0x31, 0x73, 0x9d, 0x7c, 0xac, 0x27, 0x3f, 0x4e, 0x7e, 0xc0, 0xc5, 0xc7, 0x2e,
	0x19, 0x3d, 0xfb, 0xb2, 0x4b, 0x5f, 0x96, 0x66, 0xa4, 0xef, 0xee, 0xf4, 0x9c, 0xfd, 0xb5, 0x15,
	0xd6, 0x4a, 0x19, 0x5e, 0x74, 0x83, 0x7d, 0xde, 0x15, 0x19, 0xdc, 0xa1, 0xc8, 0x79, 0x51, 0x2b,
	0xa0, 0xc7, 0xed, 0xa4, 0x7e, 0x1b, 0xdf, 0xf3, 0xc6, 0xe1, 0xe4, 0x3f, 0xcc, 0xbf, 0x9d, 0x93,
	0xe1, 0x03, 0x98, 0xd8, 0x8e, 0x1f, 0x45, 0x8e, 0xf4, 0xf4, 0x57, 0x61, 0x8b, 0x69, 0x3c, 0xce,
	0xba, 0xa0, 0xce, 0xe7, 0x96, 0xae, 0x46, 0xdb, 0x35, 0xe3, 0x9d, 0x38, 0x58, 0xf7, 0xed, 0x3f,
	0x5e, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x1d, 0xe6, 0x91, 0x6c, 0x02, 0x00, 0x00,
}
