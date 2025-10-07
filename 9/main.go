package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.
*/

func main() {
	// generates the stream - not a pipeline stage
	generator := func(done <-chan any, nums ...int) <-chan int {
		numsStream := make(chan int)
		go func() {
			defer close(numsStream)
			for _, n := range nums {
				select {
				case <-done:
					return
				case numsStream <- n:
				}
			}
		}()
		return numsStream
	}

	// multiplication pipe
	multiply := func(done <-chan any, intStream <-chan int, mult int) <-chan int {
		multStream := make(chan int)
		go func() {
			defer close(multStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multStream <- i * mult:
				}
			}
		}()
		return multStream
	}

	// printer/logger pipe
	logResults := func(done <-chan any, intStream <-chan int) <-chan any {
		finished := make(chan any)
		go func() {
			defer close(finished)
			for n := range intStream {
				select {
				case <-done:
					return
				default:
					fmt.Println(n)
				}
			}
		}()
		return finished
	}

	done := make(chan any)
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4, 5)
	<-logResults(done, multiply(done, intStream, 10))
}
