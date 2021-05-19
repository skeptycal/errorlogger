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
		logrus.FieldLogger
		LogrusCommonOptions
	}

	// LogrusCommonOptions implements several common options
	// that should be in the basic LogrusLogger interface.
	LogrusCommonOptions interface {
		SetLevel(level Level)
		GetLevel() Level
		SetFormatter(formatter logrus.Formatter)
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
		WithContext(ctx context.Context) *logrus.Entry
		WithTime(t time.Time) *logrus.Entry
		Exit(code int)
		SetNoLock()
		AddHook(hook logrus.Hook)
		IsLevelEnabled(level Level) bool
		SetReportCaller(reportCaller bool)
		ReplaceHooks(hooks logrus.LevelHooks) logrus.LevelHooks
	}

	// LogrusLogFunctions implements logrus Logrus
	// LogFunctions.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	LogrusLogFunctions interface {
		DebugFn(fn logrus.LogFunction)
		InfoFn(fn logrus.LogFunction)
		PrintFn(fn logrus.LogFunction)
		WarnFn(fn logrus.LogFunction)
		WarningFn(fn logrus.LogFunction)
		ErrorFn(fn logrus.LogFunction)
		FatalFn(fn logrus.LogFunction)
		PanicFn(fn logrus.LogFunction)
	}
)
