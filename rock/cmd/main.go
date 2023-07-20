package main

import (
	"fmt"

	"git.querycap.com/cloudchain/chaincode/rock"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(rock.Contract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
	}
}
