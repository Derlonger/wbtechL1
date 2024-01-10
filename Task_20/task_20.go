package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

// Функция reverseWords принимает строку и возвращает строку, в которой порядок слов перевернут.
func reverseWords(input string) string {
	// Разделение строки на слова
	words := strings.Fields(input)
	length := len(words)

	// Переворот порядка слов с использованием индексов i и j
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединение слов в строку с пробелами между ними
	result := strings.Join(words, " ")

	return result
}

func main() {
	// Пример использования функции reverseWords
	input := "snow dog sun"
	reversed := reverseWords(input)

	// Вывод результатов
	fmt.Println("Исходная строка: ", input)
	fmt.Println("Перевернутая строка: ", reversed)
}
