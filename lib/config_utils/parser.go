package config_utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Parse[T any](filepath string) (T, error) {
	var config T

	file, err := os.ReadFile(filepath) // nolint:gosec
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
