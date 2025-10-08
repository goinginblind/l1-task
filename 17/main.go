package main

import "fmt"

// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.

func main() {
	numsRegular := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numsMissing := []int{1, 2, 4, 6, 7, 8, 11, 15, 16}

	fmt.Printf("case 1: regular\n\texpected: %v\n\tgot: %v\n", 1, binarySearch(2, numsRegular))
	fmt.Printf("case 2: search too far right\n\texpected: %v\n\tgot: %v\n", -1, binarySearch(13, numsRegular))
	fmt.Printf("case 3: search too far left\n\texpected: %v\n\tgot: %v\n", -1, binarySearch(-1, numsRegular))
	fmt.Printf("case 4: search even amount of numbers\n\texpected: %v\n\tgot: %v\n", 2, binarySearch(3, numsRegular[:len(numsRegular)-1]))

	fmt.Printf("case 5:\n\texpected: %v\n\tgot: %v\n", 3, binarySearch(6, numsMissing))
	fmt.Printf("case 6:\n\texpected: %v\n\tgot: %v\n", 6, binarySearch(11, numsMissing))
	fmt.Printf("case 7:\n\texpected: %v\n\tgot: %v\n", -1, binarySearch(3, numsMissing))
}

func binarySearch(tgt int, nums []int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == tgt {
			return mid
		}

		if nums[mid] < tgt {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
