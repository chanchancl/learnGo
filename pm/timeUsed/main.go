package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	timeUsed := prometheus.NewHistogram(prometheus.HistogramOpts{Name: "TimeUsed"})
	prometheus.MustRegister(timeUsed)
	for i := 0; i < 101; i++ {
		timeUsed.Observe(float64(i))
	}

	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8080", nil)
}
