package main

import "fmt"

// swap in-place
// a, b = b, a

func main() {
	a, b := 7, 9
	fmt.Printf("original:\n \t%v and %v\n\n", a, b)

	// using bit ops
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("swapped via bitwise operations:\n \t%v and %v\n\n", a, b)

	// or swapped via subtraction
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("swapped with the power of MATH:\n \t%v and %v\n\n", a, b)
}
