// Package kd100   /**
package kd100

import "errors"

func newErr(text string) error {
	return errors.New(text)
}
