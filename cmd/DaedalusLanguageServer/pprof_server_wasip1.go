//go:build wasip1

package main

// WASI has no sockets, there is nothing to serve pprof on.
type pprofServer struct{}

func (*pprofServer) ChangeAddr(string) {}
func (*pprofServer) Stop()             {}
