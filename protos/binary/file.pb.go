// Code generated by protoc-gen-go. DO NOT EDIT.
// source: git.querycap.com/cloudchain/chaincode/protos/binary/file.proto

package binary // import "git.querycap.com/cloudchain/chaincode/protos/binary"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_GENERAL FileType = 0
	FileType_EXCEL   FileType = 1
)

var FileType_name = map[int32]string{
	0: "GENERAL",
	1: "EXCEL",
}

var FileType_value = map[string]int32{
	"GENERAL": 0,
	"EXCEL":   1,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}

func (FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{0}
}

type GroupChange struct {
	Changes              [][]byte `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	Deletes              []string `protobuf:"bytes,2,rep,name=deletes,proto3" json:"deletes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupChange) Reset()         { *m = GroupChange{} }
func (m *GroupChange) String() string { return proto.CompactTextString(m) }
func (*GroupChange) ProtoMessage()    {}
func (*GroupChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{0}
}
func (m *GroupChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupChange.Unmarshal(m, b)
}
func (m *GroupChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupChange.Marshal(b, m, deterministic)
}
func (dst *GroupChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupChange.Merge(dst, src)
}
func (m *GroupChange) XXX_Size() int {
	return xxx_messageInfo_GroupChange.Size(m)
}
func (m *GroupChange) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupChange.DiscardUnknown(m)
}

var xxx_messageInfo_GroupChange proto.InternalMessageInfo

func (m *GroupChange) GetChanges() [][]byte {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *GroupChange) GetDeletes() []string {
	if m != nil {
		return m.Deletes
	}
	return nil
}

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Index                []string `protobuf:"bytes,4,rep,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{1}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *File) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *File) GetIndex() []string {
	if m != nil {
		return m.Index
	}
	return nil
}

type FileIndex struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileIndex) Reset()         { *m = FileIndex{} }
func (m *FileIndex) String() string { return proto.CompactTextString(m) }
func (*FileIndex) ProtoMessage()    {}
func (*FileIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{2}
}
func (m *FileIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileIndex.Unmarshal(m, b)
}
func (m *FileIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileIndex.Marshal(b, m, deterministic)
}
func (dst *FileIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileIndex.Merge(dst, src)
}
func (m *FileIndex) XXX_Size() int {
	return xxx_messageInfo_FileIndex.Size(m)
}
func (m *FileIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_FileIndex.DiscardUnknown(m)
}

var xxx_messageInfo_FileIndex proto.InternalMessageInfo

func (m *FileIndex) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileIndex) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type FileIndexHeader struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DataHash             string               `protobuf:"bytes,2,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FileIndexHeader) Reset()         { *m = FileIndexHeader{} }
func (m *FileIndexHeader) String() string { return proto.CompactTextString(m) }
func (*FileIndexHeader) ProtoMessage()    {}
func (*FileIndexHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{3}
}
func (m *FileIndexHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileIndexHeader.Unmarshal(m, b)
}
func (m *FileIndexHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileIndexHeader.Marshal(b, m, deterministic)
}
func (dst *FileIndexHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileIndexHeader.Merge(dst, src)
}
func (m *FileIndexHeader) XXX_Size() int {
	return xxx_messageInfo_FileIndexHeader.Size(m)
}
func (m *FileIndexHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_FileIndexHeader.DiscardUnknown(m)
}

var xxx_messageInfo_FileIndexHeader proto.InternalMessageInfo

func (m *FileIndexHeader) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *FileIndexHeader) GetDataHash() string {
	if m != nil {
		return m.DataHash
	}
	return ""
}

type FileIndexTable struct {
	Header               []byte   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Indexes              [][]byte `protobuf:"bytes,2,rep,name=indexes,proto3" json:"indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileIndexTable) Reset()         { *m = FileIndexTable{} }
func (m *FileIndexTable) String() string { return proto.CompactTextString(m) }
func (*FileIndexTable) ProtoMessage()    {}
func (*FileIndexTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{4}
}
func (m *FileIndexTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileIndexTable.Unmarshal(m, b)
}
func (m *FileIndexTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileIndexTable.Marshal(b, m, deterministic)
}
func (dst *FileIndexTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileIndexTable.Merge(dst, src)
}
func (m *FileIndexTable) XXX_Size() int {
	return xxx_messageInfo_FileIndexTable.Size(m)
}
func (m *FileIndexTable) XXX_DiscardUnknown() {
	xxx_messageInfo_FileIndexTable.DiscardUnknown(m)
}

var xxx_messageInfo_FileIndexTable proto.InternalMessageInfo

func (m *FileIndexTable) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FileIndexTable) GetIndexes() [][]byte {
	if m != nil {
		return m.Indexes
	}
	return nil
}

type FileIndexTableHistory struct {
	Data                 [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileIndexTableHistory) Reset()         { *m = FileIndexTableHistory{} }
func (m *FileIndexTableHistory) String() string { return proto.CompactTextString(m) }
func (*FileIndexTableHistory) ProtoMessage()    {}
func (*FileIndexTableHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{5}
}
func (m *FileIndexTableHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileIndexTableHistory.Unmarshal(m, b)
}
func (m *FileIndexTableHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileIndexTableHistory.Marshal(b, m, deterministic)
}
func (dst *FileIndexTableHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileIndexTableHistory.Merge(dst, src)
}
func (m *FileIndexTableHistory) XXX_Size() int {
	return xxx_messageInfo_FileIndexTableHistory.Size(m)
}
func (m *FileIndexTableHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_FileIndexTableHistory.DiscardUnknown(m)
}

var xxx_messageInfo_FileIndexTableHistory proto.InternalMessageInfo

func (m *FileIndexTableHistory) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GroupList struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Files                map[string]string    `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GroupList) Reset()         { *m = GroupList{} }
func (m *GroupList) String() string { return proto.CompactTextString(m) }
func (*GroupList) ProtoMessage()    {}
func (*GroupList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{6}
}
func (m *GroupList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupList.Unmarshal(m, b)
}
func (m *GroupList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupList.Marshal(b, m, deterministic)
}
func (dst *GroupList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupList.Merge(dst, src)
}
func (m *GroupList) XXX_Size() int {
	return xxx_messageInfo_GroupList.Size(m)
}
func (m *GroupList) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupList.DiscardUnknown(m)
}

var xxx_messageInfo_GroupList proto.InternalMessageInfo

func (m *GroupList) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *GroupList) GetFiles() map[string]string {
	if m != nil {
		return m.Files
	}
	return nil
}

type GroupFile struct {
	Files                map[string][]byte `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GroupFile) Reset()         { *m = GroupFile{} }
func (m *GroupFile) String() string { return proto.CompactTextString(m) }
func (*GroupFile) ProtoMessage()    {}
func (*GroupFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd6689efd4a6b78, []int{7}
}
func (m *GroupFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupFile.Unmarshal(m, b)
}
func (m *GroupFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupFile.Marshal(b, m, deterministic)
}
func (dst *GroupFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupFile.Merge(dst, src)
}
func (m *GroupFile) XXX_Size() int {
	return xxx_messageInfo_GroupFile.Size(m)
}
func (m *GroupFile) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupFile.DiscardUnknown(m)
}

var xxx_messageInfo_GroupFile proto.InternalMessageInfo

func (m *GroupFile) GetFiles() map[string][]byte {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupChange)(nil), "binary.GroupChange")
	proto.RegisterType((*File)(nil), "binary.File")
	proto.RegisterType((*FileIndex)(nil), "binary.FileIndex")
	proto.RegisterType((*FileIndexHeader)(nil), "binary.FileIndexHeader")
	proto.RegisterType((*FileIndexTable)(nil), "binary.FileIndexTable")
	proto.RegisterType((*FileIndexTableHistory)(nil), "binary.FileIndexTableHistory")
	proto.RegisterType((*GroupList)(nil), "binary.GroupList")
	proto.RegisterMapType((map[string]string)(nil), "binary.GroupList.FilesEntry")
	proto.RegisterType((*GroupFile)(nil), "binary.GroupFile")
	proto.RegisterMapType((map[string][]byte)(nil), "binary.GroupFile.FilesEntry")
	proto.RegisterEnum("binary.FileType", FileType_name, FileType_value)
}

func init() {
	proto.RegisterFile("git.querycap.com/cloudchain/chaincode/protos/binary/file.proto", fileDescriptor_fcd6689efd4a6b78)
}

var fileDescriptor_fcd6689efd4a6b78 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5f, 0x8b, 0xd4, 0x3e,
	0x14, 0xfd, 0x65, 0xfe, 0xed, 0xf4, 0x4e, 0xf9, 0x39, 0x04, 0x95, 0x32, 0x0a, 0x96, 0x3e, 0x15,
	0x85, 0x14, 0x66, 0x11, 0x06, 0x41, 0x61, 0x77, 0xa9, 0x3b, 0xc2, 0xb0, 0x0f, 0x61, 0x1e, 0x16,
	0x5f, 0x24, 0xd3, 0x66, 0xa7, 0xc1, 0x4e, 0x53, 0x9b, 0x54, 0xec, 0xe7, 0xf2, 0x0b, 0x4a, 0xd2,
	0x3f, 0xa3, 0xe0, 0x93, 0xfb, 0x52, 0xee, 0xb9, 0x39, 0xf7, 0x9e, 0x9b, 0x73, 0x53, 0xf8, 0x70,
	0x14, 0x9a, 0x7c, 0xab, 0x79, 0xd5, 0x24, 0xac, 0x24, 0x89, 0x3c, 0x45, 0x49, 0x2e, 0xeb, 0x34,
	0xc9, 0x98, 0x28, 0x22, 0xfb, 0x4d, 0x64, 0xca, 0xa3, 0xb2, 0x92, 0x5a, 0xaa, 0xe8, 0x20, 0x0a,
	0x56, 0x35, 0xd1, 0x83, 0xc8, 0x39, 0xb1, 0x29, 0x3c, 0x6b, 0x53, 0xab, 0x57, 0x47, 0x29, 0x8f,
	0x79, 0x47, 0x3c, 0xd4, 0x0f, 0x91, 0x16, 0x27, 0xae, 0x34, 0x3b, 0x95, 0x2d, 0x31, 0xb8, 0x82,
	0xc5, 0x6d, 0x25, 0xeb, 0xf2, 0x26, 0x63, 0xc5, 0x91, 0x63, 0x0f, 0x2e, 0x12, 0x1b, 0x29, 0x0f,
	0xf9, 0xe3, 0xd0, 0xa5, 0x3d, 0x34, 0x27, 0x29, 0xcf, 0xb9, 0xe6, 0xca, 0x1b, 0xf9, 0xe3, 0xd0,
	0xa1, 0x3d, 0x0c, 0xee, 0x61, 0xf2, 0x51, 0xe4, 0x1c, 0x63, 0x98, 0x14, 0xec, 0xc4, 0x3d, 0xe4,
	0xa3, 0xd0, 0xa1, 0x36, 0x36, 0xb9, 0x94, 0x69, 0xe6, 0x8d, 0x7c, 0x14, 0xba, 0xd4, 0xc6, 0x26,
	0xa7, 0x9b, 0x92, 0x7b, 0x63, 0x1f, 0x85, 0x53, 0x6a, 0x63, 0xfc, 0x14, 0xa6, 0xa2, 0x48, 0xf9,
	0x0f, 0x6f, 0x62, 0x7b, 0xb7, 0x20, 0x78, 0x0f, 0x8e, 0xe9, 0xfc, 0xc9, 0x00, 0xbc, 0x82, 0xb9,
	0xb9, 0xe0, 0xdd, 0x59, 0x62, 0xc0, 0xe7, 0xf2, 0x91, 0x3d, 0xe8, 0xca, 0x33, 0x78, 0x32, 0x94,
	0x6f, 0x39, 0x4b, 0x79, 0x85, 0x37, 0xe0, 0x0c, 0x0e, 0xd8, 0x2e, 0x8b, 0xf5, 0x8a, 0xb4, 0x1e,
	0x91, 0xde, 0x23, 0xb2, 0xef, 0x19, 0xf4, 0x4c, 0xc6, 0x2f, 0xc0, 0x31, 0xd3, 0x7f, 0xc9, 0x98,
	0xca, 0x3a, 0x99, 0xb9, 0x49, 0x6c, 0x99, 0xca, 0x82, 0x6b, 0xf8, 0x7f, 0x50, 0xda, 0xb3, 0x43,
	0xce, 0xf1, 0x73, 0x98, 0x65, 0x56, 0xd2, 0xaa, 0xb8, 0xb4, 0x43, 0xc6, 0x46, 0x3b, 0x5c, 0x67,
	0xa3, 0x4b, 0x7b, 0x18, 0xbc, 0x81, 0x67, 0x7f, 0xf6, 0xd8, 0x0a, 0xa5, 0x65, 0xd5, 0x0c, 0x1e,
	0xb6, 0x0b, 0xb1, 0x71, 0xf0, 0x13, 0x81, 0x63, 0xf7, 0xb6, 0x13, 0x4a, 0x3f, 0xe2, 0x56, 0x6b,
	0x98, 0x1a, 0x13, 0xdb, 0x61, 0x16, 0xeb, 0x97, 0xa4, 0x7d, 0x37, 0x64, 0xe8, 0x4d, 0xcc, 0x4c,
	0x2a, 0x2e, 0x74, 0xd5, 0xd0, 0x96, 0xba, 0xda, 0x00, 0x9c, 0x93, 0x78, 0x09, 0xe3, 0xaf, 0xbc,
	0xe9, 0x36, 0x62, 0x42, 0xb3, 0x8c, 0xef, 0x2c, 0xaf, 0x79, 0xbf, 0x0c, 0x0b, 0xde, 0x8d, 0x36,
	0x28, 0x68, 0xba, 0xa1, 0xed, 0x73, 0x19, 0xa4, 0xd1, 0x5f, 0xa4, 0x0d, 0xe3, 0xb1, 0xd2, 0xee,
	0x6f, 0xd2, 0xaf, 0x03, 0x98, 0x9b, 0xca, 0xbd, 0x79, 0x6c, 0x0b, 0xb8, 0xb8, 0x8d, 0xef, 0x62,
	0x7a, 0xb5, 0x5b, 0xfe, 0x87, 0x1d, 0x98, 0xc6, 0xf7, 0x37, 0xf1, 0x6e, 0x89, 0xae, 0xdf, 0x7e,
	0xbe, 0xfc, 0x87, 0xdf, 0xee, 0x30, 0xb3, 0xf0, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4,
	0x64, 0xcf, 0x14, 0xb4, 0x03, 0x00, 0x00,
}