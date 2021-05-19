package errorlogger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (

	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	defaultLogLevel logrus.Level = logrus.InfoLevel

	// DefaultTimestampFormat is time.RFC3339
	//
	// Note that this is not the most current standard but it is the
	// most stable and recommended with the Go standard library.
	//
	// Additional notes
	//
	// The RFC822, RFC850, and RFC1123 formats should be applied only to
	// local times. Applying them to UTC times will use "UTC" as the time
	// zone abbreviation, while strictly speaking those RFCs require the
	// use of "GMT" in that case.
	//
	// In general RFC1123Z should be used instead of RFC1123 for servers
	// that insist on that format, and RFC3339 should be preferred for
	// new protocols.
	//
	// While RFC3339, RFC822, RFC822Z, RFC1123, and RFC1123Z are useful
	// for formatting, when used with time.Parse they do not accept all
	// the time formats permitted by the RFCs and they do accept time
	// formats not formally defined.
	//
	// The RFC3339Nano format removes trailing zeros from the seconds
	// field and thus may not sort correctly once formatted.
	DefaultTimestampFormat string = time.RFC3339
)

var (
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
	// The Formatter interface is used to implement a custom Formatter.
	// It takes an `Entry`. It exposes all the fields, including the
	// default ones:
	//
	// * `entry.Data["msg"]`. The message passed from Info, Warn, Error ..
	// * `entry.Data["time"]`. The timestamp.
	// * `entry.Data["level"]. The level the entry was logged at.
	//
	// Any additional fields added with `WithField` or `WithFields` are
	// also in `entry.Data`. Format is expected to return an array of
	// bytes which are then logged to `logger.Out`.
	//
	// Reference: logrus@v1.8.1 formatter.go
	// 	type Formatter interface {
	// 		Format(*Entry) ([]byte, error)
	// 	}
	Formatter interface{ logrus.Formatter }

	// logrusLogger implements the most common functionality
	// of the logging interface of the Logrus package.
	//
	// It is a compatible superset of the standard library
	// log package and a compatible subset of the ErrorLogger
	// package.
	logrusLogger interface {
		logrus.FieldLogger
		logrusCommonOptions
	}

	// logrusCommonOptions implements several common options
	// that should be in the basic LogrusLogger interface.
	logrusCommonOptions interface {
		SetLevel(level logrus.Level)
		GetLevel() logrus.Level
		SetFormatter(formatter logrus.Formatter)
		SetOutput(output io.Writer)
	}

	// logrusLoggerComplete implements the complete exported
	// interface implemented by the logrus.Logger struct.
	//
	// Most users will not need to use this. ErrorLogger
	// contains the most common functionality, including the
	// basic LogrusLogger interface.
	logrusLoggerComplete interface {
		logrusLogger
		logrusOptions
		logrusLogFunctions
	}

	// logrusOptions implements rarely used logging options.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	logrusOptions interface {
		WithContext(ctx context.Context) *logrus.Entry
		WithTime(t time.Time) *logrus.Entry
		Exit(code int)
		SetNoLock()
		AddHook(hook logrus.Hook)
		IsLevelEnabled(level logrus.Level) bool
		SetReportCaller(reportCaller bool)
		ReplaceHooks(hooks logrus.LevelHooks) logrus.LevelHooks
	}

	// logrusLogFunctions implements logrus Logrus
	// LogFunctions.
	// Instead of using this directly, create your own custom
	// interface that uses the options required.
	logrusLogFunctions interface {
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

var _ logrusLoggerComplete
