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
	// Приводим ключ и значение к объекту cacheItem
	item := cacheItem{key: key, value: value}

	// Проверяем, есть ли элемент в словаре
	_, inIt := cache.items[key]
	if inIt {
		// Если есть, то даем ему новое значение
		cache.items[key].Value = item.value
		// Двигаем в начало очереди
		cache.queue.MoveToFront(cache.items[key])
	} else {
		// Если размер очереди == размеру очереди кэша удаляем последний элемент кэша и его значение из словаря
		if cache.queue.Len() == cache.capacity {
			backItem := cache.queue.Back()
			for backItemKey, backItemValue := range cache.items {
				if backItemValue == backItem {
					delete(cache.items, backItemKey)
				}
			}
			cache.queue.Remove(cache.queue.Back())

		}
		cache.items[key] = cache.queue.PushFront(value)
	}
	return inIt
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	_, inIt := cache.items[key]
	if inIt {
		cache.queue.MoveToFront(cache.items[key])
		return cache.items[key].Value, true
	} else {
		return nil, false
	}
}

func (cache *lruCache) Clear() {
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
