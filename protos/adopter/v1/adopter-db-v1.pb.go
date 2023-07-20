// Code generated by protoc-gen-go. DO NOT EDIT.
// source: git.querycap.com/cloudchain/chaincode/protos/adopter/adopter-db-v1.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DBPayloadV1 struct {
	Payload              []*DatabaseV1 `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DBPayloadV1) Reset()         { *m = DBPayloadV1{} }
func (m *DBPayloadV1) String() string { return proto.CompactTextString(m) }
func (*DBPayloadV1) ProtoMessage()    {}
func (*DBPayloadV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{0}
}

func (m *DBPayloadV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBPayloadV1.Unmarshal(m, b)
}
func (m *DBPayloadV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBPayloadV1.Marshal(b, m, deterministic)
}
func (m *DBPayloadV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBPayloadV1.Merge(m, src)
}
func (m *DBPayloadV1) XXX_Size() int {
	return xxx_messageInfo_DBPayloadV1.Size(m)
}
func (m *DBPayloadV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DBPayloadV1.DiscardUnknown(m)
}

var xxx_messageInfo_DBPayloadV1 proto.InternalMessageInfo

func (m *DBPayloadV1) GetPayload() []*DatabaseV1 {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DatabaseV1 struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               int64    `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Hash                 string   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseV1) Reset()         { *m = DatabaseV1{} }
func (m *DatabaseV1) String() string { return proto.CompactTextString(m) }
func (*DatabaseV1) ProtoMessage()    {}
func (*DatabaseV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{1}
}

func (m *DatabaseV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseV1.Unmarshal(m, b)
}
func (m *DatabaseV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseV1.Marshal(b, m, deterministic)
}
func (m *DatabaseV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseV1.Merge(m, src)
}
func (m *DatabaseV1) XXX_Size() int {
	return xxx_messageInfo_DatabaseV1.Size(m)
}
func (m *DatabaseV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseV1.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseV1 proto.InternalMessageInfo

func (m *DatabaseV1) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *DatabaseV1) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DatabaseV1) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *DatabaseV1) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type DBStatV1 struct {
	Logs                 []*DatabaseStatV1 `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DBStatV1) Reset()         { *m = DBStatV1{} }
func (m *DBStatV1) String() string { return proto.CompactTextString(m) }
func (*DBStatV1) ProtoMessage()    {}
func (*DBStatV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{2}
}

func (m *DBStatV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBStatV1.Unmarshal(m, b)
}
func (m *DBStatV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBStatV1.Marshal(b, m, deterministic)
}
func (m *DBStatV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBStatV1.Merge(m, src)
}
func (m *DBStatV1) XXX_Size() int {
	return xxx_messageInfo_DBStatV1.Size(m)
}
func (m *DBStatV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DBStatV1.DiscardUnknown(m)
}

var xxx_messageInfo_DBStatV1 proto.InternalMessageInfo

func (m *DBStatV1) GetLogs() []*DatabaseStatV1 {
	if m != nil {
		return m.Logs
	}
	return nil
}

type DatabaseStatV1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsDdl                bool     `protobuf:"varint,2,opt,name=is_ddl,json=isDdl,proto3" json:"is_ddl,omitempty"`
	Ddl                  *DDLV1   `protobuf:"bytes,3,opt,name=ddl,proto3" json:"ddl,omitempty"`
	Dml                  *DMLV1   `protobuf:"bytes,4,opt,name=dml,proto3" json:"dml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseStatV1) Reset()         { *m = DatabaseStatV1{} }
func (m *DatabaseStatV1) String() string { return proto.CompactTextString(m) }
func (*DatabaseStatV1) ProtoMessage()    {}
func (*DatabaseStatV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{3}
}

func (m *DatabaseStatV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseStatV1.Unmarshal(m, b)
}
func (m *DatabaseStatV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseStatV1.Marshal(b, m, deterministic)
}
func (m *DatabaseStatV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseStatV1.Merge(m, src)
}
func (m *DatabaseStatV1) XXX_Size() int {
	return xxx_messageInfo_DatabaseStatV1.Size(m)
}
func (m *DatabaseStatV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseStatV1.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseStatV1 proto.InternalMessageInfo

func (m *DatabaseStatV1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabaseStatV1) GetIsDdl() bool {
	if m != nil {
		return m.IsDdl
	}
	return false
}

func (m *DatabaseStatV1) GetDdl() *DDLV1 {
	if m != nil {
		return m.Ddl
	}
	return nil
}

func (m *DatabaseStatV1) GetDml() *DMLV1 {
	if m != nil {
		return m.Dml
	}
	return nil
}

type TableStatV1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Adds                 uint64   `protobuf:"varint,2,opt,name=adds,proto3" json:"adds,omitempty"`
	Updates              uint64   `protobuf:"varint,3,opt,name=updates,proto3" json:"updates,omitempty"`
	Deletes              uint64   `protobuf:"varint,4,opt,name=deletes,proto3" json:"deletes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableStatV1) Reset()         { *m = TableStatV1{} }
func (m *TableStatV1) String() string { return proto.CompactTextString(m) }
func (*TableStatV1) ProtoMessage()    {}
func (*TableStatV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{4}
}

func (m *TableStatV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableStatV1.Unmarshal(m, b)
}
func (m *TableStatV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableStatV1.Marshal(b, m, deterministic)
}
func (m *TableStatV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableStatV1.Merge(m, src)
}
func (m *TableStatV1) XXX_Size() int {
	return xxx_messageInfo_TableStatV1.Size(m)
}
func (m *TableStatV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TableStatV1.DiscardUnknown(m)
}

var xxx_messageInfo_TableStatV1 proto.InternalMessageInfo

func (m *TableStatV1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableStatV1) GetAdds() uint64 {
	if m != nil {
		return m.Adds
	}
	return 0
}

func (m *TableStatV1) GetUpdates() uint64 {
	if m != nil {
		return m.Updates
	}
	return 0
}

func (m *TableStatV1) GetDeletes() uint64 {
	if m != nil {
		return m.Deletes
	}
	return 0
}

type DMLV1 struct {
	TableStats           []*TableStatV1 `protobuf:"bytes,1,rep,name=table_stats,json=tableStats,proto3" json:"table_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DMLV1) Reset()         { *m = DMLV1{} }
func (m *DMLV1) String() string { return proto.CompactTextString(m) }
func (*DMLV1) ProtoMessage()    {}
func (*DMLV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{5}
}

func (m *DMLV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DMLV1.Unmarshal(m, b)
}
func (m *DMLV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DMLV1.Marshal(b, m, deterministic)
}
func (m *DMLV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DMLV1.Merge(m, src)
}
func (m *DMLV1) XXX_Size() int {
	return xxx_messageInfo_DMLV1.Size(m)
}
func (m *DMLV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DMLV1.DiscardUnknown(m)
}

var xxx_messageInfo_DMLV1 proto.InternalMessageInfo

func (m *DMLV1) GetTableStats() []*TableStatV1 {
	if m != nil {
		return m.TableStats
	}
	return nil
}

type DDLV1 struct {
	TableName            string   `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Sql                  string   `protobuf:"bytes,2,opt,name=sql,proto3" json:"sql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DDLV1) Reset()         { *m = DDLV1{} }
func (m *DDLV1) String() string { return proto.CompactTextString(m) }
func (*DDLV1) ProtoMessage()    {}
func (*DDLV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e64c4ff2ad8f8f70, []int{6}
}

func (m *DDLV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DDLV1.Unmarshal(m, b)
}
func (m *DDLV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DDLV1.Marshal(b, m, deterministic)
}
func (m *DDLV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DDLV1.Merge(m, src)
}
func (m *DDLV1) XXX_Size() int {
	return xxx_messageInfo_DDLV1.Size(m)
}
func (m *DDLV1) XXX_DiscardUnknown() {
	xxx_messageInfo_DDLV1.DiscardUnknown(m)
}

var xxx_messageInfo_DDLV1 proto.InternalMessageInfo

func (m *DDLV1) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *DDLV1) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func init() {
	proto.RegisterType((*DBPayloadV1)(nil), "adopter.DBPayloadV1")
	proto.RegisterType((*DatabaseV1)(nil), "adopter.DatabaseV1")
	proto.RegisterType((*DBStatV1)(nil), "adopter.DBStatV1")
	proto.RegisterType((*DatabaseStatV1)(nil), "adopter.DatabaseStatV1")
	proto.RegisterType((*TableStatV1)(nil), "adopter.TableStatV1")
	proto.RegisterType((*DMLV1)(nil), "adopter.DMLV1")
	proto.RegisterType((*DDLV1)(nil), "adopter.DDLV1")
}

func init() {
	proto.RegisterFile("git.querycap.com/cloudchain/chaincode/protos/adopter/adopter-db-v1.proto", fileDescriptor_e64c4ff2ad8f8f70)
}

var fileDescriptor_e64c4ff2ad8f8f70 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0xab, 0x9c, 0x30,
	0x14, 0xc6, 0xb1, 0x3a, 0xaf, 0x23, 0x5c, 0x4a, 0xfa, 0x92, 0x42, 0x41, 0x5c, 0x0d, 0x94, 0x51,
	0xbc, 0x7d, 0x2e, 0x4a, 0x17, 0x83, 0x8b, 0x2e, 0xda, 0x52, 0xd2, 0xe2, 0xa2, 0x9b, 0x21, 0x9a,
	0xcc, 0x28, 0xc4, 0x89, 0x77, 0x92, 0x29, 0xdc, 0x55, 0xff, 0xf5, 0x92, 0x13, 0x9d, 0x07, 0x85,
	0x2e, 0xba, 0xd1, 0xf3, 0x7d, 0xe7, 0x23, 0xdf, 0xcf, 0x20, 0x7c, 0xda, 0xb5, 0x26, 0xbd, 0x3b,
	0x8a, 0xc3, 0x7d, 0xcd, 0xfa, 0xb4, 0x56, 0x5d, 0x56, 0x4b, 0x75, 0xe4, 0x75, 0xc3, 0xda, 0x7d,
	0x86, 0xcf, 0x5a, 0x71, 0x91, 0xf5, 0x07, 0x65, 0x94, 0xce, 0x18, 0x57, 0xbd, 0x11, 0x87, 0xf1,
	0xbd, 0xe2, 0xd5, 0xea, 0x57, 0x9e, 0xe2, 0x92, 0xcc, 0x06, 0x33, 0xf9, 0x00, 0x61, 0xb1, 0xfe,
	0xc6, 0xee, 0xa5, 0x62, 0xbc, 0xcc, 0xc9, 0x0a, 0x66, 0xbd, 0x13, 0x91, 0x17, 0xfb, 0xcb, 0xf0,
	0xf6, 0x51, 0x3a, 0x24, 0xd3, 0x82, 0x19, 0x56, 0x31, 0x2d, 0xca, 0x9c, 0x8e, 0x99, 0x44, 0x02,
	0x9c, 0x6d, 0xf2, 0x1c, 0xe6, 0xdb, 0x56, 0x8a, 0x3d, 0xeb, 0x44, 0xe4, 0xc5, 0xde, 0x72, 0x41,
	0x4f, 0x9a, 0x3c, 0x85, 0xa9, 0xda, 0x6e, 0xb5, 0x30, 0xd1, 0x83, 0xd8, 0x5b, 0xfa, 0x74, 0x50,
	0xd6, 0x97, 0x62, 0xbf, 0x33, 0x4d, 0xe4, 0x3b, 0xdf, 0x29, 0x42, 0x20, 0x68, 0x98, 0x6e, 0xa2,
	0x00, 0xcf, 0xc1, 0x39, 0x79, 0x07, 0xf3, 0x62, 0xfd, 0xdd, 0x30, 0x53, 0xe6, 0xe4, 0x25, 0x04,
	0x52, 0xed, 0xf4, 0x40, 0xf9, 0xec, 0x2f, 0x4a, 0x17, 0xa3, 0x18, 0x4a, 0x7e, 0xc3, 0xcd, 0xb5,
	0x6f, 0x8f, 0xbf, 0xc0, 0xc4, 0x99, 0x3c, 0x81, 0x69, 0xab, 0x37, 0x9c, 0x4b, 0x44, 0x9c, 0xd3,
	0x49, 0xab, 0x0b, 0x2e, 0x49, 0x0c, 0xbe, 0xf5, 0x2c, 0x5e, 0x78, 0x7b, 0x73, 0x2e, 0x2a, 0x3e,
	0x97, 0x39, 0xb5, 0x2b, 0x4c, 0x74, 0x12, 0x51, 0xaf, 0x12, 0x5f, 0x5c, 0xa2, 0x93, 0x49, 0x0b,
	0xe1, 0x0f, 0x56, 0xc9, 0x7f, 0xb5, 0x13, 0x08, 0x18, 0xe7, 0x1a, 0xbb, 0x03, 0x8a, 0x33, 0x89,
	0x60, 0x76, 0xec, 0x39, 0x33, 0x42, 0x63, 0x7d, 0x40, 0x47, 0x69, 0x37, 0x5c, 0x48, 0x61, 0x37,
	0x81, 0xdb, 0x0c, 0x32, 0xf9, 0x08, 0x13, 0x2c, 0x26, 0x6f, 0x20, 0x34, 0xb6, 0x73, 0xa3, 0x0d,
	0x33, 0xe3, 0x45, 0x3d, 0x3e, 0xd1, 0x5d, 0xf0, 0x50, 0x30, 0xa3, 0xd0, 0xc9, 0x7b, 0x98, 0xe0,
	0xa7, 0x91, 0x17, 0xe0, 0xec, 0xcd, 0x05, 0xea, 0x02, 0x9d, 0xaf, 0x96, 0xf7, 0x21, 0xf8, 0xfa,
	0xce, 0x5d, 0xd5, 0x82, 0xda, 0x71, 0xfd, 0xf6, 0xe7, 0xeb, 0xff, 0xf9, 0x3f, 0xab, 0x29, 0xea,
	0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xe3, 0x5e, 0x06, 0xde, 0x02, 0x00, 0x00,
}