package utils

import (
	"os"
)

func Dirname() (string, error) {
	ex, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return ex, nil
}
