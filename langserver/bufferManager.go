package langserver

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// BufferManager ...
type BufferManager struct {
	documents map[string]BufferedDocument
	mtx       sync.RWMutex
}

// NewBufferManager ...
func NewBufferManager() *BufferManager {
	return &BufferManager{
		documents: make(map[string]BufferedDocument),
	}
}

// DeleteBuffer ...
func (m *BufferManager) DeleteBuffer(documentURI string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.documents, documentURI)
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

// GetBuffer ...
func (m *BufferManager) GetBufferCtx(ctx context.Context, documentURI string) (BufferedDocument, error) {
	ticker := time.NewTicker(60 * time.Second) // 60s delay until request is aborted
	defer ticker.Stop()
	for {
		m.mtx.RLock()
		doc, ok := m.documents[documentURI]
		m.mtx.RUnlock()
		if ok {
			return doc, nil
		}
		select {
		case <-ticker.C:
			return "", fmt.Errorf("timeout: document %q not found", documentURI)
		case <-time.After(100 * time.Millisecond):
		case <-ctx.Done():
			return "", fmt.Errorf("document %q not found", documentURI)
		}
	}
}
