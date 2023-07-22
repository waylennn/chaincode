package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/waylennn/chaincode/adopter/forest"
)

func main() {
	err := shim.Start(new(forest.Asset))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
