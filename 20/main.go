package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает порядок слов в строке.
// Пример:
//
//	входная строка: «snow dog sun»
//	выход: «sun dog snow».
//
// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
func main() {
	cases := []string{
		"snow dog sun",
		"hello helo worldle",
		"privet mir",
		"BigGiantSingleWord",
		"Я люблю Россию",
		"M a n y S m a l l W o r d s",
	}

	for i, c := range cases {
		fmt.Printf("Case %v:\n", i)
		fmt.Printf("\tOriginal: %v\n", c)
		fmt.Printf("\tReversed: %v\n", ReverseSentence(c))
	}
}

func ReverseSentence(s string) string {
	var buf strings.Builder
	buf.Grow(len(s))
	prevWhitespace := len(s)

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' {
			buf.WriteString(s[i+1:prevWhitespace] + " ")
			prevWhitespace = i
		}
	}
	buf.WriteString(s[0:prevWhitespace]) // add the last word

	return buf.String()
}
