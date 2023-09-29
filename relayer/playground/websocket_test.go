package playground

import (
	goctx "context"
	"fmt"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"reflect"
	"testing"
	"time"
)

const (
	endpoint     = "/websocket"
	hostname     = "127.0.0.1"
	httpAddress  = "http://" + hostname + ":"
	httpsAddress = "https://" + hostname + ":"
	wsAddress    = "ws://" + hostname + ":"
	wssAddress   = "wss://" + hostname + ":"
)

func httpNewWithClient(httpRemote string, wssRemote string) (errStatus error, errSocket error) {
	http, err := jsonrpcclient.DefaultHTTPClient(httpRemote)
	if err != nil {
		panic(err)
	}
	http.Timeout = time.Duration(2) * time.Second
	// with v0.35.9 the paramter endpoint has to be removed
	// _ = endpoint
	// rpc, err := rpchttp.NewWithClient(wssRemote, http)
	rpc, err := rpchttp.NewWithClient(wssRemote+endpoint, http)
	if err != nil {
		panic(err)
	}
	return tryRequests(rpc)
}

func tryRequests(rpc *rpchttp.HTTP) (errStatus error, errSocket error) {
	errSocket = nil
	errSocket = nil
	_, errStatus = rpc.Status(goctx.Background())
	errSocket = rpc.Start()
	if errSocket == nil {
		defer func() { _ = rpc.Stop() }()
	}
	return
}

type httpNewWithClientCaseArgs struct {
	httpRemote string
	httpPort   string
	wssRemote  string
	wsPort     string
}
type httpNewWithClientCase struct {
	name          string
	args          httpNewWithClientCaseArgs
	wantStatusErr error
	wantSocketErr error
}

func PermuteTestCases(addresses map[string]string, ports []string) []httpNewWithClientCase {
	var testCases []httpNewWithClientCase
	for httpUrlName, httpUrl := range addresses {
		for _, httpPort := range ports {
			for wsUrlName, wsUrl := range addresses {
				for _, wsPort := range ports {
					testCases = append(testCases,
						httpNewWithClientCase{
							fmt.Sprintf("http client: %s:%s and ws client: %s:%s", httpUrlName, httpPort, wsUrlName, wsPort),
							httpNewWithClientCaseArgs{
								httpUrl,
								httpPort,
								wsUrl,
								wsPort,
							},
							nil,
							nil,
						})
				}
			}
		}
	}
	return testCases
}

func Test_httpNewWithClient(t *testing.T) {
	addresses := map[string]string{
		"http":  httpAddress,
		"https": httpsAddress,
		"ws":    wsAddress,
		"wss":   wssAddress,
	}
	ports := []string{"26557"}
	for _, tt := range PermuteTestCases(addresses, ports) {
		t.Run(tt.name, func(t *testing.T) {
			statusErr, socketErr := httpNewWithClient(tt.args.httpRemote+tt.args.httpPort, tt.args.wssRemote+tt.args.wsPort)
			if !reflect.DeepEqual(statusErr, tt.wantStatusErr) {
				t.Errorf("httpNewWithClient() statusErr = %v, want %v", statusErr, tt.wantStatusErr)
			}
			if !reflect.DeepEqual(socketErr, tt.wantSocketErr) {
				t.Errorf("httpNewWithClient() socketErr = %v, want %v", socketErr, tt.wantSocketErr)
			}
		})
	}
}
