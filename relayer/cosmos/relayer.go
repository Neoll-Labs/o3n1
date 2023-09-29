package cosmos

import (
	"context"
	"github.com/nelsonstr/o3n1/ether-state/relayer/infura"
	"log"
	"time"
)

func RelayerEthCosmos(client Client) {

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {

		log.Println("wait")
		time.Sleep(10 * time.Second)
		log.Println("QueryAllEthereumAddressRequest")

		for address := range ethereumAddresss.address {
			log.Printf("address: %s \n", address)

			blocknumber, lb, err := infura.GetEthLastBlock()
			if err != nil {
				continue
			}

			state, err := infura.GetStorageAt(address, lb)
			if err != nil {
				continue
			}

			nonce, err := infura.GetTransactionCount(address, lb)
			if err != nil {
				continue
			}

			err = client.SaveEthereumAddressState(context.Background(), Alice, address, uint64(blocknumber), uint64(nonce), uint64(state))
			if err != nil {
				log.Printf("save state error %s\n", err)
			}
		}
	}
}
