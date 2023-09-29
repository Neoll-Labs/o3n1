package cosmos

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/nelsonstr/o3n1/ether-state/relayer/infura"
)

const ethereumBlockTime = 10 * time.Second

var (
	ethereumAddresss = struct {
		sync.Mutex
		address map[string]bool
	}{
		address: make(map[string]bool),
	}

	// Create a channel to receive messages from WebSocket connections
	EnableAddressChannel  = make(chan string)
	DisableAddressChannel = make(chan string)
)

// LoadActiveAddresses loads the active addresses from the ethereum network only once in the startup.
func LoadActiveAddresses(client Client) {
	log.Println("QueryAllEthereumAddressRequest")
	resp, err := client.QueryAllEthereumAddress(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, address := range resp {
		log.Printf("address: %s, enable: %v\n", address.Index, address.Active)
		if address.Active {
			EnableAddressChannel <- address.Index
			continue
		}
	}

	for v, _ := range ethereumAddresss.address {
		log.Println("addresses actives ", v)
	}
}

// ProcessEnableAddressMessages add addresses to list of active address to monitor the state.
func ProcessEnableAddressMessages() {
	for {
		msg := <-EnableAddressChannel
		go func() {
			log.Printf("add address %s\n", msg)
			ethereumAddresss.Mutex.Lock()
			ethereumAddresss.address[msg] = true

			defer ethereumAddresss.Mutex.Unlock()
		}()
	}
}

// ProcessDisableAddressMessages removes an address from the list of active addresses.
func ProcessDisableAddressMessages() {
	for {
		msg := <-DisableAddressChannel
		go func() {
			log.Printf("remove address %s\n", msg)
			ethereumAddresss.Mutex.Lock()
			delete(ethereumAddresss.address, msg)

			defer ethereumAddresss.Mutex.Unlock()
		}()
	}
}

// RelayerEthCosmos fetch the data from Ethereum and save in blockchain.
// all the data is taken from the same block number.
func RelayerEthCosmos(client Client) {

	ticker := time.NewTicker(ethereumBlockTime)
	defer ticker.Stop()

	for range ticker.C {

		time.Sleep(ethereumBlockTime)
		log.Println("save state for: ", ethereumAddresss.address)
		for address := range ethereumAddresss.address {
			log.Printf("address: %s \n", address)

			blocknumber, lb, err := infura.GetEthLastBlock()
			if err != nil {
				log.Printf("infura.GetEthLastBlock error %w\n", err)
				continue
			}

			state, err := infura.GetStorageAt(address, lb)
			if err != nil {
				log.Printf("infura.GetStorageAt error %w\n", err)
				continue
			}

			nonce, err := infura.GetTransactionCount(address, lb)
			if err != nil {
				log.Printf("infura.GetTransactionCount error %w\n", err)
				continue
			}

			log.Printf("save state %s\n", address)
			err = client.SaveEthereumAddressState(context.Background(), Alice, address, uint64(blocknumber), uint64(nonce), uint64(state))
			if err != nil {
				log.Printf("save state error %s\n", err)
			}
		}
	}
}
