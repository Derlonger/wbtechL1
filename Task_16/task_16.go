package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

// Функция partition разбивает массив на две части относительно опорного элемента
// и возвращает отсортированный массив и индекс опорного элемента
func partition(arr []int, start, end int) ([]int, int) {
	pivot := arr[end]
	i := start

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return arr, i
}

// Функция quickSort выполняет сортировку массива методом быстрой сортировки
func quickSort(arr []int, start, end int) []int {
	if start < end {
		var p int
		arr, p = partition(arr, start, end)
		arr = quickSort(arr, start, p-1)
		arr = quickSort(arr, p+1, end)
	}
	return arr
}

func main() {
	// Пример использования быстрой сортировки
	nums := []int{7, 5, 4, 3, 2, 34, 5, 6, 7, 1}
	fmt.Println("Неотсортированный список:", nums)
	fmt.Println("Отсортированный список:", quickSort(nums, 0, len(nums)-1))
}
