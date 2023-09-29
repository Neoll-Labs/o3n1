package infura

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const (
	smartContractSddress = "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8"
)

func TestInfuraCalls(t *testing.T) {

	latestBlock, lb, err := GetEthLastBlock()
	log.Println("latestBlock", latestBlock, lb)
	assert.NoError(t, err)

	data, err := GetStorageAt(smartContractSddress, lb)
	log.Println("data", data)
	assert.NoError(t, err)

	nonce, err := GetTransactionCount(smartContractSddress, lb)
	log.Println("transactions for address", nonce)
	assert.NoError(t, err)

}
