package main

import (
	"time"
)

type StatusResponse struct {
	Voltage      int    `json:"voltage"`
	Mode         string `json:"mode"`
	BatteryLevel int    `json:"battery_level"`
	Charging     string `json:"charging"`
	StartupTime  struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
		Hour  int `json:"hour"`
		Min   int `json:"min"`
		Sec   int `json:"sec"`
	} `json:"startup_time"`
}

type Status struct {
	BatteryPercentage   float64
	BatteryVoltageVolts float64
	Charging            float64
	UptimeSeconds       float64
}

func GetStatus() (*Status, error) {
	result, err := ParseUrl[StatusResponse](endpoint + "/get/status")
	if err != nil {
		return nil, err
	}

	var charging float64
	if result.Charging == "charging" {
		charging = 1
	} else {
		charging = 0
	}

	startupTime := result.StartupTime
	startupDate := time.Date(startupTime.Year, time.Month(startupTime.Month), startupTime.Day, startupTime.Hour, startupTime.Min, startupTime.Sec, 0, time.Local)
	now := time.Now()
	diff := now.Sub(startupDate)

	returnValue := &Status{
		BatteryPercentage:   float64(result.BatteryLevel) / float64(100),
		BatteryVoltageVolts: float64(result.Voltage) / float64(1000),
		Charging:            charging,
		UptimeSeconds:       diff.Seconds(),
	}

	return returnValue, nil
}
