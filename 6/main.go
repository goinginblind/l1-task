package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления,
// через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

const (
	workDelay    = time.Millisecond * 500
	exampleDelay = time.Second * 3
)

func main() {
	// Выход через условие
	go func() {
		for i := range 10 {
			if i == 5 {
				fmt.Println("Goroutine I shuts down...")
				return
			}
			fmt.Println("Goroutine I does job #", i)
			time.Sleep(workDelay)
		}
	}()
	time.Sleep(exampleDelay)

	// Выход через уведомление из другого канала
	done := make(chan any)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Goroutine II shuts down...")
				return
			default:
				fmt.Println("Goroutine II does some work")
				time.Sleep(workDelay)
			}
		}
	}()
	time.Sleep(exampleDelay)
	close(done)

	// Выход через контекст
	ctx, cancel := context.WithTimeout(context.Background(), workDelay*4)
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine III shuts down...")
				return
			default:
				fmt.Println("Goroutine III does some work")
				time.Sleep(workDelay)
			}
		}
	}(ctx)
	time.Sleep(exampleDelay)

	// Выход через runtime.Goexit()
	go func() {
		fmt.Println("Goroutine IV says \"Hi!\"")
		runtime.Goexit()
		fmt.Println("I'm dead, I can't talk.")
	}()
	time.Sleep(exampleDelay)
}
