// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/services.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	protos/services.proto

It has these top-level messages:
	Message
	Beat
	Chunk
	Status
	Checksum
	PictureLocation
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Success StatusCode = 1
	StatusCode_Failed  StatusCode = 2
)

var StatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Success",
	2: "Failed",
}
var StatusCode_value = map[string]int32{
	"Unknown": 0,
	"Success": 1,
	"Failed":  2,
}

func (x StatusCode) String() string {
	return proto1.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	TextMessage        string   `protobuf:"bytes,1,opt,name=TextMessage" json:"TextMessage,omitempty"`
	MessageDestination []string `protobuf:"bytes,2,rep,name=MessageDestination" json:"MessageDestination,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto1.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetTextMessage() string {
	if m != nil {
		return m.TextMessage
	}
	return ""
}

func (m *Message) GetMessageDestination() []string {
	if m != nil {
		return m.MessageDestination
	}
	return nil
}

type Beat struct {
	Beat []byte `protobuf:"bytes,1,opt,name=Beat,proto3" json:"Beat,omitempty"`
}

func (m *Beat) Reset()                    { *m = Beat{} }
func (m *Beat) String() string            { return proto1.CompactTextString(m) }
func (*Beat) ProtoMessage()               {}
func (*Beat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Beat) GetBeat() []byte {
	if m != nil {
		return m.Beat
	}
	return nil
}

type Chunk struct {
	// Types that are valid to be assigned to Content_:
	//	*Chunk_Content
	//	*Chunk_Location
	//	*Chunk_TextPost
	Content_ isChunk_Content_ `protobuf_oneof:"content"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto1.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isChunk_Content_ interface{ isChunk_Content_() }

type Chunk_Content struct {
	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3,oneof"`
}
type Chunk_Location struct {
	Location string `protobuf:"bytes,2,opt,name=Location,oneof"`
}
type Chunk_TextPost struct {
	TextPost string `protobuf:"bytes,3,opt,name=TextPost,oneof"`
}

func (*Chunk_Content) isChunk_Content_()  {}
func (*Chunk_Location) isChunk_Content_() {}
func (*Chunk_TextPost) isChunk_Content_() {}

func (m *Chunk) GetContent_() isChunk_Content_ {
	if m != nil {
		return m.Content_
	}
	return nil
}

func (m *Chunk) GetContent() []byte {
	if x, ok := m.GetContent_().(*Chunk_Content); ok {
		return x.Content
	}
	return nil
}

func (m *Chunk) GetLocation() string {
	if x, ok := m.GetContent_().(*Chunk_Location); ok {
		return x.Location
	}
	return ""
}

func (m *Chunk) GetTextPost() string {
	if x, ok := m.GetContent_().(*Chunk_TextPost); ok {
		return x.TextPost
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Chunk) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _Chunk_OneofMarshaler, _Chunk_OneofUnmarshaler, _Chunk_OneofSizer, []interface{}{
		(*Chunk_Content)(nil),
		(*Chunk_Location)(nil),
		(*Chunk_TextPost)(nil),
	}
}

func _Chunk_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*Chunk)
	// content
	switch x := m.Content_.(type) {
	case *Chunk_Content:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		b.EncodeRawBytes(x.Content)
	case *Chunk_Location:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		b.EncodeStringBytes(x.Location)
	case *Chunk_TextPost:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		b.EncodeStringBytes(x.TextPost)
	case nil:
	default:
		return fmt.Errorf("Chunk.Content_ has unexpected type %T", x)
	}
	return nil
}

func _Chunk_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*Chunk)
	switch tag {
	case 1: // content.Content
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Content_ = &Chunk_Content{x}
		return true, err
	case 2: // content.Location
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Content_ = &Chunk_Location{x}
		return true, err
	case 3: // content.TextPost
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Content_ = &Chunk_TextPost{x}
		return true, err
	default:
		return false, nil
	}
}

func _Chunk_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*Chunk)
	// content
	switch x := m.Content_.(type) {
	case *Chunk_Content:
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(len(x.Content)))
		n += len(x.Content)
	case *Chunk_Location:
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(len(x.Location)))
		n += len(x.Location)
	case *Chunk_TextPost:
		n += proto1.SizeVarint(3<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(len(x.TextPost)))
		n += len(x.TextPost)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Status struct {
	Message string     `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	Code    StatusCode `protobuf:"varint,2,opt,name=Code,enum=proto.StatusCode" json:"Code,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto1.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Unknown
}

type Checksum struct {
	ChecksumValue int64 `protobuf:"varint,1,opt,name=ChecksumValue" json:"ChecksumValue,omitempty"`
}

func (m *Checksum) Reset()                    { *m = Checksum{} }
func (m *Checksum) String() string            { return proto1.CompactTextString(m) }
func (*Checksum) ProtoMessage()               {}
func (*Checksum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Checksum) GetChecksumValue() int64 {
	if m != nil {
		return m.ChecksumValue
	}
	return 0
}

type PictureLocation struct {
	Location string `protobuf:"bytes,1,opt,name=Location" json:"Location,omitempty"`
}

func (m *PictureLocation) Reset()                    { *m = PictureLocation{} }
func (m *PictureLocation) String() string            { return proto1.CompactTextString(m) }
func (*PictureLocation) ProtoMessage()               {}
func (*PictureLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PictureLocation) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto1.RegisterType((*Message)(nil), "proto.Message")
	proto1.RegisterType((*Beat)(nil), "proto.Beat")
	proto1.RegisterType((*Chunk)(nil), "proto.Chunk")
	proto1.RegisterType((*Status)(nil), "proto.Status")
	proto1.RegisterType((*Checksum)(nil), "proto.Checksum")
	proto1.RegisterType((*PictureLocation)(nil), "proto.PictureLocation")
	proto1.RegisterEnum("proto.StatusCode", StatusCode_name, StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Messaging service

type MessagingClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	CheckMessageStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	ChangeMessageStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	RetrieveMessages(ctx context.Context, in *Message, opts ...grpc.CallOption) (Messaging_RetrieveMessagesClient, error)
}

type messagingClient struct {
	cc *grpc.ClientConn
}

func NewMessagingClient(cc *grpc.ClientConn) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Messaging/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) CheckMessageStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Messaging/CheckMessageStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) ChangeMessageStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Messaging/ChangeMessageStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) RetrieveMessages(ctx context.Context, in *Message, opts ...grpc.CallOption) (Messaging_RetrieveMessagesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Messaging_serviceDesc.Streams[0], c.cc, "/proto.Messaging/RetrieveMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingRetrieveMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messaging_RetrieveMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messagingRetrieveMessagesClient struct {
	grpc.ClientStream
}

func (x *messagingRetrieveMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Messaging service

type MessagingServer interface {
	SendMessage(context.Context, *Message) (*Status, error)
	CheckMessageStatus(context.Context, *Message) (*Status, error)
	ChangeMessageStatus(context.Context, *Message) (*Status, error)
	RetrieveMessages(*Message, Messaging_RetrieveMessagesServer) error
}

func RegisterMessagingServer(s *grpc.Server, srv MessagingServer) {
	s.RegisterService(&_Messaging_serviceDesc, srv)
}

func _Messaging_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Messaging/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_CheckMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).CheckMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Messaging/CheckMessageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).CheckMessageStatus(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_ChangeMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).ChangeMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Messaging/ChangeMessageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).ChangeMessageStatus(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_RetrieveMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServer).RetrieveMessages(m, &messagingRetrieveMessagesServer{stream})
}

type Messaging_RetrieveMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messagingRetrieveMessagesServer struct {
	grpc.ServerStream
}

func (x *messagingRetrieveMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Messaging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Messaging_SendMessage_Handler,
		},
		{
			MethodName: "CheckMessageStatus",
			Handler:    _Messaging_CheckMessageStatus_Handler,
		},
		{
			MethodName: "ChangeMessageStatus",
			Handler:    _Messaging_ChangeMessageStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RetrieveMessages",
			Handler:       _Messaging_RetrieveMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/services.proto",
}

// Client API for Commenting service

type CommentingClient interface {
	PostComment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	DeleteComment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	EditComment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
}

type commentingClient struct {
	cc *grpc.ClientConn
}

func NewCommentingClient(cc *grpc.ClientConn) CommentingClient {
	return &commentingClient{cc}
}

func (c *commentingClient) PostComment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Commenting/PostComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentingClient) DeleteComment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Commenting/DeleteComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentingClient) EditComment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Commenting/EditComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Commenting service

type CommentingServer interface {
	PostComment(context.Context, *Message) (*Status, error)
	DeleteComment(context.Context, *Message) (*Status, error)
	EditComment(context.Context, *Message) (*Status, error)
}

func RegisterCommentingServer(s *grpc.Server, srv CommentingServer) {
	s.RegisterService(&_Commenting_serviceDesc, srv)
}

func _Commenting_PostComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentingServer).PostComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commenting/PostComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentingServer).PostComment(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commenting_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentingServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commenting/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentingServer).DeleteComment(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commenting_EditComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentingServer).EditComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Commenting/EditComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentingServer).EditComment(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Commenting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Commenting",
	HandlerType: (*CommentingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostComment",
			Handler:    _Commenting_PostComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Commenting_DeleteComment_Handler,
		},
		{
			MethodName: "EditComment",
			Handler:    _Commenting_EditComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/services.proto",
}

// Client API for Heartbeating service

type HeartbeatingClient interface {
	TradeHeartbeat(ctx context.Context, in *Beat, opts ...grpc.CallOption) (*Status, error)
}

type heartbeatingClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatingClient(cc *grpc.ClientConn) HeartbeatingClient {
	return &heartbeatingClient{cc}
}

func (c *heartbeatingClient) TradeHeartbeat(ctx context.Context, in *Beat, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.Heartbeating/TradeHeartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Heartbeating service

type HeartbeatingServer interface {
	TradeHeartbeat(context.Context, *Beat) (*Status, error)
}

func RegisterHeartbeatingServer(s *grpc.Server, srv HeartbeatingServer) {
	s.RegisterService(&_Heartbeating_serviceDesc, srv)
}

func _Heartbeating_TradeHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Beat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatingServer).TradeHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Heartbeating/TradeHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatingServer).TradeHeartbeat(ctx, req.(*Beat))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeating_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Heartbeating",
	HandlerType: (*HeartbeatingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TradeHeartbeat",
			Handler:    _Heartbeating_TradeHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/services.proto",
}

// Client API for PictureUploading service

type PictureUploadingClient interface {
	UploadPicture(ctx context.Context, opts ...grpc.CallOption) (PictureUploading_UploadPictureClient, error)
	SendChecksum(ctx context.Context, in *Checksum, opts ...grpc.CallOption) (*Status, error)
}

type pictureUploadingClient struct {
	cc *grpc.ClientConn
}

func NewPictureUploadingClient(cc *grpc.ClientConn) PictureUploadingClient {
	return &pictureUploadingClient{cc}
}

func (c *pictureUploadingClient) UploadPicture(ctx context.Context, opts ...grpc.CallOption) (PictureUploading_UploadPictureClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PictureUploading_serviceDesc.Streams[0], c.cc, "/proto.PictureUploading/UploadPicture", opts...)
	if err != nil {
		return nil, err
	}
	x := &pictureUploadingUploadPictureClient{stream}
	return x, nil
}

type PictureUploading_UploadPictureClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type pictureUploadingUploadPictureClient struct {
	grpc.ClientStream
}

func (x *pictureUploadingUploadPictureClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pictureUploadingUploadPictureClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pictureUploadingClient) SendChecksum(ctx context.Context, in *Checksum, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.PictureUploading/SendChecksum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PictureUploading service

type PictureUploadingServer interface {
	UploadPicture(PictureUploading_UploadPictureServer) error
	SendChecksum(context.Context, *Checksum) (*Status, error)
}

func RegisterPictureUploadingServer(s *grpc.Server, srv PictureUploadingServer) {
	s.RegisterService(&_PictureUploading_serviceDesc, srv)
}

func _PictureUploading_UploadPicture_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PictureUploadingServer).UploadPicture(&pictureUploadingUploadPictureServer{stream})
}

type PictureUploading_UploadPictureServer interface {
	SendAndClose(*Status) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type pictureUploadingUploadPictureServer struct {
	grpc.ServerStream
}

func (x *pictureUploadingUploadPictureServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pictureUploadingUploadPictureServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PictureUploading_SendChecksum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Checksum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PictureUploadingServer).SendChecksum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PictureUploading/SendChecksum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PictureUploadingServer).SendChecksum(ctx, req.(*Checksum))
	}
	return interceptor(ctx, in, info, handler)
}

var _PictureUploading_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PictureUploading",
	HandlerType: (*PictureUploadingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChecksum",
			Handler:    _PictureUploading_SendChecksum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadPicture",
			Handler:       _PictureUploading_UploadPicture_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/services.proto",
}

// Client API for PictureDownloading service

type PictureDownloadingClient interface {
	DownloadPicture(ctx context.Context, in *PictureLocation, opts ...grpc.CallOption) (PictureDownloading_DownloadPictureClient, error)
}

type pictureDownloadingClient struct {
	cc *grpc.ClientConn
}

func NewPictureDownloadingClient(cc *grpc.ClientConn) PictureDownloadingClient {
	return &pictureDownloadingClient{cc}
}

func (c *pictureDownloadingClient) DownloadPicture(ctx context.Context, in *PictureLocation, opts ...grpc.CallOption) (PictureDownloading_DownloadPictureClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PictureDownloading_serviceDesc.Streams[0], c.cc, "/proto.PictureDownloading/DownloadPicture", opts...)
	if err != nil {
		return nil, err
	}
	x := &pictureDownloadingDownloadPictureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PictureDownloading_DownloadPictureClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type pictureDownloadingDownloadPictureClient struct {
	grpc.ClientStream
}

func (x *pictureDownloadingDownloadPictureClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PictureDownloading service

type PictureDownloadingServer interface {
	DownloadPicture(*PictureLocation, PictureDownloading_DownloadPictureServer) error
}

func RegisterPictureDownloadingServer(s *grpc.Server, srv PictureDownloadingServer) {
	s.RegisterService(&_PictureDownloading_serviceDesc, srv)
}

func _PictureDownloading_DownloadPicture_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PictureLocation)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PictureDownloadingServer).DownloadPicture(m, &pictureDownloadingDownloadPictureServer{stream})
}

type PictureDownloading_DownloadPictureServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type pictureDownloadingDownloadPictureServer struct {
	grpc.ServerStream
}

func (x *pictureDownloadingDownloadPictureServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _PictureDownloading_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PictureDownloading",
	HandlerType: (*PictureDownloadingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadPicture",
			Handler:       _PictureDownloading_DownloadPicture_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/services.proto",
}

func init() { proto1.RegisterFile("protos/services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x9b, 0x36, 0x69, 0x26, 0x5f, 0x66, 0x10, 0x28, 0xb2, 0x38, 0x44, 0x16, 0x48, 0x11,
	0x12, 0x26, 0x32, 0x02, 0x89, 0x0b, 0x87, 0x3a, 0xa0, 0x20, 0x81, 0xa8, 0x9c, 0x96, 0x0b, 0xa7,
	0xad, 0x3d, 0x4a, 0x4c, 0x92, 0xdd, 0xca, 0xbb, 0x6e, 0xf9, 0x35, 0xfc, 0x38, 0x7e, 0x09, 0xb2,
	0x77, 0xd7, 0x34, 0x29, 0x52, 0xdb, 0x93, 0xf7, 0xbd, 0xf1, 0x9b, 0x37, 0x6f, 0x3f, 0xe0, 0xc9,
	0x65, 0x2e, 0x94, 0x90, 0xaf, 0x25, 0xe5, 0x57, 0x59, 0x42, 0x32, 0xa8, 0x30, 0x1e, 0x55, 0x1f,
	0xff, 0x07, 0xb4, 0xbf, 0x92, 0x94, 0x6c, 0x49, 0x38, 0x86, 0xee, 0x19, 0xfd, 0x52, 0x06, 0x8e,
	0x9c, 0xb1, 0x33, 0xe9, 0xc4, 0x37, 0x29, 0x0c, 0x00, 0xcd, 0x72, 0x46, 0x52, 0x65, 0x9c, 0xa9,
	0x4c, 0xf0, 0xd1, 0xc1, 0xb8, 0x39, 0xe9, 0xc4, 0xff, 0xa9, 0xf8, 0x1e, 0x1c, 0x9e, 0x10, 0x53,
	0x88, 0xfa, 0x5b, 0xb5, 0xec, 0xc5, 0xd5, 0xda, 0xff, 0x09, 0x47, 0xd1, 0xaa, 0xe0, 0x6b, 0xf4,
	0xa0, 0x1d, 0x09, 0xae, 0x88, 0x9b, 0xfa, 0xbc, 0x11, 0x5b, 0x02, 0x9f, 0xc1, 0xf1, 0x17, 0x91,
	0x58, 0x1b, 0x67, 0xd2, 0x99, 0x37, 0xe2, 0x9a, 0x29, 0xab, 0xe5, 0x74, 0xa7, 0x42, 0xaa, 0x51,
	0xd3, 0x56, 0x2d, 0x73, 0xd2, 0x81, 0x76, 0xa2, 0xdb, 0xf8, 0x9f, 0xa1, 0xb5, 0x50, 0x4c, 0x15,
	0x12, 0x47, 0x75, 0x5c, 0x93, 0xaf, 0x4e, 0xff, 0x02, 0x0e, 0x23, 0x91, 0x52, 0x65, 0x33, 0x08,
	0x1f, 0xe9, 0x5d, 0x0a, 0xb4, 0xac, 0x2c, 0xc4, 0x55, 0xd9, 0x9f, 0xc2, 0x71, 0xb4, 0xa2, 0x64,
	0x2d, 0x8b, 0x2d, 0x3e, 0x87, 0xbe, 0x5d, 0x7f, 0x67, 0x9b, 0x42, 0xb7, 0x6c, 0xc6, 0xbb, 0xa4,
	0xff, 0x0a, 0x86, 0xa7, 0x59, 0xa2, 0x8a, 0x9c, 0xea, 0xc1, 0xbd, 0x1b, 0xb1, 0xf4, 0x18, 0x35,
	0x7e, 0x19, 0x02, 0xfc, 0x33, 0xc5, 0x2e, 0xb4, 0xcf, 0xf9, 0x9a, 0x8b, 0x6b, 0xee, 0x36, 0x4a,
	0xb0, 0x28, 0x92, 0x84, 0xa4, 0x74, 0x1d, 0x04, 0x68, 0x7d, 0x62, 0xd9, 0x86, 0x52, 0xf7, 0x20,
	0xfc, 0xe3, 0x40, 0x47, 0xe7, 0xc8, 0xf8, 0x12, 0x03, 0xe8, 0x2e, 0x88, 0xa7, 0x36, 0xd8, 0xc0,
	0x44, 0x31, 0xd8, 0xeb, 0xef, 0x44, 0xf3, 0x1b, 0xf8, 0x16, 0xb0, 0x9a, 0xd8, 0xfc, 0x60, 0x76,
	0xea, 0x4e, 0xd9, 0x3b, 0x78, 0x1c, 0xad, 0x18, 0x5f, 0xd2, 0x83, 0x75, 0x6e, 0x4c, 0x2a, 0xcf,
	0xe8, 0xca, 0x2a, 0x6f, 0x8b, 0xf6, 0xb0, 0xdf, 0x98, 0x3a, 0xe1, 0x6f, 0x07, 0x20, 0x12, 0xdb,
	0x2d, 0x71, 0x65, 0x52, 0x96, 0xc7, 0x6c, 0x98, 0xbb, 0x6d, 0xa7, 0xd0, 0x9f, 0xd1, 0x86, 0x14,
	0xdd, 0x5b, 0x11, 0x40, 0xf7, 0x63, 0x9a, 0xdd, 0xdb, 0x21, 0xfc, 0x00, 0xbd, 0x39, 0xb1, 0x5c,
	0x5d, 0x10, 0x33, 0x13, 0x0e, 0xce, 0x72, 0x96, 0x52, 0x4d, 0x62, 0xd7, 0x48, 0xca, 0x07, 0x70,
	0x5b, 0x9f, 0x83, 0x6b, 0x2e, 0xca, 0xf9, 0xe5, 0x46, 0xb0, 0x54, 0xf7, 0xe8, 0x6b, 0x60, 0x2a,
	0xd8, 0x33, 0xaa, 0xea, 0xed, 0xec, 0xf5, 0x98, 0x38, 0x18, 0x40, 0xaf, 0x3c, 0xfb, 0xfa, 0x8a,
	0x0e, 0xeb, 0xdf, 0x35, 0xb1, 0xa7, 0x08, 0xbf, 0x01, 0x9a, 0xce, 0x33, 0x71, 0xcd, 0xad, 0xeb,
	0x7b, 0x18, 0x5a, 0x68, 0x7d, 0x9f, 0x1a, 0xdd, 0xde, 0x55, 0xf6, 0x76, 0xe6, 0x99, 0x3a, 0x17,
	0xad, 0x0a, 0xbe, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xc0, 0x62, 0x63, 0x76, 0x04, 0x00,
	0x00,
}