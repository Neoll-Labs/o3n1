package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"ether-state/app"
	"ether-state/cmd/ether-stated/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
