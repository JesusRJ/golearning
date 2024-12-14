package main

import (
	"errors"
	"fmt"
)

type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

func boom() error {
	return errors.New("boom!")
}

func getBoom() error {
	err := boom()
	return Wrap(err, "getBoom")
}

func main() {
	err := getBoom()
	err = Wrap(err, "main")
	fmt.Println(err)
}
