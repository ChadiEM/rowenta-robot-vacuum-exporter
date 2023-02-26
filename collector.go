package main

import "github.com/prometheus/client_golang/prometheus"

type Collector struct {
	batteryPercentage *prometheus.Desc
	batteryVoltage    *prometheus.Desc
	charging          *prometheus.Desc
	uptime            *prometheus.Desc
	distanceDriven    *prometheus.Desc
	cleaningTime      *prometheus.Desc
	totalRuns         *prometheus.Desc
}

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
		distanceDriven: prometheus.NewDesc("rowenta_distance_driven_meters",
			"Distance travelled in meters",
			nil, nil,
		),
		cleaningTime: prometheus.NewDesc("rowenta_clean_time_seconds",
			"Time spent cleaning time in seconds",
			nil, nil,
		),
		totalRuns: prometheus.NewDesc("rowenta_runs_total",
			"Total number of runs",
			nil, nil,
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.batteryPercentage
	ch <- collector.batteryVoltage
	ch <- collector.charging
	ch <- collector.uptime
	ch <- collector.distanceDriven
	ch <- collector.cleaningTime
	ch <- collector.totalRuns
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	status, err := GetStatus()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(
			prometheus.NewDesc("rowenta_status_error",
				"Error getting rowenta status", nil, nil),
			err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.batteryPercentage, prometheus.GaugeValue, status.BatteryPercentage)
		ch <- prometheus.MustNewConstMetric(collector.batteryVoltage, prometheus.GaugeValue, status.BatteryVoltageVolts)
		ch <- prometheus.MustNewConstMetric(collector.charging, prometheus.GaugeValue, status.Charging)
		ch <- prometheus.MustNewConstMetric(collector.uptime, prometheus.GaugeValue, status.UptimeSeconds)
	}

	statistics, err := GetPermanentStatistics()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(
			prometheus.NewDesc("rowenta_statistics_error",
				"Error getting rowenta statistics", nil, nil),
			err)
	} else {
		ch <- prometheus.MustNewConstMetric(collector.distanceDriven, prometheus.CounterValue, statistics.TotalDistanceDrivenMeters)
		ch <- prometheus.MustNewConstMetric(collector.cleaningTime, prometheus.CounterValue, statistics.TotalCleaningTimeSeconds)
		ch <- prometheus.MustNewConstMetric(collector.totalRuns, prometheus.CounterValue, statistics.TotalNumberOfCleaningRuns)
	}
}
