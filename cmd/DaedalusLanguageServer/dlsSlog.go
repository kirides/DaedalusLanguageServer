package main

import (
	"log"

	"log/slog"

	dls "github.com/kirides/DaedalusLanguageServer"
)

type dlsSlog struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func newDlsSlog(handler slog.Handler) *dlsSlog {
	return &dlsSlog{
		debug: slog.NewLogLogger(handler, slog.LevelDebug),
		info:  slog.NewLogLogger(handler, slog.LevelInfo),
		warn:  slog.NewLogLogger(handler, slog.LevelWarn),
		error: slog.NewLogLogger(handler, slog.LevelError),
	}
}

func (d *dlsSlog) Debugf(template string, args ...interface{}) { d.debug.Printf(template, args...) }
func (d *dlsSlog) Infof(template string, args ...interface{})  { d.info.Printf(template, args...) }
func (d *dlsSlog) Warnf(template string, args ...interface{})  { d.warn.Printf(template, args...) }
func (d *dlsSlog) Errorf(template string, args ...interface{}) { d.error.Printf(template, args...) }

var _ dls.Logger = (*dlsSlog)(nil)
