package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

// Определение структуры Point

type Point struct {
	x, y float64
}

// NewPoint Конструктор для создания нового экземпляра Point
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// DistanceTo Метод для вычисления расстояние между двумя точками с помощью теоремы пифагора
func (p Point) DistanceTo(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	//Вычисление расстояние между точками
	distance := point1.DistanceTo(point2)

	// Вывод результаты
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
