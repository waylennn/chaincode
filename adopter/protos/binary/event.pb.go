// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/waylennn/chaincode/adopter/protos/binary/event.proto

package binary // import "github.com/waylennn/chaincode/adopter/protos/binary"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileUploadType int32

const (
	FileUploadType_FILE_UPLOAD_NEW    FileUploadType = 0
	FileUploadType_FILE_UPLOAD_UPDATE FileUploadType = 1
)

var FileUploadType_name = map[int32]string{
	0: "FILE_UPLOAD_NEW",
	1: "FILE_UPLOAD_UPDATE",
}

var FileUploadType_value = map[string]int32{
	"FILE_UPLOAD_NEW":    0,
	"FILE_UPLOAD_UPDATE": 1,
}

func (x FileUploadType) String() string {
	return proto.EnumName(FileUploadType_name, int32(x))
}

func (FileUploadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b4b773f61f03b9, []int{0}
}

type FileUploadEvent struct {
	Type                 FileUploadType `protobuf:"varint,1,opt,name=type,proto3,enum=binary.FileUploadType" json:"type,omitempty"`
	DataHash             string         `protobuf:"bytes,2,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FileUploadEvent) Reset()         { *m = FileUploadEvent{} }
func (m *FileUploadEvent) String() string { return proto.CompactTextString(m) }
func (*FileUploadEvent) ProtoMessage()    {}
func (*FileUploadEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b4b773f61f03b9, []int{0}
}
func (m *FileUploadEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileUploadEvent.Unmarshal(m, b)
}
func (m *FileUploadEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileUploadEvent.Marshal(b, m, deterministic)
}
func (dst *FileUploadEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileUploadEvent.Merge(dst, src)
}
func (m *FileUploadEvent) XXX_Size() int {
	return xxx_messageInfo_FileUploadEvent.Size(m)
}
func (m *FileUploadEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_FileUploadEvent.DiscardUnknown(m)
}

var xxx_messageInfo_FileUploadEvent proto.InternalMessageInfo

func (m *FileUploadEvent) GetType() FileUploadType {
	if m != nil {
		return m.Type
	}
	return FileUploadType_FILE_UPLOAD_NEW
}

func (m *FileUploadEvent) GetDataHash() string {
	if m != nil {
		return m.DataHash
	}
	return ""
}

func init() {
	proto.RegisterType((*FileUploadEvent)(nil), "binary.FileUploadEvent")
	proto.RegisterEnum("binary.FileUploadType", FileUploadType_name, FileUploadType_value)
}

func init() {
	proto.RegisterFile("github.com/waylennn/chaincode/adopter/protos/binary/event.proto", fileDescriptor_a0b4b773f61f03b9)
}

var fileDescriptor_a0b4b773f61f03b9 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0x2c, 0xd1,
	0x2b, 0x2c, 0x4d, 0x2d, 0xaa, 0x4c, 0x4e, 0x2c, 0xd0, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xc9,
	0x2f, 0x4d, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0x93, 0xc9, 0xf9, 0x29, 0xa9, 0xfa, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0xc5, 0xfa, 0x49, 0x99, 0x79, 0x89, 0x45, 0x95, 0xfa, 0xa9, 0x65, 0xa9,
	0x79, 0x25, 0x7a, 0x60, 0x31, 0x21, 0x36, 0x88, 0x98, 0x52, 0x14, 0x17, 0xbf, 0x5b, 0x66, 0x4e,
	0x6a, 0x68, 0x41, 0x4e, 0x7e, 0x62, 0x8a, 0x2b, 0x48, 0x81, 0x90, 0x16, 0x17, 0x4b, 0x49, 0x65,
	0x41, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x98, 0x1e, 0x44, 0xa5, 0x1e, 0x42, 0x59,
	0x48, 0x65, 0x41, 0x6a, 0x10, 0x58, 0x8d, 0x90, 0x34, 0x17, 0x67, 0x4a, 0x62, 0x49, 0x62, 0x7c,
	0x46, 0x62, 0x71, 0x86, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x07, 0x48, 0xc0, 0x23, 0xb1,
	0x38, 0x43, 0xcb, 0x96, 0x8b, 0x0f, 0x55, 0x93, 0x90, 0x30, 0x17, 0xbf, 0x9b, 0xa7, 0x8f, 0x6b,
	0x7c, 0x68, 0x80, 0x8f, 0xbf, 0xa3, 0x4b, 0xbc, 0x9f, 0x6b, 0xb8, 0x00, 0x83, 0x90, 0x18, 0x97,
	0x10, 0xb2, 0x60, 0x68, 0x80, 0x8b, 0x63, 0x88, 0xab, 0x00, 0xa3, 0x93, 0x69, 0x94, 0x31, 0x19,
	0xbe, 0x4c, 0x62, 0x03, 0x73, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x85, 0x7e, 0x45,
	0x23, 0x01, 0x00, 0x00,
}
