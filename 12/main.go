package main

import "fmt"

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	var (
		words   = []string{"cat", "dog", "dog", "hotdog", "cat", "bear", "pineapple"}
		seen    = make(map[string]struct{})
		uniques = []string{}
	)

	for _, w := range words {
		if _, ok := seen[w]; !ok {
			seen[w] = struct{}{}
			uniques = append(uniques, w)
		}
	}

	fmt.Println(uniques)
}
