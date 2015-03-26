package util

import (
	"errors"
	"time"
)

// Do a thing, or timeout in `timeout` seconds
func Timeout(todo func(chan bool), timeout int) error {
	to := make(chan bool, 1)
	done := make(chan bool, 1)
	stop := make(chan bool, 1)
	go func() {
		todo(stop)
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
		stop <- true
		return errors.New("Timeout")
	}
}
