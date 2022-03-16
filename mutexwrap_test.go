package errorlogger

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"sync"
	"testing"
)

// An Enabler represents an object that can be enabled or disabled.
//
// Reference: http://github.com/skeptycal/types
type Enabler interface {
	Enable()
	Disable()
}

type mutexWrapWriter struct {
	mu MutexWrap
	io.Writer
}

type mutexEnableWriter struct {
	mu MutexEnable
	io.Writer
}

var writerTestList = []struct {
	name string
	w    io.Writer
}{
	{"mutexWrapWriter", mutexWrapWriter{MutexWrap{}, NopWriter{}}},
	{"mutexEnableWriter", mutexEnableWriter{MutexEnable{}, NopWriter{}}},
	{"nopWriter", NopWriter{}},
	{"lenWriter", LenWriter{}},
}

func BenchmarkWriters(b *testing.B) {
	for _, bb := range writerTestList {
		i := 2
		// for i := 0; i < 4; i++ {
		name := fmt.Sprintf("%v (size: %d)", bb.name, i)
		crazyWriterLoop(b, name, bb.w, 1<<i)

	}
	// }
}

func crazyWriterLoop(b *testing.B, name string, w io.Writer, size int) {
	var loopsize = 32
	r := rand.Reader
	no := make([]byte, 0, size)
	buf := bytes.NewBuffer(no)

	// do a lot of time wasting reading and writing ...
	b.Run(name, func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			for j := 0; j < loopsize; j++ {

				// enable and disable randomly and often ...
				if v, ok := w.(Enabler); ok {
					n := mathrand.Intn(10000)
					if n&1 == 0 {
						v.Enable()
					} else {
						v.Disable()
					}
				}

				// lock and unlock if available
				if v, ok := w.(sync.Locker); ok {
					v.Lock()
					defer v.Unlock()
				}

				buf.Reset()
				n, err := io.Copy(buf, r)
				if err != nil {
					b.Logf("io.Copy failed (%v bytes): %v", n, err)
				}
				n, err = io.Copy(w, buf)
				if err != nil {
					b.Logf("io.Copy failed (%v bytes): %v", n, err)
				}
			}
		}
	})
}
