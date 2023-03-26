package util

import (
	"fmt"
	"os"
)

// Config returns environment variable exported via godotenv
func Config(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("config for key %s do not exist", key)
	} else {
		return val, nil
	}
}
