package config_utils

import (
	"errors"
	"flag"
)

var ErrConfigNotFound = errors.New("config not found")

// GetConfigPath ...
func GetConfigPath() (string, error) {
	path := flag.String("c", "", "config path")
	flag.Parse()

	if path == nil {
		return "", ErrConfigNotFound
	}

	return *path, nil
}
