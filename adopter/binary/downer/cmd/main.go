package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/waylennn/chaincode/adopter/binary/downer"
)

func main() {
	err := shim.Start(new(downer.Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
