package main

import (
	"context"
	"fmt"
	"github.com/nelsonstr/o3n1/ether-state/relayer/cosmos"
	"github.com/nelsonstr/o3n1/ether-state/relayer/websocket"

	"os"
	"os/signal"
	"syscall"
)

func main() {

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go cosmos.ProcessEnableAddressMessages()
	go cosmos.ProcessDisableAddressMessages()

	go websocket.Subscribe("etherstate.etherstate.MsgDisableEthAddress.address",
		"message.action = '/etherstate.etherstate.MsgDisableEthAddress'",
		cosmos.DisableAddressChannel)
	go websocket.Subscribe("etherstate.etherstate.MsgEnableEthAddress.address",
		"message.action = '/etherstate.etherstate.MsgEnableEthAddress'",
		cosmos.EnableAddressChannel)

	client := cosmos.NewEtherStateClient(context.Background())
	cosmos.LoadActiveAddresses(client)

	go cosmos.RelayerEthCosmos(client)

	fmt.Println("Press Ctrl+C to exit.")
	<-exit

	fmt.Println("Shutting down...")

	os.Exit(0)
}
