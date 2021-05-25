package errorlogger

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// test_info provides samples and test cases for the tests
// and benchmarks in this package.

var (
	testDefaultLogger ErrorLogger    = New()
	testLogrusLogger  *logrus.Logger = logrus.New()
	errFake           error          = errors.New("fake")
	fakeSysCallError  error          = os.NewSyscallError("fake syscall error", errors.New("fake syscall error"))
)

var (
	e = newTestStruct(true, nil, nil, nil)
	w = newTestStruct(true, fakeSysCallError, nil, nil)

	errFuncTestList = []struct {
		name string
		fn   func(err error) error
	}{
		{"noErr", e.noErr},
		{"yesErr", e.yesErr},
		{"noErr", w.noErr},
		{"yesErr", w.yesErr},
	}

	// errorloggerTests provide a set of instantiated errorloggers
	// used for tests.
	// input uses type interface{} in order to allow testing with
	// a variety of types that may or may not implement ErrorLogger.
	//
	// If ErrorLogger is not implemented, the wantErr bool is true.
	errorloggerTests = []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		// control
		{"global ErrorLogger", testDefaultLogger, false},

		// Check for false positive and false negative errors
		// Test New() should pass and nil should fail
		{"New()", New(), false},
		{"nil", nil, true},

		// NewWithOptions() is also tested here
		{"NewWithOptions(false, nil, nil, nil) (should pass)", NewWithOptions(false, nil, nil, nil), false},
		{"logrus logger alone", logrus.Logger{}, true},

		{"NewWithOptions(true, nil, nil, nil)", NewWithOptions(true, nil, nil, nil), false},
		{"NewWithOptions(true, nil, nil, string)", NewWithOptions(true, nil, nil, "fake"), false},
		{"NewWithOptions(true, nil, nil, integer)", NewWithOptions(true, nil, nil, 42), false},
		{"NewWithOptions(all defaults ...)", NewWithOptions(true, DefaultLogFunc, DefaultErrWrap, defaultlogger), false},
		{"NewWithOptions(false, DefaultLogFunc, nil)", NewWithOptions(true, DefaultLogFunc, nil, nil), false},

		// Various tests using private struct
		{"logrus logger in errorLogger (not public)", &errorLogger{Logger: &logrus.Logger{}}, false},
		{"default ErrorLogger with nil wrapper (not public)", &errorLogger{wrap: nil}, false},
		// Do not need a check for this in the constructor since errorLogger is not exported
		// But something to be aware of ...
		// {"ErrorLogger with nil logger (should fail)", &errorLogger{Logger: nil}, true},
	}
)

func newTestStruct(enabled bool, wrap error, fn func(args ...interface{}), logger *logrus.Logger) *errorLogger {
	if fn == nil {
		fn = DefaultLogFunc
	}

	if wrap == nil {
		// the defaultErrWrap is actually nil ... so this is not needed.
		// However, if the default is later changed to a package-wide
		// wrapper, this will be a valid check
		wrap = DefaultErrWrap
	}

	if logger == nil {
		logger = defaultlogger
	}

	e := &errorLogger{
		wrap:    wrap,
		logFunc: fn,
		Logger:  logger,
	}
	if enabled {
		e.Enable()
		return e
	}
	e.Disable()
	return e
}
