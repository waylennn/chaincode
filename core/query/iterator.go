package query

import (
	"git.querycap.com/cloudchain/chaincode/common"
	"github.com/golang/protobuf/proto"
	commonledger "github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/plugins/gscc/protos"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"github.com/pkg/errors"
)

var logger = shim.NewLogger("binary/core/query")

type resultType uint8

// consts
const (
	StateQueryResult resultType = iota + 1
	HistoryQueryResult
)

// NewHistoryQueryIterator ...
func NewHistoryQueryIterator(stub shim.ChaincodeStubInterface, ccid, key string) (*HistoryQueryIterator, error) {
	invokeArgs := [][]byte{[]byte(common.GSCCGetHistoryForKey), []byte(ccid), []byte(key)}
	response := stub.InvokeChaincode(common.GSCC, invokeArgs, stub.GetChannelID())
	if response.Status != shim.OK {
		logger.Error("Failed to get history for key %s, error %s", key, response.Message)
		return nil, errors.New(response.Message)
	}

	payload := &protos.RangeQueryResponse{}
	err := proto.Unmarshal(response.Payload, payload)
	if err != nil {
		logger.Errorf("unmarshall error: %s", err)
		return nil, errors.Errorf("unmarshall error: %s", err)
	}

	return &HistoryQueryIterator{CommonIterator: &CommonIterator{payload, 0, stub}}, nil
}

// HistoryQueryIterator allows a chaincode to iterator over a set of
// key/value pairs returned by a history query.
type HistoryQueryIterator struct {
	*CommonIterator
}

// Next returns the next key and value in the history query iterator.
func (iter *HistoryQueryIterator) Next() (*queryresult.KeyModification, error) {
	result, err := iter.nextResult(HistoryQueryResult)
	if err != nil {
		return nil, err
	}
	return result.(*queryresult.KeyModification), nil
}

// CommonIterator allows a chaincode to check whether any more result
// to be fetched from an iterator and close it when done.
type CommonIterator struct {
	response   *protos.RangeQueryResponse
	currentLoc int
	stub       shim.ChaincodeStubInterface
}

func (iter *CommonIterator) nextResult(rType resultType) (commonledger.QueryResult, error) {
	if iter.currentLoc < len(iter.response.Results) {
		queryResult, err := iter.getResultFromBytes(iter.response.Results[iter.currentLoc], rType)
		if err != nil {
			logger.Errorf("Failed to decode query results: %s", err)
			return nil, err
		}
		iter.currentLoc++

		if iter.currentLoc == len(iter.response.Results) && iter.response.HasMore {
			if err = iter.fetchNextQueryResult(); err != nil {
				logger.Errorf("Failed to fetch next results: %s", err)
				return nil, err
			}
		}

		return queryResult, nil
	} else if !iter.response.HasMore {
		return nil, errors.New("no such key")
	}

	return nil, errors.New("invalid iterator state")
}

func (iter *CommonIterator) fetchNextQueryResult() error {
	invokeArgs := [][]byte{[]byte(common.GSCCQueryStateNext), []byte(iter.response.TxId), []byte(iter.response.Id)}
	response := iter.stub.InvokeChaincode(common.GSCC, invokeArgs, iter.stub.GetChannelID())
	if response.Status == shim.OK {
		queryResponse := &protos.RangeQueryResponse{}
		err := proto.Unmarshal(response.Payload, queryResponse)
		if err != nil {
			return err
		}
		iter.currentLoc = 0
		iter.response = queryResponse
		return nil
	}

	return errors.New(response.Message)
}

func (iter *CommonIterator) getResultFromBytes(queryResultBytes *protos.RangeQueryResultBytes, rType resultType) (commonledger.QueryResult, error) {
	if rType == StateQueryResult {
		return nil, errors.New("not implemented")
	} else if rType == HistoryQueryResult {
		historyQueryResult := &queryresult.KeyModification{}
		if err := proto.Unmarshal(queryResultBytes.ResultBytes, historyQueryResult); err != nil {
			return nil, err
		}
		return historyQueryResult, nil
	}
	return nil, errors.New("wrong result type")
}

// HasNext returns true if the range query iterator contains additional keys
// and values
func (iter *CommonIterator) HasNext() bool {
	if iter.currentLoc < len(iter.response.Results) || iter.response.HasMore {
		return true
	}
	return false
}

// Close closes the iterator.
func (iter *CommonIterator) Close() error {
	invokeArgs := [][]byte{[]byte(common.GSCCQueryStateClose), []byte(iter.response.TxId), []byte(iter.response.Id)}
	response := iter.stub.InvokeChaincode(common.GSCC, invokeArgs, iter.stub.GetChannelID())
	if response.Status != shim.OK {
		return errors.New(response.Message)
	}
	return nil
}
