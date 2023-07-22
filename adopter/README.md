# chaincode

链码程序，用户数据上下链操作

./peer.sh chaincode list --installed

./peer.sh chaincode list --instantiated -C cloudchain

./peer.sh chaincode install -p github.com/waylennn/chaincode/adopter/github.com/waylennn/chaincode/adopter/cmd -n adopter -v 0.6.0

./peer.sh chaincode upgrade -o orderer0.cloudchain-dev.querycap.com:7050 -C cloudchain -n adopter -l golang -v 0.6.0 -c '{"Args":[]}' --tls true --cafile /data/gopath/src/git.querycap.com/cloudchain/docker-compose-files/fabric/crypto-config/ordererOrganizations/cloudchain-dev.querycap.com/orderers/orderer0.cloudchain-dev.querycap.com/msp/tlscacerts/tlsca.cloudchain-dev.querycap.com-cert.pem
