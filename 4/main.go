package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.*/

type Job struct {
	id   int
	task string
}

func (j Job) Process() {
	log.Printf("Job %d: did \"%s\".", j.id, j.task)
}

func worker(ctx context.Context, jobCh chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobCh:
			if !ok {
				return
			}
			j.Process()
			time.Sleep(time.Second * 1)
		}
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

	jobCh := make(chan Job, 4)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// create workers
	for range *workerCount {
		wg.Add(1)
		go worker(ctx, jobCh, &wg)
	}

	// dispense jobs
	go func() {
		for id, task := range tasks {
			select {
			case <-ctx.Done():
				return
			case jobCh <- Job{id: id, task: task}:
			}
		}
		close(jobCh) // normal termination
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	cancel()
	wg.Wait()
}
