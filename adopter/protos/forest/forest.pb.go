// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forest.proto

package forest

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

type PutStateRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutStateRequest) Reset()         { *m = PutStateRequest{} }
func (m *PutStateRequest) String() string { return proto.CompactTextString(m) }
func (*PutStateRequest) ProtoMessage()    {}
func (*PutStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d1b8996afffe6d, []int{0}
}

func (m *PutStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutStateRequest.Unmarshal(m, b)
}
func (m *PutStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutStateRequest.Marshal(b, m, deterministic)
}
func (m *PutStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutStateRequest.Merge(m, src)
}
func (m *PutStateRequest) XXX_Size() int {
	return xxx_messageInfo_PutStateRequest.Size(m)
}
func (m *PutStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutStateRequest proto.InternalMessageInfo

func (m *PutStateRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *PutStateRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetStateRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateRequest) Reset()         { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()    {}
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d1b8996afffe6d, []int{1}
}

func (m *GetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateRequest.Unmarshal(m, b)
}
func (m *GetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateRequest.Marshal(b, m, deterministic)
}
func (m *GetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateRequest.Merge(m, src)
}
func (m *GetStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetStateRequest.Size(m)
}
func (m *GetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateRequest proto.InternalMessageInfo

func (m *GetStateRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetHistoryRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryRequest) Reset()         { *m = GetHistoryRequest{} }
func (m *GetHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetHistoryRequest) ProtoMessage()    {}
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d1b8996afffe6d, []int{2}
}

func (m *GetHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryRequest.Unmarshal(m, b)
}
func (m *GetHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryRequest.Merge(m, src)
}
func (m *GetHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetHistoryRequest.Size(m)
}
func (m *GetHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryRequest proto.InternalMessageInfo

func (m *GetHistoryRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetHistoryResponse struct {
	KeyModifications     []*KeyModification `protobuf:"bytes,1,rep,name=key_modifications,json=keyModifications,proto3" json:"key_modifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetHistoryResponse) Reset()         { *m = GetHistoryResponse{} }
func (m *GetHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetHistoryResponse) ProtoMessage()    {}
func (*GetHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d1b8996afffe6d, []int{3}
}

func (m *GetHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryResponse.Unmarshal(m, b)
}
func (m *GetHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryResponse.Merge(m, src)
}
func (m *GetHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetHistoryResponse.Size(m)
}
func (m *GetHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryResponse proto.InternalMessageInfo

func (m *GetHistoryResponse) GetKeyModifications() []*KeyModification {
	if m != nil {
		return m.KeyModifications
	}
	return nil
}

type KeyModification struct {
	TxId                 string   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsDelete             bool     `protobuf:"varint,4,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyModification) Reset()         { *m = KeyModification{} }
func (m *KeyModification) String() string { return proto.CompactTextString(m) }
func (*KeyModification) ProtoMessage()    {}
func (*KeyModification) Descriptor() ([]byte, []int) {
	return fileDescriptor_47d1b8996afffe6d, []int{4}
}

func (m *KeyModification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyModification.Unmarshal(m, b)
}
func (m *KeyModification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyModification.Marshal(b, m, deterministic)
}
func (m *KeyModification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyModification.Merge(m, src)
}
func (m *KeyModification) XXX_Size() int {
	return xxx_messageInfo_KeyModification.Size(m)
}
func (m *KeyModification) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyModification.DiscardUnknown(m)
}

var xxx_messageInfo_KeyModification proto.InternalMessageInfo

func (m *KeyModification) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *KeyModification) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyModification) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *KeyModification) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func init() {
	proto.RegisterType((*PutStateRequest)(nil), "forest.PutStateRequest")
	proto.RegisterType((*GetStateRequest)(nil), "forest.GetStateRequest")
	proto.RegisterType((*GetHistoryRequest)(nil), "forest.GetHistoryRequest")
	proto.RegisterType((*GetHistoryResponse)(nil), "forest.GetHistoryResponse")
	proto.RegisterType((*KeyModification)(nil), "forest.KeyModification")
}

func init() { proto.RegisterFile("forest.proto", fileDescriptor_47d1b8996afffe6d) }

var fileDescriptor_47d1b8996afffe6d = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x59, 0xd3, 0x96, 0x66, 0x2c, 0xc4, 0x8e, 0x82, 0x0b, 0x7a, 0x58, 0x02, 0x62, 0x4e,
	0x3d, 0xe8, 0xd1, 0x6b, 0xa1, 0x8a, 0x08, 0xb2, 0xde, 0xbc, 0x84, 0x68, 0xa6, 0xb0, 0xa4, 0xe9,
	0xc6, 0xcc, 0x44, 0x9b, 0x7f, 0x2f, 0x26, 0x82, 0x51, 0x44, 0x6f, 0xf3, 0xde, 0x7c, 0x0f, 0xe6,
	0x0d, 0xcc, 0xd6, 0xbe, 0x26, 0x96, 0x45, 0x55, 0x7b, 0xf1, 0x38, 0xe9, 0x55, 0x7c, 0x05, 0xd1,
	0x7d, 0x23, 0x0f, 0x92, 0x09, 0x59, 0x7a, 0x69, 0x88, 0x05, 0x11, 0x46, 0x05, 0xb5, 0xac, 0x95,
	0x09, 0x92, 0xd0, 0x76, 0x33, 0x1e, 0xc1, 0xf8, 0x35, 0xdb, 0x34, 0xa4, 0xf7, 0x8c, 0x4a, 0x66,
	0xb6, 0x17, 0xf1, 0x19, 0x44, 0x2b, 0xfa, 0x37, 0x1c, 0x9f, 0xc3, 0x7c, 0x45, 0x72, 0xed, 0x58,
	0x7c, 0xdd, 0xfe, 0x05, 0x3e, 0x02, 0x0e, 0x41, 0xae, 0xfc, 0x96, 0x09, 0x97, 0x30, 0x2f, 0xa8,
	0x4d, 0x4b, 0x9f, 0xbb, 0xb5, 0x7b, 0xce, 0xc4, 0xf9, 0x6d, 0x1f, 0xdb, 0xbf, 0x38, 0x5e, 0x7c,
	0x96, 0xba, 0xa5, 0xf6, 0x6e, 0xb0, 0xb7, 0x07, 0xc5, 0x77, 0x83, 0xe3, 0x37, 0x88, 0x7e, 0x40,
	0x78, 0x08, 0x63, 0xd9, 0xa5, 0x2e, 0xd7, 0xca, 0xa8, 0x8f, 0x1b, 0x64, 0x77, 0x93, 0xff, 0xde,
	0x14, 0x4f, 0x21, 0x14, 0x57, 0x12, 0x4b, 0x56, 0x56, 0x3a, 0x30, 0x2a, 0x09, 0xec, 0x97, 0x81,
	0x27, 0x10, 0x3a, 0x4e, 0x73, 0xda, 0x90, 0x90, 0x1e, 0x19, 0x95, 0x4c, 0xed, 0xd4, 0xf1, 0xb2,
	0xd3, 0x4f, 0x93, 0xee, 0xe1, 0x97, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x45, 0x79, 0x6c,
	0x80, 0x01, 0x00, 0x00,
}
