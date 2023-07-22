package rock

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	Namespace    = "Rock~"
	FnPutState   = "PutState"
	FnGetState   = "GetState"
	FnGetHistory = "GetHistory"
)

var logger = shim.NewLogger(Namespace)

type Contract struct{}

func (c *Contract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Infof("contract init")
	return shim.Success(nil)
}

func (c *Contract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	logger.Infof("contract invoke: fn = %s", fn)
	switch fn {
	case FnPutState:
		return putState(stub, args)
	case FnGetState:
		return getState(stub, args)
	case FnGetHistory:
		return getHistory(stub, args)
	default:
		return shim.Error("unsupported fn")
	}
}

func putState(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("failed to put state: incorrect number of arguments: %v", len(args)))
	}
	req := &PutStateRequest{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put state, failed to unmarshal request: %v", err))
	}
	key, err := stub.CreateCompositeKey(Namespace, req.Keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put state, failed to create composite key: %v", err))
	}
	err = stub.PutState(key, []byte(req.Value))
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put state, failed to call stub.PutState: %v", err))
	}
	return shim.Success(nil)
}

func getState(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("failed to get state, incorrect number of arguments: %v", len(args)))
	}
	req := &GetStateRequest{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get state, failed to unmarshal request: %v", err))
	}
	key, err := stub.CreateCompositeKey(Namespace, req.Keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get state, failed to create composite key: %v", err))
	}
	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get value, failed to call stub.GetState: %v", err))
	}
	payload := value
	return shim.Success(payload)
}

func getHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("failed to get history, incorrect number of arguments: %v", len(args)))
	}
	req := &GetHistoryRequest{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to unmarshal request: %v", err))
	}
	key, err := stub.CreateCompositeKey(Namespace, req.Keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to create composite key: %v", err))
	}
	iter, err := stub.GetHistoryForKey(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to call stub.GetHistoryForKey: %v", err))
	}
	defer iter.Close()
	resp := &GetHistoryResponse{}
	for iter.HasNext() {
		km, err := iter.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to get history, failed to iterate next: %v", err))
		}
		keyModification := &KeyModification{
			TxId:      km.TxId,
			Value:     string(km.Value),
			Timestamp: km.Timestamp.GetSeconds(),
			IsDelete:  km.IsDelete,
		}
		resp.KeyModifications = append(resp.KeyModifications, keyModification)
	}
	payload, err := json.Marshal(resp)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to marshal response: %v", err))
	}
	return shim.Success(payload)
}
