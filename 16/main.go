package main

import "fmt"

// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.
func main() {
	nums := []int{1, 2, 7, 5, 3, 6, 8, 4}
	QuickSort(nums)
	fmt.Println(nums)
}

// QuickSort() sorts a slice in-place using quick sorting.
func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	pi := partition(nums, 0, len(nums)-1)

	QuickSort(nums[:pi])
	QuickSort(nums[pi+1:])
}

// partition() partitions the slice, picking its last element as
// the pivot point and returning the pivots new index.
func partition(nums []int, low, high int) int {
	var (
		pivot = nums[high]
		i     = low - 1
	)

	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[high] = nums[high], nums[i+1]

	return i + 1
}
