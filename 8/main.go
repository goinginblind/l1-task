package main

import (
	"flag"
	"fmt"
)

/* Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
Подсказка: используйте битовые операции (|, &^). */

// example use: go run . -v 5 -b 0 -o false --> forces the 0th bit to be 0, prints 4
func main() {
	var (
		v = flag.Int64("v", 5, "the value to modify")
		b = flag.Uint("b", 0, "0-based index of the bit to change")
		o = flag.Bool("o", false, "set to zero - false, to one - true")
	)
	flag.Parse()

	var (
		val    = *v
		bitIdx = *b
		toOne  = *o
	)

	mask := int64(1) << bitIdx

	if toOne {
		val = val | mask
	} else {
		val = val &^ mask
	}

	fmt.Println(val)
}
