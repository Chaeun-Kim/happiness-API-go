package main

import (
	"encoding/json"
	"io/ioutil"
	"lightcast/happiness/constants"
	"lightcast/happiness/model"
	"log/slog"
	"os"
	"path/filepath"
)

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
