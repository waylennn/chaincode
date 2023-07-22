package binary

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/waylennn/chaincode/adopter/core/avl"
)

// Compare implements the method of interface `avl.Entry`
func (index *FileIndex) Compare(other avl.Entry) int {
	otherMe := other.(*FileIndex)
	if index.FileName > otherMe.FileName {
		return 1
	}

	if index.FileName < otherMe.FileName {
		return -1
	}

	return 0
}

// ToBytes serializes the instance to bytes
func (index *FileIndex) ToBytes() ([]byte, error) {
	return proto.Marshal(index)
}
