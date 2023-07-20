package main

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/forest"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(forest.Asset))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
