package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/gorilla/websocket"
	"github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

var done chan interface{}
var exit chan os.Signal

// receiveHandler is the handler for receiving messages from the filtered events.
func receiveHandler(addressKey, query string, manageAddress chan string, connection *websocket.Conn) {
	//defer close(done)
	for {
		messageType, msg, err := connection.ReadMessage()
		_ = messageType
		if msg == nil {
			continue
		}
		var e types.RPCResponse
		err = json.Unmarshal(msg, &e)
		if err != nil {
			log.Println("Error in receive:", err, msg)
		}

		var r coretypes.ResultEvent
		err = json.Unmarshal(e.Result, &r)
		if err != nil {
			log.Println("Error in receive:", err, e.Result)
		}

		if v, exists := r.Events[addressKey]; exists {
			const ethAddressPosition = 0
			log.Println("manage ", query, "  ", v[ethAddressPosition])
			manageAddress <- v[ethAddressPosition]
		}
	}
}

// Subscribe subscribes events from chain.
func Subscribe(addressKey, query string, addressChn chan string) {
	done = make(chan interface{}) // Channel to indicate that the receiverHandler is done
	exit = make(chan os.Signal)   // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(exit, os.Interrupt) // Notify the interrupt channel for SIGINT

	socketUrl := "ws://localhost:26657/websocket"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}

	defer conn.Close()

	go receiveHandler(addressKey, query, addressChn, conn)

	body := fmt.Sprintf(`{"jsonrpc": "2.0","method": "subscribe", "id": 0,"params": {"query": "%s"}}`, query)
	err = conn.WriteMessage(websocket.TextMessage, []byte(body))
	if err != nil {
		log.Println("Error during writing to websocket:", err)
		return
	}

	for {
		select {
		case <-exit:
			// SIGINT (Ctrl + C) Terminate gracefully.
			log.Println("Received SIGINT interrupt signal.")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("Error during closing websocket: %w\n", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}
