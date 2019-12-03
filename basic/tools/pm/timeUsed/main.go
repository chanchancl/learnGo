package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
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

	fmt.Println("Listening!")
	http.Handle("/metrics", prometheus.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))

}
