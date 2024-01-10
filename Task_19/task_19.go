package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func reverseString(input string) string {
	runes := []rune(input)
	fmt.Println(runes)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "snow dog sun"
	reversed := reverseString(input)

	fmt.Println("Исходная строка: ", input)
	fmt.Println("Перевернутая строка: ", reversed)
}
