package common

import (
	"fmt"
)

const (
	systemKeyPrefix   = "cloudchain"
	fileIndexTableKey = "file_index_table"
	fileStatKey       = "file_stat"
	dbMaxSeqKey       = "max_sequence"
)

// DBStatKey generates db stat key in ledger
func DBStatKey(seq uint64) string {
	return fmt.Sprintf("stat_%d", seq)
}

// DBMaxSeqKey generates db max sequence number key in ledger
func DBMaxSeqKey() string {
	return generateSystemKey(dbMaxSeqKey)
}

// BinlogKey generates db binlog key in ledger
func BinlogKey(seq uint64) string {
	return fmt.Sprintf("binlog_%d", seq)
}

// FileStatKey generates file stat key in ledger
func FileStatKey() string {
	return generateSystemKey(fileStatKey)
}

// FileIndexTableKey generates file index table key in ledger
func FileIndexTableKey() string {
	return generateSystemKey(fileIndexTableKey)
}

func generateSystemKey(key string) string {
	return fmt.Sprintf("%s_%s", systemKeyPrefix, key)
}
