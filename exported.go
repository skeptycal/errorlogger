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

func (e *errorLogger) SetOptions(o Options) error {
	// TODO - stuff
	return nil
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

func SetTextFormatter() {
	log.SetFormatter(defaultTextFormatter)
}
