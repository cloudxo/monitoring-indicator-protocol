// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shard_group_reader.proto

package logcache_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import loggregator_v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SetShardGroupRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SubGroup *GroupedSourceIds `protobuf:"bytes,2,opt,name=sub_group,json=subGroup" json:"sub_group,omitempty"`
	// local_only is used for internals only. A client should not set this.
	LocalOnly bool `protobuf:"varint,3,opt,name=local_only,json=localOnly" json:"local_only,omitempty"`
}

func (m *SetShardGroupRequest) Reset()                    { *m = SetShardGroupRequest{} }
func (m *SetShardGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*SetShardGroupRequest) ProtoMessage()               {}
func (*SetShardGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *SetShardGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetShardGroupRequest) GetSubGroup() *GroupedSourceIds {
	if m != nil {
		return m.SubGroup
	}
	return nil
}

func (m *SetShardGroupRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type GroupedSourceIds struct {
	SourceIds []string `protobuf:"bytes,1,rep,name=source_ids,json=sourceIds" json:"source_ids,omitempty"`
	// arg are given back to the requester when they recieve the given
	// sub_group.
	Arg string `protobuf:"bytes,2,opt,name=arg" json:"arg,omitempty"`
}

func (m *GroupedSourceIds) Reset()                    { *m = GroupedSourceIds{} }
func (m *GroupedSourceIds) String() string            { return proto.CompactTextString(m) }
func (*GroupedSourceIds) ProtoMessage()               {}
func (*GroupedSourceIds) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *GroupedSourceIds) GetSourceIds() []string {
	if m != nil {
		return m.SourceIds
	}
	return nil
}

func (m *GroupedSourceIds) GetArg() string {
	if m != nil {
		return m.Arg
	}
	return ""
}

type SetShardGroupResponse struct {
}

func (m *SetShardGroupResponse) Reset()                    { *m = SetShardGroupResponse{} }
func (m *SetShardGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*SetShardGroupResponse) ProtoMessage()               {}
func (*SetShardGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type ShardGroupReadRequest struct {
	Name          string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RequesterId   uint64         `protobuf:"varint,2,opt,name=requester_id,json=requesterId" json:"requester_id,omitempty"`
	StartTime     int64          `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime       int64          `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	Limit         int64          `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	EnvelopeTypes []EnvelopeType `protobuf:"varint,6,rep,packed,name=envelope_types,json=envelopeTypes,enum=logcache.v1.EnvelopeType" json:"envelope_types,omitempty"`
	// local_only is used for internals only. A client should not set this.
	LocalOnly bool `protobuf:"varint,7,opt,name=local_only,json=localOnly" json:"local_only,omitempty"`
}

func (m *ShardGroupReadRequest) Reset()                    { *m = ShardGroupReadRequest{} }
func (m *ShardGroupReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ShardGroupReadRequest) ProtoMessage()               {}
func (*ShardGroupReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ShardGroupReadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShardGroupReadRequest) GetRequesterId() uint64 {
	if m != nil {
		return m.RequesterId
	}
	return 0
}

func (m *ShardGroupReadRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ShardGroupReadRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *ShardGroupReadRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ShardGroupReadRequest) GetEnvelopeTypes() []EnvelopeType {
	if m != nil {
		return m.EnvelopeTypes
	}
	return nil
}

func (m *ShardGroupReadRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type ShardGroupReadResponse struct {
	Envelopes *loggregator_v2.EnvelopeBatch `protobuf:"bytes,1,opt,name=envelopes" json:"envelopes,omitempty"`
	// args are given back to the requester when they recieve the given
	// sub_group.
	Args []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *ShardGroupReadResponse) Reset()                    { *m = ShardGroupReadResponse{} }
func (m *ShardGroupReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ShardGroupReadResponse) ProtoMessage()               {}
func (*ShardGroupReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ShardGroupReadResponse) GetEnvelopes() *loggregator_v2.EnvelopeBatch {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func (m *ShardGroupReadResponse) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ShardGroupRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// local_only is used for internals only. A client should not set this.
	LocalOnly bool `protobuf:"varint,2,opt,name=local_only,json=localOnly" json:"local_only,omitempty"`
}

func (m *ShardGroupRequest) Reset()                    { *m = ShardGroupRequest{} }
func (m *ShardGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*ShardGroupRequest) ProtoMessage()               {}
func (*ShardGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *ShardGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShardGroupRequest) GetLocalOnly() bool {
	if m != nil {
		return m.LocalOnly
	}
	return false
}

type ShardGroupResponse struct {
	SubGroups    []*GroupedSourceIds `protobuf:"bytes,1,rep,name=sub_groups,json=subGroups" json:"sub_groups,omitempty"`
	RequesterIds []uint64            `protobuf:"varint,2,rep,packed,name=requester_ids,json=requesterIds" json:"requester_ids,omitempty"`
	Args         []string            `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
}

func (m *ShardGroupResponse) Reset()                    { *m = ShardGroupResponse{} }
func (m *ShardGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*ShardGroupResponse) ProtoMessage()               {}
func (*ShardGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ShardGroupResponse) GetSubGroups() []*GroupedSourceIds {
	if m != nil {
		return m.SubGroups
	}
	return nil
}

func (m *ShardGroupResponse) GetRequesterIds() []uint64 {
	if m != nil {
		return m.RequesterIds
	}
	return nil
}

func (m *ShardGroupResponse) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*SetShardGroupRequest)(nil), "logcache.v1.SetShardGroupRequest")
	proto.RegisterType((*GroupedSourceIds)(nil), "logcache.v1.GroupedSourceIds")
	proto.RegisterType((*SetShardGroupResponse)(nil), "logcache.v1.SetShardGroupResponse")
	proto.RegisterType((*ShardGroupReadRequest)(nil), "logcache.v1.ShardGroupReadRequest")
	proto.RegisterType((*ShardGroupReadResponse)(nil), "logcache.v1.ShardGroupReadResponse")
	proto.RegisterType((*ShardGroupRequest)(nil), "logcache.v1.ShardGroupRequest")
	proto.RegisterType((*ShardGroupResponse)(nil), "logcache.v1.ShardGroupResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShardGroupReader service

type ShardGroupReaderClient interface {
	SetShardGroup(ctx context.Context, in *SetShardGroupRequest, opts ...grpc.CallOption) (*SetShardGroupResponse, error)
	Read(ctx context.Context, in *ShardGroupReadRequest, opts ...grpc.CallOption) (*ShardGroupReadResponse, error)
	ShardGroup(ctx context.Context, in *ShardGroupRequest, opts ...grpc.CallOption) (*ShardGroupResponse, error)
}

type shardGroupReaderClient struct {
	cc *grpc.ClientConn
}

func NewShardGroupReaderClient(cc *grpc.ClientConn) ShardGroupReaderClient {
	return &shardGroupReaderClient{cc}
}

func (c *shardGroupReaderClient) SetShardGroup(ctx context.Context, in *SetShardGroupRequest, opts ...grpc.CallOption) (*SetShardGroupResponse, error) {
	out := new(SetShardGroupResponse)
	err := grpc.Invoke(ctx, "/logcache.v1.ShardGroupReader/SetShardGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardGroupReaderClient) Read(ctx context.Context, in *ShardGroupReadRequest, opts ...grpc.CallOption) (*ShardGroupReadResponse, error) {
	out := new(ShardGroupReadResponse)
	err := grpc.Invoke(ctx, "/logcache.v1.ShardGroupReader/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardGroupReaderClient) ShardGroup(ctx context.Context, in *ShardGroupRequest, opts ...grpc.CallOption) (*ShardGroupResponse, error) {
	out := new(ShardGroupResponse)
	err := grpc.Invoke(ctx, "/logcache.v1.ShardGroupReader/ShardGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShardGroupReader service

type ShardGroupReaderServer interface {
	SetShardGroup(context.Context, *SetShardGroupRequest) (*SetShardGroupResponse, error)
	Read(context.Context, *ShardGroupReadRequest) (*ShardGroupReadResponse, error)
	ShardGroup(context.Context, *ShardGroupRequest) (*ShardGroupResponse, error)
}

func RegisterShardGroupReaderServer(s *grpc.Server, srv ShardGroupReaderServer) {
	s.RegisterService(&_ShardGroupReader_serviceDesc, srv)
}

func _ShardGroupReader_SetShardGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetShardGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardGroupReaderServer).SetShardGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.ShardGroupReader/SetShardGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardGroupReaderServer).SetShardGroup(ctx, req.(*SetShardGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardGroupReader_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardGroupReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardGroupReaderServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.ShardGroupReader/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardGroupReaderServer).Read(ctx, req.(*ShardGroupReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardGroupReader_ShardGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardGroupReaderServer).ShardGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.ShardGroupReader/ShardGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardGroupReaderServer).ShardGroup(ctx, req.(*ShardGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShardGroupReader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.v1.ShardGroupReader",
	HandlerType: (*ShardGroupReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetShardGroup",
			Handler:    _ShardGroupReader_SetShardGroup_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ShardGroupReader_Read_Handler,
		},
		{
			MethodName: "ShardGroup",
			Handler:    _ShardGroupReader_ShardGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shard_group_reader.proto",
}

func init() { proto.RegisterFile("shard_group_reader.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0xe3, 0xb4, 0x8d, 0x27, 0x49, 0x95, 0xae, 0xda, 0x7e, 0x6e, 0x3e, 0x52, 0x52, 0xe7,
	0x12, 0x2e, 0xb1, 0x1a, 0x6e, 0x85, 0x03, 0x02, 0x01, 0xea, 0x09, 0x69, 0xd3, 0xbb, 0xb5, 0x89,
	0x47, 0x8e, 0x25, 0xc7, 0x6b, 0x76, 0x37, 0x91, 0x22, 0xc4, 0x01, 0x24, 0x4e, 0x1c, 0xf9, 0x69,
	0xfc, 0x05, 0xfe, 0x05, 0x07, 0x90, 0xd7, 0x4e, 0xe2, 0x98, 0x24, 0x70, 0xdb, 0x9d, 0x19, 0xcf,
	0x7b, 0xf3, 0xe6, 0x79, 0xc1, 0x96, 0x53, 0x26, 0x7c, 0x2f, 0x10, 0x7c, 0x9e, 0x78, 0x02, 0x99,
	0x8f, 0x62, 0x90, 0x08, 0xae, 0x38, 0xa9, 0x47, 0x3c, 0x98, 0xb0, 0xc9, 0x14, 0x07, 0x8b, 0xdb,
	0xf6, 0xa3, 0x80, 0xf3, 0x20, 0x42, 0x97, 0x25, 0xa1, 0xcb, 0xe2, 0x98, 0x2b, 0xa6, 0x42, 0x1e,
	0xcb, 0xac, 0xb4, 0x7d, 0xb6, 0x18, 0xba, 0x18, 0x2f, 0x30, 0xe2, 0x09, 0xe6, 0xa1, 0x06, 0x06,
	0x02, 0x65, 0x5e, 0xe0, 0x7c, 0x31, 0xe0, 0x7c, 0x84, 0x6a, 0x94, 0x62, 0xbd, 0x4d, 0xa1, 0x28,
	0xbe, 0x9f, 0xa3, 0x54, 0x84, 0x40, 0x35, 0x66, 0x33, 0xb4, 0x8d, 0xae, 0xd1, 0xb7, 0xa8, 0x3e,
	0x93, 0x3b, 0xb0, 0xe4, 0x7c, 0x9c, 0x51, 0xb2, 0x2b, 0x5d, 0xa3, 0x5f, 0x1f, 0x76, 0x06, 0x05,
	0x32, 0x03, 0xdd, 0x01, 0xfd, 0x11, 0x9f, 0x8b, 0x09, 0xde, 0xfb, 0x92, 0xd6, 0xe4, 0x7c, 0xac,
	0x83, 0xa4, 0x03, 0x10, 0xf1, 0x09, 0x8b, 0x3c, 0x1e, 0x47, 0x4b, 0xdb, 0xec, 0x1a, 0xfd, 0x1a,
	0xb5, 0x74, 0xe4, 0x5d, 0x1c, 0x2d, 0x9d, 0x57, 0xd0, 0x2a, 0x7f, 0x9c, 0x7e, 0x22, 0xf5, 0xc5,
	0x0b, 0x7d, 0x69, 0x1b, 0x5d, 0xb3, 0x6f, 0x51, 0x4b, 0xae, 0xd3, 0x2d, 0x30, 0x99, 0x08, 0x34,
	0x0f, 0x8b, 0xa6, 0x47, 0xe7, 0x3f, 0xb8, 0x28, 0xcd, 0x22, 0x13, 0x1e, 0x4b, 0x74, 0x7e, 0x19,
	0x70, 0x51, 0x0c, 0x33, 0xff, 0xd0, 0x98, 0x37, 0xd0, 0x10, 0x59, 0x1a, 0x85, 0x17, 0xfa, 0x1a,
	0xa1, 0x4a, 0xeb, 0xeb, 0xd8, 0xbd, 0xaf, 0xa9, 0x29, 0x26, 0x94, 0xa7, 0xc2, 0x19, 0xea, 0x69,
	0x4c, 0x6a, 0xe9, 0xc8, 0x43, 0x38, 0x43, 0x72, 0x05, 0x35, 0x8c, 0xfd, 0x2c, 0x59, 0xd5, 0xc9,
	0x13, 0x8c, 0x7d, 0x9d, 0x3a, 0x87, 0xa3, 0x28, 0x9c, 0x85, 0xca, 0x3e, 0xd2, 0xf1, 0xec, 0x42,
	0x5e, 0xc0, 0xe9, 0x6a, 0x4d, 0x9e, 0x5a, 0x26, 0x28, 0xed, 0xe3, 0xae, 0xd9, 0x3f, 0x1d, 0x5e,
	0x6d, 0xc9, 0xfb, 0x3a, 0x2f, 0x79, 0x58, 0x26, 0x48, 0x9b, 0x58, 0xb8, 0xc9, 0x92, 0xbe, 0x27,
	0x65, 0x7d, 0x43, 0xb8, 0x2c, 0x0b, 0x90, 0x69, 0x43, 0x9e, 0x81, 0xb5, 0xea, 0x24, 0xb5, 0x0c,
	0xf9, 0x52, 0x03, 0x81, 0x01, 0x53, 0x5c, 0x0c, 0x16, 0xc3, 0x35, 0xf0, 0x4b, 0xa6, 0x26, 0x53,
	0xba, 0xa9, 0x4f, 0xe5, 0x63, 0x22, 0x90, 0x76, 0x45, 0x2f, 0x47, 0x9f, 0x9d, 0x37, 0x70, 0xf6,
	0x6f, 0x76, 0xda, 0xa6, 0x5c, 0x29, 0x53, 0xfe, 0x6a, 0x00, 0xf9, 0x73, 0x97, 0xe4, 0x39, 0xc0,
	0xda, 0x84, 0x99, 0x2b, 0xfe, 0xea, 0x42, 0x6b, 0xe5, 0x42, 0x49, 0x7a, 0xd0, 0x2c, 0xee, 0x36,
	0x63, 0x5e, 0xa5, 0x8d, 0xc2, 0x72, 0x37, 0x53, 0x99, 0x9b, 0xa9, 0x86, 0x3f, 0x2b, 0xd0, 0xda,
	0x56, 0x10, 0x05, 0xf9, 0x64, 0x40, 0x73, 0xcb, 0x71, 0xe4, 0x66, 0x8b, 0xc9, 0xae, 0x3f, 0xab,
	0xed, 0x1c, 0x2a, 0xc9, 0x0d, 0xfb, 0xe4, 0xf3, 0xf7, 0x1f, 0xdf, 0x2a, 0xbd, 0xf6, 0xa5, 0xbb,
	0xb8, 0x75, 0x0b, 0x0f, 0x81, 0xfb, 0x21, 0x55, 0xee, 0xe3, 0xdd, 0xe6, 0x3f, 0x24, 0x1c, 0xaa,
	0x29, 0x1b, 0x52, 0x6a, 0xbb, 0xcb, 0xed, 0xed, 0xde, 0xc1, 0x9a, 0x1c, 0xfb, 0x5a, 0x63, 0xdb,
	0x64, 0x0f, 0x36, 0x11, 0x00, 0x85, 0x81, 0xaf, 0xf7, 0xb6, 0xcc, 0x20, 0x1f, 0xef, 0xcd, 0xe7,
	0x70, 0x3d, 0x0d, 0xd7, 0x21, 0xff, 0xef, 0x86, 0x73, 0x67, 0xa8, 0xd8, 0xf8, 0x58, 0xbf, 0x56,
	0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x98, 0xca, 0x78, 0x15, 0x05, 0x00, 0x00,
}