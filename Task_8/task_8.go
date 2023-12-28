package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(value int64, position uint, bitValue uint) int64 {
	// Проверка введенных данных
	if position > 63 {
		fmt.Println("Ошибка: Позиция бита должна быть от 0 до 63.")
		return 0
	}
	if bitValue > 1 {
		fmt.Println("Ошибка: Значение бита должно быть 0 или 1.")
		return 0
	}

	// Очистить бит на указанной позиции
	mask := ^(int64(1) << position)
	value &= mask

	// Установить новое значение бита
	value |= int64(bitValue) << position

	return value
}

func main() {
	var number int64
	var position uint
	var bitValue uint

	// Ввод данных
	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)

	fmt.Print("Введите позицию бита (от 0 до 63): ")
	fmt.Scan(&position)

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&bitValue)

	// Установка бита
	result := setBit(number, position, bitValue)

	// Вывод результата
	fmt.Printf("Результат после установки бита: %d\n", result)
}
