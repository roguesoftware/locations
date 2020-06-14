// Code generated by protoc-gen-go. DO NOT EDIT.
// source: story.proto

package proto

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

type Author struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc0550edf122997, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StoryVotes struct {
	Up                   int32    `protobuf:"varint,1,opt,name=up,proto3" json:"up,omitempty"`
	Down                 int32    `protobuf:"varint,2,opt,name=down,proto3" json:"down,omitempty"`
	Net                  int32    `protobuf:"varint,3,opt,name=net,proto3" json:"net,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoryVotes) Reset()         { *m = StoryVotes{} }
func (m *StoryVotes) String() string { return proto.CompactTextString(m) }
func (*StoryVotes) ProtoMessage()    {}
func (*StoryVotes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc0550edf122997, []int{1}
}

func (m *StoryVotes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryVotes.Unmarshal(m, b)
}
func (m *StoryVotes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryVotes.Marshal(b, m, deterministic)
}
func (m *StoryVotes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryVotes.Merge(m, src)
}
func (m *StoryVotes) XXX_Size() int {
	return xxx_messageInfo_StoryVotes.Size(m)
}
func (m *StoryVotes) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryVotes.DiscardUnknown(m)
}

var xxx_messageInfo_StoryVotes proto.InternalMessageInfo

func (m *StoryVotes) GetUp() int32 {
	if m != nil {
		return m.Up
	}
	return 0
}

func (m *StoryVotes) GetDown() int32 {
	if m != nil {
		return m.Down
	}
	return 0
}

func (m *StoryVotes) GetNet() int32 {
	if m != nil {
		return m.Net
	}
	return 0
}

type StoryItem struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationId           string      `protobuf:"bytes,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Title                string      `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Story                string      `protobuf:"bytes,6,opt,name=story,proto3" json:"story,omitempty"`
	Author               *Author     `protobuf:"bytes,8,opt,name=author,proto3" json:"author,omitempty"`
	Votes                *StoryVotes `protobuf:"bytes,9,opt,name=votes,proto3" json:"votes,omitempty"`
	Created              string      `protobuf:"bytes,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string      `protobuf:"bytes,14,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted              string      `protobuf:"bytes,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StoryItem) Reset()         { *m = StoryItem{} }
func (m *StoryItem) String() string { return proto.CompactTextString(m) }
func (*StoryItem) ProtoMessage()    {}
func (*StoryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc0550edf122997, []int{2}
}

func (m *StoryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryItem.Unmarshal(m, b)
}
func (m *StoryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryItem.Marshal(b, m, deterministic)
}
func (m *StoryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryItem.Merge(m, src)
}
func (m *StoryItem) XXX_Size() int {
	return xxx_messageInfo_StoryItem.Size(m)
}
func (m *StoryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryItem.DiscardUnknown(m)
}

var xxx_messageInfo_StoryItem proto.InternalMessageInfo

func (m *StoryItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoryItem) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func (m *StoryItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StoryItem) GetStory() string {
	if m != nil {
		return m.Story
	}
	return ""
}

func (m *StoryItem) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *StoryItem) GetVotes() *StoryVotes {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *StoryItem) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *StoryItem) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *StoryItem) GetDeleted() string {
	if m != nil {
		return m.Deleted
	}
	return ""
}

type StoryRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	LocationIds          []string `protobuf:"bytes,3,rep,name=location_ids,json=locationIds,proto3" json:"location_ids,omitempty"`
	UserIds              []string `protobuf:"bytes,5,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoryRequest) Reset()         { *m = StoryRequest{} }
func (m *StoryRequest) String() string { return proto.CompactTextString(m) }
func (*StoryRequest) ProtoMessage()    {}
func (*StoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc0550edf122997, []int{3}
}

func (m *StoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryRequest.Unmarshal(m, b)
}
func (m *StoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryRequest.Marshal(b, m, deterministic)
}
func (m *StoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryRequest.Merge(m, src)
}
func (m *StoryRequest) XXX_Size() int {
	return xxx_messageInfo_StoryRequest.Size(m)
}
func (m *StoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoryRequest proto.InternalMessageInfo

func (m *StoryRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *StoryRequest) GetLocationIds() []string {
	if m != nil {
		return m.LocationIds
	}
	return nil
}

func (m *StoryRequest) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type StoryReply struct {
	Stories              []*StoryItem `protobuf:"bytes,1,rep,name=stories,proto3" json:"stories,omitempty"`
	Count                int32        `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StoryReply) Reset()         { *m = StoryReply{} }
func (m *StoryReply) String() string { return proto.CompactTextString(m) }
func (*StoryReply) ProtoMessage()    {}
func (*StoryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc0550edf122997, []int{4}
}

func (m *StoryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoryReply.Unmarshal(m, b)
}
func (m *StoryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoryReply.Marshal(b, m, deterministic)
}
func (m *StoryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoryReply.Merge(m, src)
}
func (m *StoryReply) XXX_Size() int {
	return xxx_messageInfo_StoryReply.Size(m)
}
func (m *StoryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StoryReply.DiscardUnknown(m)
}

var xxx_messageInfo_StoryReply proto.InternalMessageInfo

func (m *StoryReply) GetStories() []*StoryItem {
	if m != nil {
		return m.Stories
	}
	return nil
}

func (m *StoryReply) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Author)(nil), "proto.Author")
	proto.RegisterType((*StoryVotes)(nil), "proto.StoryVotes")
	proto.RegisterType((*StoryItem)(nil), "proto.StoryItem")
	proto.RegisterType((*StoryRequest)(nil), "proto.StoryRequest")
	proto.RegisterType((*StoryReply)(nil), "proto.StoryReply")
}

func init() {
	proto.RegisterFile("story.proto", fileDescriptor_6fc0550edf122997)
}

var fileDescriptor_6fc0550edf122997 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x3f, 0xef, 0xd3, 0x30,
	0x10, 0x55, 0x12, 0x92, 0x34, 0x97, 0xb6, 0xb4, 0x86, 0xc1, 0xb0, 0x50, 0x22, 0x21, 0x2a, 0x84,
	0x3a, 0x14, 0x36, 0x26, 0x2a, 0x24, 0xd4, 0x85, 0xc1, 0x95, 0x18, 0x18, 0x40, 0x21, 0x3e, 0x09,
	0x4b, 0x69, 0x1c, 0x62, 0xa7, 0xa8, 0x5f, 0x80, 0xcf, 0x8d, 0x7c, 0x76, 0xd4, 0xa2, 0xdf, 0xe4,
	0x7b, 0xef, 0xdd, 0x1f, 0xbf, 0x3b, 0x28, 0x8d, 0xd5, 0xc3, 0x75, 0xd7, 0x0f, 0xda, 0x6a, 0x96,
	0xd2, 0x53, 0xbd, 0x85, 0xec, 0xe3, 0x68, 0x7f, 0xe9, 0x81, 0x2d, 0x21, 0x56, 0x92, 0x47, 0x9b,
	0x68, 0x5b, 0x88, 0x58, 0x49, 0xc6, 0xe0, 0x51, 0x57, 0x9f, 0x91, 0x27, 0xc4, 0x50, 0x5c, 0x1d,
	0x00, 0x4e, 0xae, 0xc7, 0x57, 0x6d, 0xd1, 0xb8, 0x8a, 0xb1, 0xa7, 0x8a, 0x54, 0xc4, 0x63, 0xef,
	0x2a, 0xa4, 0xfe, 0xd3, 0xf1, 0x98, 0x18, 0x8a, 0xd9, 0x0a, 0x92, 0x0e, 0x2d, 0x35, 0x49, 0x85,
	0x0b, 0xab, 0xbf, 0x31, 0x14, 0xd4, 0xe4, 0x68, 0xf1, 0xfc, 0x60, 0xea, 0x0b, 0x28, 0x5b, 0xdd,
	0xd4, 0x56, 0xe9, 0xee, 0x87, 0x92, 0x61, 0x38, 0x4c, 0xd4, 0x51, 0xb2, 0xa7, 0x90, 0x5a, 0x65,
	0x5b, 0xe4, 0x29, 0x49, 0x1e, 0x38, 0x96, 0xcc, 0xf1, 0xcc, 0xb3, 0x04, 0xd8, 0x2b, 0xc8, 0x6a,
	0x32, 0xc7, 0x67, 0x9b, 0x68, 0x5b, 0xee, 0x17, 0xde, 0xfb, 0xce, 0x3b, 0x16, 0x41, 0x64, 0xaf,
	0x21, 0xbd, 0x38, 0x43, 0xbc, 0xa0, 0xac, 0x75, 0xc8, 0xba, 0x39, 0x15, 0x5e, 0x67, 0x1c, 0xf2,
	0x66, 0xc0, 0xda, 0xa2, 0xe4, 0x0b, 0x9a, 0x33, 0x41, 0xa7, 0x8c, 0xbd, 0x24, 0x65, 0xe9, 0x95,
	0x00, 0x9d, 0x22, 0xb1, 0x45, 0xa7, 0x3c, 0xf6, 0x4a, 0x80, 0xd5, 0x77, 0x98, 0xd3, 0x08, 0x81,
	0xbf, 0x47, 0x34, 0xd6, 0xad, 0x4a, 0x49, 0xc3, 0xa3, 0x4d, 0xb2, 0x2d, 0x84, 0x0b, 0xd9, 0x4b,
	0x98, 0xdf, 0x2d, 0xc3, 0xf0, 0x84, 0xa4, 0xf2, 0xb6, 0x0d, 0xc3, 0x9e, 0xc1, 0x6c, 0x34, 0x38,
	0x90, 0x9c, 0x92, 0x9c, 0x3b, 0x7c, 0x94, 0xa6, 0xfa, 0x12, 0x8e, 0x25, 0xb0, 0x6f, 0xaf, 0xec,
	0x0d, 0xe4, 0x6e, 0x29, 0x0a, 0xfd, 0x84, 0x72, 0xbf, 0xba, 0xb7, 0xe9, 0x6e, 0x21, 0xa6, 0x04,
	0xb7, 0xcd, 0x46, 0x8f, 0xdd, 0x74, 0x36, 0x0f, 0xf6, 0x9f, 0xc2, 0x7f, 0x4f, 0x38, 0x5c, 0x54,
	0x83, 0xec, 0x3d, 0xc0, 0x67, 0xb4, 0xa7, 0x50, 0xf3, 0xe4, 0xbe, 0x5d, 0xb0, 0xf4, 0x7c, 0xfd,
	0x3f, 0xd9, 0xb7, 0xd7, 0x43, 0xf1, 0x2d, 0xdf, 0x7d, 0x20, 0xf6, 0x67, 0x46, 0xcf, 0xbb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xab, 0x2e, 0x58, 0x98, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoryServiceClient is the client API for StoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoryServiceClient interface {
	GetStories(ctx context.Context, in *StoryRequest, opts ...grpc.CallOption) (*StoryReply, error)
}

type storyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoryServiceClient(cc grpc.ClientConnInterface) StoryServiceClient {
	return &storyServiceClient{cc}
}

func (c *storyServiceClient) GetStories(ctx context.Context, in *StoryRequest, opts ...grpc.CallOption) (*StoryReply, error) {
	out := new(StoryReply)
	err := c.cc.Invoke(ctx, "/proto.StoryService/GetStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoryServiceServer is the server API for StoryService service.
type StoryServiceServer interface {
	GetStories(context.Context, *StoryRequest) (*StoryReply, error)
}

// UnimplementedStoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoryServiceServer struct {
}

func (*UnimplementedStoryServiceServer) GetStories(ctx context.Context, req *StoryRequest) (*StoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStories not implemented")
}

func RegisterStoryServiceServer(s *grpc.Server, srv StoryServiceServer) {
	s.RegisterService(&_StoryService_serviceDesc, srv)
}

func _StoryService_GetStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).GetStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StoryService/GetStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).GetStories(ctx, req.(*StoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StoryService",
	HandlerType: (*StoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStories",
			Handler:    _StoryService_GetStories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "story.proto",
}
