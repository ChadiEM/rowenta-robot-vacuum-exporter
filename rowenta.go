package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var endpoint = os.Getenv("ROWENTA_ENDPOINT")

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
	BatteryPercentage float64
	BatteryVoltage    float64
	Charging          float64
	Uptime            float64
}

func GetStatus() (*Status, error) {
	resp, err := http.Get(endpoint + "/get/status")
	if err != nil {
		fmt.Println("No response from request")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	var result StatusResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
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
		BatteryPercentage: float64(result.BatteryLevel) / float64(100),
		BatteryVoltage:    float64(result.Voltage) / float64(1000),
		Charging:          charging,
		Uptime:            diff.Seconds(),
	}

	return returnValue, nil
}
