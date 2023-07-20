package lottery

import (
	"encoding/json"
	"fmt"
	"strconv"

	"git.querycap.com/cloudchain/chaincode/protos/lottery"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// These are function names from Invoke first parameter
const (
	AddParticipator     = "AddParticipator"
	NewLottery          = "NewLottery"
	StartLottery        = "StartLottery"
	QueryParticipator   = "QueryParticipator"
	DelParticipator     = "DelParticipator"
	QueryLotteryData    = "QueryLotteryData"
	QueryLotteryHistory = "QueryLotteryHistory"
)

// Key prefix
const (
	ParticipatorPollPrefix = "Party"
	LotteryPoolPrefix      = "Lottery"
	LotteryHistoryPrefix   = "History"
)

var logger = shim.NewLogger("lottery")

// Chaincode
type Chaincode struct{}

// Init 实现 Chaincode interface Init 方法
func (s *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Init chaincode adopter")
	return shim.Success(nil)
}

// Invoke 实现 Chaincode interface Invoke 方法
func (s *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()

	logger.Infof("Invoke function: %s with args %+v ", fn, args)

	switch fn {
	case AddParticipator:
		return addParticipator(stub, args)
	case NewLottery:
		return newLottery(stub, args)
	case StartLottery:
		return startLottery(stub, args)
	case QueryParticipator:
		return queryParticipator(stub, args)
	case DelParticipator:
		return delParticipator(stub, args)
	case QueryLotteryData:
		return queryLotteryData(stub, args)
	case QueryLotteryHistory:
		return queryLotteryHistory(stub, args)

	}

	return shim.Error(fmt.Sprintf("Requested function %s not found.", fn))
}

// addParticipator func add participator for lottery, and calculate hash code for every participator according the input
// args[0] is id string
// args[1] is name string
// args[2] is blessing word string
// args[3] is time string
func addParticipator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error(fmt.Sprintf("need 4 args"))
	}

	idStr := args[0]
	name := args[1]
	word := args[2]
	timeStr := args[3]

	lotteryCode := generateKey(word + timeStr + idStr + name)

	data := &lottery.Participator{
		Id:          idStr,
		Name:        name,
		Word:        word,
		Time:        timeStr,
		LotteryCode: lotteryCode,
	}
	body, err := proto.Marshal(data)
	if err != nil {
		return shim.Error(fmt.Sprintf("add user %s fail: proto marshal participator fail %s", idStr, err.Error()))
	}

	key, err := stub.CreateCompositeKey(ParticipatorPollPrefix, []string{idStr})
	if err != nil {
		return shim.Error(fmt.Sprintf("add user %s fail: create composite key error %s", idStr, err.Error()))
	}
	err = stub.PutState(key, body)
	if err != nil {
		return shim.Error(fmt.Sprintf("add user %s fail:  put state error %s", idStr, err.Error()))
	}
	return shim.Success([]byte(lotteryCode))
}

// newLottery func create a new lottery
// args[0] is the number of lottery
func newLottery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("need 1 args"))
	}

	lotteryNumStr := args[0]

	_, err := strconv.Atoi(lotteryNumStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s new lottery fail: %s", lotteryNumStr, err.Error()))
	}

	ParticipatorGroup, err := stub.GetStateByPartialCompositeKey(ParticipatorPollPrefix, []string{})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s new lottery fail: get state by composite key error %s", lotteryNumStr, err.Error()))
	}
	defer ParticipatorGroup.Close()

	tmp := make(map[string]*lottery.PartyLottery)

	for ParticipatorGroup.HasNext() {
		participatorKV, err := ParticipatorGroup.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("%s new lottery fail: next error %s", lotteryNumStr, err.Error()))
		}
		participator := &lottery.Participator{}
		err = proto.Unmarshal(participatorKV.Value, participator)
		if err != nil {
			return shim.Error(fmt.Sprintf("%s new lottery fail: promo unmarshal error %s", lotteryNumStr, err.Error()))
		}

		partyLottery := &lottery.PartyLottery{
			Id: participator.Id,
		}

		tmp[participator.LotteryCode] = partyLottery
	}

	// Parties is the pool parties for lottery, the party's code is the key
	// PriorityParties is the priority pool parties for the assign round lottery
	//    only when the current round lottery get more one parties, then put those parties to this map;key is the round number
	// WinLottery is the map for win lottery
	lotteryPool := &lottery.LotteryStoreData{
		Parties:         convertTOArrayForParty(tmp),
		PriorityParties: make([]*lottery.PriorityPartyKV, 0),
		WinLottery:      make([]*lottery.WinLotteryKV, 0),
	}

	value, err := proto.Marshal(lotteryPool)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s new lottery fail: proto marshal error %s", lotteryNumStr, err.Error()))
	}

	key, err := stub.CreateCompositeKey(LotteryPoolPrefix, []string{lotteryNumStr})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s new lottery fail: create composite key error %s", lotteryNumStr, err.Error()))
	}

	err = stub.PutState(key, value)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s new lottery fail: put state error %s", lotteryNumStr, err.Error()))
	}

	return shim.Success(nil)
}

// startLottery func draw a lottery by assign the lottery number and the round number, and a word for win
// args[0] is the lottery number
// args[1] is the round number
// args[2] is a word for win
func startLottery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error(fmt.Sprintf("need 3 args"))
	}

	lotteryNumStr := args[0]
	roundNumStr := args[1]
	winWord := args[2]

	winCode := generateKey(winWord)

	// lottery number must be a int
	_, err := strconv.Atoi(lotteryNumStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	// round number must be a int
	_, err = strconv.Atoi(roundNumStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	key, err := stub.CreateCompositeKey(LotteryPoolPrefix, []string{lotteryNumStr})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	if value == nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: has no parties", lotteryNumStr, roundNumStr, winCode))
	}

	lotteryStoreData := &lottery.LotteryStoreData{}
	err = proto.Unmarshal(value, lotteryStoreData)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	lotteryData := &lottery.LotteryData{
		Parties:         convertTOMapForParty(lotteryStoreData.Parties),
		PriorityParties: convertTOMapForPriority(lotteryStoreData.PriorityParties),
		WinLottery:      convertTOMapForWin(lotteryStoreData.WinLottery),
	}

	// init lotteryData nil elements
	if lotteryData.Parties == nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: parties is empty", lotteryNumStr, roundNumStr, winCode))
	}
	if lotteryData.PriorityParties == nil {
		lotteryData.PriorityParties = make(map[string]*lottery.ListOfRound, 0)
	}
	if lotteryData.WinLottery == nil {
		lotteryData.WinLottery = make(map[string]*lottery.WinLottery, 0)
	}

	// select lottery parties
	// priority select from PriorityParties if the round number is exist,
	// because when get not only one win lottery, set those parties to the PriorityParties
	partyCodes := make([]string, 0)
	if _, exist := lotteryData.PriorityParties[roundNumStr]; exist {
		count := len(lotteryData.PriorityParties[roundNumStr].Round)
		if count != 0 {
			partyCodes = lotteryData.PriorityParties[roundNumStr].Round[count-1].Value
		}
	} else {
		for lotteryCode := range lotteryData.Parties {
			partyCodes = append(partyCodes, lotteryCode)
		}
	}

	if len(partyCodes) == 0 {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: parties is empty", lotteryNumStr, roundNumStr, winCode))
	}

	// lottery calculate
	sortCodes, sortCodeValues, err := sortHash(winCode, partyCodes)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: calculate lottery error %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	winPartyCodes, err := selectWinLottery(sortCodes, sortCodeValues)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: calculate lottery error %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	// record lottery
	err = recordLottery(stub, lotteryNumStr, roundNumStr, winWord, winCode, sortCodes, sortCodeValues)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: record lottery error %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	resp := make([]lottery.WinLottery, 0)

	if len(winPartyCodes) == 0 {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: get zero win party ", lotteryNumStr, roundNumStr, winCode))
	} else if len(winPartyCodes) == 1 {
		// get only one win party, add the party to WinLottery,
		// and delete the party from Parties,
		winLotteryCode := winPartyCodes[0]
		_, exist := lotteryData.Parties[winLotteryCode]
		if !exist {
			return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: win party %s not exist", lotteryNumStr, roundNumStr, winCode, winLotteryCode))
		}

		// add the party to WinLottery
		win := &lottery.WinLottery{
			Id:          lotteryData.Parties[winLotteryCode].Id,
			LotteryCode: winLotteryCode,
			WinWord:     winWord,
			WinCode:     winCode,
		}
		lotteryData.WinLottery[roundNumStr] = win

		// delete the party from Parties
		delete(lotteryData.Parties, winLotteryCode)

		resp = append(resp, *win)
	} else {
		// get not only one win party, set those parties to PriorityParties
		for _, winLotteryCode := range winPartyCodes {
			_, exist := lotteryData.Parties[winLotteryCode]
			if !exist {
				return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: win party %s not exist", lotteryNumStr, roundNumStr, winCode, winLotteryCode))
			}

			resp = append(resp, lottery.WinLottery{
				Id:          lotteryData.Parties[winLotteryCode].Id,
				LotteryCode: winLotteryCode,
				WinWord:     winWord,
				WinCode:     winCode,
			})
		}

		if _, exist := lotteryData.PriorityParties[roundNumStr]; !exist {
			lotteryData.PriorityParties[roundNumStr] = &lottery.ListOfRound{
				Round: make([]*lottery.ListOfString, 0),
			}
		}

		lotteryData.PriorityParties[roundNumStr].Round = append(lotteryData.PriorityParties[roundNumStr].Round,
			&lottery.ListOfString{
				Value: winPartyCodes,
			})
	}

	newLotteryStoreData := &lottery.LotteryStoreData{
		Parties:         convertTOArrayForParty(lotteryData.Parties),
		PriorityParties: convertTOArrayForPriority(lotteryData.PriorityParties),
		WinLottery:      convertTOArrayForWin(lotteryData.WinLottery),
	}

	// restore lotteryData
	newValue, err := proto.Marshal(newLotteryStoreData)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	err = stub.PutState(key, newValue)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s-%s-%s start lottery fail: json marshal error %s", lotteryNumStr, roundNumStr, winCode, err.Error()))
	}
	return shim.Success(respBody)
}

// recordLottery record the process for every round lottery
func recordLottery(stub shim.ChaincodeStubInterface, lotteryNumStr, roundNumStr, winWord, winCode string, sortCode []string, sortValue []float64) error {
	data := &lottery.LotteryRecord{
		WinWord:   winWord,
		WinCode:   winCode,
		SortCode:  sortCode,
		SortValue: sortValue,
	}

	key, err := stub.CreateCompositeKey(LotteryHistoryPrefix, []string{lotteryNumStr, roundNumStr})
	if err != nil {
		return err
	}

	value, err := stub.GetState(key)
	if err != nil {
		return err
	}

	his := &lottery.LotteryHistory{}
	if value == nil {
		his.LotteryHistory = []*lottery.LotteryRecord{data}
	} else {
		err := proto.Unmarshal(value, his)
		if err != nil {
			return err
		}
		his.LotteryHistory = append(his.LotteryHistory, data)
	}

	body, err := proto.Marshal(his)
	if err != nil {
		return err
	}

	err = stub.PutState(key, body)
	if err != nil {
		return err
	}
	return nil
}

// queryParticipator func query party info
// args[0] is the party id string
// if the party id string is empty, then query all parties
func queryParticipator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("need 1 args"))
	}

	idStr := args[0]

	resp := make([]lottery.Participator, 0)
	if idStr != "" {
		key, err := stub.CreateCompositeKey(ParticipatorPollPrefix, []string{idStr})
		if err != nil {
			return shim.Error(fmt.Sprintf("query user %s fail: create composite key error %s", idStr, err.Error()))
		}
		value, err := stub.GetState(key)
		if err != nil {
			return shim.Error(fmt.Sprintf("query user %s fail:  put state error %s", idStr, err.Error()))
		}
		if value == nil {
			return shim.Success(nil)
		}

		participator := &lottery.Participator{}
		err = proto.Unmarshal(value, participator)
		if err != nil {
			return shim.Error(fmt.Sprintf("query user %s fail:  proto unmarshal error %s", idStr, err.Error()))
		}
		resp = append(resp, *participator)
	} else {
		ParticipatorGroup, err := stub.GetStateByPartialCompositeKey(ParticipatorPollPrefix, []string{})
		if err != nil {
			return shim.Error(fmt.Sprintf("query user %s fail: get state by composite key error %s", idStr, err.Error()))
		}
		defer ParticipatorGroup.Close()

		for ParticipatorGroup.HasNext() {
			participatorKV, err := ParticipatorGroup.Next()
			if err != nil {
				return shim.Error(fmt.Sprintf("query user %s fail:  next error %s", idStr, err.Error()))
			}
			participator := &lottery.Participator{}
			err = proto.Unmarshal(participatorKV.Value, participator)
			if err != nil {
				return shim.Error(fmt.Sprintf("query user %s fail:  promo unmarshal error %s", idStr, err.Error()))
			}
			resp = append(resp, *participator)
		}
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		return shim.Error(fmt.Sprintf("query user %s fail:  json marshal error %s", idStr, err.Error()))
	}

	return shim.Success(respBody)
}

// delParticipator func delete party
// args[0] is the party id string
// if the party id string is empty, then delete all parties
func delParticipator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("need 1 args"))
	}

	idStr := args[0]

	if idStr != "" {
		key, err := stub.CreateCompositeKey(ParticipatorPollPrefix, []string{idStr})
		if err != nil {
			return shim.Error(fmt.Sprintf("del user %s fail: create composite key error %s", idStr, err.Error()))
		}
		err = stub.DelState(key)
		if err != nil {
			return shim.Error(fmt.Sprintf("del user %s fail:  put state error %s", idStr, err.Error()))
		}
	} else {
		ParticipatorGroup, err := stub.GetStateByPartialCompositeKey(ParticipatorPollPrefix, []string{})
		if err != nil {
			return shim.Error(fmt.Sprintf("del user %s fail: get state by composite key error %s", idStr, err.Error()))
		}
		defer ParticipatorGroup.Close()

		for ParticipatorGroup.HasNext() {
			participatorKV, err := ParticipatorGroup.Next()
			if err != nil {
				return shim.Error(fmt.Sprintf("del user %s fail:  next error %s", idStr, err.Error()))
			}
			err = stub.DelState(participatorKV.Key)
			if err != nil {
				return shim.Error(fmt.Sprintf("del user %s fail:  put state error %s", idStr, err.Error()))
			}
		}
	}
	return shim.Success(nil)
}

// queryLotteryData func query the lottery current data for the assign lottery number
// args[0] is the lottery number
func queryLotteryData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("need 1 args"))
	}
	lotteryNumStr := args[0]

	_, err := strconv.Atoi(lotteryNumStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s lottery fail: %s", lotteryNumStr, err.Error()))
	}

	key, err := stub.CreateCompositeKey(LotteryPoolPrefix, []string{lotteryNumStr})
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s lottery fail: %s", lotteryNumStr, err.Error()))
	}
	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s lottery fail: %s", lotteryNumStr, err.Error()))
	}
	if value == nil {
		return shim.Success(nil)
	}

	lotteryData := &lottery.LotteryStoreData{}
	err = proto.Unmarshal(value, lotteryData)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s lottery fail: proto unmarshal error %s", lotteryNumStr, err.Error()))
	}

	resp, err := json.Marshal(lotteryData)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s lottery fail: json marshal error %s", lotteryNumStr, err.Error()))
	}
	return shim.Success(resp)
}

// queryLotteryHistory func query the lottery history process for the assign lottery number and round number
// args[0] is the lottery number
// args[1] is the round number
// the resp include the calculate result
func queryLotteryHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error(fmt.Sprintf("need 1 args"))
	}
	lotteryNumStr := args[0]
	roundNumStr := args[1]

	_, err := strconv.Atoi(lotteryNumStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s-%s lottery histroy fail: %s", lotteryNumStr, roundNumStr, err.Error()))
	}

	_, err = strconv.Atoi(roundNumStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s-%s lottery histroy fail: %s", lotteryNumStr, roundNumStr, err.Error()))
	}

	key, err := stub.CreateCompositeKey(LotteryHistoryPrefix, []string{lotteryNumStr, roundNumStr})
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s-%s lottery histroy fail: %s", lotteryNumStr, roundNumStr, err.Error()))
	}
	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s-%s lottery histroy fail: %s", lotteryNumStr, roundNumStr, err.Error()))
	}
	if value == nil {
		return shim.Success(nil)
	}

	lotteryHistory := &lottery.LotteryHistory{}
	err = proto.Unmarshal(value, lotteryHistory)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s-%s lottery history fail: proto unmarshal error %s", lotteryNumStr, roundNumStr, err.Error()))
	}

	resp, err := json.Marshal(lotteryHistory)
	if err != nil {
		return shim.Error(fmt.Sprintf("query %s-%s lottery history fail: json marshal error %s", lotteryNumStr, roundNumStr, err.Error()))
	}
	return shim.Success(resp)
}
