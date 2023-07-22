package downer

import (
	"fmt"

	"github.com/waylennn/chaincode/adopter/core/query"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"github.com/waylennn/chaincode/adopter/common"
	"github.com/waylennn/chaincode/adopter/protos/binary"
	"github.com/waylennn/chaincode/adopter/utils"
)

func getGroupFileList(stub shim.ChaincodeStubInterface) peer.Response {

	table, err := getFileIndexTable(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	header, err := utils.GetFileIndexHeader(table.Header)
	if err != nil {
		return shim.Error(err.Error())
	}

	list := &binary.GroupList{Timestamp: header.Timestamp}
	m := make(map[string]string)

	for _, indexBytes := range table.Indexes {
		index, err := utils.GetFileIndex(indexBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		m[index.FileName] = index.Index
	}
	list.Files = m

	listBytes, err := utils.Marshal(list)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(listBytes)
}

func getFileIndexTable(stub shim.ChaincodeStubInterface) (*binary.FileIndexTable, error) {
	invokeArgs := [][]byte{[]byte(common.GSCCGetState), []byte(common.UpperChaincodeID), []byte(common.FileIndexTableKey())}
	logger.Debugf("Start invoke gscc system chaincode")
	response := stub.InvokeChaincode(common.GSCC, invokeArgs, stub.GetChannelID())
	if response.Status != shim.OK {
		return nil, errors.Errorf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
	}
	if response.Payload == nil {
		return nil, errors.New("You have not uploaded any files yet")
	}

	logger.Debugf("Success fetching file index table")
	return utils.GetFileIndexTable(response.Payload)
}

func getGroupDown(stub shim.ChaincodeStubInterface, subcmd string) peer.Response {
	switch subcmd {
	case GroupDownFile:
		return getGroupDownFile(stub)
	case CloudchainDownStat:
		return getGroupDownStat(stub)
	}

	return shim.Error(fmt.Sprintf("Requested function %s not found for %s", subcmd, GroupDown))
}

func getGroupDownFile(stub shim.ChaincodeStubInterface) peer.Response {

	table, err := getFileIndexTable(stub)
	if err != nil {
		return shim.Error(errors.WithMessage(err, "Get file index table error").Error())
	}

	if len(table.Indexes) == 0 {
		return shim.Error("There are no files can be downloaded")
	}

	m, err := getFileDataMap(stub, table)
	if err != nil {
		return shim.Error(err.Error())
	}

	files := &binary.GroupFile{Files: m}
	fileBytes, err := utils.Marshal(files)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(fileBytes)
}

func getFileDataMap(stub shim.ChaincodeStubInterface, table *binary.FileIndexTable) (map[string][]byte, error) {
	m := make(map[string][]byte)
	for _, indexBytes := range table.Indexes {
		index, err := utils.GetFileIndex(indexBytes)
		if err != nil {
			logger.Errorf("Error unmarshaling FileIndex: %s", err)
			return nil, err
		}

		data, err := getFileFromIndex(stub, index.Index)
		if err != nil {
			return nil, errors.Wrapf(err, "Bad file (%s) index: %s", index.FileName, index.Index)
		}

		m[index.FileName] = data
	}

	return m, nil
}

func getFileFromIndex(stub shim.ChaincodeStubInterface, index string) ([]byte, error) {
	invokeArgs := [][]byte{[]byte(common.GSCCGetState), []byte(common.UpperChaincodeID), []byte(index)}
	response := stub.InvokeChaincode(common.GSCC, invokeArgs, stub.GetChannelID())
	if response.Status != shim.OK {
		logger.Errorf("Error getting state with key: %s, %s", index, string(response.Payload))
		return nil, errors.Errorf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
	}
	if response.Payload == nil {
		logger.Errorf("File with index %s is missing", index)
		return nil, errors.Errorf("File %s data is missing", index)
	}

	return response.Payload, nil
}

func getGroupDownStat(stub shim.ChaincodeStubInterface) peer.Response {

	iter, err := query.NewHistoryQueryIterator(stub, common.UpperChaincodeID, common.FileStatKey())
	if err != nil {
		logger.Errorf("Error getting history for key: %s, %s", common.FileStatKey(), err)
		return shim.Error(err.Error())
	}
	defer iter.Close()

	sets := &binary.FileStatsSets{}
	for iter.HasNext() {
		result, err := iter.Next()

		switch {
		case err != nil:
			logger.Errorf("Error moving next item: %s", err)
			return shim.Error(err.Error())
		default:
			sets.Data = append(sets.Data, result.Value)
		}
	}

	setsBytes, err := utils.Marshal(sets)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(setsBytes)
}

func getFileDown(stub shim.ChaincodeStubInterface, hash string) peer.Response {

	data, err := getFileFromIndex(stub, hash)
	if err != nil {
		return shim.Error(err.Error())
	}
	if data == nil {
		return shim.Error(errors.Errorf("Bad file index: %s", hash).Error())
	}

	return shim.Success(data)
}

func getGroupHistoryFiles(stub shim.ChaincodeStubInterface, hash string) peer.Response {

	iter, err := query.NewHistoryQueryIterator(stub, common.UpperChaincodeID, common.FileIndexTableKey())
	if err != nil {
		logger.Errorf("Error getting history for key: %s, %s", common.FileIndexTableKey(), err)
		return shim.Error(err.Error())
	}
	defer iter.Close()

	var found *binary.FileIndexTable

Loop:
	for iter.HasNext() {
		result, err := iter.Next()
		switch {
		case err != nil:
			logger.Errorf("Error moving next item: %s", err)
			return shim.Error(err.Error())
		default:
			table, err := utils.GetFileIndexTable(result.Value)
			if err != nil {
				return shim.Error(err.Error())
			}

			header, err := utils.GetFileIndexHeader(table.Header)
			if err != nil {
				return shim.Error(err.Error())
			}

			if header.DataHash == hash {
				found = table
				break Loop
			}
		}
	}

	if found == nil {
		return shim.Error(fmt.Sprintf("Can't find any history for commit id: %s", hash))
	}

	m, err := getFileDataMap(stub, found)
	if err != nil {
		return shim.Error(err.Error())
	}

	files := &binary.GroupFile{Files: m}
	filesBytes, err := utils.Marshal(files)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(filesBytes)
}
