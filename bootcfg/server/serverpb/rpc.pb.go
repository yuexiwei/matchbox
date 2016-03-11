// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	GroupGetRequest
	GroupListRequest
	GroupGetResponse
	GroupListResponse
	ProfileGetRequest
	ProfileGetResponse
	ProfileListRequest
	ProfileListResponse
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import storagepb "github.com/coreos/coreos-baremetal/bootcfg/storage/storagepb"

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
const _ = proto.ProtoPackageIsVersion1

type GroupGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GroupGetRequest) Reset()                    { *m = GroupGetRequest{} }
func (m *GroupGetRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupGetRequest) ProtoMessage()               {}
func (*GroupGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GroupListRequest struct {
}

func (m *GroupListRequest) Reset()                    { *m = GroupListRequest{} }
func (m *GroupListRequest) String() string            { return proto.CompactTextString(m) }
func (*GroupListRequest) ProtoMessage()               {}
func (*GroupListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GroupGetResponse struct {
	Group *storagepb.Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *GroupGetResponse) Reset()                    { *m = GroupGetResponse{} }
func (m *GroupGetResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupGetResponse) ProtoMessage()               {}
func (*GroupGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GroupGetResponse) GetGroup() *storagepb.Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GroupListResponse struct {
	Groups []*storagepb.Group `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
}

func (m *GroupListResponse) Reset()                    { *m = GroupListResponse{} }
func (m *GroupListResponse) String() string            { return proto.CompactTextString(m) }
func (*GroupListResponse) ProtoMessage()               {}
func (*GroupListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupListResponse) GetGroups() []*storagepb.Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ProfileGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ProfileGetRequest) Reset()                    { *m = ProfileGetRequest{} }
func (m *ProfileGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileGetRequest) ProtoMessage()               {}
func (*ProfileGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ProfileGetResponse struct {
	Profile *storagepb.Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
}

func (m *ProfileGetResponse) Reset()                    { *m = ProfileGetResponse{} }
func (m *ProfileGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileGetResponse) ProtoMessage()               {}
func (*ProfileGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ProfileGetResponse) GetProfile() *storagepb.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ProfileListRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ProfileListRequest) Reset()                    { *m = ProfileListRequest{} }
func (m *ProfileListRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileListRequest) ProtoMessage()               {}
func (*ProfileListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ProfileListResponse struct {
	Profiles []*storagepb.Profile `protobuf:"bytes,1,rep,name=profiles" json:"profiles,omitempty"`
}

func (m *ProfileListResponse) Reset()                    { *m = ProfileListResponse{} }
func (m *ProfileListResponse) String() string            { return proto.CompactTextString(m) }
func (*ProfileListResponse) ProtoMessage()               {}
func (*ProfileListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ProfileListResponse) GetProfiles() []*storagepb.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupGetRequest)(nil), "serverpb.GroupGetRequest")
	proto.RegisterType((*GroupListRequest)(nil), "serverpb.GroupListRequest")
	proto.RegisterType((*GroupGetResponse)(nil), "serverpb.GroupGetResponse")
	proto.RegisterType((*GroupListResponse)(nil), "serverpb.GroupListResponse")
	proto.RegisterType((*ProfileGetRequest)(nil), "serverpb.ProfileGetRequest")
	proto.RegisterType((*ProfileGetResponse)(nil), "serverpb.ProfileGetResponse")
	proto.RegisterType((*ProfileListRequest)(nil), "serverpb.ProfileListRequest")
	proto.RegisterType((*ProfileListResponse)(nil), "serverpb.ProfileListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Groups service

type GroupsClient interface {
	// Get a machine Group by id.
	GroupGet(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupGetResponse, error)
	// List all machine Groups.
	GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error)
}

type groupsClient struct {
	cc *grpc.ClientConn
}

func NewGroupsClient(cc *grpc.ClientConn) GroupsClient {
	return &groupsClient{cc}
}

func (c *groupsClient) GroupGet(ctx context.Context, in *GroupGetRequest, opts ...grpc.CallOption) (*GroupGetResponse, error) {
	out := new(GroupGetResponse)
	err := grpc.Invoke(ctx, "/serverpb.Groups/GroupGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error) {
	out := new(GroupListResponse)
	err := grpc.Invoke(ctx, "/serverpb.Groups/GroupList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Groups service

type GroupsServer interface {
	// Get a machine Group by id.
	GroupGet(context.Context, *GroupGetRequest) (*GroupGetResponse, error)
	// List all machine Groups.
	GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error)
}

func RegisterGroupsServer(s *grpc.Server, srv GroupsServer) {
	s.RegisterService(&_Groups_serviceDesc, srv)
}

func _Groups_GroupGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupsServer).GroupGet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Groups_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupsServer).GroupList(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Groups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Groups",
	HandlerType: (*GroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupGet",
			Handler:    _Groups_GroupGet_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Groups_GroupList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Profiles service

type ProfilesClient interface {
	// Get a Profile by id.
	ProfileGet(ctx context.Context, in *ProfileGetRequest, opts ...grpc.CallOption) (*ProfileGetResponse, error)
	// List all Profiles.
	ProfileList(ctx context.Context, in *ProfileListRequest, opts ...grpc.CallOption) (*ProfileListResponse, error)
}

type profilesClient struct {
	cc *grpc.ClientConn
}

func NewProfilesClient(cc *grpc.ClientConn) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) ProfileGet(ctx context.Context, in *ProfileGetRequest, opts ...grpc.CallOption) (*ProfileGetResponse, error) {
	out := new(ProfileGetResponse)
	err := grpc.Invoke(ctx, "/serverpb.Profiles/ProfileGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ProfileList(ctx context.Context, in *ProfileListRequest, opts ...grpc.CallOption) (*ProfileListResponse, error) {
	out := new(ProfileListResponse)
	err := grpc.Invoke(ctx, "/serverpb.Profiles/ProfileList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profiles service

type ProfilesServer interface {
	// Get a Profile by id.
	ProfileGet(context.Context, *ProfileGetRequest) (*ProfileGetResponse, error)
	// List all Profiles.
	ProfileList(context.Context, *ProfileListRequest) (*ProfileListResponse, error)
}

func RegisterProfilesServer(s *grpc.Server, srv ProfilesServer) {
	s.RegisterService(&_Profiles_serviceDesc, srv)
}

func _Profiles_ProfileGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProfileGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProfilesServer).ProfileGet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Profiles_ProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProfileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProfilesServer).ProfileList(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Profiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfileGet",
			Handler:    _Profiles_ProfileGet_Handler,
		},
		{
			MethodName: "ProfileList",
			Handler:    _Profiles_ProfileList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x4b, 0x4e, 0xf3, 0x30,
	0x10, 0xfe, 0xf3, 0x23, 0x4a, 0x3a, 0x95, 0xa0, 0x1d, 0x36, 0x10, 0x40, 0x02, 0x83, 0x50, 0x17,
	0xe0, 0x4a, 0x65, 0x87, 0xc4, 0x06, 0x04, 0x15, 0xa8, 0x0b, 0x94, 0x1b, 0xd4, 0xc1, 0x0d, 0x91,
	0x5a, 0x6c, 0x6c, 0x97, 0x9b, 0x70, 0x02, 0x2e, 0x4a, 0x71, 0xec, 0x60, 0x9a, 0xc0, 0x2a, 0xd6,
	0x7c, 0xaf, 0x79, 0x04, 0xda, 0x4a, 0x66, 0x54, 0x2a, 0x61, 0x04, 0xc6, 0x9a, 0xab, 0x37, 0xae,
	0x24, 0x4b, 0x1e, 0xf2, 0xc2, 0x3c, 0x2f, 0x18, 0xcd, 0xc4, 0x7c, 0x90, 0x09, 0xc5, 0x85, 0x76,
	0x9f, 0x73, 0x36, 0x51, 0x7c, 0xce, 0xcd, 0x64, 0x36, 0x60, 0x42, 0x98, 0x6c, 0x9a, 0x0f, 0xb4,
	0x11, 0x6a, 0x92, 0x73, 0xff, 0x95, 0xcc, 0xbf, 0x4a, 0x57, 0x72, 0x04, 0x5b, 0x23, 0x25, 0x16,
	0x72, 0xc4, 0x4d, 0xca, 0x5f, 0x17, 0x5c, 0x1b, 0xdc, 0x84, 0xff, 0xc5, 0xd3, 0x4e, 0x74, 0x18,
	0xf5, 0xdb, 0xe9, 0xf2, 0x45, 0x10, 0xba, 0x96, 0x32, 0x2e, 0xb4, 0xe7, 0x90, 0x4b, 0x57, 0xb3,
	0x32, 0x2d, 0xc5, 0x8b, 0xe6, 0x78, 0x0a, 0xeb, 0xf9, 0x57, 0xcd, 0x4a, 0x3b, 0xc3, 0x2e, 0xad,
	0x32, 0xa9, 0xe5, 0xa6, 0x25, 0x4c, 0xae, 0xa0, 0x17, 0xf8, 0x39, 0x71, 0x1f, 0x5a, 0x16, 0xd5,
	0x4b, 0xf5, 0x5a, 0xa3, 0xda, 0xe1, 0xe4, 0x18, 0x7a, 0x8f, 0x4a, 0x4c, 0x8b, 0x19, 0xff, 0xa3,
	0xe7, 0x6b, 0xc0, 0x90, 0xe4, 0x42, 0xce, 0x60, 0x43, 0x96, 0x55, 0xd7, 0x23, 0x06, 0x29, 0x8e,
	0x9f, 0x7a, 0x0a, 0x39, 0xa9, 0x3c, 0x82, 0xc9, 0x6b, 0x49, 0xb7, 0xb0, 0xfd, 0x83, 0xe5, 0xa2,
	0x28, 0xc4, 0xce, 0xc7, 0x4f, 0xd4, 0x94, 0x55, 0x71, 0x86, 0xef, 0x11, 0xb4, 0xec, 0x9c, 0x1a,
	0x6f, 0x20, 0xf6, 0xbb, 0xc5, 0x5d, 0xea, 0xaf, 0x4e, 0x57, 0xce, 0x94, 0x24, 0x4d, 0x50, 0x99,
	0x4e, 0xfe, 0xe1, 0x1d, 0xb4, 0xab, 0x25, 0xe3, 0x2a, 0x35, 0x98, 0x27, 0xd9, 0x6b, 0xc4, 0xbc,
	0xcf, 0xf0, 0x23, 0x82, 0xd8, 0x75, 0xab, 0xf1, 0x1e, 0xe0, 0x7b, 0xab, 0x18, 0x28, 0x6b, 0x07,
	0x49, 0xf6, 0x9b, 0xc1, 0xaa, 0xbf, 0x31, 0x74, 0x82, 0xb5, 0x61, 0x9d, 0x1e, 0xf6, 0x78, 0xf0,
	0x0b, 0xea, 0xdd, 0x58, 0xcb, 0xfe, 0xcc, 0x17, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xa6,
	0x0e, 0xd9, 0x2f, 0x03, 0x00, 0x00,
}