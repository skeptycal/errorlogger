// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import (
	"errors"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	testDefaultLogger ErrorLogger    = New()
	testLogrusLogger  *logrus.Logger = logrus.New()
	fakeError         error          = errors.New("fake")
	fakeSysCallError  error          = os.NewSyscallError("fake syscall error", errors.New("fake syscall error"))
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
		wrap:    DefaultErrWrap,
		logFunc: DefaultLogFunc,
		Logger:  defaultlogger,
	}
	if enabled {
		e.Enable()
		return e
	}
	e.Disable()
	return e
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		// control
		{"global ErrorLogger", testDefaultLogger, false},
		{"nil", nil, true},

		// Test New() (should pass) and nil (should fail)
		{"New()", New(), false},

		// Also tests NewWithOptions()
		{"NewWithOptions(false, nil, nil) (should pass)", NewWithOptions(false, nil, nil), false},
		{"logrus logger alone", logrus.Logger{}, true},

		{"NewWithOptions(true, nil, nil)", NewWithOptions(true, nil, nil), false},
		{"NewWithOptions(true, DefaultLogFunc, DefaultErrWrap)", NewWithOptions(true, DefaultLogFunc, DefaultErrWrap), false},
		{"NewWithOptions(false, DefaultLogFunc, nil)", NewWithOptions(true, DefaultLogFunc, nil), false},

		// Various tests using private struct
		{"logrus logger in errorLogger (not public)", &errorLogger{Logger: &logrus.Logger{}}, false},
		{"default ErrorLogger with nil wrapper (not public)", &errorLogger{wrap: nil}, false},
		// Do not need a check for this in the constructor since errorLogger is not exported
		// But something to be aware of ...
		// {"ErrorLogger with nil logger (should fail)", &errorLogger{Logger: nil}, true},
	}

	wantStruct := &errorLogger{Logger: defaultlogger}
	var want ErrorLogger = wantStruct

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, ok := tt.input.(ErrorLogger)
			if !ok && !tt.wantErr {
				t.Errorf("New(%s) does not implement ErrorLogger: got %T, want %T", tt.name, got, want)

			}

			switch got.(type) {
			case ErrorLogger:
				if tt.wantErr {
					t.Errorf("New(%s) implements ErrorLogger: got %T, want %T", tt.name, got, want)
				}
			default:
				if !tt.wantErr {
					t.Errorf("New(%s) does not implement ErrorLogger: got %T, want %T", tt.name, got, want)
				}
			}
		})
	}
}

func Test_errorLogger_SetErrorWrap(t *testing.T) {
	tests := []struct {
		name  string
		input error
		wrap  error
	}{
		{"fakeError", fakeError, fakeSysCallError},
		{"nil", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e.SetErrorWrap(tt.wrap)
			got := e.yesErr(tt.input)
			if errors.Is(got, fakeSysCallError) {
				t.Errorf("SetErrorWrap(%s) did not wrap error: got %v, want %v", tt.name, got, tt.wrap)

			}
		})
	}
}
