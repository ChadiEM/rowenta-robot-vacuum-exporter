package main

import (
	"log"
	"net/http"
	"os"
	"rowenta-robot-vacuum-exporter/internal/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var endpoint = os.Getenv("ROWENTA_ENDPOINT")

func main() {
	col := collector.New(endpoint)
	prometheus.MustRegister(col)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9101", nil))
}
