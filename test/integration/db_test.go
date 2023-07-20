package integration

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.querycap.com/cloudchain/chaincode/binary"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const dbupParamsStr = `CmYIARIuCAEQgoDgxrLZjswFIiAKBXRlc3QxEgAaFWNyZWF0ZSBkYXRhYmFzZSB0ZXN0MRoyCgoItO5dEPC1u6gDEiQKBXRlc3QxEhsSABoXChVjcmVhdGUgZGF0YWJhc2UgdGVzdDEKZggCEi4IARCCgMDstNmOzAUiIAoFdGVzdDISABoVY3JlYXRlIGRhdGFiYXNlIHRlc3QyGjIKCgi07l0QoO3KqQMSJAoFdGVzdDISGxIAGhcKFWNyZWF0ZSBkYXRhYmFzZSB0ZXN0Mg==`

func TestDBUp(t *testing.T) {
	cc := binary.NewChaincode()
	stub := shim.NewMockStub("cloudchain", cc)

	bytes, err := base64.StdEncoding.DecodeString(dbupParamsStr)
	assert.Nil(t, err)

	resp := stub.MockInvoke("123", [][]byte{[]byte("DBUP"), bytes})
	assert.Equal(t, int32(shim.OK), resp.Status)
	if resp.Status != int32(shim.OK) {
		t.Log(resp.Message)
	}
}
