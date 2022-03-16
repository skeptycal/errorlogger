package errorlogger

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// test_info provides samples and test cases for the tests
// and benchmarks in this package.

var (
	testLogrusLogger  *logrus.Logger = logrus.New()
	testDefaultLogger ErrorLogger    = NewWithOptions(true, "", nil, nil, testLogrusLogger)
	errFake           error          = errors.New("fake")
	fakeSysCallError  error          = os.NewSyscallError("fake syscall error", fmt.Errorf("fake syscall error"))
)

func newTestStruct(enabled bool, msg string, wrap error, fn func(args ...interface{}), logger *logrus.Logger) *errorLogger {
	if logger == nil {
		logger = defaultlogger
	}

	e := errorLogger{
		msg:    msg,
		Logger: logger,
	}

	if enabled {
		e.Enable()
	} else {
		e.Disable()
	}

	// e.Logger = logger
	e.logFunc = e.Error

	if wrap == nil {
		// the defaultErrWrap is actually nil ... so this is not needed.
		// However, if the default is later changed to a package-wide
		// wrapper, this will be a valid check
		wrap = DefaultErrWrap
	}
	e.wrap = wrap
	return &e
}
