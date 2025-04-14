package config_utils

import (
	"fmt"
	"os"
)

func ReadSecret(key string) string {
	file := os.Getenv(key)
	if file == "" {
		panic(fmt.Sprintf("path to secret not found in environment variable %s", key))
	}
	secret, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(secret)
}
