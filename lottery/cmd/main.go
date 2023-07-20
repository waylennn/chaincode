package main

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/lottery"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(lottery.Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
