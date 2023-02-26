package main

import (
	"time"
)

type PermanentStatisticsResponse struct {
	TotalDistanceDriven       int `json:"total_distance_driven"`
	TotalCleaningTime         int `json:"total_cleaning_time"`
	TotalAreaCleaned          int `json:"total_area_cleaned"`
	TotalNumberOfCleaningRuns int `json:"total_number_of_cleaning_runs"`
}

type PermanentStatistics struct {
	TotalDistanceDrivenMeters float64
	TotalCleaningTimeSeconds  float64
	TotalNumberOfCleaningRuns float64
}

func GetPermanentStatistics() (*PermanentStatistics, error) {
	result, err := ParseUrl[PermanentStatisticsResponse](endpoint + "/get/permanent_statistics")
	if err != nil {
		return nil, err
	}

	returnValue := &PermanentStatistics{
		TotalDistanceDrivenMeters: float64(result.TotalDistanceDriven) / 100.0,
		TotalCleaningTimeSeconds:  (time.Duration(result.TotalCleaningTime) * time.Minute).Seconds(),
		TotalNumberOfCleaningRuns: float64(result.TotalNumberOfCleaningRuns),
	}

	return returnValue, nil
}
