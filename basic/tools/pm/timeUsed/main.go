package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	timeUsed := prometheus.NewHistogram(prometheus.HistogramOpts{Name: "TimeUsed"})
	timeUsedNorm := prometheus.NewHistogram(prometheus.HistogramOpts{Name: "TimeUsedNorm"})
	prometheus.MustRegister(timeUsed)
	prometheus.MustRegister(timeUsedNorm)
	go func() {
		for {
			timeUsed.Observe(rand.Float64() * 10)
			timeUsedNorm.Observe(rand.NormFloat64() * 10)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on : http://localhost:8080/metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
