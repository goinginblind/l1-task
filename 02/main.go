package main

import (
	"fmt"
	"sync"
)

/* Написать программу, которая конкурентно рассчитает значения квадратов чисел,
взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат. */

func main() {
	nums := []int{2, 4, 6, 8, 10}
	res := make([]int, len(nums))

	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(1)
		go func(i, num int) {
			defer wg.Done()
			res[i] = num * num
		}(i, n)
	}

	wg.Wait()

	fmt.Println(res)
}
