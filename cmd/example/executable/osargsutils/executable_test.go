package osargsutils

import "testing"

/*
hereMe2() got =
/private/var/folders/f9/fw_fbd1x00s3yhkldmcfz28w0000gn/T/go-build505778189/b001
want - from HereMe() using path.Split() instead of of Dir() and Base()
/private/var/folders/f9/fw_fbd1x00s3yhkldmcfz28w0000gn/T/go-build505778189/b001/
*/

// Benchmark Results
/*
/// It does not matter ... in fact, simply calling the argzero function to get the initial
/// path is what takes up all of the time.

/// Will any of this matter? Not at all ... this function is likely to be called only once
/// during the runtime of the program. a few hundred nanoseconds will not matter ...

/// however, if 100,000 of these are started up in a scaling architecture, run quickly,
/// and exit nearly immediately ... well then the startup code has a huge effect.

BenchmarkHereMe2-8            	   20613	     54626 ns/op	    3680 B/op	      42 allocs/op
BenchmarkHereMe-8             	   23355	     60157 ns/op	    3680 B/op	      42 allocs/op
BenchmarkZeroOsExecutable-8   	   22617	     51969 ns/op	    3680 B/op	      42 allocs/op
BenchmarkZeroOsArgs-8         	   24009	     53699 ns/op	    3680 B/op	      42 allocs/op
BenchmarkRawOsArgsZero-8      	   23281	     49985 ns/op	    3680 B/op	      42 allocs/op
*/

func BenchmarkHereMe2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hereMe2()
	}
}

func BenchmarkHereMe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HereMe()
	}
}

func BenchmarkZeroOsExecutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroOsExecutable()
	}
}

func BenchmarkZeroOsArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroOsArgs()
	}
}

func BenchmarkRawOsArgsZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zeroOsArgs()
	}
}

func Test_hereMe2(t *testing.T) {

	t.Run("hereMe2", func(t *testing.T) {
		got, got1, err := hereMe2()
		want, want1, _ := HereMe()
		wantErr := false
		if (err != nil) != wantErr {
			t.Errorf("hereMe2() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got != want {
			t.Errorf("hereMe2() got = %v, want %v", got, want)
		}
		if got1 != want1 {
			t.Errorf("hereMe2() got1 = %v, want %v", got1, want1)
		}
	})
}
