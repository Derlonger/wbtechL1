package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

// binarySearch выполняет бинарный поиск в отсортированном массиве.
// Возвращает индекс найденного элемента и true, если элемент найден,
// иначе возвращает -1 и false.
func binarySearch(arr []int, target int) (int, bool) {
	// Сортировка массива перед началом поиска
	sort.Ints(arr)

	low := 0
	high := len(arr) - 1

	// Пока нижняя граница не станет больше или равна верхней
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// Проверка, был ли найден элемент
	if low == len(arr) || arr[low] != target {
		return -1, false
	}
	return low, true
}

func main() {
	arr := []int{3, 6, 7, 8, 12, 17, 18, 95, 103, 111}
	target := 95

	// Выполнение бинарного поиска
	index, ok := binarySearch(arr, target)

	// Вывод результата
	if ok == false {
		fmt.Println("Элемент не найден.")
	} else {
		fmt.Println("Индекс найден успешно и равен:", index)
	}
}
