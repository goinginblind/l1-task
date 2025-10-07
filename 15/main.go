package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
Приведите корректный пример реализации.

	var justString string

	func someFunc() {
		  v := createHugeString(1 &lt;&lt; 10)
		  justString = v[:100] <--
	джастСтринг теперь хранит референс на огромный массив (ненужный) под маленьким слайсом байтов которые действительно используются.
	Этот массив не сможет собрать GC потому что он всё ещё считается используемым целиком.
	}

	func main() {
		  someFunc()
	}
*/

var justString string

func createHugeString(n int) string {
	var buf strings.Builder
	buf.Grow(n)
	for range n {
		buf.WriteRune(rune('A' + rand.Intn(100)))
	}

	return buf.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	// strings.Clone() гарантирует что для новой строки будет выделен новый массив.
	// Массив из-под старой огромной строки соберет GC.
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
