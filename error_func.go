// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import "github.com/pkg/errors"

// Disable disables logging and sets a no-op function for
// Err() to prevent slowdowns while logging is disabled.
func (e *errorLogger) Disable() {
	e.enabled = false
	e.errFunc = e.noErr
}

// Enable enables logging and restores the Err() logging functionality.
func (e *errorLogger) Enable() {
	e.enabled = true
	e.errFunc = e.yesErr
}

// Err logs an error to the provided logger, if it is enabled,
// and returns the error unchanged.
func (e *errorLogger) Err(err error) error {
	return e.errFunc(err)
}

// noErr is a no-op placeholder for Err
func (e *errorLogger) noErr(err error) error {
	return err
}

// Err logs errors and passes them through unchanged.
func (e *errorLogger) yesErr(err error) error {
	if err != nil {
		if e.wrap != nil {
			err = errors.Wrap(err, e.wrap.Error())
		}
		e.logFunc(err)
	}
	return err
}