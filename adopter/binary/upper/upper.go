package upper

import (
	"bytes"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	systemKeyPrefix    = "cloudchain"
	ownerMSPIDKey      = "cloudchain_owner_mspid"
	groupFileUploadEID = "Event:GroupFileUpload"
)

var logger = shim.NewLogger("binary/upper")

// Chaincode 链码
type Chaincode struct{}

// Init 实现 ChainCode 接口的Init方法
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Info("Chaincode Init")

	function, _ := stub.GetFunctionAndParameters()

	switch function {
	case "init":
		msp, err := cid.GetMSPID(stub)
		if err != nil {
			return shim.Error(fmt.Sprintf("GetMSPID failed: %s", err))
		}
		logger.Debugf("Get MSPID: %s", msp)

		err = stub.PutState(ownerMSPIDKey, []byte(msp))
		if err != nil {
			return shim.Error(fmt.Sprintf("Save MSPID failed: %s", err))
		}

		return shim.Success(nil)

	case "upgrade":
		if !t.checkMSPID(stub) {
			logger.Errorf("Check invoker msp id failed")
			return shim.Error("Access denied: You are not the owner of the channel")
		}

		return shim.Success(nil)

	default:
		return shim.Error(fmt.Sprintf("Unknow init function: %s", function))
	}
}

// Invoke 实现 ChainCode 接口的Invoke方法
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	functionAndParameters := stub.GetArgs()
	if len(functionAndParameters) < 1 {
		return shim.Error("No parameters provided")
	}
	function := string(functionAndParameters[0])
	args := functionAndParameters[1:]

	logger.Debugf("Invoke function: %s on chain: %s", function, stub.GetChannelID())

	switch function {
	case "GROUPUP":
		if len(args) != 1 {
			logger.Error("Incorrect number of arguments. Expecting 1")
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		if !t.checkMSPID(stub) {
			logger.Errorf("Check invoker msp id failed")
			return shim.Error("Access denied: You are not the owner of the channel")
		}

		params := args[0]
		if params == nil {
			logger.Error("The arguments is nil")
			return shim.Error("The arguments is nil")
		}

		return t.saveFiles(stub, params)

	case "DBUP":
		if len(args) != 1 {
			logger.Error("Incorrect number of arguments. Expecting 1")
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		if !t.checkMSPID(stub) {
			logger.Errorf("Check invoker msp id failed")
			return shim.Error("Access denied: You are not the owner of the channel")
		}

		return t.dbUp(stub, args[0])

	default:
		logger.Errorf("Unknow chaincode action: %s", function)
		return shim.Error("Unkonw action")
	}
}

func (t *Chaincode) checkMSPID(stub shim.ChaincodeStubInterface) bool {
	ownerMSPID, err := stub.GetState(ownerMSPIDKey)
	if err != nil {
		logger.Errorf("Get owner msp id failed: %s", err)
		return false
	}

	msp, err := cid.GetMSPID(stub)
	if err != nil {
		logger.Errorf("Get invoker msp id failed: %s", err)
		return false
	}

	if bytes.Compare(ownerMSPID, []byte(msp)) != 0 {
		logger.Warningf("Invoker msp id (%s) is different from owner msp id (%s)", msp, string(ownerMSPID))
		return false
	}

	return true
}
