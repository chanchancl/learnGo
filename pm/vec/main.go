package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	vec = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "CounterVec"},
		[]string{"count"})
)

func init() {
	prometheus.MustRegister(vec)
}

func main() {
	vec.WithLabelValues("success").Add(0)
	vec.WithLabelValues("failure").Add(0)
	vec.WithLabelValues("total").Inc()
	
	http.Handle("/metrics", prometheus.Handler())
	fmt.Println("Listening!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
