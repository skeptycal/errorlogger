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

// benchmark results
/*
Benchmark_errorLogger_noErr_yesErr/noErr(fake_error)-8         	358696474	         3.501 ns/op	       0 B/op	       0 allocs/op
Benchmark_errorLogger_noErr_yesErr/yesErr(fake_error)-8        	  468074	      2491 ns/op	     456 B/op	      15 allocs/op
Benchmark_errorLogger_noErr_yesErr/noErr(fake_error)#01-8      	334853433	         3.277 ns/op	       0 B/op	       0 allocs/op
Benchmark_errorLogger_noErr_yesErr/yesErr(fake_error)#01-8     	  244546	      4559 ns/op	     995 B/op	      23 allocs/op
Benchmark_errorLogger_noErr_yesErr/noErr(nil)-8                	348771727	         3.332 ns/op	       0 B/op	       0 allocs/op
Benchmark_errorLogger_noErr_yesErr/yesErr(nil)-8               	199257543	         6.007 ns/op	       0 B/op	       0 allocs/op
Benchmark_errorLogger_noErr_yesErr/noErr(nil)#01-8             	378662402	         3.169 ns/op	       0 B/op	       0 allocs/op
Benchmark_errorLogger_noErr_yesErr/yesErr(nil)#01-8            	199490152	         6.470 ns/op	       0 B/op	       0 allocs/op
*/

func Benchmark_errorLogger_noErr_yesErr(b *testing.B) {
	tests := []struct {
		name    string
		errName string
		input   error
		want    error
	}{
		// TODO: Add test cases.
		{"yesnologger", "", errFake, errFake}, // return an error unchanged
		{"default", "nil", nil, nil},          // return nil unchanged
	}

	for _, bb := range tests {
		for _, test := range privateStructTests {

			errFuncTests := []struct {
				name string
				fn   func(err error) error
			}{
				{"noErr", test.e.noErr},
				{"noErr", test.e.noErr1},
				{"noErr", test.e.yesErr},
			}

			for _, fnTest := range errFuncTests {
				b.Run(bb.name+"."+fnTest.name+"("+bb.name+")", func(b *testing.B) {
					var fake error
					for i := 0; i < b.N; i++ {
						fake = fnTest.fn(bb.input)
					}
					fakeOuter = fake
				})
			}
		}
	}
}

// func Test_errorLogger_noErr_yesErr(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input error
// 		want  error
// 	}{
// 		// TODO: Add test cases.
// 		{"fake error", errFake, errFake}, // return an error unchanged
// 		{"nil", nil, nil},                // return nil unchanged
// 	}

// 	for _, tt := range tests {
// 		for _, test := range errFuncTestList {
// 			t.Run(test.name+"("+tt.name+")", func(t *testing.T) {
// 				if got := test.fn(tt.input); got != tt.want {
// 					t.Errorf("errorLogger.noErr() error = got %v, want %v", got, tt.want)
// 				}
// 			})
// 		}
// 	}
// }
