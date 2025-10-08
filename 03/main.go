package main

import (
	"flag"
	"fmt"
	"sync"
)

/* Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров. */

/* Можно поставить флажок, чтобы поменять количество воркеров, по умолчанию их 4.
Пример использования: go run ./3 --w 8 */

type Job struct {
	id   int
	task string
}

func (j Job) Process() {
	fmt.Printf("Job %d processed: %s\n", j.id, j.task)
}

func worker(jobCh chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobCh {
		job.Process()
	}
}

func main() {
	workerCount := flag.Int("w", 4, "Specify the amount of workers")
	flag.Parse()
	tasks := []string{
		"Buy fruit",
		"Eat ice cream",
		"Code something",
		"Learn about concurrency",
		"Attend GopherCon",
		"Learn about GC",
		"Play Elden Ring",
		"Go on a walk",
	}

	jobCh := make(chan Job, 2)
	var wg sync.WaitGroup

	// worker pool
	for id := 1; id <= *workerCount; id++ {
		wg.Add(1)
		go worker(jobCh, &wg)
	}

	// dispense jobs
	for id, task := range tasks {
		jobCh <- Job{id, task}
	}

	close(jobCh)
	wg.Wait()
}
