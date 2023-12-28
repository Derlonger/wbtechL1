package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// Исходные множества
var set1 = []int{4, 2, 6, 5, 3, 1, 3, 4, 5}
var set2 = []int{1, 8, 61, 5, 33, 7, 2, 42, 9}

// Результирующий срез
var result []int

func main() {
	// Перебираем элементы первого множества
	for _, i := range set1 {
		// Перебираем элементы второго множества
		for _, j := range set2 {
			// Если элементы равны, добавляем его в результирующий срез
			if i == j {
				result = append(result, i)
			}
		}
	}

	// Убираем дубликаты из результирующего среза
	result = removeDuplicates(result)

	// Выводим результат
	fmt.Println("Пересечение множеств:", result)
}

// Функция для удаления дубликатов из среза
func removeDuplicates(input []int) []int {
	encountered := map[int]bool{}
	var result []int

	// Перебираем элементы входного среза
	for _, value := range input {
		// Если элемент еще не встречался, добавляем его в результирующий срез
		if !encountered[value] {
			encountered[value] = true
			result = append(result, value)
		}
	}

	return result
}
