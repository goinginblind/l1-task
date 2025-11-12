package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.
// Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
// (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
func main() {
	s := []string{"apple", "banana", "cat", "dog"}
	fmt.Printf("before: %v\n\tlen: %v\n\tcap: %v\n\n", s, len(s), cap(s))

	s = DeleteIthElement(s, 3)
	fmt.Printf("after i=3 deletion: %v\n\tlen: %v\n\tcap: %v\n\n", s, len(s), cap(s))

	s = DeleteIthElement(s, 1)
	fmt.Printf("after i=1 deletion: %v\n\tlen: %v\n\tcap: %v\n\n", s, len(s), cap(s))

	s = append(s, "tiger")
	fmt.Printf("after appending: %v\n\tlen: %v\n\tcap: %v\n\n", s, len(s), cap(s))
}

// DeleteIthElement takes a slice s and copies all the values
// starting at indexes i+1 into slice s[i:].
func DeleteIthElement[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}

	copy(s[i:], s[i+1:])
	clear(s[len(s)-1:]) // this one is for the GC to pick up underlying memory, if T is a reference type
	return s[:len(s)-1]
}
