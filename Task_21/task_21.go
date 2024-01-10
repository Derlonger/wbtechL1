package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Animal Target - интерфейс, который мы хотим использовать в клиентском коде.
type Animal interface {
	MakeSound()
}

// Dog Adaptee - класс, который мы хотим адаптировать.
type Dog struct{}

func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав. Хозяин, дай пожрать.")
}

// DogAdapter Adapter - адаптер, реализующий интерфейс Target и использующий Adaptee.
type DogAdapter struct {
	*Dog
}

func (adapter *DogAdapter) MakeSound() {
	adapter.WoofWoof()
}

// NewDogAdapter Функция для создания экземпляра адаптера.
func NewDogAdapter(dog *Dog) Animal {
	return &DogAdapter{dog}
}

func main() {
	fmt.Println("\nВы останавливаетесь перед дверью и вставляете в ухо адаптер с двумя чипами")
	myDog := []Animal{NewDogAdapter(&Dog{})}

	fmt.Println("Открываете дверь и заходите домой")
	for _, member := range myDog {
		member.MakeSound()
	}
}
