package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Human Определение структуры
type Human struct {
	Name string
	Age  int
	// Дополнительные поля и методы могут быть добавлены по неообходимости
}

// Action Определение структуры и  встраиванием Human
type Action struct {
	Human // Встраивание структуры Human в структуру Action
	// Дополнительные поля и методы могут быть добавлены по необходимости
}

func (a *Action) DoSomething() {
	fmt.Printf("%s is doing something!\n", a.Name)
}

func main() {
	// Создание экземпляра структуры Action
	person := Action{
		Human: Human{
			Name: "Pit",
			Age:  22,
		},
	}

	// Использование методов из струкутры Human и Action
	fmt.Printf("%s, тебе %d года.\n", person.Name, person.Age)
	person.DoSomething()
}
