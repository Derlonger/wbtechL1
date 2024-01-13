package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.

Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func isUnique(str string) bool {
	// Инициализируем карту для отслеживания встреченных символов
	charMap := make(map[rune]bool)

	// Приводим строку к нижнему регистру
	strLower := strings.ToLower(str)

	// Проходим по каждому символу строки
	for _, char := range strLower {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("Строка '%s' уникальна: %t\n", str1, isUnique(str1))
	fmt.Printf("Строка '%s' уникальна: %t\n", str2, isUnique(str2))
	fmt.Printf("Строка '%s' уникальна: %t\n", str3, isUnique(str3))
}
