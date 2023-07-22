module github.com/waylennn/chaincode/adopter


replace (
	github.com/Shopify/sarama v1.40.0 => github.com/IBM/sarama v1.40.0
	github.com/hyperledger/fabric v2.1.1+incompatible => github.com/hyperledger/fabric v1.4.7
)

go 1.13


