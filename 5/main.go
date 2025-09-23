package main

import (
	"flag"
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер для ограничения времени работы.

func main() {
	timeLimit := flag.Duration("N", time.Second*3, "Execution time limit: e.g. '5s' is 5 seconds")
	flag.Parse()
	done := time.After(*timeLimit)

	sender := func() <-chan int {
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for i := range 100 {
				select {
				case <-done:
					fmt.Printf("Time limit exceeded: %v\n", *timeLimit)
					return
				case valueStream <- i:
					time.Sleep(time.Second * 1)
				}
			}
		}()
		return valueStream
	}

	receiver := sender()
	for i := range receiver {
		fmt.Println("Received value:", i)
	}

	fmt.Println("Shutting down...")
}
