package main

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/binary/upper"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(upper.Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
