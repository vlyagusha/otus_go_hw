package hw04lrucache

import "sync"

type Key string

type cacheItem struct {
	key   Key
	value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type CacheMultithreading struct {
	cache Cache
	mutex sync.Mutex
}

func (c *CacheMultithreading) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	was := c.cache.Set(key, value)
	defer c.mutex.Unlock()

	return was
}

func (c *CacheMultithreading) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	value, exists := c.cache.Get(key)
	defer c.mutex.Unlock()

	return value, exists
}

func (c *CacheMultithreading) Clear() {
	c.mutex.Lock()
	c.cache.Clear()
	defer c.mutex.Unlock()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	item, exists := l.items[key]
	if exists {
		item.Value = cacheItem{
			key:   key,
			value: value,
		}
		l.queue.MoveToFront(item)

		return true
	}

	if l.queue.Len() >= l.capacity {
		delete(l.items, l.queue.Back().Value.(cacheItem).key)
		l.queue.Remove(l.queue.Back())
	}
	listItem := l.queue.PushFront(cacheItem{
		key:   key,
		value: value,
	})
	l.items[key] = listItem

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, exists := l.items[key]

	if !exists {
		return nil, false
	}

	l.queue.MoveToFront(item)

	return item.Value.(cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &CacheMultithreading{
		cache: &lruCache{
			capacity: capacity,
			queue:    NewList(),
			items:    make(map[Key]*ListItem, capacity),
		},
		mutex: sync.Mutex{},
	}
}
