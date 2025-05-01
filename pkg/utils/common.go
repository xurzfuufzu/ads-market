package utils

import (
	"Ads-marketplace/internal/domain"
	"encoding/json"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err := fn(); err != nil {
			time.Sleep(delay)
			attempts--

		}
		return nil
	}
	return
}

func SerializePlatforms(platforms []domain.Platform) (string, error) {
	platformsJson, err := json.Marshal(platforms)
	if err != nil {
		return "", err
	}
	return string(platformsJson), nil
}

func DeserializePlatforms(data string) ([]domain.Platform, error) {
	var platforms []domain.Platform
	err := json.Unmarshal([]byte(data), &platforms)
	if err != nil {
		return nil, err
	}
	return platforms, nil
}
