package lottery

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"sort"

	"git.querycap.com/cloudchain/chaincode/protos/lottery"
	"github.com/sirupsen/logrus"
)

func generateKey(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

func selectWinLottery(sortCodes []string, sortCodeValues []float64) ([]string, error) {
	if len(sortCodes) < 2 {
		return sortCodes, nil
	}
	for _, item := range sortCodeValues {
		if item < 0 || item > 1 {
			return nil, fmt.Errorf("sortCodeValues:%+v has invalid value:%+v", sortCodeValues, item)
		}
	}
	matchKey := 1
	val := int64(sortCodeValues[0] * math.Pow(10, 16))

	for matchKey < len(sortCodeValues) {
		tmp := int64(sortCodeValues[matchKey] * math.Pow(10, 16))
		if tmp != val {
			break
		}
		matchKey++
	}
	return sortCodes[:matchKey], nil
}

func sortHash(target string, input []string) ([]string, []float64, error) {
	var tmp []float64
	m := make(map[float64][]string)

	if len(input) < 1 {
		return nil, nil, fmt.Errorf("empty input")
	}

	for _, v := range input {
		cs, err := cos(b2f(target), b2f(v))
		if err != nil {
			logrus.Errorf("calc similarity between %s and %s fail: %v\n", target, v, err)
			return nil, nil, err
		}
		if _, ok := m[cs]; ok {
			m[cs] = append(m[cs], v)
		} else {
			m[cs] = []string{v}
			tmp = append(tmp, cs)
		}
	}

	var resultS []string
	var resultF []float64
	sort.Float64s(tmp)
	for i := len(tmp) - 1; i >= 0; i-- {
		for _, v := range m[tmp[i]] {
			resultS = append(resultS, v)
			resultF = append(resultF, tmp[i])
		}
	}

	return resultS, resultF, nil
}

func cos(a []float64, b []float64) (cosSimilarity float64, err error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("vectors should not be nil (all zeros)")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

// byte to float64
func b2f(s string) []float64 {
	b := []byte(s)
	var f []float64
	for _, v := range b {
		f = append(f, float64(v))
	}
	return f
}

func convertTOMapForParty(data []*lottery.PartyKV) map[string]*lottery.PartyLottery {
	resp := make(map[string]*lottery.PartyLottery, 0)
	for _, v := range data {
		resp[v.Key] = v.Value
	}
	return resp
}

func convertTOMapForPriority(data []*lottery.PriorityPartyKV) map[string]*lottery.ListOfRound {
	resp := make(map[string]*lottery.ListOfRound, 0)
	for _, v := range data {
		resp[v.Key] = v.Value
	}
	return resp
}

func convertTOMapForWin(data []*lottery.WinLotteryKV) map[string]*lottery.WinLottery {
	resp := make(map[string]*lottery.WinLottery, 0)
	for _, v := range data {
		resp[v.Key] = v.Value
	}
	return resp
}

func convertTOArrayForParty(data map[string]*lottery.PartyLottery) []*lottery.PartyKV {
	sortStr := make([]string, 0)
	for k := range data {
		sortStr = append(sortStr, k)
	}
	sort.Strings(sortStr)
	resp := make([]*lottery.PartyKV, 0)
	for _, v := range sortStr {
		resp = append(resp, &lottery.PartyKV{
			Key:   v,
			Value: data[v],
		})
	}

	return resp
}

func convertTOArrayForPriority(data map[string]*lottery.ListOfRound) []*lottery.PriorityPartyKV {
	sortStr := make([]string, 0)
	for k := range data {
		sortStr = append(sortStr, k)
	}
	sort.Strings(sortStr)
	resp := make([]*lottery.PriorityPartyKV, 0)
	for _, v := range sortStr {
		resp = append(resp, &lottery.PriorityPartyKV{
			Key:   v,
			Value: data[v],
		})
	}

	return resp
}

func convertTOArrayForWin(data map[string]*lottery.WinLottery) []*lottery.WinLotteryKV {
	sortStr := make([]string, 0)
	for k := range data {
		sortStr = append(sortStr, k)
	}
	sort.Strings(sortStr)
	resp := make([]*lottery.WinLotteryKV, 0)
	for _, v := range sortStr {
		resp = append(resp, &lottery.WinLotteryKV{
			Key:   v,
			Value: data[v],
		})
	}

	return resp
}
