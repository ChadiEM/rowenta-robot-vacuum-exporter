package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type Collector struct {
	batteryPercentage *prometheus.Desc
	batteryVoltage    *prometheus.Desc
	charging          *prometheus.Desc
	uptime            *prometheus.Desc
}

// Declare your exporter metrics here. Referred to as "collectors"
func newCollector() *Collector {
	return &Collector{
		batteryPercentage: prometheus.NewDesc("rowenta_battery_level",
			"Battery Percentage",
			nil, nil,
		),
		batteryVoltage: prometheus.NewDesc("rowenta_battery_voltage_volts",
			"Battery Voltage in volts",
			nil, nil,
		),
		charging: prometheus.NewDesc("rowenta_charging",
			"Charging state",
			nil, nil,
		),
		uptime: prometheus.NewDesc("rowenta_uptime_seconds",
			"Robot uptime in seconds",
			nil, nil,
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.batteryPercentage
	ch <- collector.batteryVoltage
	ch <- collector.charging
	ch <- collector.uptime
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	status, err := GetStatus()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(
			prometheus.NewDesc("rowenta_error",
				"Error getting rowenta status", nil, nil),
			err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.batteryPercentage, prometheus.GaugeValue, status.BatteryPercentage)
		ch <- prometheus.MustNewConstMetric(collector.batteryVoltage, prometheus.GaugeValue, status.BatteryVoltage)
		ch <- prometheus.MustNewConstMetric(collector.charging, prometheus.GaugeValue, status.Charging)
		ch <- prometheus.MustNewConstMetric(collector.uptime, prometheus.GaugeValue, status.Uptime)
	}
}

func main() {
	collector := newCollector()
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9101", nil))
}
