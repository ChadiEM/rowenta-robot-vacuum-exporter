package rowenta

import (
	"rowenta-robot-vacuum-exporter/internal/parser"
)

type PermanentStatisticsResponse struct {
	TotalDistanceDriven       int `json:"total_distance_driven"`
	TotalCleaningTime         int `json:"total_cleaning_time"`
	TotalAreaCleaned          int `json:"total_area_cleaned"`
	TotalNumberOfCleaningRuns int `json:"total_number_of_cleaning_runs"`
}

type PermanentStatistics struct {
	TotalDistanceDrivenMeters float64
	TotalAreaCleanedMeters2   float64
	TotalCleaningTimeSeconds  float64
	TotalNumberOfCleaningRuns float64
}

func GetPermanentStatistics(endpoint string) (*PermanentStatistics, error) {
	result, err := parser.ParseURL[PermanentStatisticsResponse](endpoint + "/get/permanent_statistics")
	if err != nil {
		return nil, err
	}

	returnValue := &PermanentStatistics{
		TotalDistanceDrivenMeters: float64(result.TotalDistanceDriven) / 128.0,
		TotalAreaCleanedMeters2:   float64(result.TotalAreaCleaned) / 64.0,
		TotalCleaningTimeSeconds:  float64(result.TotalCleaningTime) * 3600.0 / 64.0,
		TotalNumberOfCleaningRuns: float64(result.TotalNumberOfCleaningRuns),
	}

	return returnValue, nil
}
