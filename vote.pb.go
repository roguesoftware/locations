// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vote.proto

package vote

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VoteItem struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Context              string   `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	ContextId            string   `protobuf:"bytes,6,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	Value                string   `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	Created              string   `protobuf:"bytes,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string   `protobuf:"bytes,14,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              string   `protobuf:"bytes,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteItem) Reset()         { *m = VoteItem{} }
func (m *VoteItem) String() string { return proto.CompactTextString(m) }
func (*VoteItem) ProtoMessage()    {}
func (*VoteItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d31c94b62a6ac7, []int{0}
}

func (m *VoteItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteItem.Unmarshal(m, b)
}
func (m *VoteItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteItem.Marshal(b, m, deterministic)
}
func (m *VoteItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteItem.Merge(m, src)
}
func (m *VoteItem) XXX_Size() int {
	return xxx_messageInfo_VoteItem.Size(m)
}
func (m *VoteItem) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteItem.DiscardUnknown(m)
}

var xxx_messageInfo_VoteItem proto.InternalMessageInfo

func (m *VoteItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VoteItem) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VoteItem) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *VoteItem) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *VoteItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *VoteItem) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *VoteItem) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *VoteItem) GetDeleted() string {
	if m != nil {
		return m.Deleted
	}
	return ""
}

type VoteRequest struct {
	ContextId            string   `protobuf:"bytes,1,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteRequest) Reset()         { *m = VoteRequest{} }
func (m *VoteRequest) String() string { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()    {}
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d31c94b62a6ac7, []int{1}
}

func (m *VoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteRequest.Unmarshal(m, b)
}
func (m *VoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteRequest.Marshal(b, m, deterministic)
}
func (m *VoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRequest.Merge(m, src)
}
func (m *VoteRequest) XXX_Size() int {
	return xxx_messageInfo_VoteRequest.Size(m)
}
func (m *VoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRequest proto.InternalMessageInfo

func (m *VoteRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *VoteRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type VoteReply struct {
	Votes                []*VoteItem `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	Count                int32       `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VoteReply) Reset()         { *m = VoteReply{} }
func (m *VoteReply) String() string { return proto.CompactTextString(m) }
func (*VoteReply) ProtoMessage()    {}
func (*VoteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_21d31c94b62a6ac7, []int{2}
}

func (m *VoteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteReply.Unmarshal(m, b)
}
func (m *VoteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteReply.Marshal(b, m, deterministic)
}
func (m *VoteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteReply.Merge(m, src)
}
func (m *VoteReply) XXX_Size() int {
	return xxx_messageInfo_VoteReply.Size(m)
}
func (m *VoteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteReply.DiscardUnknown(m)
}

var xxx_messageInfo_VoteReply proto.InternalMessageInfo

func (m *VoteReply) GetVotes() []*VoteItem {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *VoteReply) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*VoteItem)(nil), "proto.VoteItem")
	proto.RegisterType((*VoteRequest)(nil), "proto.VoteRequest")
	proto.RegisterType((*VoteReply)(nil), "proto.VoteReply")
}

func init() {
	proto.RegisterFile("vote.proto", fileDescriptor_21d31c94b62a6ac7)
}

var fileDescriptor_21d31c94b62a6ac7 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0x94, 0x4d, 0x93, 0x29, 0xb6, 0xb2, 0x08, 0x2e, 0x82, 0x50, 0x02, 0x42, 0x4f,
	0x41, 0xea, 0xd1, 0x83, 0x20, 0x88, 0xe6, 0x1a, 0xa1, 0x07, 0x2f, 0x52, 0xb3, 0x73, 0x08, 0xc4,
	0x6e, 0x4c, 0x76, 0x83, 0xfd, 0x97, 0xfe, 0x24, 0xd9, 0xd9, 0x2d, 0xa4, 0x39, 0x25, 0x6f, 0x3e,
	0x66, 0xf6, 0xbd, 0x07, 0x30, 0x68, 0x83, 0x79, 0xdb, 0x69, 0xa3, 0x05, 0xa7, 0x4f, 0xf6, 0xc7,
	0x20, 0xd9, 0x69, 0x83, 0x85, 0xc1, 0x6f, 0xb1, 0x84, 0xa8, 0x56, 0x92, 0xad, 0xd9, 0x26, 0x2d,
	0xa3, 0x5a, 0x89, 0x6b, 0x98, 0xdb, 0x1e, 0xbb, 0xcf, 0x5a, 0xc9, 0x19, 0x0d, 0x63, 0x27, 0x0b,
	0x25, 0x24, 0xcc, 0x2b, 0x7d, 0x30, 0xf8, 0x6b, 0x24, 0x27, 0x70, 0x92, 0xe2, 0x16, 0x20, 0xfc,
	0xba, 0xad, 0x98, 0x60, 0x1a, 0x26, 0x85, 0x12, 0x57, 0xc0, 0x87, 0x7d, 0x63, 0x51, 0x26, 0x44,
	0xbc, 0xa0, 0x73, 0x1d, 0xee, 0x0d, 0x2a, 0x79, 0x11, 0xce, 0x79, 0xe9, 0x88, 0x6d, 0x15, 0x91,
	0xa5, 0x27, 0x41, 0x3a, 0xa2, 0xb0, 0x41, 0x47, 0x56, 0x9e, 0x04, 0x99, 0xbd, 0xc0, 0xc2, 0x25,
	0x2a, 0xf1, 0xc7, 0x62, 0x3f, 0x75, 0xc4, 0xa6, 0x8e, 0x46, 0x19, 0xa3, 0x71, 0xc6, 0xec, 0x0d,
	0x52, 0x7f, 0xa6, 0x6d, 0x8e, 0xe2, 0x0e, 0xb8, 0xeb, 0xae, 0x97, 0x6c, 0x3d, 0xdb, 0x2c, 0xb6,
	0x2b, 0x5f, 0x62, 0x7e, 0x6a, 0xae, 0xf4, 0xd4, 0xc5, 0xab, 0xb4, 0x3d, 0x18, 0xaa, 0x8b, 0x97,
	0x5e, 0x6c, 0x9f, 0xbc, 0xa1, 0x77, 0xec, 0x86, 0xba, 0x42, 0x71, 0x0f, 0xc9, 0x2b, 0x9a, 0x1d,
	0x2d, 0x88, 0xd1, 0xa1, 0x60, 0xf8, 0xe6, 0xf2, 0x6c, 0xd6, 0x36, 0xc7, 0xe7, 0xe4, 0x23, 0xce,
	0x1f, 0xdd, 0x0b, 0x5f, 0x31, 0xa1, 0x87, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x90, 0xa3,
	0xca, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VoteServiceClient is the client API for VoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VoteServiceClient interface {
	GetVotes(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error)
}

type voteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteServiceClient(cc grpc.ClientConnInterface) VoteServiceClient {
	return &voteServiceClient{cc}
}

func (c *voteServiceClient) GetVotes(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error) {
	out := new(VoteReply)
	err := c.cc.Invoke(ctx, "/proto.VoteService/GetVotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServiceServer is the server API for VoteService service.
type VoteServiceServer interface {
	GetVotes(context.Context, *VoteRequest) (*VoteReply, error)
}

// UnimplementedVoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVoteServiceServer struct {
}

func (*UnimplementedVoteServiceServer) GetVotes(ctx context.Context, req *VoteRequest) (*VoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotes not implemented")
}

func RegisterVoteServiceServer(s *grpc.Server, srv VoteServiceServer) {
	s.RegisterService(&_VoteService_serviceDesc, srv)
}

func _VoteService_GetVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VoteService/GetVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetVotes(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VoteService",
	HandlerType: (*VoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVotes",
			Handler:    _VoteService_GetVotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vote.proto",
}