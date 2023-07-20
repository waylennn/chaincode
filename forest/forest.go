package forest

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/protos/forest"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	FnPutState   = "PutState"
	FnGetState   = "GetState"
	FnGetHistory = "GetHistory"
)

const (
	KeyPrefix = "Forest~"
)

var logger = shim.NewLogger("forest")

type Asset struct{}

func (s *Asset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Init asset: forest")
	return shim.Success(nil)
}

func (s *Asset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case FnPutState:
		return putState(stub, args)
	case FnGetState:
		return getState(stub, args)
	case FnGetHistory:
		return getHistory(stub, args)
	default:
		return shim.Error(fmt.Sprintf("Incorrect function: %s", fn))
	}
}

func putState(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Failed to put state, incorrect number of arguments: %d", len(args)))
	}

	req := &forest.PutStateRequest{}
	err := proto.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put state, unmarshal request error: %s", err))
	}

	key, err := stub.CreateCompositeKey(KeyPrefix, req.Keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put state, create composite key error: %s", err))
	}

	err = stub.PutState(key, req.Value)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to put state, put state error: %s", err))
	}

	return shim.Success(nil)
}

func getState(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Failed to get state, incorrect number of arguments: %d", len(args)))
	}

	req := &forest.GetStateRequest{}
	err := proto.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get state, unmarshal request error: %s", err))
	}

	key, err := stub.CreateCompositeKey(KeyPrefix, req.Keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get state, create composite key error: %s", err))
	}

	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get value, get state error: %s", err))
	}

	payload := value
	return shim.Success(payload)
}

func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Failed to get history, incorrect number of arguments: %d", len(args)))
	}

	req := &forest.GetHistoryRequest{}
	err := proto.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get history, unmarshal request error: %s", err))
	}

	key, err := stub.CreateCompositeKey(KeyPrefix, req.Keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get history, create composite key error: %s", err))
	}

	iter, err := stub.GetHistoryForKey(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get history, get history error: %s", err))
	}
	defer iter.Close()

	resp := &forest.GetHistoryResponse{}
	for iter.HasNext() {
		km, err := iter.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to get history, iterate error: %s", err))
		}
		keyModification := &forest.KeyModification{
			TxId:      km.TxId,
			Value:     km.Value,
			Timestamp: km.Timestamp.GetSeconds(),
			IsDelete:  km.IsDelete,
		}
		resp.KeyModifications = append(resp.KeyModifications, keyModification)
	}

	payload, err := proto.Marshal(resp)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get history, marshal response error: %s", err))
	}
	return shim.Success(payload)
}
