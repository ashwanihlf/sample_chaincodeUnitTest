# Hyperledger Fabric chaincode Unit Testing Sample

Pre Requisite - Go Lang is installed, GOROOT and GOPATH environment variable are set
Note : This chaincode is developed on Fabric 1.4 based shim package, so make sure you have right package, alternatively
you can do following before executing the actual steps

> mkdir -p $GOPATH/src/github.com/hyperledger
> cd $GOPATH/src/github.com/hyperledger
> git clone -b release-1.4 https://github.com/hyperledger/fabric.git

# Following are the steps to run the setup
1. create a working folder, change directory to working folder
2. git clone https://github.com/ashwanihlf/sample_chaincodeUnitTest.git
3. cd sample_chaincodeUnitTest  
5. go test -v --tags nopkcs11

you should see Run and Pass /Fail information of chaincode
