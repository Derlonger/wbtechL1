package main

import "fmt"

/*
Интерфейсы в программировании представляют собой контракты, определяющие,
какие методы должен реализовать тип данных. В языке программирования Go интерфейсы
играют важную роль и используются для достижения полиморфизма без необходимости явного наследования.

В Go интерфейс определяется набором методов, но не содержит их реализации.
Тип данных автоматически считается реализующим интерфейс,
если он предоставляет все необходимые методы.
*/

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	printArea(circle) // Вывод: Area: 78.5
}

// Shape Определение интерфейса с методом
type Shape interface {
	Area() float64
}

// Circle Реализация интерфейса для круга
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
