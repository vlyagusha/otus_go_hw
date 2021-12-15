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
}

var mutex = sync.Mutex{}

func (l *lruCache) Set(key Key, value interface{}) bool {
	mutex.Lock()

	item, exists := l.items[key]
	if exists {
		item.Value = cacheItem{
			Key:   key,
			Value: value,
		}

		l.queue.MoveToFront(item)
		mutex.Unlock()

		return true
	}

	listItem := l.queue.PushFront(cacheItem{
		Key:   key,
		Value: value,
	})
	l.items[key] = listItem

	mutex.Unlock()

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	mutex.Lock()
	item, exists := l.items[key]

	if !exists {
		mutex.Unlock()

		return nil, false
	}

	l.queue.MoveToFront(item)
	mutex.Unlock()

	return item.Value.(cacheItem).Value, true
}

func (l *lruCache) Clear() {
	mutex.Lock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
	mutex.Unlock()
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
