package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func removeElement(slice []int, index int) []int {
	// Проверка на валидность индекса
	if index < 0 || index >= len(slice) {
		fmt.Println("Неверный индекс для удаления элемента.")
		return slice
	}

	// Удаление элемента из слайса
	// Так как функции delete в го мы нет, просто будем сращивать слайсы до и после index
	slice = append(slice[:index], slice[index+1:]...)

	return slice
}

func main() {
	// Пример использования
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Исходный слайс:", mySlice)

	indexToRemove := 3
	mySlice = removeElement(mySlice, indexToRemove)

	fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, mySlice)
}
