package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.

// with sync.RWMMutex
type SafeCounter struct {
	counter int
	mu      sync.RWMutex // RW in case we wanted to read and the writes were still on-going
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	sc.counter++
}

func (sc *SafeCounter) Show() int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()

	return sc.counter
}

// with atomics <- this one is the better version
type AtomicCounter struct {
	counter atomic.Int64
}

func (ac *AtomicCounter) Increment() {
	ac.counter.Add(1)
}

func (ac *AtomicCounter) Show() int64 {
	return ac.counter.Load()
}

func main() {
	const (
		goroutineNum = 1000
		operations   = 100
		opsTotal     = goroutineNum * operations
	)
	var (
		uc int // this one will have a data race
		sc SafeCounter
		ac AtomicCounter
		wg sync.WaitGroup
	)

	for range goroutineNum {
		wg.Go(func() {
			for range operations {
				ac.Increment()
				sc.Increment()
				uc++
			}
		})
	}

	wg.Wait()
	fmt.Printf("safe:\n\twanted %v\n\tgot: %v\n\n", opsTotal, sc.Show())
	fmt.Printf("atomic:\n\twanted %v\n\tgot: %v\n\n", opsTotal, ac.Show())
	fmt.Printf("unsafe (regular int):\n\twanted %v\n\tgot: %v\n\n", opsTotal, uc)
}
