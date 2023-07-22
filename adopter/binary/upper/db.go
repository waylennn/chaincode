package upper

import (
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/waylennn/chaincode/adopter/common"
	"github.com/waylennn/chaincode/adopter/utils"
)

func (t *Chaincode) dbUp(stub shim.ChaincodeStubInterface, args []byte) peer.Response {
	if args == nil {
		return shim.Error("DBUP binlog's parameter cannot be nil")
	}

	binlog, err := utils.GetBinlog(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	var max uint64
	for _, dataBytes := range binlog.Data {
		data, err := utils.GetBinlogData(dataBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		if data.Sequence >= max {
			max = data.Sequence
		}

		err = stub.PutState(common.BinlogKey(data.Sequence), data.Log)
		if err != nil {
			return shim.Error(err.Error())
		}

		statsBytes, err := utils.Marshal(data.Stat)
		if err != nil {
			return shim.Error(err.Error())
		}

		if logger.IsEnabledFor(shim.LogDebug) {
			statTime, err := ptypes.Timestamp(data.Stat.Timestamp)
			if err != nil {
				logger.Warningf("Timestamp converts a google.protobuf.Timestamp proto to a time.Time failed: %s\n", err)
			}
			logger.Debugf("Stat time: %s\n", statTime.Format(time.RFC3339))

		}

		err = stub.PutState(common.DBStatKey(data.Sequence), statsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	err = stub.PutState(common.DBMaxSeqKey(), []byte(strconv.FormatUint(max, 10)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
