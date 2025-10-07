package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е.
// вывести элементы, присутствующие и в первом, и во втором.
//
// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func main() {
	var (
		A, B = []int{1, 2, 3, 5}, []int{2, 4, 5}

		seen = make(map[int]bool)
		res  = make([]int, 0)
	)

	for _, a := range A {
		seen[a] = true
	}

	for _, b := range B {
		if seen[b] {
			res = append(res, b)
		}
	}

	fmt.Println(res)
}
