package downer

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/waylennn/chaincode/adopter/common"
	"github.com/waylennn/chaincode/adopter/protos/binary"
	"github.com/waylennn/chaincode/adopter/utils"
)

func getDBDown(stub shim.ChaincodeStubInterface, subcmd string, seqStr string) peer.Response {
	seq, err := strconv.ParseUint(seqStr, 10, 64)
	if err != nil {
		logger.Errorf("DB binlog sequence number format error: %s, %s", seqStr, err)
		return shim.Error(err.Error())
	}

	switch subcmd {
	case DBDownBinlog:
		return getDBDownBinlog(stub, seq)
	case CloudchainDownStat:
		return getDBDownStat(stub, seq)
	}

	return shim.Error(fmt.Sprintf("Requested function %s not found for %s", subcmd, DBDown))
}

func getDBDownStat(stub shim.ChaincodeStubInterface, seq uint64) peer.Response {
	invokeArgs := [][]byte{[]byte(common.GSCCGetState), []byte(common.UpperChaincodeID), []byte(common.DBMaxSeqKey())}
	response := stub.InvokeChaincode(common.GSCC, invokeArgs, stub.GetChannelID())
	if response.Status != shim.OK {
		logger.Errorf("Error getting state with key: %s, %s", common.DBMaxSeqKey(), response.Message)
		return shim.Error(response.Message)
	}
	if response.Payload == nil {
		logger.Warning("The max sequence number is missing")
		return shim.Success(nil)
	}

	max, err := strconv.ParseUint(string(response.Payload), 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	stats := make([]*binary.Stat, 0)
	for i := uint64(0); i < seq; i++ {
		max = max - i
		if max < 0 {
			break
		}

		args := [][]byte{[]byte(common.GSCCGetState), []byte(common.UpperChaincodeID), []byte(common.DBStatKey(max))}
		resp := stub.InvokeChaincode(common.GSCC, args, stub.GetChannelID())
		if resp.Status != shim.OK {
			logger.Errorf("Error getting stat with key: %s, %s", common.DBStatKey(max), resp.Message)
			return shim.Error(resp.Message)
		}
		if resp.Payload == nil {
			logger.Warningf("DB stat with sequence number %d is missing", max)
			continue
		}

		stat, err := utils.GetStat(resp.Payload)
		if err != nil {
			logger.Errorf("Error unmarshaling Stat: %s", err)
			return shim.Error(err.Error())
		}

		stats = append(stats, stat)
	}

	if len(stats) == 0 {
		logger.Warning("No database stat found")
		return shim.Success(nil)
	}

	statsBytes, err := utils.Marshal(&binary.Stats{Data: stats})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(statsBytes)
}

func getDBDownBinlog(stub shim.ChaincodeStubInterface, seq uint64) peer.Response {
	invokeArgs := [][]byte{[]byte(common.GSCCGetState), []byte(common.UpperChaincodeID), []byte(common.BinlogKey(seq))}
	response := stub.InvokeChaincode(common.GSCC, invokeArgs, stub.GetChannelID())
	if response.Status != shim.OK {
		logger.Errorf("Error getting stat for key: %s, %s", common.BinlogKey(seq), response.Message)
		return shim.Error(response.Message)
	}

	if response.Payload == nil {
		logger.Errorf("DB binlog with sequence number %d was nil", seq)
		return shim.Error(fmt.Sprintf("binlog does not exist - %d", seq))
	}

	return shim.Success(response.Payload)
}
