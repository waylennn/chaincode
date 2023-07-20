proto-forest:
	protoc --proto_path=protos/forest --go_out=protos/forest forest.proto

build-rock:
	go build -o cmd-rock git.querycap.com/cloudchain/chaincode/rock/cmd

build-forest:
	go build -o cmd-forest git.querycap.com/cloudchain/chaincode/forest/cmd

build-adopter:
	go build -o cmd-adopter git.querycap.com/cloudchain/chaincode/adopter/cmd
