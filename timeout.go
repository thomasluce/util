package util

import (
	"errors"
	"time"
)

// Do a thing, or timeout in `timeout` seconds
func Timeout(todo func(), timeout int) error {
	to := make(chan bool, 1)
	done := make(chan bool, 1)
	go func() {
		todo()
		done <- true
	}()

	go func() {
		time.Sleep(time.Duration(timeout * int(time.Second)))
		to <- true
	}()

	select {
	case <-done:
		return nil
	case <-to:
		return errors.New("Timeout")
	}
}
