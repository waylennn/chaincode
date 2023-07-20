package main

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/binary/downer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(downer.Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
