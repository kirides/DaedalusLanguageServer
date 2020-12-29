package langserver

import (
	"sync"
)

// BufferManager ...
type BufferManager struct {
	mtx       sync.RWMutex
	documents map[string]BufferedDocument
}

// NewBufferManager ...
func NewBufferManager() *BufferManager {
	return &BufferManager{
		documents: make(map[string]BufferedDocument),
	}
}

// UpdateBuffer ...
func (m *BufferManager) UpdateBuffer(documentURI string, buf string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	d := BufferedDocument(buf)
	m.documents[documentURI] = d
}

// GetBuffer ...
func (m *BufferManager) GetBuffer(documentURI string) BufferedDocument {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	if doc, ok := m.documents[documentURI]; ok {
		return doc
	}
	return ""
}
