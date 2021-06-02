package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	mu                   sync.Mutex
	rsp                  []byte
	CurrentRequestNumber = int32(0)
	RequestCount         = int32(0)
)

type TestHandler struct{}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	atomic.AddInt32(&CurrentRequestNumber, 1)
	defer atomic.AddInt32(&CurrentRequestNumber, -1)
	atomic.AddInt32(&RequestCount, 1)

	// log.Println(RequestCount)

	if RequestCount%100 == 0 {
		// log.Message(log.Info, "Current Request Number", "TotalRequestCount", RequestCount, "CurrentWaitingRequestNumber", CurrentRequestNumber)
		log.Printf("CurrentRequestNumber : %v, Total Request : %v\n", CurrentRequestNumber, RequestCount)
	}

	time.Sleep(500 * time.Millisecond)

	st := strconv.Itoa(int(RequestCount))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(st))
}

func main() {
	// rsp = make([]byte, 1000)
	// for i := 0; i < 1000; i++ {
	// 	rsp[i] = byte(int('a') + i%26)
	// }

	h2s := &http2.Server{}
	port := "9001"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(&TestHandler{}, h2s),
	}

	server.ListenAndServe()

}
