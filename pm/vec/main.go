package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	vec = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "CounterVec", Help: " "},
		[]string{"count"})
)

func init() {
	prometheus.MustRegister(vec)
}

func main() {
	vec.WithLabelValues("success").Inc()
	vec.WithLabelValues("failure").Inc()
	vec.WithLabelValues("total").Inc()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
