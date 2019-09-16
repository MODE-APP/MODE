// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	services.proto

It has these top-level messages:
	Message
	Beat
	Chunk
	Status
	PictureLocation
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	// Types that are valid to be assigned to ChunkContent:
	//	*Chunk_Content
	//	*Chunk_Location
	//	*Chunk_TextPost
	//	*Chunk_Checksum
	ChunkContent isChunk_ChunkContent `protobuf_oneof:"ChunkContent"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto1.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isChunk_ChunkContent interface{ isChunk_ChunkContent() }

type Chunk_Content struct {
	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3,oneof"`
}
type Chunk_Location struct {
	Location string `protobuf:"bytes,2,opt,name=Location,oneof"`
}
type Chunk_TextPost struct {
	TextPost string `protobuf:"bytes,3,opt,name=TextPost,oneof"`
}
type Chunk_Checksum struct {
	Checksum int64 `protobuf:"varint,4,opt,name=Checksum,oneof"`
}

func (*Chunk_Content) isChunk_ChunkContent()  {}
func (*Chunk_Location) isChunk_ChunkContent() {}
func (*Chunk_TextPost) isChunk_ChunkContent() {}
func (*Chunk_Checksum) isChunk_ChunkContent() {}

func (m *Chunk) GetChunkContent() isChunk_ChunkContent {
	if m != nil {
		return m.ChunkContent
	}
	return nil
}

func (m *Chunk) GetContent() []byte {
	if x, ok := m.GetChunkContent().(*Chunk_Content); ok {
		return x.Content
	}
	return nil
}

func (m *Chunk) GetLocation() string {
	if x, ok := m.GetChunkContent().(*Chunk_Location); ok {
		return x.Location
	}
	return ""
}

func (m *Chunk) GetTextPost() string {
	if x, ok := m.GetChunkContent().(*Chunk_TextPost); ok {
		return x.TextPost
	}
	return ""
}

func (m *Chunk) GetChecksum() int64 {
	if x, ok := m.GetChunkContent().(*Chunk_Checksum); ok {
		return x.Checksum
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Chunk) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _Chunk_OneofMarshaler, _Chunk_OneofUnmarshaler, _Chunk_OneofSizer, []interface{}{
		(*Chunk_Content)(nil),
		(*Chunk_Location)(nil),
		(*Chunk_TextPost)(nil),
		(*Chunk_Checksum)(nil),
	}
}

func _Chunk_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*Chunk)
	// ChunkContent
	switch x := m.ChunkContent.(type) {
	case *Chunk_Content:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		b.EncodeRawBytes(x.Content)
	case *Chunk_Location:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		b.EncodeStringBytes(x.Location)
	case *Chunk_TextPost:
		b.EncodeVarint(3<<3 | proto1.WireBytes)
		b.EncodeStringBytes(x.TextPost)
	case *Chunk_Checksum:
		b.EncodeVarint(4<<3 | proto1.WireVarint)
		b.EncodeVarint(uint64(x.Checksum))
	case nil:
	default:
		return fmt.Errorf("Chunk.ChunkContent has unexpected type %T", x)
	}
	return nil
}

func _Chunk_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*Chunk)
	switch tag {
	case 1: // ChunkContent.Content
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ChunkContent = &Chunk_Content{x}
		return true, err
	case 2: // ChunkContent.Location
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ChunkContent = &Chunk_Location{x}
		return true, err
	case 3: // ChunkContent.TextPost
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ChunkContent = &Chunk_TextPost{x}
		return true, err
	case 4: // ChunkContent.Checksum
		if wire != proto1.WireVarint {
			return true, proto1.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ChunkContent = &Chunk_Checksum{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Chunk_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*Chunk)
	// ChunkContent
	switch x := m.ChunkContent.(type) {
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
	case *Chunk_Checksum:
		n += proto1.SizeVarint(4<<3 | proto1.WireVarint)
		n += proto1.SizeVarint(uint64(x.Checksum))
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

type PictureLocation struct {
	Location string `protobuf:"bytes,1,opt,name=Location" json:"Location,omitempty"`
}

func (m *PictureLocation) Reset()                    { *m = PictureLocation{} }
func (m *PictureLocation) String() string            { return proto1.CompactTextString(m) }
func (*PictureLocation) ProtoMessage()               {}
func (*PictureLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto1.RegisterType((*PictureLocation)(nil), "proto.PictureLocation")
	proto1.RegisterEnum("proto.StatusCode", StatusCode_name, StatusCode_value)
}

func init() { proto1.RegisterFile("services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xcd, 0x6c, 0xbb, 0xad, 0xbd, 0x69, 0xb3, 0x71, 0x04, 0x09, 0xc1, 0x87, 0x12, 0x10, 0x82,
	0x60, 0x28, 0x11, 0x05, 0x5f, 0x7c, 0x68, 0xaa, 0x54, 0x50, 0x5c, 0xd2, 0xdd, 0x27, 0x9f, 0x66,
	0x93, 0x4b, 0x1b, 0xda, 0xce, 0x2c, 0x99, 0xc9, 0xae, 0x3f, 0xe1, 0x2f, 0xf8, 0x71, 0x7e, 0x89,
	0x24, 0x99, 0x89, 0xbb, 0x5d, 0x61, 0xeb, 0x53, 0xe6, 0xdc, 0x33, 0xe7, 0xdc, 0x33, 0x37, 0x17,
	0x1c, 0x89, 0xe5, 0x4d, 0x91, 0xa1, 0x8c, 0xae, 0x4b, 0xa1, 0x04, 0x3d, 0x6d, 0x3e, 0xc1, 0x77,
	0x18, 0x7e, 0x45, 0x29, 0xd9, 0x1a, 0xe9, 0x14, 0xec, 0x0b, 0xfc, 0xa1, 0x34, 0xf4, 0xc8, 0x94,
	0x84, 0xa3, 0xf4, 0x6e, 0x89, 0x46, 0x40, 0xf5, 0x71, 0x81, 0x52, 0x15, 0x9c, 0xa9, 0x42, 0x70,
	0xef, 0x64, 0xda, 0x0b, 0x47, 0xe9, 0x3f, 0x98, 0xc0, 0x87, 0xfe, 0x1c, 0x99, 0xa2, 0xb4, 0xfd,
	0x36, 0x96, 0xe3, 0xb4, 0x39, 0x07, 0x3f, 0x09, 0x9c, 0x26, 0x9b, 0x8a, 0x6f, 0xa9, 0x0f, 0xc3,
	0x44, 0x70, 0x85, 0x5c, 0x5f, 0x58, 0x5a, 0xa9, 0x29, 0xd0, 0x17, 0xf0, 0xe4, 0x8b, 0xc8, 0x4c,
	0x1f, 0x12, 0x8e, 0x96, 0x56, 0xda, 0x55, 0x6a, 0xb6, 0x8e, 0x77, 0x2e, 0xa4, 0xf2, 0x7a, 0x86,
	0x35, 0x95, 0x9a, 0x4d, 0x36, 0x98, 0x6d, 0x65, 0xb5, 0xf7, 0xfa, 0x53, 0x12, 0xf6, 0x6a, 0xd6,
	0x54, 0xe6, 0x0e, 0x8c, 0x9b, 0xf6, 0xba, 0x53, 0xf0, 0x19, 0x06, 0x2b, 0xc5, 0x54, 0x25, 0xa9,
	0xd7, 0x8d, 0x44, 0xcf, 0xa0, 0x9b, 0xd0, 0x4b, 0xe8, 0x27, 0x22, 0xc7, 0x26, 0x89, 0x13, 0x3f,
	0x6d, 0x27, 0x19, 0xb5, 0xb2, 0x9a, 0x48, 0x1b, 0x3a, 0x78, 0x0d, 0x67, 0xe7, 0x45, 0xa6, 0xaa,
	0x12, 0xbb, 0xa4, 0xfe, 0x9d, 0x77, 0xb4, 0xa6, 0x1d, 0x7e, 0x15, 0x03, 0xfc, 0xb5, 0xa0, 0x36,
	0x0c, 0x2f, 0xf9, 0x96, 0x8b, 0x5b, 0xee, 0x5a, 0x35, 0x58, 0x55, 0x59, 0x86, 0x52, 0xba, 0x84,
	0x02, 0x0c, 0x3e, 0xb1, 0x62, 0x87, 0xb9, 0x7b, 0x12, 0xff, 0x26, 0x30, 0x6a, 0x53, 0x15, 0x7c,
	0x4d, 0x23, 0xb0, 0x57, 0xc8, 0x73, 0x13, 0xd3, 0xd1, 0xc1, 0x34, 0xf6, 0x27, 0xf7, 0x82, 0x06,
	0x16, 0x7d, 0x0b, 0xb4, 0x99, 0x83, 0xbe, 0xa0, 0xdf, 0xfd, 0xa8, 0xec, 0x1d, 0x3c, 0x4b, 0x36,
	0x8c, 0xaf, 0xf1, 0xbf, 0x75, 0x6e, 0x8a, 0xaa, 0x2c, 0xf0, 0xc6, 0x28, 0x1f, 0x8a, 0x0e, 0x70,
	0x60, 0xcd, 0x48, 0xfc, 0x8b, 0x00, 0x24, 0x62, 0xbf, 0x47, 0xae, 0xf4, 0x2b, 0xeb, 0xff, 0xaa,
	0x2b, 0x8f, 0xb7, 0x9d, 0xc1, 0x64, 0x81, 0x3b, 0x54, 0x78, 0xb4, 0x22, 0x02, 0xfb, 0x63, 0x5e,
	0x1c, 0xdd, 0x21, 0xfe, 0x00, 0xe3, 0x25, 0xb2, 0x52, 0x5d, 0x21, 0xd3, 0x09, 0x9d, 0x8b, 0x92,
	0xe5, 0xd8, 0x15, 0xa9, 0xad, 0x25, 0xf5, 0xca, 0x3f, 0xd4, 0xcf, 0xc1, 0xd5, 0x8b, 0x72, 0x79,
	0xbd, 0x13, 0x2c, 0x6f, 0x3d, 0x26, 0x2d, 0xd0, 0x0c, 0x1d, 0x6b, 0x55, 0xb3, 0xad, 0x07, 0x1e,
	0x21, 0x89, 0xbf, 0x01, 0xd5, 0x37, 0x17, 0xe2, 0x96, 0x1b, 0x97, 0xf7, 0x70, 0x66, 0xa0, 0xf1,
	0x79, 0xae, 0x95, 0x07, 0xab, 0xe9, 0xdf, 0xf3, 0x9f, 0x91, 0xab, 0x41, 0x03, 0xdf, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0x20, 0x6a, 0xd9, 0x31, 0x04, 0x00, 0x00,
}
