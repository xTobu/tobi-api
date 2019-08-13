package utils

import (
	"fmt"
	"time"
	"tobi-api/lib/log"
)

type stop struct {
	error
}

// Retry ...
func Retry(attempts int, times int, fn func() error) error {
	sleep := time.Duration(times) * time.Second
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			log.Info(fmt.Sprintf("retry func error: %s. attemp #%d after %s.", err.Error(), attempts, sleep))
			time.Sleep(sleep)
			return Retry(attempts, times, fn)
		}
		return err
	}
	return nil
}
