package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Выталкивание элементов из-за размера очереди

		c := NewCache(3)

		wasInCacheFirst := c.Set("first", 100)
		require.False(t, wasInCacheFirst)

		wasInCacheSecond := c.Set("second", 200)
		require.False(t, wasInCacheSecond)

		wasInCacheThird := c.Set("third", 300)
		require.False(t, wasInCacheThird)

		wasInCacheFourth := c.Set("fourth", 400)
		require.False(t, wasInCacheFourth)

		item, getFirst := c.Get("first")
		require.Equal(t, nil, item)
		require.False(t, getFirst)

	})

	t.Run("purge LRU logic", func(t *testing.T) {
		// Выталкивание давно используемых элементов

		c := NewCache(3)
		// Добавляем 3 объекта в кэш
		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300)
		require.False(t, wasInCache)

		// Два из них запрашиваем
		value, inIt := c.Get("aaa")
		require.Equal(t, 100, value)
		require.True(t, inIt)

		value, inIt = c.Get("ccc")
		require.Equal(t, 300, value)
		require.True(t, inIt)

		// Добавляем новый объект в кэш, он переполняется -> элемент, с которым дольше всего не работали, удаляется из кэша
		wasInCache = c.Set("ddd", 400)
		require.False(t, wasInCache)

		// Запрашиваем удаленный элемент
		item, inIt := c.Get("bbb")
		require.False(t, inIt)
		require.Equal(t, nil, item)

	})

	t.Run("Clear cache", func(t *testing.T) {
		c := NewCache(3)
		c.Set("first", 100)
		c.Set("second", 200)
		c.Set("third", 300)
		c.Clear()

		first, inIt := c.Get("first")
		require.False(t, inIt)
		require.Equal(t, nil, first)

		second, inIt := c.Get("second")
		require.False(t, inIt)
		require.Equal(t, nil, second)

		third, inIt := c.Get("third")
		require.False(t, inIt)
		require.Equal(t, nil, third)

	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
