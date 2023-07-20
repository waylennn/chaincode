package utils

import (
	"git.querycap.com/cloudchain/chaincode/protos/adopter"
	"git.querycap.com/cloudchain/chaincode/protos/binary"
	"github.com/golang/protobuf/proto"
)

// Marshal serializes a protobuf message.
func Marshal(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

// GetAdopterStat ...
func GetAdopterStat(bytes []byte) (*adopter.Stat, error) {
	s := &adopter.Stat{}
	err := proto.Unmarshal(bytes, s)
	return s, err
}

// GetAdopterGroup ...
func GetAdopterGroup(bytes []byte) (*adopter.Group, error) {
	g := &adopter.Group{}
	err := proto.Unmarshal(bytes, g)
	return g, err
}

// GetAdopterFilePayloadV1 ...
func GetAdopterFilePayloadV1(bytes []byte) (*adopter.FilePayloadV1, error) {
	payload := &adopter.FilePayloadV1{}
	err := proto.Unmarshal(bytes, payload)
	return payload, err
}

// GetAdopterGMetadataV1 ...
func GetAdopterGMetadataV1(bytes []byte) (*adopter.GMetadataV1, error) {
	m := &adopter.GMetadataV1{}
	err := proto.Unmarshal(bytes, m)
	return m, err
}

// GetAdopterSMetadataV1 ...
func GetAdopterSMetadataV1(bytes []byte) (*adopter.SMetadataV1, error) {
	m := &adopter.SMetadataV1{}
	err := proto.Unmarshal(bytes, m)
	return m, err
}

// GetFile gets `File` from bytes
func GetFile(bytes []byte) (*binary.File, error) {
	file := &binary.File{}
	err := proto.Unmarshal(bytes, file)
	return file, err
}

// GetFileIndexTableHistory gets `FileIndexTableHistory` from bytes
func GetFileIndexTableHistory(bytes []byte) (*binary.FileIndexTableHistory, error) {
	his := &binary.FileIndexTableHistory{}
	err := proto.Unmarshal(bytes, his)
	return his, err
}

// GetGroupFile gets `GroupFile` from bytes
func GetGroupFile(bytes []byte) (*binary.GroupFile, error) {
	file := &binary.GroupFile{}
	err := proto.Unmarshal(bytes, file)
	return file, err
}

// GetStat gets `Stat` from bytes
func GetStat(bytes []byte) (*binary.Stat, error) {
	stat := &binary.Stat{}
	err := proto.Unmarshal(bytes, stat)
	return stat, err
}

// GetBinlog gets `Binlog` from bytes
func GetBinlog(bytes []byte) (*binary.Binlog, error) {
	binlog := &binary.Binlog{}
	err := proto.Unmarshal(bytes, binlog)
	return binlog, err
}

// GetBinlogData gets `BinlogData` from bytes
func GetBinlogData(bytes []byte) (*binary.BinlogData, error) {
	binlogData := &binary.BinlogData{}
	err := proto.Unmarshal(bytes, binlogData)
	return binlogData, err
}

// GetGroupList gets `GroupList` from bytes
func GetGroupList(bytes []byte) (*binary.GroupList, error) {
	list := &binary.GroupList{}
	err := proto.Unmarshal(bytes, list)
	return list, err
}

// GetGroupChange gets `GroupChange` from bytes
func GetGroupChange(bytes []byte) (*binary.GroupChange, error) {
	changes := &binary.GroupChange{}
	err := proto.Unmarshal(bytes, changes)
	return changes, err
}

// GetFileIndexHeader gets `FileIndexHeader` from bytes
func GetFileIndexHeader(bytes []byte) (*binary.FileIndexHeader, error) {
	header := &binary.FileIndexHeader{}
	err := proto.Unmarshal(bytes, header)
	return header, err
}

// GetFileIndexTable gets `FileIndexTable` from bytes
func GetFileIndexTable(bytes []byte) (*binary.FileIndexTable, error) {
	table := &binary.FileIndexTable{}
	err := proto.Unmarshal(bytes, table)
	return table, err
}

// GetFileIndex gets `FileIndex` from bytes
func GetFileIndex(bytes []byte) (*binary.FileIndex, error) {
	index := &binary.FileIndex{}
	err := proto.Unmarshal(bytes, index)
	return index, err
}
