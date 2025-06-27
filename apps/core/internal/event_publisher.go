package internal

import (
	"sync"

	"github.com/piksar-eu/webapp/apps/core/pkg/events"
)

func NewSyncEventPublisher() events.EventPublisher {
	return &SyncEventPublisher{
		listeners: make(map[string][]events.EventListener),
	}
}

type SyncEventPublisher struct {
	mu        sync.RWMutex
	listeners map[string][]events.EventListener
}

func (p *SyncEventPublisher) Publish(events ...events.Event) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, event := range events {
		if ls, ok := p.listeners[event.Type]; ok {
			for _, listener := range ls {
				listener(event)
			}
		}
	}
}

func (p *SyncEventPublisher) Register(eventType string, listener events.EventListener) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.listeners[eventType] = append(p.listeners[eventType], listener)
}
