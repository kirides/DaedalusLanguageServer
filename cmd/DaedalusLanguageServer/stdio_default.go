//go:build !wasip1

package main

import (
	"io"
	"os"
)

func stdio() (io.ReadCloser, io.WriteCloser) { return os.Stdin, os.Stdout }
