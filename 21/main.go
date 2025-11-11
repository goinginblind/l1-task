package main

import (
	"container/list"
	"context"
	"errors"
	"sync"
)

/*
   Реализовать паттерн проектирования «Адаптер» на любом примере.
   Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.
   Продемонстрируйте на простом примере в Go:
   у вас есть существующий интерфейс (или структура) и другой,
   несовместимый по интерфейсу потребитель — напишите адаптер,
   который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.
   Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.
*/

// Cache - интерфейс, который ожидает наш клиент. Например, если мы пишем тесты но у нас нет возможности
// воспользоваться настоящим редисом, но есть вот такой вот простой ручной кэш, то мы можем воспользоваться
// адпатером. В общем, адпатер применяется в следующих ситуациях:
//
//  1. Замена одной зависимости на другую (например для тестов)
//  2. Две структуры с разными методами можно обернуть в струкутры, у которых методы подходят под один интерфейс
//     (например, интерфейс ObjectStorage который будет совместим с адаптером для S3 и GCS,
//     тут ключевое что это могут быть два файловых хранилища, у которых методы по сути одинаковые,
//     но различаются сигнатуры этих методов).
//  3. Интеграция внешних пакетов друг с другом - мы не можем поменять ни один из них, тогда тут можно использовать адпатеры
//     (например, пакет А ожидает логгер обычный, а мы используем зап - пишем адпатер зап-обычный логгер)
//  4. Интеграция существующего кода с новым, например, если мы меняем один пакет на другой - весь проект всё ещё ожидает
//     интерфейс старого типа, а наш новый пакет не удовлетворяет такому интерфейсу - тут можно пользоваться адпатером.
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
}

type SimpleCacheAdapter struct {
	c *SimpleLRUCache
}

func (sca *SimpleCacheAdapter) Get(ctx context.Context, key string) (string, error) {
	val, ok := sca.c.Get(key)
	if !ok {
		return "", errors.New("key not found")
	}

	return val, nil
}

func (sca *SimpleCacheAdapter) Set(ctx context.Context, key, value string) error {
	sca.c.Insert(key, value)
	return nil
}

// SimpleLRUCache is a cache with keys being held in a map with values stored as pointers
// to list.Elements. Entries themselves are held in the doubly linked list. The cache can be used
// concurrently.
type SimpleLRUCache struct {
	keys    map[string]*list.Element
	entries *list.List

	entriesCap int

	mu sync.RWMutex
}

// cacheEntry holds a key-value pair,
// which allows us to easily remove cache entries
// from the linked list of SimpleLRUCache
type cacheEntry struct {
	key string
	val string
}

// Get retrieves a value for a key, if it doesn't exist it returns
// an empty string and a 'false', if it exists, then the value and 'true'.
// It is concurrency-safe, uses a read-lock.
func (slc *SimpleLRUCache) Get(key string) (string, bool) {
	slc.mu.RLock()
	defer slc.mu.RUnlock()

	listElem, ok := slc.keys[key]
	if !ok {
		return "", false
	}

	entry := listElem.Value.(cacheEntry)
	return entry.val, true
}

// Insert puts a key value pair into cache, its concurrency-safe and
// if the amount of cache entries exceeds the slc.entriesCap, it then
// evicts the least recently used entry. If the key already exists, Insert
// updates the entry.
func (slc *SimpleLRUCache) Insert(key, value string) {
	slc.mu.Lock()
	defer slc.mu.Unlock()

	var (
		// that's the linked list that holds
		// actual cache entries
		ll = slc.entries

		// cacheEntry to be inserted into the
		// linked list
		ce = cacheEntry{
			key: key,
			val: value,
		}
	)

	listElem, ok := slc.keys[key]
	if !ok {
		elem := ll.PushFront(ce)
		slc.keys[key] = elem
	} else {
		listElem.Value = ce
		ll.MoveToFront(listElem)
	}

	if ll.Len() > slc.entriesCap {
		lruElem := ll.Back()

		if lruElem != nil {
			lruEntry := lruElem.Value.(cacheEntry)
			delete(slc.keys, lruEntry.key)
			ll.Remove(lruElem)
		}
	}
}

// NewManualCache creates a new instance of a simple lru cache
// and returns a pointer to it. If the argument provided is less
// than 1, it sets the entriesCap of the cache to 1.
func NewManualCache(entriesCap int) *SimpleLRUCache {
	if entriesCap < 1 {
		entriesCap = 1
	}
	return &SimpleLRUCache{
		keys:       make(map[string]*list.Element, entriesCap),
		entries:    list.New(),
		entriesCap: entriesCap,
	}
}

func main() {
}
