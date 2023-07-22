package upper

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"github.com/waylennn/chaincode/adopter/common"
	"github.com/waylennn/chaincode/adopter/core/avl"
	"github.com/waylennn/chaincode/adopter/protos/binary"
	"github.com/waylennn/chaincode/adopter/utils"
)

func (t *Chaincode) saveFiles(stub shim.ChaincodeStubInterface, changes []byte) peer.Response {
	logger.Debug("Starting save files")
	groupChange, err := utils.GetGroupChange(changes)
	if err != nil {
		logger.Errorf("Error unmarshaling GroupChange: %s\n", err)
		return shim.Error(err.Error())
	}

	data, err := stub.GetState(common.FileIndexTableKey())
	if err != nil {
		logger.Errorf("Error get file index table: %s\n", err)
		return shim.Error(err.Error())
	}

	var hash string
	var etype binary.FileUploadType
	if data == nil {
		logger.Debug("File index table is nil, save new")

		etype = binary.FileUploadType_FILE_UPLOAD_NEW
		hash, err = t.saveNewFiles(stub, groupChange)
		if err != nil {
			logger.Errorf("Error saving new files: %s\n", err)
			return shim.Error(err.Error())
		}
	} else {
		logger.Debug("File index table already exists, save update")

		etype = binary.FileUploadType_FILE_UPLOAD_UPDATE
		table, err := utils.GetFileIndexTable(data)
		if err != nil {
			logger.Errorf("Error unmarshaling FileIndexTable: %s\n", err)
			return shim.Error(err.Error())
		}

		if len(table.GetIndexes()) == 0 {
			logger.Debug("File index talbe is empty, save new")
			hash, err = t.saveNewFiles(stub, groupChange)
			if err != nil {
				logger.Errorf("Error saving new files: %s\n", err)
				return shim.Error(err.Error())
			}
		} else {
			logger.Debug("File index table is not empty, save update")

			hash, err = t.saveChangeFiles(stub, table, groupChange)
			if err != nil {
				logger.Errorf("Error saving update files: %s\n", err)
				return shim.Error(err.Error())
			}
		}
	}

	logger.Debugf("Group hash: %s", hash)
	if hash != "" {
		payload := &binary.FileUploadEvent{
			Type:     etype,
			DataHash: hash,
		}

		pbytes, err := utils.Marshal(payload)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.SetEvent(groupFileUploadEID, pbytes)
		if err != nil {
			return shim.Error("Unable to set group file update event")
		}
	}

	logger.Debug("Group up Successful")
	return shim.Success(nil)
}

func (t *Chaincode) saveChangeFiles(stub shim.ChaincodeStubInterface, table *binary.FileIndexTable, groupChange *binary.GroupChange) (string, error) {
	if len(groupChange.Changes) == 0 && len(groupChange.Deletes) == 0 {
		return "", errors.New("File changes are empty")
	}

	avlTree := avl.NewImmutable()
	for _, indexBytes := range table.GetIndexes() {
		index, err := utils.GetFileIndex(indexBytes)
		if err != nil {
			return "", err
		}

		avlTree.Insert(index)
	}

	statSetDataBytes := make([][]byte, 0)
	for _, fileBytes := range groupChange.Changes {
		file, err := utils.GetFile(fileBytes)
		if err != nil {
			return "", err
		}
		if file.Name == "" {
			return "", errors.New("File name can't be empty")
		}
		if file.Data == nil {
			return "", errors.New("File data can't be nil")
		}

		index := &binary.FileIndex{FileName: file.Name}
		entry := avlTree.Get(index)

		// entry 为空表示文件是新增
		if entry == nil {
			sha1, err := t.saveNewFile(stub, file.Data)
			if err != nil {
				return "", errors.Errorf("%s: %s", file.Name, err)
			}

			statBytes, err := t.makeNewFileStat(file)
			if err != nil {
				return "", err
			}
			statSetDataBytes = append(statSetDataBytes, statBytes)

			index.Index = sha1
			avlTree.Insert(index)
		} else {
			// 如果程序运行到此表示文件是更新
			exists := entry.(*binary.FileIndex)
			oldIndex := exists.Index
			sha1, change, err := t.updateExistsFile(stub, oldIndex, file.Data)
			if err != nil {
				return "", errors.Errorf("%s: %s", file.Name, err)
			}
			if !change {
				continue
			}
			exists.Index = sha1

			var oldFile *binary.File
			if file.Type == int32(binary.FileType_EXCEL) {
				old, err := stub.GetState(oldIndex)
				if err != nil {
					return "", err
				}
				if old == nil {
					return "", errors.Errorf("Unable to find the last file content with index: %s", oldIndex)
				}

				oldFile = &binary.File{Name: file.Name, Data: old}
			}

			statBytes, err := t.makeUpdateFileStat(file, oldFile)
			if err != nil {
				return "", err
			}
			statSetDataBytes = append(statSetDataBytes, statBytes)
		}
	}

	for _, filename := range groupChange.Deletes {
		if filename == "" {
			continue
		}

		index := &binary.FileIndex{FileName: filename}
		found := avlTree.Delete(index)
		if found {
			statBytes, err := t.makeDeleteFileStat(filename)
			if err != nil {
				return "", err
			}
			statSetDataBytes = append(statSetDataBytes, statBytes)
		}
	}

	if len(statSetDataBytes) == 0 {
		return "", nil
	}

	header, err := utils.GetFileIndexHeader(table.GetHeader())
	if err != nil {
		return "", err
	}

	indexes, err := avlTree.InOrder()
	if err != nil {
		return "", err
	}

	hashBytes := make([]byte, 0)
	for _, index := range indexes {
		hashBytes = append(hashBytes, index...)
	}

	now := ptypes.TimestampNow()
	hash, err := utils.ComputeSHA1String(hashBytes)
	if err != nil {
		return "", err
	}

	header.Timestamp = now
	header.DataHash = hash
	headerBytes, err := utils.Marshal(header)
	if err != nil {
		return "", err
	}

	newTable := &binary.FileIndexTable{Header: headerBytes, Indexes: indexes}
	newTableBytes, err := utils.Marshal(newTable)
	if err != nil {
		return "", err
	}

	err = stub.PutState(common.FileIndexTableKey(), newTableBytes)
	if err != nil {
		return "", err
	}

	fileStatSet := &binary.FileStatSet{Timestamp: now, DataHash: hash, Data: statSetDataBytes}
	err = t.saveFileStatSet(stub, fileStatSet)
	return hash, err
}

func (t *Chaincode) fileShouldDelete(filename string, deletes []string) bool {
	for _, fname := range deletes {
		if fname == filename {
			return true
		}
	}
	return false
}

func (t *Chaincode) updateExistsFile(stub shim.ChaincodeStubInterface, oldSha1Str string, filedata []byte) (string, bool, error) {
	sha1, err := utils.ComputeSHA1String(filedata)
	if err != nil {
		return sha1, false, errors.Wrap(err, "Failed to calculate file fingerprint")
	}

	if sha1 == oldSha1Str {
		return sha1, false, nil
	}

	err = stub.PutState(sha1, filedata)
	if err != nil {
		return sha1, false, errors.Wrap(err, "Failed to save file content")
	}

	return sha1, true, nil
}

func (t *Chaincode) saveNewFile(stub shim.ChaincodeStubInterface, filedata []byte) (string, error) {
	sha1, err := utils.ComputeSHA1String(filedata)
	if err != nil {
		return sha1, errors.Wrap(err, "Failed to calculate file fingerprint")
	}
	bs, err := stub.GetState(sha1)
	if err != nil {
		return sha1, errors.Wrap(err, "Failed to get file content")
	}
	if bs == nil {
		err := stub.PutState(sha1, filedata)
		if err != nil {
			return sha1, errors.Wrap(err, "Failed to save file content")
		}
	}
	return sha1, nil
}

func (t *Chaincode) saveNewFiles(stub shim.ChaincodeStubInterface, groupChange *binary.GroupChange) (string, error) {
	if len(groupChange.Changes) == 0 {
		return "", errors.New("File changes are empty")
	}

	indexes := make([][]byte, 0)
	statSetDataBytes := make([][]byte, 0)
	hashBytes := make([]byte, 0)

	for _, fileBytes := range groupChange.Changes {
		file, err := utils.GetFile(fileBytes)
		if err != nil {
			return "", err
		}
		if file.Name == "" {
			return "", errors.New("File name can't be empty")
		}
		if file.Data == nil {
			return "", errors.New("File data can't be nil")
		}

		sha1, err := utils.ComputeSHA1String(file.Data)
		if err != nil {
			return "", errors.Wrapf(err, "Save file %s failed", file.Name)
		}

		indexBytes, err := createFileIndex(file.Name, sha1)
		if err != nil {
			return "", errors.Wrapf(err, "Save file %s failed", file.Name)
		}

		err = stub.PutState(sha1, file.Data)
		if err != nil {
			return "", errors.Wrapf(err, "Save file %s failed", file.Name)
		}

		statBytes, err := t.makeNewFileStat(file)

		if err != nil {
			return "", err
		}

		indexes = append(indexes, indexBytes)
		hashBytes = append(hashBytes, indexBytes...)
		statSetDataBytes = append(statSetDataBytes, statBytes)
	}

	now := ptypes.TimestampNow()
	hash, err := utils.ComputeSHA1String(hashBytes)
	if err != nil {
		return "", err
	}

	header := &binary.FileIndexHeader{Timestamp: now, DataHash: hash}
	headerBytes, err := utils.Marshal(header)
	if err != nil {
		return "", err
	}

	table := &binary.FileIndexTable{Header: headerBytes, Indexes: indexes}
	tableBytes, err := utils.Marshal(table)
	if err != nil {
		return "", err
	}

	err = stub.PutState(common.FileIndexTableKey(), tableBytes)
	if err != nil {
		return "", err
	}

	fileStatSet := &binary.FileStatSet{Timestamp: now, DataHash: hash, Data: statSetDataBytes}
	err = t.saveFileStatSet(stub, fileStatSet)
	return hash, err
}

func createFileIndex(filename string, sha1 string) ([]byte, error) {
	index := &binary.FileIndex{FileName: filename, Index: sha1}
	return utils.Marshal(index)
}
