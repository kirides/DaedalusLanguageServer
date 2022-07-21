package langserver

import "sync"

type concurrentMap[K comparable, V any] struct {
	m sync.Map
}

func (c *concurrentMap[K, V]) Store(key K, value V) {
	c.m.Store(key, value)
}

func (c *concurrentMap[K, V]) Load(key K) (V, bool) {
	var def V
	v, loaded := c.m.Load(key)

	if !loaded {
		return def, false
	}
	return v.(V), true
}

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
