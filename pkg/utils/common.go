package utils

import (
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
