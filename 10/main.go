package main

import (
	"fmt"
	"math"
	"slices"
)

// Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить эти значения в группы с шагом 10 градусов.
// Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
// Пояснение: диапазон -20 включает значения от -20 до -29.9, диапазон 10 – от 10 до 19.9, и т.д. Порядок в подмножествах не важен.

func main() {
	round := func(num float64) float64 {
		if num >= 0 {
			return math.Floor(num/10) * 10
		}
		return math.Ceil(num/10) * 10
	}

	groups, temps := []float64{}, []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := make(map[float64][]float64)

	for _, temp := range temps {
		group := round(temp)
		if _, ok := tempMap[group]; !ok {
			groups = append(groups, group)
		}

		tempMap[group] = append(tempMap[group], temp)
	}
	slices.Sort(groups)

	for _, group := range groups {
		fmt.Printf("%v: %v\n", group, tempMap[group])
	}
}
