package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (

	// Vec 相较之下就是多了可变的Label
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "this_Counter_Success",
	})
	counterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "this_CounterVec_Failed",
	}, []string{
		"code",
		"api",
	})

	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "this_Gauge_OngoingRequest",
	})

	gaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "this_GaugeVec_OngoingRequest",
	}, []string{})

	histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "this_Histogram_MathScoreInOneClass",
		Buckets: prometheus.LinearBuckets(0, 10, 10),
	})

	summary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "this_Summary_MathScoreInOneClass",
		Objectives: map[float64]float64{0.5: 0.01, 0.8: 0.01, 0.9: 0.01, 0.95: 0.001},
	})
)

func init() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(counterVec)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(gaugeVec)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)
}

func OperateCounter() {
	for range 10 {
		counter.Add(1)
	}

	for range 15 {
		counterVec.WithLabelValues("200", "/v1/index.html").Add(1)
	}
	for range 24 {
		counterVec.WithLabelValues("201", "/v2/index.html").Add(1)
	}
}

func OperateGauge() {
	gauge.Add(10)
	gauge.Add(-2)

	gaugeVec.WithLabelValues().Add(10)
	gaugeVec.WithLabelValues().Add(-2)
}

func OperateHistogram() {
	// 65 students, 40 of them are 0 - 60, 25 of them are 60 - 100
	for i := range 65 {
		if i < 40 {
			v := float64(int(80 * rand.Float64()))
			histogram.Observe(v)
			summary.Observe(v)
		} else {
			v := float64(int(80 + 20*rand.Float64()))
			histogram.Observe(v)
			summary.Observe(v)
		}
	}
}

func main() {
	/*
		# HELP this_CounterVec_Failed
		# TYPE this_CounterVec_Failed counter
		this_CounterVec_Failed{api="/v1/index.html",code="200"} 15
		this_CounterVec_Failed{api="/v2/index.html",code="201"} 24

		# HELP this_Counter_Success
		# TYPE this_Counter_Success counter
		this_Counter_Success{api="/index.html"} 10

		# HELP this_Gauge_OngoingRequest
		# TYPE this_Gauge_OngoingRequest gauge
		this_Gauge_OngoingRequest 8

		# HELP this_Histogram_MathScoreInOneClass
		# TYPE this_Histogram_MathScoreInOneClass histogram
		this_Histogram_MathScoreInOneClass_bucket{le="0"} 1
		this_Histogram_MathScoreInOneClass_bucket{le="10"} 5
		this_Histogram_MathScoreInOneClass_bucket{le="20"} 12
		this_Histogram_MathScoreInOneClass_bucket{le="30"} 21
		this_Histogram_MathScoreInOneClass_bucket{le="40"} 25
		this_Histogram_MathScoreInOneClass_bucket{le="50"} 31
		this_Histogram_MathScoreInOneClass_bucket{le="60"} 34
		this_Histogram_MathScoreInOneClass_bucket{le="70"} 39
		this_Histogram_MathScoreInOneClass_bucket{le="80"} 41
		this_Histogram_MathScoreInOneClass_bucket{le="90"} 53
		this_Histogram_MathScoreInOneClass_bucket{le="+Inf"} 65
		this_Histogram_MathScoreInOneClass_sum 3579
		this_Histogram_MathScoreInOneClass_count 65

		# HELP this_Summary_MathScoreInOneClass
		# TYPE this_Summary_MathScoreInOneClass summary
		this_Summary_MathScoreInOneClass{quantile="0.5"} 55
		this_Summary_MathScoreInOneClass{quantile="0.8"} 89
		this_Summary_MathScoreInOneClass{quantile="0.9"} 96
		this_Summary_MathScoreInOneClass{quantile="0.95"} 96
		this_Summary_MathScoreInOneClass_sum 3579
		this_Summary_MathScoreInOneClass_count 65
	*/
	OperateCounter()
	OperateGauge()
	OperateHistogram()
	log.Println("Listen on http://localhost:8089/metrics")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8089", nil))
}
