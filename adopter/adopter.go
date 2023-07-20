package adopter

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/protos/adopter"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// These are function names from Invoke first parameter
const (
	SetResource  = "SetResource"
	SetAuthority = "SetAuthority"
)

// Key prefix
const (
	MetaKeyPrefix = "Meta~"
	AuthKeyPrefix = "Auth~"
)

var logger = shim.NewLogger("adopter")

// Chaincode 链码
type Chaincode struct{}

// Init 实现 Chaincode interface Init 方法
func (s *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Init chaincode adopter")
	return shim.Success(nil)
}

// Invoke 实现 Chaincode interface Invoke 方法
func (s *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetArgs()
	if len(args) < 2 {
		return shim.Error(fmt.Sprintf("Incorrect number of arguments, %d", len(args)))
	}

	fname := string(args[0])
	rid := string(args[1])

	if len(args) < 3 {
		return shim.Error(fmt.Sprintf("Missing 3rd argument for %s", fname))
	}

	logger.Infof("Invoke function: %s with resource id: %s", fname, rid)

	switch fname {
	case SetResource:
		return setResource(stub, rid, args[2])
	case SetAuthority:
		return setAuthority(stub, rid, args[2])
	}

	return shim.Error(fmt.Sprintf("Requested function %s not found.", fname))
}

func setResource(stub shim.ChaincodeStubInterface, rid string, reqBytes []byte) pb.Response {
	req := &adopter.SetResourceRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set resource with id %s, proto unmarshal error: %s", rid, err))
	}

	key, err := stub.CreateCompositeKey(MetaKeyPrefix, []string{rid, req.Commit})
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set resource with id %s, create composite key error %s", rid, err))
	}
	err = stub.PutState(key, req.Payload)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set resource with id %s, put state error %s", rid, err))
	}
	return shim.Success(nil)
}

func setAuthority(stub shim.ChaincodeStubInterface, rid string, reqBytes []byte) pb.Response {
	req := &adopter.SetAuthorityRequest{}
	err := proto.Unmarshal(reqBytes, req)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set authority with id %s, proto unmarshal error: %s", rid, err))
	}

	key, err := stub.CreateCompositeKey(AuthKeyPrefix, []string{rid, req.AuthOrgId})
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set authority with id %s, create composite key error %s", rid, err))
	}
	err = stub.PutState(key, req.Payload)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to set authority with id %s, put state error %s", rid, err))
	}
	return shim.Success(nil)
}
