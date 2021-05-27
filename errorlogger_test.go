// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {

	wantStruct := &errorLogger{Logger: defaultlogger}
	var want ErrorLogger = wantStruct

	for _, tt := range errorloggerTests {
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
		{"fakeError", errFake, fakeSysCallError},
		{"nil", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorLoggerTestStruct.SetErrorWrap(tt.wrap)
			got := errorLoggerTestStruct.yesErr(tt.input)
			if errors.Is(got, fakeSysCallError) {
				t.Errorf("SetErrorWrap(%s) did not wrap error: got %v, want %v", tt.name, got, tt.wrap)

			}
		})
	}
}
