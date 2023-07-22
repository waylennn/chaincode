package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/waylennn/chaincode/adopter/rock"
)

func main() {
	err := shim.Start(new(rock.Contract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
