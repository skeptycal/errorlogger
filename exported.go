// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Level type values: Panic, Fatal, Error, Warn, Info, Debug, Trace
type Level = logrus.Level

// LoggerFunc defines the function signature used when logging errors.
type LoggerFunc = func(args ...interface{})

var (
	// Log is the default global ErrorLogger. It implements the
	// ErrorLogger interface as well as the logrus.Logger interface,
	// which is compatible with the standard library "log" package.
	Log = New()

	// Err is the logging function for the global ErrorLogger.
	Err = Log.Err

	// Log = EL

	// log exports the default logrus error logger as a convenience.
	// The default parameters are
	//  Out: os.Stderr
	//  Formatter: TextFormatter
	//  Hooks: logrus default
	//  Level: Error
	//
	// Reference: https://github.com/sirupsen/logrus/
	log = &logrus.Logger{

		Out: os.Stderr,

		Formatter: defaultTextFormatter,

		Hooks: make(logrus.LevelHooks),

		Level: defaultLogLevel,
	}
)

// ErrInvalidWriter is returned when a log output writer is
// passed that does not implement io.Writer.
var ErrInvalidWriter = os.ErrInvalid

// New returns a new ErrorLogger with default options and logging enabled.
// Most users will not need to call this, since the default global
// ErrorLogger EL is provided.
func New() ErrorLogger {
	return NewWithOptions(defaultEnabled, defaultLogFunc, defaultErrWrap)
}

// NewWithOptions returns a new ErrorLogger with options determined
// by parameters.
func NewWithOptions(enabled bool, fn LoggerFunc, wrap error) ErrorLogger {
	e := &errorLogger{}
	if enabled {
		e.Enable()
	} else {
		e.Disable()
	}
	e.Logger = log

	e.SetLoggerFunc(fn)
	e.SetErrorWrap(wrap)

	return e
}

// SetLogOutput sets the output writer for logging.
// The default is os.Stderr. Any io.Writer can be setup
// to receive messages.
func SetLogOutput(w io.Writer) error {
	switch v := w.(type) {
	case io.Writer:
		log.SetOutput(v)
		return nil
	default:
		return Err(ErrInvalidWriter)
	}
}

// SetTextFormatter sets the log formatter to the default text
// log formatter for use with TTY logging. Use
//  Log.SetJSONFormatter(defaultJSONFormatter)
// to return to set JSON formatting of logs.
//
// Many third party logging formatters are available.
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
func SetTextFormatter() {
	log.SetFormatter(defaultTextFormatter)
}

// SetJSONFormatter sets the log formatter to a logrus JSON
// formatter with default characteristics. Use
//  Log.SetJSONFormatter(defaultTextFormatter)
// to return to default text formatting of logs.
//
// Many third party logging formatters are available.
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
func SetJSONFormatter() {
	log.SetFormatter(defaultJSONFormatter)
}
