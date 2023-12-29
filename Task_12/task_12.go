package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

func main() {
	// Изначальный список
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	// Создаем мапу для предоставления множества
	set := make(map[string]bool)

	// Добавляем каждую строку в множество
	for _, value := range strings {
		set[value] = true
	}

	// Вывод
	fmt.Println("Список строк: ", strings)
	fmt.Println("Множество строк: ", set)
}
