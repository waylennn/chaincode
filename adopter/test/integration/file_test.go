package integration

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"github.com/waylennn/chaincode/adopter/binary"
	pb "github.com/waylennn/chaincode/adopter/protos/binary"
	"github.com/waylennn/chaincode/adopter/utils"
)

func TestChaincodeInit(t *testing.T) {
	cc := binary.NewChaincode()
	stub := shim.NewMockStub("cloudchain", cc)

	checkInit(t, stub, [][]byte{[]byte("init"), []byte("initialize")})
}

func TestChaincodeGroupUpNew(t *testing.T) {
	cc := binary.NewChaincode()
	stub := shim.NewMockStub("cloudchain", cc)

	changes := make([][]byte, 3)
	docxFilePath := path.Join(GetTestFixturePath(), "sample.docx")
	docxData, err := ioutil.ReadFile(docxFilePath)
	assert.Nil(t, err)

	bytes, err := utils.Marshal(&pb.File{Name: "sample.docx", Data: docxData})
	assert.Nil(t, err)
	changes[0] = bytes

	xlsxFilePath := path.Join(GetTestFixturePath(), "sample.xlsx")
	xlsxData, err := ioutil.ReadFile(xlsxFilePath)
	assert.Nil(t, err)

	bytes, err = utils.Marshal(&pb.File{Name: "sample.xlsx", Data: xlsxData})
	assert.Nil(t, err)
	changes[1] = bytes

	pptxFilePath := path.Join(GetTestFixturePath(), "sample.pptx")
	pptxData, err := ioutil.ReadFile(pptxFilePath)
	assert.Nil(t, err)

	bytes, err = utils.Marshal(&pb.File{Name: "sample.pptx", Data: pptxData})
	assert.Nil(t, err)
	changes[2] = bytes

	gc := &pb.GroupChange{
		Changes: changes,
		Deletes: []string{},
	}

	gcBytes, err := utils.Marshal(gc)
	assert.Nil(t, err)

	checkInvoke(t, stub, [][]byte{[]byte("GROUPUP"), gcBytes})
}

func TestGoupListEmpty(t *testing.T) {
	cc := binary.NewChaincode()
	stub := shim.NewMockStub("cloudchain", cc)

	args := [][]byte{[]byte("GROUPLIST")}
	res := stub.MockInvoke("222", args)

	assert.Equal(t, int32(shim.ERROR), res.Status)
	assert.Contains(t, string(res.Message), "You have not uploaded any files yet")
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	assert.Equal(t, int32(shim.OK), res.Status)
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	assert.Equal(t, int32(shim.OK), res.Status)
}

func checkQuery(t *testing.T, stub *shim.MockStub, args [][]byte) []byte {
	res := stub.MockInvoke("22", args)
	assert.Equal(t, int32(shim.OK), res.Status)
	assert.NotNil(t, res.Payload)

	return res.Payload
}

func GetTestFixturePath() string {
	return path.Join(utils.GoPath(), "src", "adopter", "test/fixtures")
}
