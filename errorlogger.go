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
	"io"
)

const defaultEnabled bool = true

// Defaults for ErrorLogger
var (
	// defaultLogFunc is Log.Error, which will log messages
	// of level ErrorLevel or higher.
	defaultLogFunc LoggerFunc = defaultlogger.Error

	// defaultErrWrap is the default error used to wrap
	// errors processed with Err. A <nil> value disables
	// error wrapping.
	defaultErrWrap error = nil
)

// ErrorLogger implements error logging to a logrus log
// (or a standard library log) by providing convenience
// methods, advanced formatting options, more automated
// logging, a more efficient way to log errors within
// code, and methods to temporarily disable/enable
// logging, such as in the case of performance
// optimization or during critical code blocks.
type ErrorLogger interface {

	// Disable disables logging and sets a no-op function for
	// Err() to prevent slowdowns while logging is disabled.
	Disable()

	// Enable enables logging and restores the Err() logging functionality.
	Enable()

	// EnableText enables text formatting of log errors (default)
	SetText()

	// EnableJSON enables JSON formatting of log errors
	SetJSON()

	// SetOptions accepts an Options set and adjust the
	// ErrorLogger options accordingly. Any options that are not included are ignored. The Options struct has methods for managing, saving and loading Options sets.
	SetOptions(o Options) error

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

	LogrusLogger
}

// Options is Pretty options
type Options struct {
	// Width is an max column width for single line arrays
	// Default is 80
	Width int
	// Prefix is a prefix for all lines
	// Default is an empty string
	Prefix string
	// Indent is the nested indentation
	// Default is two spaces
	Indent string
	// SortKeys will sort the keys alphabetically
	// Default is false
	SortKeys bool
}

// errorLogger implements ErrorLogger with logrus or the
// standard library log package.
type errorLogger struct {
	enabled bool                  // `default:"true"`
	wrap    error                 // `default:"nil"` // nil = disabled
	errFunc func(err error) error // `default:"()yesErr"`
	logFunc LoggerFunc            // `default:"logrus.New()"`
	*Logger
}

// SetErrorType allows ErrorLogger to wrap errors in a specified custom message.
// Setting wrap == "" will disable wrapping of errors.
func (e *errorLogger) SetErrorWrap(wrap error) {
	e.wrap = wrap
}

// SetJSON sets the log format to JSON. The JSON output conforms
// to RFC 7159 (https://www.rfc-editor.org/rfc/rfc7159.html) from
// March 2014.
//
// It should be noted that this format has been obsoleted the
// latest version of the JSON standard from December 2017,
// RFC 8259 (https://datatracker.ietf.org/doc/html/rfc8259)
//
// The default is compact "ugly" json. A "pretty" format can be
// selected with
//  Log.SetOptions()
//
// Use
//  Log.SetText()
// to return to the default Text formatter.
//
// In general,
//  Log.Setformatter(myformatter)
// can be used to set any custom formatter.
//
// Many other third party logging formatters are available.
//
// - FluentdFormatter. Formats entries that can be parsed by Kubernetes and Google Container Engine.
//
// - GELF. Formats entries so they comply to Graylog's GELF 1.1 specification.
//
// - logstash. Logs fields as Logstash Events.
//
// - prefixed. Displays log entry source along with alternative layout.
//
// - zalgo. Invoking the Power of Zalgo.
//
// - nested-logrus-formatter. Converts logrus fields to a nested structure.
//
// - powerful-logrus-formatter. get fileName, log's line number and the latest function's name when print log; Sava log to files.
//
// - caption-json-formatter. logrus's message json formatter with human-readable caption added.
//
// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#JSONFormatter
func (e *errorLogger) SetJSON() {
	e.SetFormatter(defaultJSONFormatter)
}

// SetLoggerFunc sets the logger function that is used to
// write log messages. This allows rapid switching between loggers
// as well as turning the logging off and on regularly.
//
// The default is Log.Error(err), which is compatible with
// the standard library log package and logrus. Setting this to
// a no-op function allows fast pass through of logging
// information that is not useful to record permanently.
//
// The function signature must be of type LoggerFunc:
//  func(args ...interface{}).
func (e *errorLogger) SetLoggerFunc(fn LoggerFunc) {
	e.logFunc = fn
}

// SetLogLevel converts lvl to a compatible log level and sets the log level.
//
// Allowed values: Panic, Fatal, Error, Warn, Info, Debug, Trace
func (e *errorLogger) SetLogLevel(lvl string) error {
	level, err := ParseLevel(lvl)
	if err != nil {
		return Err(err)
	}
	e.Logger.SetLevel(level)
	return nil
}

// SetLogOutput sets the output writer for logging.
// The default is os.Stderr. Any io.Writer can be setup
// to receive messages.
func (e *errorLogger) SetLogOutput(w io.Writer) error {
	switch v := w.(type) {
	case io.Writer:
		e.SetOutput(v)
		return nil
	default:
		return Err(ErrInvalidWriter)
	}
}

func (e *errorLogger) SetOptions(o Options) error {
	// TODO - stuff
	return nil
}

// SetText sets the log format to Text. This is the default
// formatter.
//
// It provides ANSI colorized (if available) TTY output to os.Stderr.
//
// Use
//  Log.SetJSON()
// to switch to the JSON formatter.
//
// In general,
//  Log.Setformatter(myformatter)
// can be used to set any custom formatter.
//
// Many other third party logging formatters are available.
//
// - FluentdFormatter. Formats entries that can be parsed by Kubernetes and Google Container Engine.
//
// - GELF. Formats entries so they comply to Graylog's GELF 1.1 specification.
//
// - logstash. Logs fields as Logstash Events.
//
// - prefixed. Displays log entry source along with alternative layout.
//
// - zalgo. Invoking the Power of Zalgo.
//
// - nested-logrus-formatter. Converts logrus fields to a nested structure.
//
// - powerful-logrus-formatter. get fileName, log's line number and the latest function's name when print log; Sava log to files.
//
// - caption-json-formatter. logrus's message json formatter with human-readable caption added.
// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#TextFormatter
func (e *errorLogger) SetText() {
	e.SetFormatter(defaultTextFormatter)
}
