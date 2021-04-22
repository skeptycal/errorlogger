// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package errorlogger implements error logging to a logrus log
// (or a standard library log) by providing a convenient way to
// log errors and to temporarily disable/enable logging.
//
// A global Log and Err with default behaviors are supplied that
// may be aliased if you wish:
//  Log = errorlogger.Log
//  Err = errorlogger.Err
//
// If you do not intend to use any options or disable the logger,
// it may be more convenient to use a function alias to call the
// most common method, Err(), like this:
//  var Err = errorlogger.New().Err
// then, just call the function:
//  err := someProcess(stuff)
//  if err != nil {
//   return Err(err)
//  }
//
// Either way, the default ErrorLogger is enabled and ready to go:
//  EL := errorlogger.New() // enabled by default
//  Err := EL.Err
//
// If a private ErrorLogger is desired, or if name collisions with
// Err cause conflicts, you may implement your own.
//  myErr := errorlogger.New()
//  err := myErr.Err
//
// Example:
//  f, err := os.Open("somefile.txt")
//  if err != nil {
// 	 return nil, e.Err(err) // avoids additional logging steps
//  }
//  e.Disable() // can be disabled and enabled as desired
package errorlogger

import (
	"github.com/sirupsen/logrus"
)

const (
	defaultLogLevel logrus.Level = logrus.InfoLevel
	defaultEnabled  bool         = true
)

// Defaults for ErrorLogger
var (
	// defaultLogFunc is Log.Error, which will log messages
	// of level ErrorLevel or higher.
	defaultLogFunc LoggerFunc = log.Error

	// defaultErrWrap is the default error used to wrap
	// errors processed with Err. A <nil> value disables
	// error wrapping.
	defaultErrWrap error = nil

	// defaultTextFormatter is the default log formatter. Use
	//  Log.SetFormatter()
	// to change to another logrus formatter or
	//  Log.SetJSONFormatter(defaultTextFormatter)
	// to return to default text formatting of logs.
	//
	// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#TextFormatter
	defaultTextFormatter logrus.Formatter = new(logrus.TextFormatter)

	// defaultJSONFormatter is the a JSON formatter with
	// default characteristics. Use
	//  Log.SetJSONFormatter(defaultJSONFormatter)
	// to enable JSON logging.
	//
	// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#JSONFormatter
	defaultJSONFormatter logrus.Formatter = new(logrus.JSONFormatter)
)

// ErrorLogger implements error logging to a logrus log
// (or a standard library log) by providing convenience
// methods, advanced formatting options, more automated
// logging, a more efficient way to log errors within
// code, and methods to temporarily disable/enable
// logging, such as in the case of performance optimization
// or during critical code blocks.
type ErrorLogger interface {

	// Disable disables logging and sets a no-op function for
	// Err() to prevent slowdowns while logging is disabled.
	Disable()

	// Enable enables logging and restores the Err() logging functionality.
	Enable()

	// EnableText enables text formatting of log errors (default)
	EnableText()

	// EnableJSON enables JSON formatting of log errors
	EnableJSON()

	// LogLevel sets the logging level from a string value.
	// Allowed values: Panic, Fatal, Error, Warn, Info, Debug, Trace
	SetLogLevel(lvl string) error

	// Err logs an error to the provided logger, if it is enabled,
	// and returns the error unchanged.
	Err(err error) error

	// SetLoggerFunc allows setting of the logger function.
	// The default is log.Error(), which is compatible with
	// the standard library log package and logrus.
	SetLoggerFunc(fn LoggerFunc)

	// SetErrorWrap allows ErrorLogger to wrap errors in a
	// specified custom type. For example, if you want all errors
	// returned to be of type *os.PathError
	SetErrorWrap(wrap error)

	logrus.Ext1FieldLogger
}

// errorLogger implements ErrorLogger with logrus or the
// standard library log package.
type errorLogger struct {
	enabled bool                  // `default:"true"`
	wrap    error                 // `default:"nil"` // nil = disabled
	errFunc func(err error) error // `default:"()yesErr"`
	logFunc LoggerFunc            // `default:"logrus.New()"`
	*logrus.Logger
}

// SetErrorType allows ErrorLogger to wrap errors in a specified custom message.
// Setting wrap == "" will disable wrapping of errors.
func (e *errorLogger) SetErrorWrap(wrap error) {
	e.wrap = wrap
}

// EnableText enables text formatting of log errors (default)
func (e *errorLogger) EnableText() {
	e.Logger.SetFormatter(defaultTextFormatter)
}

// EnableJSON enables JSON formatting of log errors
func (e *errorLogger) EnableJSON() {
	e.Logger.SetFormatter(defaultJSONFormatter)
}

func (e *errorLogger) SetLogLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return Err(err)
	}
	e.Logger.SetLevel(level)
	return nil
}

// SetLoggerFunc allows setting of the logger function.
// The default is Log.Error(err), which is compatible with
// the standard library log package and logrus.
//
// The function signature must be of type LoggerFunc:
//  func(args ...interface{}).
func (e *errorLogger) SetLoggerFunc(fn LoggerFunc) {
	e.logFunc = fn
}
