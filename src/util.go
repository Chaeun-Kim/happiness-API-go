package main

import (
	"encoding/json"
	"fmt"
	"github.com/montanaflynn/stats"
	"io/ioutil"
	"lightcast/happiness/constants"
	"lightcast/happiness/model"
	"log/slog"
	"os"
	"path/filepath"
)

func ComputeMetrics(data []model.IndexMapping, metrics []string) ([]model.MetricMapping, error) {
	values := make([]float64, len(data))
	for i, item := range data {
		values[i] = item.Value
	}

	metricFuncs := map[string]func(stats.Float64Data) (float64, error){
		constants.AVERAGE_METRIC: stats.Mean,
		constants.MEDIAN_METRIC:  stats.Median,
		constants.P25_METRIC:     func(values stats.Float64Data) (float64, error) { return stats.Percentile(values, 25.0) },
		constants.P75_METRIC:     func(values stats.Float64Data) (float64, error) { return stats.Percentile(values, 75.0) },
		constants.MIN_METRIC:     stats.Min,
		constants.MAX_METRIC:     stats.Max,
		constants.STDDEV_METRIC:  stats.StandardDeviation,
	}

	metricMappings := []model.MetricMapping{}
	for _, metric := range metrics {
		if computeFunc, exists := metricFuncs[metric]; exists {
			metricValue, err := computeFunc(values)
			if err != nil {
				slog.Error("Failed to produce metric %v", err)
				return nil, fmt.Errorf("Failed to calculate metric '%s'", metric)
			}
			metricMappings = append(metricMappings, model.MetricMapping{Name: metric, Value: metricValue})
		}
	}

	return metricMappings, nil
}

func InitHappinessIndexData() model.HappinessIndexMap {
	bytes, err := LoadJsonBytes(constants.HAPPINESS_INDEX_DATA_PATH)
	if err != nil {
		slog.Info("Failed to load JSON data: %v", err)
		os.Exit(1)
	}

	var rawData []map[string]float64
	err = json.Unmarshal(bytes, &rawData)
	if err != nil {
		slog.Info("Failed to unmarshall JSON data: %v", err)
		os.Exit(1)
	}

	data := make(model.HappinessIndexMap)
	for _, item := range rawData {
		for key, value := range item {
			data[key] = value
		}
	}

	return data
}

func LoadJsonBytes(filename string) ([]byte, error) {
	file_path, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return bytes, err
}
