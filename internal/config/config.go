package config

import (
	"encoding/json"
	"io"
	"movie-reservation-system/internal/models"
	"os"
)

func GetConfig() (models.Config, error) {
	configFile, err := os.Open("./internal/config/config.json")
	if err != nil {
		return models.Config{}, err
	}
	defer configFile.Close()

	byteValue, _ := io.ReadAll(configFile)

	var config models.Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return models.Config{}, err
	}

	return config, nil
}
