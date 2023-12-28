package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

var numbers = make(map[int]int, 100) // Создание мапы на 100 слотов
var mu sync.Mutex                    // Мютекс для синхронизации доступа к мап
var wg = sync.WaitGroup{}            // WaitGroup для создание очереди горутин

func main() {
	wg.Add(2) // Добавляем 2 горутины в WaitGroup

	// Write
	go func() {
		defer wg.Done() // Уменьшаем счетчик горутин при завершении
		for i := 0; i < 100; i++ {
			mu.Lock() // Захватывем мютекс перед доступом к мапе
			numbers[i] = i
			mu.Unlock() // Освобождаем мютекс после доступа к мапе
		}
	}()

	// Ожидание для того что бы данные успели записаться в мапу
	//time.Sleep(1 * time.Millisecond)

	// Read
	go func() {
		defer wg.Done() // Уменьшаем счетчик горутин при завершении
		for i := 0; i < 100; i++ {
			mu.Lock() // Захватываем мютекс перед доступом к мапе
			fmt.Println(numbers[i])
			mu.Unlock() // Освобождаем мютекс после доступа к мап
		}
	}()

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println(numbers)
}
