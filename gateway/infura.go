/*
 license x
*/

package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

const (
	infuraURL = "https://goerli.infura.io/v3/c10ab2d73dc64469b269c4b3235c8130" // nls
)

type Request struct {
	JsonRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type response struct {
	ID     int        `json:"id"`
	Result string     `json:"result"`
	Error  *jsonError `json:"error"`
}

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInfuraCall          = errors.New("error calling infura service")
	ErrUnmarshallResponse  = errors.New("error unmarshall response")
	ErrInfuraErrorResponse = errors.New("infura return error")
)

func GetEthLastBlock() (float32, string, error) {
	response, err := rpcCall("eth_blockNumber", nil)
	if err != nil {
		return 0, "", ErrInfuraCall
	}

	blockNumber := new(big.Int)
	blockNumber.SetString(response.Result[2:], 16)
	return float32(blockNumber.Int64()), response.Result, nil
}

func GetTransactionCount(ethAddress, blockNumber string) (int, error) {
	response, err := rpcCall("eth_getTransactionCount", []interface{}{ethAddress, blockNumber})
	if err != nil {
		return 0, err
	}

	// Convert the hexadecimal balance to decimal
	result := new(big.Int)
	result.SetString(response.Result[2:], 16)

	return strconv.Atoi(result.String())
}

// GetStorageAt Returns the value from a storage position at a given address.
func GetStorageAt(smartContractAddress string, blockNumber string) (int, error) {
	response, err := rpcCall("eth_getStorageAt", []interface{}{smartContractAddress, "0x0", blockNumber})
	if err != nil {
		return 0, err
	}

	// Convert the hexadecimal balance to decimal
	result := new(big.Int)
	result.SetString(response.Result[2:], 16)

	return strconv.Atoi(result.String())
}

func rpcCall(method string, params []interface{}) (response, error) {
	request := &Request{
		JsonRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return response{}, ErrInfuraCall
	}
	// Create an HTTP POST request to the JSON-RPC server
	resp, err := http.Post(infuraURL, "application/json", bytes.NewReader(requestData))
	if err != nil {
		log.Println("Error sending JSON-RPC request:", err)
		return response{}, ErrInfuraCall
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	// Decode the JSON-RPC response
	var data response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("Error decoding JSON-RPC response:", err)

		return response{}, ErrUnmarshallResponse
	}

	// Check for errors in the response
	if data.Error != nil {
		log.Printf("JSON-RPC error %d: %s\n", data.Error.Code, data.Error.Message)

		return response{}, ErrInfuraErrorResponse
	}
	return data, nil
}
