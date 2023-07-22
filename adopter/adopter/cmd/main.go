package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(adopter.Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
