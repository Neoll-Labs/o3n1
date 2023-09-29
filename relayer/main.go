package main

import (
	"context"
	"fmt"
	"github.com/nelsonstr/o3n1/ether-state/relayer/cosmos"

	"os"
	"os/signal"
	"syscall"
)

func main() {

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()
	c := cosmos.NewEtherStateClient(ctx)

	go cosmos.PoolingAddress(c)

	go cosmos.ProcessEnableAddressMessages()
	go cosmos.ProcessDisableAddressMessages()

	go cosmos.RelayerEthCosmos(c)

	//client, _ := rpchttp.New("ws://localhost:26657/websocket")
	//err := client.Start()
	//if err != nil {
	//	log.Println(err)
	//}
	//defer client.Stop()
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//query := "tm.event = 'Tx' AND tx.height = 3"
	//txs, err := client.Subscribe(ctx, "test-client", query)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//go func() {
	//	for e := range txs {
	//		fmt.Println("got ", e.Data.(types.EventDataTx))
	//	}
	//}()

	fmt.Println("Press Ctrl+C to exit.")
	<-exit

	fmt.Println("Shutting down...")

	os.Exit(0)
}
