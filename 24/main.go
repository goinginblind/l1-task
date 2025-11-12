package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками на плоскости.
// Точки представлены в виде структуры Point с инкапсулированными (приватными)
// полями x, y (типа float64) и конструктором.
// Расстояние рассчитывается по формуле между координатами двух точек.
// Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))
}

func main() {
	var (
		// A = NewPoint(0, 0)
		B = NewPoint(4.23, 1)
		C = NewPoint(0.17, 3.12)
	)

	fmt.Printf("%.03f", B.Distance(C))
}
