package langserver

import "sync"

type concurrentSet[K comparable] struct {
	m sync.Map
}

func (c *concurrentSet[K]) Store(key K) {
	c.m.Store(key, struct{}{})
}

func (c *concurrentSet[K]) Contains(key K) bool {
	_, loaded := c.m.Load(key)
	return loaded
}
