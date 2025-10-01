package main

import (
	"fmt"
	"sync"
)

// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

type SafeMap[T1 comparable, T2 any] struct {
	data map[T1]T2
	mu   sync.RWMutex
}

func NewSafeMap[T1 comparable, T2 any]() *SafeMap[T1, T2] {
	return &SafeMap[T1, T2]{
		data: make(map[T1]T2),
	}
}

func (sm *SafeMap[T1, T2]) Read(key T1) (T2, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	val, ok := sm.data[key]
	return val, ok
}

func (sm *SafeMap[T1, T2]) Write(key T1, val T2) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.data[key] = val
}

func (sm *SafeMap[T1, T2]) Length() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return len(sm.data)
}

func main() {
	myMap := NewSafeMap[int, string]()
	var wg sync.WaitGroup
	for i := range 1000 {
		wg.Go(func() {
			myMap.Write(i, "some value")
		})
	}
	wg.Go(func() {
		for k := range 1000 {
			if k%333 == 0 {
				val, ok := myMap.Read(k)
				fmt.Printf("Key: %d; Exists: %v; Value: %v\n", k, ok, val)
				fmt.Printf("Current length: %d\n\n", myMap.Length())
			}
		}
	})
	wg.Wait()
}
