// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import (
	"testing"
)

var (
	e = newTestStruct(true, nil, nil, nil)
	w = newTestStruct(true, fakeSysCallError, nil, nil)

	errFuncList = []struct {
		name string
		fn   func(err error) error
	}{
		{"noErr", e.noErr},
		{"yesErr", e.yesErr},
		{"noErr", w.noErr},
		{"yesErr", w.yesErr},
	}
)

func Test_errorLogger_noErr_yesErr(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  error
	}{
		// TODO: Add test cases.
		{"fake error", fakeError, fakeError}, // return an error unchanged
		{"nil", nil, nil},                    // return nil unchanged
	}

	for _, tt := range tests {
		for _, test := range errFuncList {
			t.Run(test.name+"("+tt.name+")", func(t *testing.T) {
				if got := test.fn(tt.input); got != tt.want {
					t.Errorf("errorLogger.noErr() error = got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
