package infura

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	smartContractSddress = "0xe8aCaaB95d1102D099F82F03f6106289ee19abA8"
)

func TestInfuraGetEthLastBlock(t *testing.T) {

	latestBlock, lb, err := GetEthLastBlock()
	t.Log("latestBlock", latestBlock, "latest")

	assert.NoError(t, err)
	assert.NotNil(t, lb)
	assert.NotNil(t, latestBlock)
}

func TestInfuraGetStorageAt(t *testing.T) {
	data, err := GetStorageAt(smartContractSddress, "latest")
	t.Log("data", data)

	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestInfuraGetTransactionCount(t *testing.T) {
	nonce, err := GetTransactionCount(smartContractSddress, "latest")
	t.Log("transactions for address", nonce)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, nonce, 0)
}
