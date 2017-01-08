/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)


const   DONOR      =  "donot"
const   CHARITY   =  "charity"
const   ACCEPTOR =  "acceptor"

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

err := stub.PutState("initial_amount", []byte("10000"))
    if err != nil {
        return nil, err
    }
	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "fund_charity"{
		res, err := t.fund_charity(stub, args)
													//lets make sure all open trades are still valid
		return res, err
	} else if function == "set_user"{
		res, err := t.set_user(stub,args)
		return res, err
	} else if function == "distribute_charity"{
		res, err := t.distribute_charity(stub,args)
		return res, err
	}
	
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	 if function == "isValid" {

		return t.isValid(stub,args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) isValid(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var value1, value2 ,jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	value1 = args[0] //fund_amount
	value2 = args[1] //distribute_amount
	val1Asbytes, err := stub.GetState(value1)
	fmt.Println(" fun amount is:"+ value1)
	val2Asbytes, err := stub.GetState(value2)
	fmt.Println(" distribute amount is:"+ value2)
	val2str := string(val2Asbytes)
	fmt.Println(val2str)
										//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + value1 + "\"}"
		return nil, errors.New(jsonResp)
	}

	return val1Asbytes, nil													//send it onward
}

func (t *SimpleChaincode) fund_charity(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name, value string
    var err error
    fmt.Println("running write()")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    name = args[0]                            //fund_amount
    value = args[1]								//actual fund amount	
    err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}

func (t *SimpleChaincode) distribute_charity(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name, value string
    var err error
    fmt.Println("running write()")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }

    name = args[0]                            //distribute_amount
    value = args[1]								//actual distribute amount	
    err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}
// -------- set user with initial amount-------------
func (t *SimpleChaincode) set_user(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var name, value string
	/*
	//   0       1
	// "name", "bob" , "amount" , "10000"
	if len(args) < 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	fmt.Println("- start set user")
	fmt.Println(args[0] + " - " + args[1] + " - "+ args[2] + " - " + args[3])
	marbleAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New("Failed to get thing")
	}
	str := {"user:"+ args[1],"amount:"+args[3] }
	json.Unmarshal(marbleAsBytes, &str)										//un stringify it aka JSON.parse()
	res.User = args[1]														//change the user
	
	jsonAsBytes, _ := json.Marshal(res)
	err = stub.PutState(args[0], []byte(str))								//rewrite the marble with id as key
	if err != nil {
		return nil, err
	}
	*/
	if len(args) < 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	fmt.Println("- start set user")
	name = args[0] // initial_amount
	value = args[1]

	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}