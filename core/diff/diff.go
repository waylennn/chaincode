package diff

import (
	"encoding/json"
	"reflect"
	"sort"

	"git.querycap.com/cloudchain/chaincode/core/avl"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

// Do computes the (row oriented) modifications of the xlsx file.
func Do(current, old []byte, pk []string) (adds, updates, deletes uint32, err error) {
	oldXlsx, err := xlsx.OpenBinary(old)
	if err != nil {
		return 0, 0, 0, err
	}

	oldPKIndex, err := findPKIndex(oldXlsx.Sheets[0], pk)
	if err != nil {
		return 0, 0, 0, err
	}

	oldOutput, err := oldXlsx.ToSlice()
	if err != nil {
		return 0, 0, 0, err
	}

	tree, err := makeAVL(oldOutput[0][1:], oldPKIndex)
	if err != nil {
		return 0, 0, 0, err
	}

	currXlsx, err := xlsx.OpenBinary(current)
	if err != nil {
		return 0, 0, 0, err
	}

	newPKIndex, err := findPKIndex(currXlsx.Sheets[0], pk)
	if err != nil {
		return 0, 0, 0, err
	}

	currOutput, err := currXlsx.ToSlice()
	if err != nil {
		return 0, 0, 0, err
	}

	dic, err := makeMap(currOutput[0][1:], newPKIndex)
	if err != nil {
		return 0, 0, 0, err
	}

	a, u, d := diff(dic, tree)
	return a, u, d, nil
}

func diff(dic map[string][]string, tree *avl.Immutable) (adds, updates, deletes uint32) {
	var same uint32
	for key, value := range dic {
		ele := tree.Get(&entry{PK: []byte(key)})
		if ele == nil {
			adds = adds + 1
		} else {
			entry := ele.(*entry)
			if !reflect.DeepEqual(value, entry.Value) {
				updates = updates + 1
			} else {
				same = same + 1
			}
		}
	}

	deletes = uint32(tree.Len()) - updates - same
	return
}

func makeMap(data [][]string, pkIndex []int) (map[string][]string, error) {
	output := make(map[string][]string)
	for _, row := range data {
		pkarray := make([]string, len(pkIndex))
		value := row[:]

		for i, v := range pkIndex {
			pkarray[i] = row[v]
			value = remove(value, v)
		}

		pk, err := json.Marshal(pkarray)
		if err != nil {
			return nil, err
		}

		output[string(pk)] = value
	}

	return output, nil
}

func makeAVL(data [][]string, pkIndex []int) (*avl.Immutable, error) {

	tree := avl.NewImmutable()

	for _, row := range data {
		pkarray := make([]string, len(pkIndex))
		value := row[:]

		for i, v := range pkIndex {
			pkarray[i] = row[v]
			value = remove(value, v)
		}

		pk, err := json.Marshal(pkarray)
		if err != nil {
			return nil, err
		}

		ele := &entry{PK: pk, Value: value}
		tree.Insert(ele)
	}

	return tree, nil
}

func findPKIndex(sheet *xlsx.Sheet, pk []string) ([]int, error) {
	if len(sheet.Rows) == 0 {
		return nil, errors.Errorf("Sheet %s does not contain any rows", sheet.Name)
	}

	index := make([]int, 0)
	for _, key := range pk {
		for i, cell := range sheet.Rows[0].Cells {
			if cell.Value == key {
				index = append(index, i)
				break
			}
		}
	}

	if len(index) != len(pk) {
		return nil, errors.Errorf("Unable to find all unique indexes (%v) in the sheet", pk)
	}

	sort.Ints(index)
	return index, nil
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
