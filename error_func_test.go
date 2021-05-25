// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import (
	"testing"
)

var (
	fakeOuter   error
	yesnologger = New()
)

type nopWriter struct{}

func (nopWriter) Write(b []byte) (n int, err error) {
	return 0, nil
}

func init() {
	yesnologger.SetOutput(nopWriter{})
}

// noErr is a no-op errorFunc for disabling logging without
// constant repetitive flag checks or other hacks.
// https://en.wikipedia.org/wiki/NOP_(code)
func (e *errorLogger) noErr1(err error) error {
	// TODO does this really need to return an error?
	// TODO does the compiler remove this?
	// TODO would a pointer be better here? *(&err)
	// TODO does the compiler remove this?
	return *(&err)
}

func Benchmark_errorLogger_noErr_yesErr(b *testing.B) {
	tests := []struct {
		name  string
		input error
		fn    LoggerFunc
		want  error
	}{
		// TODO: Add test cases.
		{"fake error", errFake, errFake}, // return an error unchanged
		{"nil", nil, nil},                // return nil unchanged
	}

	for _, bb := range tests {
		for _, test := range errFuncTestList {
			b.Run(test.name+"("+bb.name+")", func(b *testing.B) {
				var fake error
				for i := 0; i < b.N; i++ {
					fake = test.fn(bb.input)

				}
				fakeOuter = fake
			})
		}
	}
}

func Test_errorLogger_noErr_yesErr(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  error
	}{
		// TODO: Add test cases.
		{"fake error", errFake, errFake}, // return an error unchanged
		{"nil", nil, nil},                // return nil unchanged
	}

	for _, tt := range tests {
		for _, test := range errFuncTestList {
			t.Run(test.name+"("+tt.name+")", func(t *testing.T) {
				if got := test.fn(tt.input); got != tt.want {
					t.Errorf("errorLogger.noErr() error = got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
