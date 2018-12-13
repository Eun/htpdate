// +build !cgo

package main

import (
	"errors"
	"time"
)

func setTime(t time.Time) error {
	return errors.New("unable to set time natively, use the --command parameter")
}
