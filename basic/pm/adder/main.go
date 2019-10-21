package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
var timeCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "timeCounter"})

func init() {
	prometheus.MustRegister(timeCounter)
}

func main() {
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		for {
			timeCounter.Inc()
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("Listening !")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
