package diff

import (
	"bytes"

	"github.com/waylennn/chaincode/adopter/core/avl"
)

type entry struct {
	PK    []byte
	Value []string
}

func (e *entry) Compare(other avl.Entry) int {
	otherMe := other.(*entry)
	return bytes.Compare(e.PK, otherMe.PK)
}

func (e *entry) ToBytes() ([]byte, error) {
	return e.PK[:], nil
}
