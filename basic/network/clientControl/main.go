package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
)

func main() {

	// mp := make(map[string]http.Client)
	// nfInstanceId -> client[ endpoints ]
	//

	client := &http.Client{
		Transport: &http.Transport{},
	}

	// per endpoint one client

	doRequest(client)
	// client.CloseIdleConnections()
	doRequest(client)
}

func doRequest(client *http.Client) {
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn, reused : %v, idleTime : %v, wasIdle : %v, conn : %v\n", connInfo.Reused, connInfo.IdleTime, connInfo.WasIdle, &connInfo.Conn)
		},
		ConnectStart: func(network, addr string) {
			fmt.Printf("Dial start %v, %v\n", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("Dial done %v, %v, %v\n", network, addr, err)
		},
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
		WroteRequest: func(wr httptrace.WroteRequestInfo) {
			fmt.Println("Wrote request", wr)
		},
	}
	req, _ := http.NewRequest("GET", "https://www.taobao.com", nil)
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rsp.Body.Close()
	bt := make([]byte, 100)
	n, err := rsp.Body.Read(bt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Response:\n%v\n", string(bt))
	fmt.Printf("总共读取了 %v 字符\n", n)

}
