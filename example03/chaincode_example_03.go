package main

import (
        "errors"
        "fmt"
        "strconv"
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        )

type ChaincodeEX3 struct {
}

type Baggage struct {
    Item string
    Position string
    Temperature int
}


func (t *ChaincodeEX3) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if len(args) != 0 {
        return nil, errors.New("Incorrect number of arguments. Expecting 0")
    }
    
    return nil, nil
}

func (t *ChaincodeEX3) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var key string
    var err error
    
    if len(args) != 4 {
        return nil, errors.New("Incorrect number of arguments. Expecting 4")
    }
    
    key = args[0]
    item := args[1]
    position := args[2]
    temperature, err := strconv.Atoi(args[3])
    if err != nil {
        return nil, err
    }
    
    value := Baggage{item, position, temperature}

    valbytes, err := json.Marshal(value)
    if err != nil {
        return nil, err
    }
    
    
    err = stub.PutState(key, valbytes)
    if err != nil {
        return nil, err
    }
    
    
    return nil, nil
}


func (t *ChaincodeEX3) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    if function != "query" {
        return nil, errors.New("Invalid query function name. Expecting \"query\"")
    }
    var key string
    var value Baggage
    var err error
    
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
    }
    
    key = args[0]
    
    valbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    if valbytes == nil {
        jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
    
    err = json.Unmarshal(valbytes, &value)
    if err != nil {
        return nil, errors.New("Error")
    }
    
    item := value.Item
    position := value.Position
    temperature := strconv.Itoa(value.Temperature)
    if err != nil {
        return nil, errors.New("Error")
    }
    
    message := "{item:" + item + ", position:" + position + ", temperature:" + temperature + "}"
    return []byte(message), nil
}

func main() {
    err := shim.Start(new(ChaincodeEX3))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}





