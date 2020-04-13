package main

import (
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


// Shows the use of shim.MockInit method to invoke init method of chaincode
func Test_MockInit(t *testing.T) {

	fmt.Println("--------Entering Test_MockInit--------")
	sampleCC := new(SampleChaincode)
	mockStub := shim.NewMockStub("mockstub", sampleCC)
	
	txId := "mockTxID"

	// First argument is uuid as String and second is argument, passing nil 
	response := mockStub.MockInit(txId,nil)
	
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Init test failed")
		t.FailNow()
	}
	fmt.Println("--------Exiting Test_MockInit--------")
}

 
// Shows the use of direct calling of init method of chaincode 
func Test_Init(t *testing.T) {

	fmt.Println("--------Entering Test_Init--------")
	sampleCC := new(SampleChaincode)
	mockStub := shim.NewMockStub("mockstub", sampleCC)
	txId := "mockTxID"

	mockStub.MockTransactionStart(txId)
	// Calling Init directly on chaincode then invoking through MockInit
	response := sampleCC.Init(mockStub)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Init test failed")
		t.FailNow()
	}
	fmt.Println("--------Exiting Test_Init--------")
}


func Test_InvokeInitCar(t *testing.T) {

    fmt.Println("--------Entering Test_InvokeInitCar--------")

    // Instantiate mockStub using SampleChainCode as the target chaincode to unit test
    stub := shim.NewMockStub("mockStub", new(SampleChaincode))
    if stub == nil {
        t.Fatalf("MockStub creation failed")
    }



    // Here we perform a "mock invoke" to invoke the function "initVehiclePart" method with associated parameters
    // parameters are uuid and arguments,
    result := stub.MockInvoke("001",
        [][]byte{[]byte("initCar"),
            []byte("Ashwani"),
            []byte("Blue"),
            []byte("BMW")})

    // We expect a shim.ok if all goes well
    if result.Status != shim.OK {
        t.Fatalf("Error Invoking method")
    }

    // calling the readCar on 'Ashwani' as we just persisted that 
	response := stub.MockInvoke("002",
        [][]byte{[]byte("readCar"),
            []byte("Ashwani")})
	
    if response.Status != shim.OK {
        t.Fatalf("Error in readCar")
    }
	
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("--------Exiting Test_InvokeInitCar--------")
}

// Calling initCar directly than MockInvoke
func Test_initCar(t *testing.T) {
	fmt.Println("--------Entering Test_initCar--------")
	sampleCC := new(SampleChaincode)
	mockStub := shim.NewMockStub("mockstub", sampleCC)
	txId := "mockTxID"

	args := []string{"Ashwani", "Blue", "BMW"}
	mockStub.MockTransactionStart(txId)
	response := sampleCC.initCar(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	

	if s := response.GetStatus(); s != 200 {
		fmt.Println("initCar test failed")
		t.FailNow()
	}
	
	fmt.Println("--------Entering Test_initCar--------")
}


