package errorlogger

import (
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
	mu *MutexWrap
	w  io.Writer
}

func (w *mutexWrapWriter) Enable()                          { w.mu.Enable() }
func (w *mutexWrapWriter) Disable()                         { w.mu.Disable() }
func (w *mutexWrapWriter) Lock()                            { w.mu.Lock() }
func (w *mutexWrapWriter) Unlock()                          { w.mu.Unlock() }
func (w mutexWrapWriter) Write(p []byte) (n int, err error) { return w.w.Write(p) }

type mutexEnableWriter struct {
	mu *MutexEnable
	w  io.Writer
}

func (w *mutexEnableWriter) Enable()                          { w.mu.Enable() }
func (w *mutexEnableWriter) Disable()                         { w.mu.Disable() }
func (w *mutexEnableWriter) Lock()                            { w.mu.Lock() }
func (w *mutexEnableWriter) Unlock()                          { w.mu.Unlock() }
func (w mutexEnableWriter) Write(p []byte) (n int, err error) { return w.w.Write(p) }

var writerTestList = []struct {
	name string
	w    io.Writer
}{
	{"mutexWrapWriter", mutexWrapWriter{&MutexWrap{}, NopWriter{}}},
	{"mutexEnableWriter", mutexEnableWriter{&MutexEnable{}, NopWriter{}}},
	{"nopWriter", NopWriter{}},
	{"lenWriter", LenWriter{}},
	// {"os.Stderr", os.Stderr},
}

func BenchmarkWriters(b *testing.B) {
	for _, bb := range writerTestList {
		for i := 0; i < 8; i++ {
			name := fmt.Sprintf("%v (size: %d)", bb.name, i)
			crazyWriterLoop(b, name, bb.w, 1<<i)
		}
	}
}

func flip() bool { return mathrand.Intn(10000)&1 == 1 }

func crazyWriterLoop(b *testing.B, name string, a interface{}, size int) {
	var w io.Writer
	var e Enabler
	var l sync.Locker
	var ok bool
	if w, ok = a.(io.Writer); !ok {
		return
	}

	if e, ok = a.(Enabler); !ok {
		e = nil
	}

	if l, ok = a.(sync.Locker); !ok {
		l = nil
	}

	// b.Logf("Writer: %v", w)
	// b.Logf("Enabler: %v", e)
	// b.Logf("Locker: %v", l)

	var loopsize = 4
	r := rand.Reader
	buf := make([]byte, 0, size)

	// do a lot of time wasting reading and writing ...
	b.Run(name, func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			for j := 0; j < loopsize; j++ {

				// enable and disable randomly and often ...
				if e != nil {
					b.Log("enable writer")
					if flip() {
						a.(Enabler).Enable()
					} else {
						a.(Enabler).Disable()
					}
				}

				// lock and unlock if available
				if l != nil {
					b.Log("lock writer")
					l.Lock()
					defer l.Unlock()
				}

				r.Read(buf)
				n, err := w.Write(buf)

				if err != nil {
					b.Logf("write failed (%v bytes): %v", n, err)
				}
			}
		}
	})
}
