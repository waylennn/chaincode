package downer

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// These are function names from Invoke first parameter
const (
	GroupList    string = "GROUPLIST"
	GroupDown    string = "GROUPDOWN"
	FileDown     string = "FILEDOWN"
	GroupHistory string = "GROUPHISTORY"
	DBDown       string = "DBDOWN"

	GroupDownFile      string = "file"
	DBDownBinlog       string = "binlog"
	CloudchainDownStat string = "stat"
)

var logger = shim.NewLogger("binary/downer")

// Chaincode 链码
type Chaincode struct{}

// Init 实现 ChainCode 接口的Init方法
func (s *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Info("Init binary/downer")

	return shim.Success(nil)
}

// Invoke 实现 ChainCode 接口的Invoke方法
func (s *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetArgs()

	if len(args) < 1 {
		return shim.Error(fmt.Sprintf("Incorrect number of arguments, %d", len(args)))
	}
	fname := string(args[0])

	logger.Debugf("Invoke function: %s on chain: %s", fname, stub.GetChannelID())

	switch fname {
	case GroupList:
		return getGroupFileList(stub)

	case GroupDown:
		if len(args) < 2 {
			return shim.Error(fmt.Sprintf("missing 2nd argument for %s", fname))
		}

		subcmd := string(args[1])
		logger.Infof("Group down sub command: %s", subcmd)

		return getGroupDown(stub, subcmd)

	case FileDown:
		if len(args) < 2 {
			return shim.Error(fmt.Sprintf("missing 2nd argument for %s", fname))
		}

		return getFileDown(stub, string(args[1]))

	case GroupHistory:
		if len(args) < 2 {
			return shim.Error(fmt.Sprintf("missing 2nd argument for %s", fname))
		}

		return getGroupHistoryFiles(stub, string(args[1]))

	case DBDown:
		if len(args) < 3 {
			return shim.Error(fmt.Sprintf("missing 3rd argument for %s", fname))
		}

		subcmd := string(args[1])
		logger.Infof("DB download sub command: %s", subcmd)

		return getDBDown(stub, subcmd, string(args[2]))
	}

	return shim.Error(fmt.Sprintf("Requested function %s not found", fname))
}
