package cosmos

import (
	"context"
	"log"
	"sync"
	"time"
)

var (
	ethereumAddresss = struct {
		sync.Mutex
		address map[string]bool
	}{
		address: make(map[string]bool),
	}

	// Create a channel to receive messages from WebSocket connections
	EnableAddressChannel   = make(chan string)
	DisablewAddressChannel = make(chan string)
)

func PoolingAddress(client Client) {

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
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

			DisablewAddressChannel <- address.Index
		}
	}
}

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

func ProcessDisableAddressMessages() {
	for {
		msg := <-DisablewAddressChannel
		go func() {
			log.Printf("remove address %s\n", msg)
			ethereumAddresss.Mutex.Lock()
			delete(ethereumAddresss.address, msg)
			defer ethereumAddresss.Mutex.Unlock()
		}()
	}
}
