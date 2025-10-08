package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.), то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {
	cases := []string{
		"Hello World",
		"dlroW olleH",

		"главрыба",
		"абырвалг",

		"你好世界",
		"界世好你",

		"🍕🐹",
		"🐹🍕",

		"глав🐟, 🐹con",
		"noc🐹 ,🐟валг",
	}

	for i := 0; i < len(cases)-1; i += 2 {
		orig, want := cases[i], cases[i+1]
		got := ReverseStringUTF8(orig)

		fmt.Printf("case: %v\n\twant: %v\n\tgot: %v\n\n", orig, want, got)
	}

}

func ReverseStringUTF8(s string) string {
	var (
		buf      strings.Builder
		prevRune = len(s)
	)
	buf.Grow(len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if utf8.RuneStart(s[i]) {
			buf.WriteString(s[i:prevRune])
			prevRune = i
		}
	}

	return buf.String()
}

func ReverseStringRunes(s string) string {
	runes := []rune(s)

	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}

	return string(runes)
}
