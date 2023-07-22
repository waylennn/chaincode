module github.com/waylennn/chaincode/adopter

go 1.13

replace (
	github.com/Shopify/sarama v1.40.0 => github.com/IBM/sarama v1.40.0
	github.com/hyperledger/fabric v2.1.1+incompatible => github.com/rkcloudchain/fabric v1.4.4-0.20190902033921-b3f844e6eb44
)

require (
	github.com/Knetic/govaluate v3.0.0+incompatible // indirect
	github.com/Shopify/sarama v1.40.0 // indirect
	github.com/fsouza/go-dockerclient v1.9.7 // indirect
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hyperledger/fabric v2.1.1+incompatible
	github.com/hyperledger/fabric-amcl v0.0.0-20230602173724-9e02669dceb2 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.8 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/viper v1.16.0 // indirect
	github.com/stretchr/testify v1.8.4
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	github.com/tealeg/xlsx v1.0.5
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	google.golang.org/grpc v1.56.2 // indirect
)
