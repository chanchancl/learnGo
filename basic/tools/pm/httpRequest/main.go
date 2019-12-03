package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	httpReqs := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)
	prometheus.MustRegister(httpReqs)

	httpReqs.WithLabelValues("404", "POST").Add(42)

	m := httpReqs.WithLabelValues("200", "GET")
	for i := 0; i < 1000000; i++ {
		m.Inc()
	}

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Listening!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
