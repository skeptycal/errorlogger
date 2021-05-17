// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type (
	// Level type values: Panic, Fatal, Error, Warn, Info, Debug, Trace
	Level = logrus.Level

	// LoggerFunc defines the function signature used when logging errors.
	LoggerFunc = func(args ...interface{})
)

var (
	// Log is the default global ErrorLogger. It implements
	// the ErrorLogger interface as well as the basic
	// logrus.Logger interface, which is compatible with the
	// standard library "log" package.
	//
	// In the case of name collisions with 'Log', use an alias
	// instead of creating a new instance. For example:
	//  var mylogthatwontmessthingsup = errorlogger.Log
	Log = New()

	// Err is the logging function for the global ErrorLogger.
	Err = Log.Err

	// ErrInvalidWriter is returned when an output writer is
	// nil or does not implement io.Writer.
	ErrInvalidWriter = os.ErrInvalid
)

// New returns a new ErrorLogger with default options and
// logging enabled.
// Most users will not need to call this, since the default
// global ErrorLogger 'Log' is provided.
//
// In the case of name collisions with 'Log', use an alias
// instead of creating a new instance. For example:
//  var mylogthatwontmessthingsup = errorlogger.Log
func New() ErrorLogger {
	return NewWithOptions(defaultEnabled, defaultLogFunc, defaultErrWrap)
}

// NewWithOptions returns a new ErrorLogger with options determined
// by parameters.
//
// - enabled: defines the initial logging state.
//
// - fn: defines a custom logging function used to log information.
//
// - wrap: defines a custom error type to wrap all errors in.
func NewWithOptions(enabled bool, fn LoggerFunc, wrap error) ErrorLogger {
	e := errorLogger{}
	if enabled {
		e.Enable()
	} else {
		e.Disable()
	}
	e.Logger = defaultlogger

	e.SetLoggerFunc(fn)
	e.SetErrorWrap(wrap)

	return &e
}
