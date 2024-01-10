package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

// Div выполняет деление двух big.Int чисел и возвращает результат.
func div(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Div(a, b)
}

// Add выполняет сложение двух big.Int чисел и возвращает результат.
func add(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Add(a, b)
}

// Mul выполняет умножение двух big.Int чисел и возвращает результат.
func mul(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Mul(a, b)
}

// Sub выполняет вычитание двух big.Int чисел и возвращает результат.
func sub(a, b *big.Int) *big.Int {
	z := new(big.Int)
	return z.Sub(a, b)
}

func main() {
	// Установка значения переменной 'a' больше чем 2^20.
	a := new(big.Int)
	a.SetString("5233252525125426576586798078967845643523435353251231243547658769567567844588425346787845734", 10)

	// Установка значения переменной 'b' больше чем 2^20.
	b := new(big.Int)
	b.SetString("32415123418128481248184182481248184182481248", 10)

	// Вывод результатов операций.
	fmt.Printf("%d + %d = \n%d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = \n%d\n", a, b, sub(a, b))
	fmt.Printf("%d / %d = \n%d\n", a, b, div(a, b))
	fmt.Printf("%d * %d = \n%d\n", a, b, mul(a, b))
}
