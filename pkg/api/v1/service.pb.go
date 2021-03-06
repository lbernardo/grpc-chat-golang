// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v1

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

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserFrom             *User    `protobuf:"bytes,3,opt,name=userFrom,proto3" json:"userFrom,omitempty"`
	UserTo               *User    `protobuf:"bytes,4,opt,name=userTo,proto3" json:"userTo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetUserFrom() *User {
	if m != nil {
		return m.UserFrom
	}
	return nil
}

func (m *Message) GetUserTo() *User {
	if m != nil {
		return m.UserTo
	}
	return nil
}

type User struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Close struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "v1.Message")
	proto.RegisterType((*User)(nil), "v1.User")
	proto.RegisterType((*Close)(nil), "v1.Close")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0xad, 0xdb, 0xee, 0x94, 0xf5, 0x30, 0xa7, 0xd0, 0x53, 0x09, 0x0a, 0x8b, 0x87,
	0xe2, 0xae, 0x3f, 0x61, 0xc1, 0x93, 0x5e, 0xb6, 0x7a, 0xf4, 0x10, 0x9b, 0x41, 0x0a, 0x36, 0x23,
	0x93, 0xd8, 0x8b, 0x7f, 0x5e, 0x12, 0xba, 0x8b, 0x78, 0x7b, 0xdf, 0x7b, 0xe1, 0x85, 0x37, 0xb0,
	0xf1, 0x24, 0xf3, 0x38, 0x50, 0xf7, 0x25, 0x1c, 0x18, 0xf3, 0x79, 0xa7, 0x7f, 0xa0, 0x7c, 0x26,
	0xef, 0xcd, 0x07, 0xe1, 0x35, 0xe4, 0xa3, 0x55, 0x59, 0x9b, 0x6d, 0xd7, 0xc7, 0x7c, 0xb4, 0xa8,
	0xa0, 0x1c, 0xd8, 0x05, 0x72, 0x41, 0xe5, 0xc9, 0x3c, 0x21, 0xde, 0x40, 0xf5, 0xed, 0x49, 0x1e,
	0x85, 0x27, 0x75, 0xd9, 0x66, 0xdb, 0x7a, 0x5f, 0x75, 0xf3, 0xae, 0x7b, 0xf5, 0x24, 0xc7, 0x73,
	0x82, 0x2d, 0xac, 0xa2, 0x7e, 0x61, 0x55, 0xfc, 0x7b, 0xb3, 0xf8, 0xba, 0x81, 0x22, 0x32, 0x22,
	0x14, 0xce, 0x4c, 0xb4, 0x7c, 0x93, 0xb4, 0x2e, 0xe1, 0xea, 0xf0, 0xc9, 0x9e, 0xf6, 0x6f, 0x50,
	0xf5, 0xc1, 0xc8, 0x13, 0x8b, 0xc5, 0x3b, 0xd8, 0x1c, 0xd8, 0x39, 0x1a, 0x42, 0x1f, 0x84, 0xcc,
	0x84, 0xe7, 0xce, 0xa6, 0x8e, 0x6a, 0x99, 0xa2, 0x2f, 0xee, 0x33, 0xbc, 0x85, 0xba, 0x27, 0x67,
	0x4f, 0xeb, 0xfe, 0xe6, 0xcd, 0x3a, 0x42, 0xaa, 0x7f, 0x5f, 0xa5, 0x5b, 0x3c, 0xfc, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x66, 0x12, 0xf7, 0xff, 0x1c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StarLordClient is the client API for StarLord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StarLordClient interface {
	ConnectStream(ctx context.Context, in *User, opts ...grpc.CallOption) (StarLord_ConnectStreamClient, error)
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error)
}

type starLordClient struct {
	cc grpc.ClientConnInterface
}

func NewStarLordClient(cc grpc.ClientConnInterface) StarLordClient {
	return &starLordClient{cc}
}

func (c *starLordClient) ConnectStream(ctx context.Context, in *User, opts ...grpc.CallOption) (StarLord_ConnectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StarLord_serviceDesc.Streams[0], "/v1.StarLord/ConnectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &starLordConnectStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StarLord_ConnectStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type starLordConnectStreamClient struct {
	grpc.ClientStream
}

func (x *starLordConnectStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *starLordClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/v1.StarLord/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarLordServer is the server API for StarLord service.
type StarLordServer interface {
	ConnectStream(*User, StarLord_ConnectStreamServer) error
	SendMessage(context.Context, *Message) (*Close, error)
}

// UnimplementedStarLordServer can be embedded to have forward compatible implementations.
type UnimplementedStarLordServer struct {
}

func (*UnimplementedStarLordServer) ConnectStream(req *User, srv StarLord_ConnectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectStream not implemented")
}
func (*UnimplementedStarLordServer) SendMessage(ctx context.Context, req *Message) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterStarLordServer(s *grpc.Server, srv StarLordServer) {
	s.RegisterService(&_StarLord_serviceDesc, srv)
}

func _StarLord_ConnectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StarLordServer).ConnectStream(m, &starLordConnectStreamServer{stream})
}

type StarLord_ConnectStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type starLordConnectStreamServer struct {
	grpc.ServerStream
}

func (x *starLordConnectStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _StarLord_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarLordServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.StarLord/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarLordServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _StarLord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.StarLord",
	HandlerType: (*StarLordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _StarLord_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectStream",
			Handler:       _StarLord_ConnectStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
