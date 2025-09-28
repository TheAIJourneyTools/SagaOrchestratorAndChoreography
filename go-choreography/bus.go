package main

import "sync"

type Event struct {
	Name    string
	Payload map[string]string
}

type Bus struct {
	subs map[string][]chan Event
	mu   sync.RWMutex
}

func NewBus() *Bus {
	return &Bus{subs: map[string][]chan Event{}}
}

func (b *Bus) Publish(evt Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if chans, ok := b.subs[evt.Name]; ok {
		for _, ch := range chans {
			c := ch
			go func() { c <- evt }()
		}
	}
}

func (b *Bus) Subscribe(eventName string) <-chan Event {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan Event, 10)
	b.subs[eventName] = append(b.subs[eventName], ch)
	return ch
}
