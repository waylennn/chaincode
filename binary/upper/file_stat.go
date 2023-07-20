package upper

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/common"
	"git.querycap.com/cloudchain/chaincode/core/diff"
	"git.querycap.com/cloudchain/chaincode/protos/binary"
	"git.querycap.com/cloudchain/chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

func (t *Chaincode) makeDeleteFileStat(name string) ([]byte, error) {
	stat := &binary.FileStat{Log: fmt.Sprintf("删除文件: %s", name), HasDetail: false}
	statBytes, err := utils.Marshal(stat)
	return statBytes, err
}

func (t *Chaincode) makeNewFileStat(file *binary.File) ([]byte, error) {
	if file.Type != int32(binary.FileType_EXCEL) {
		stat := &binary.FileStat{Log: fmt.Sprintf("新增文件: %s", file.Name), HasDetail: false}
		statBytes, err := utils.Marshal(stat)
		return statBytes, err
	}

	excel, err := xlsx.OpenBinary(file.Data)
	if err != nil {
		return nil, err
	}

	sheets, err := excel.ToSlice()
	if err != nil {
		return nil, err
	}

	stat := &binary.FileStat{Log: fmt.Sprintf("新增文件: %s", file.Name), HasDetail: true, Detail: &binary.FileStatDetail{Adds: uint32(len(sheets[0])), Updates: 0, Deletes: 0}}
	return utils.Marshal(stat)
}

func (t *Chaincode) makeUpdateFileStat(newFile *binary.File, oldFile *binary.File) ([]byte, error) {
	if oldFile == nil {
		stat := &binary.FileStat{Log: fmt.Sprintf("更新文件: %s", newFile.Name), HasDetail: false}
		statBytes, err := utils.Marshal(stat)
		return statBytes, err
	}

	if len(newFile.Index) == 0 {
		return nil, errors.Errorf("Excel file %s does not contains any unique index", newFile.Name)
	}

	a, u, d, err := diff.Do(newFile.Data, oldFile.Data, newFile.Index)
	if err != nil {
		return nil, err
	}

	stat := &binary.FileStat{Log: fmt.Sprintf("更新文件: %s", newFile.Name), HasDetail: true, Detail: &binary.FileStatDetail{Adds: a, Updates: u, Deletes: d}}
	statBytes, err := utils.Marshal(stat)
	return statBytes, err
}

func (t *Chaincode) saveFileStatSet(stub shim.ChaincodeStubInterface, data *binary.FileStatSet) error {
	setBytes, err := utils.Marshal(data)
	if err != nil {
		return err
	}

	err = stub.PutState(common.FileStatKey(), setBytes)
	return err
}
