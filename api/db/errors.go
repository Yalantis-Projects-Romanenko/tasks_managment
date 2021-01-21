package db

import (
	"fmt"
	"strings"
)

// Errors raised by package x.
var (
	ErrInvalidParameters = wrapError{msg: "invalid parameters"}
)

type wrapError struct {
	err error
	msg string
}

func (err wrapError) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%s: %v", err.msg, err.err)
	}
	return err.msg
}

func (err wrapError) Wrap(inner error) error {
	return wrapError{msg: err.msg, err: inner}
}

func (err wrapError) Unwrap() error {
	return err.err
}

func (err wrapError) Is(target error) bool {
	ts := target.Error()
	return ts == err.msg || strings.HasPrefix(ts, err.msg+": ")
}
