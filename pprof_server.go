package main

import (
	"context"
	"fmt"
	"net/http"
)

type pprofServer struct {
	pprofAddr string

	server  *http.Server
	running bool
}

func (p *pprofServer) ChangeAddr(addr string) {
	if p.pprofAddr == addr {
		return
	}

	var a, b, c, d byte
	var port uint16
	if _, err := fmt.Sscanf(addr, "%d.%d.%d.%d:%d", &a, &b, &c, &d, &port); err != nil {
		return
	}
	if port == 0 {
		return
	}
	addr = fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, port)

	p.pprofAddr = addr
	p.Start()
}

func (p *pprofServer) Start() {
	if p.server == nil {
		p.server = &http.Server{Handler: http.DefaultServeMux}
		p.running = false
	}
	if p.running {
		return
	}
	if p.pprofAddr == "" {
		return
	}

	defer recover()
	p.server.Addr = p.pprofAddr
	p.running = true

	go func() {
		p.server.ListenAndServe()
		p.running = false
	}()
}
func (p *pprofServer) Stop() {
	if p.server != nil {
		p.server.Shutdown(context.Background())
	}
	p.server = nil
}
