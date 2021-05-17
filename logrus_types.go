package errorlogger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const defaultLogLevel logrus.Level = logrus.InfoLevel

var (
	// defaultTextFormatter is the default log formatter. Use
	//  Log.SetText()
	// or
	//  Log.SetFormatter(defaultTextFormatter)
	// to return to default text formatting of logs.
	//
	// To change to another logrus formatter, use
	//  Log.SetFormatter(myFormatter)
	//
	// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#TextFormatter
	defaultTextFormatter logrus.Formatter = new(logrus.TextFormatter)

	// defaultJSONFormatter is the a JSON formatter with
	// default characteristics. Use
	//  Log.SetJSON()
	// or
	//  Log.SetJSONFormatter(defaultJSONFormatter)
	// to enable JSON logging.
	//
	// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#JSONFormatter
	defaultJSONFormatter logrus.Formatter = new(logrus.JSONFormatter)

	// ParseLevel takes a string level and returns the Logrus log level constant.
	ParseLevel = logrus.ParseLevel

	// defaultlogger initializes a default logrus logger.
	// Reference: https://github.com/sirupsen/logrus/
	defaultlogger = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: defaultTextFormatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     defaultLogLevel,
	}
)

type (
	// LogrusLogger implements the most common functionality
	// of the logging interface of the Logrus package.
	//
	// It is a compatible superset of the standard library
	// log package and a compatible subset of the ErrorLogger
	// package.
	LogrusLogger interface {
		LogrusFieldLogger
		LogrusCommonOptions
	}

	// The FieldLogger interface generalizes the Entry and
	// Logger types
	LogrusFieldLogger = logrus.FieldLogger

	// LogrusCommonOptions implements several common options
	// that should be in the basic LogrusLogger interface.
	LogrusCommonOptions interface {
		SetLevel(level Level)
		GetLevel() Level
		SetFormatter(formatter Formatter)
		SetOutput(output io.Writer)
	}

	// LogrusLoggerComplete implements the complete exported
	// interface implemented by the logrus.Logger struct.
	//
	// Most users will not need to use this. ErrorLogger
	// contains the most common functionality, including the
	// basic LogrusLogger interface.
	LogrusLoggerComplete interface {
		LogrusLogger
		LogrusOptions
		LogrusLogFunctions
	}

	// LogrusOptions implements rarely used logging options.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	LogrusOptions interface {
		WithContext(ctx context.Context) *Entry
		WithTime(t time.Time) *Entry
		Exit(code int)
		SetNoLock()
		AddHook(hook Hook)
		IsLevelEnabled(level Level) bool
		SetReportCaller(reportCaller bool)
		ReplaceHooks(hooks LevelHooks) LevelHooks
	}

	// LogrusLogFunctions implements logrus Logrus
	// LogFunctions.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	LogrusLogFunctions interface {
		DebugFn(fn LogFunction)
		InfoFn(fn LogFunction)
		PrintFn(fn LogFunction)
		WarnFn(fn LogFunction)
		WarningFn(fn LogFunction)
		ErrorFn(fn LogFunction)
		FatalFn(fn LogFunction)
		PanicFn(fn LogFunction)
	}

	// Ext1FieldLogger (the first extension to FieldLogger)
	// is superfluous, it is here for consistancy. Do not use.
	Ext1FieldLogger = logrus.Ext1FieldLogger
)

// These are type aliases for logrus types.
type (
	Logger      = logrus.Logger
	Entry       = logrus.Entry
	Fields      = logrus.Fields
	LogFunction = logrus.LogFunction
	Hook        = logrus.Hook
	Formatter   = logrus.Formatter
	LevelHooks  = logrus.LevelHooks
)
