//go:build wasip1

package main

import "context"

// WASI has no signals. Waiting for one would also keep the single thread of
// the module busy: the Go runtime waits for signals in a spin loop there.
func notifyContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}
