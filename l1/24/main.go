package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками на плоскости.
	Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
	Расстояние рассчитывается по формуле между координатами двух точек.

	Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.
*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(0, 0)

	fmt.Printf("Distance: %.2f\n", p1.Distance(p2))
}
