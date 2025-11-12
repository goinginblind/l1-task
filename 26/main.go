package main

import (
	"fmt"
	"unicode"
)

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз
// (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения.
// Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
//
// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
//
// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

func IsUnique(s string) bool {
	seen := make(map[rune]struct{})
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}

	return true
}

// Won't work for most of the chars besides english letters.
func IsUniqueWithMask(s string) bool {
	var mask uint64
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 32
		}
		bit := r - 'a'

		if mask&(1<<bit) != 0 {
			return false
		}
		mask |= 1 << bit
	}

	return true
}

func main() {
	funcs := map[string]func(string) bool{
		"map":  IsUnique,
		"mask": IsUniqueWithMask,
	}
	examples := []struct {
		s    string
		want bool
	}{
		{
			s:    "abcd",
			want: true,
		},
		{
			s:    "abCdefAaf",
			want: false,
		},
		{
			s:    "aabcd",
			want: false,
		},
	}

	for fname, foo := range funcs {
		fmt.Printf("|=============== '%v' ===============|\n", fname)
		for _, ex := range examples {
			got := foo(ex.s)
			fmt.Printf(" result for '%v':\n\twant - '%v'\n\t got - '%v'\n\n", ex.s, ex.want, got)
		}
	}
}
