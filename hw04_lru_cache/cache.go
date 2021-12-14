package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.RWMutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mutex.RLock()
	item, exists := l.items[key]
	l.mutex.RUnlock()

	if exists {
		item.Value = value

		l.mutex.Lock()
		l.queue.MoveToFront(item)
		l.mutex.Unlock()
	} else {
		l.mutex.Lock()
		listItem := l.queue.PushFront(value)
		l.items[key] = listItem
		l.mutex.Unlock()
	}

	return exists
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.RLock()
	item, exists := l.items[key]
	l.mutex.RUnlock()

	if !exists {
		return nil, false
	}

	l.mutex.Lock()
	l.queue.MoveToFront(item)
	l.mutex.Unlock()

	return item.Value, true
}

func (l *lruCache) Clear() {
	l.capacity = 0
	l.queue = NewList()
	l.items = make(map[Key]*ListItem)
	l.mutex = sync.RWMutex{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    sync.RWMutex{},
	}
}
