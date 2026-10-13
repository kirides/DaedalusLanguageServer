//go:build wasip1

package main

import (
	"io"
	"os"
	"syscall"
	"time"
	"unsafe"
)

// stdio returns stdin and stdout for a WASI module.
//
// A WASI module is single-threaded, so a blocking read from stdin would also
// stop every other goroutine (parsing, request handling, timers) until the
// client sends its next message. Handing stdin to Go's poller does not help
// either: it needs the runtime to allow non-blocking stdin (wasmtime does not)
// and to implement poll_oneoff well enough for the Go scheduler (wazero fails
// at EOF). So stdin is read by asking poll_oneoff, once, whether data is
// available and otherwise yielding to the other goroutines for a moment.
func stdio() (io.ReadCloser, io.WriteCloser) {
	return &pollingStdin{}, os.Stdout
}

//go:wasmimport wasi_snapshot_preview1 poll_oneoff
func pollOneoff(in, out unsafe.Pointer, nsubscriptions uint32, nevents unsafe.Pointer) uint32

const (
	eventTypeClock  = 0
	eventTypeFdRead = 1

	subscriptionSize = 48
	eventSize        = 32
)

// stdinReadable reports whether a read from stdin would not block.
// ok is false if the runtime cannot tell.
func stdinReadable() (readable, ok bool) {
	var subs [2 * subscriptionSize / 8]uint64
	var events [2 * eventSize / 8]uint64
	var n uint32

	// subscription: userdata u64 | tag u8 (+pad) | union
	// fd_read:      union = file_descriptor u32
	subs[0] = 1 // userdata
	subs[1] = eventTypeFdRead
	subs[2] = 0 // fd 0

	// clock:        union = id u32 (+pad) | timeout u64 | precision u64 | flags u16
	// A relative timeout of zero turns the poll into a readiness check.
	const base = subscriptionSize / 8
	subs[base+0] = 2 // userdata
	subs[base+1] = eventTypeClock
	subs[base+2] = 1 // CLOCK_MONOTONIC
	subs[base+3] = 0 // timeout (ns)
	subs[base+4] = 0 // precision
	subs[base+5] = 0 // flags: relative

	if errno := pollOneoff(unsafe.Pointer(&subs[0]), unsafe.Pointer(&events[0]), 2, unsafe.Pointer(&n)); errno != 0 {
		return false, false
	}
	// event: userdata u64 | error u16 | type u8 (+pad) | fd_readwrite
	for i := uint32(0); i < n && int(i) < 2; i++ {
		if events[i*eventSize/8] == 1 { // userdata of the fd_read subscription
			return true, true
		}
	}
	return false, true
}

type pollingStdin struct{}

func (*pollingStdin) Read(p []byte) (int, error) {
	const (
		minDelay = 500 * time.Microsecond
		maxDelay = 20 * time.Millisecond
	)
	delay := minDelay
	for {
		readable, ok := stdinReadable()
		if readable || !ok {
			n, err := syscall.Read(0, p)
			if err == syscall.EAGAIN {
				// non-blocking stdin that had nothing after all
				n, err = 0, nil
			} else if err == nil && n == 0 && len(p) > 0 {
				return 0, io.EOF
			}
			if n > 0 || err != nil {
				return n, err
			}
		}
		time.Sleep(delay)
		if delay *= 2; delay > maxDelay {
			delay = maxDelay
		}
	}
}

func (*pollingStdin) Close() error { return nil }
