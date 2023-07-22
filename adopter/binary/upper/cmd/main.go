package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/waylennn/chaincode/adopter/binary/upper"
)

func main() {
	err := shim.Start(new(upper.Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
