package hw04lrucache

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

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	if node, inIt := cache.items[key]; inIt {
		node.Value.(*cacheItem).value = value
		cache.queue.MoveToFront(node)
		return true
	}

	if cache.queue.Len() == cache.capacity {
		cache.Clear()
	}

	item := &cacheItem{key: key, value: value}
	node := cache.queue.PushFront(item)
	cache.items[item.key] = node

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if _, inIt := cache.items[key]; inIt {
		cache.queue.MoveToFront(cache.items[key])
		cache.items[key] = cache.queue.Front()
		return cache.items[key].Value.(*cacheItem).value, true
	} else {
		return nil, false
	}
}

func (cache *lruCache) Clear() {
	item := cache.queue.Back().Value.(*cacheItem)
	delete(cache.items, item.key)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
