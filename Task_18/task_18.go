package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкриминироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Counter Структура-счетчик
type Counter struct {
	count int        // Поле count хранит текущее значение счетчика
	mutex sync.Mutex // Mutex используется для обеспечения безопасного доступа к общим данным
}

// Increment Метод для увеличения счетчика
func (c *Counter) Increment() {
	c.mutex.Lock()         // блокируем доступ к общим данным
	defer c.mutex.Unlock() // гарантируем, что мьютекс будет разблокирован после завершения функции
	c.count++              // увеличиваем значение счетчика
}

// GetCount Метод для получения текущего значения счетчика
func (c *Counter) GetCount() int {
	c.mutex.Lock()         // блокируем доступ к общим данным
	defer c.mutex.Unlock() // гарантируем, что мьютекс будет разблокирован после завершения функции
	return c.count         // возвращаем текущее значение счетчика
}

func main() {
	// Создаем экземпляр структуры-счетчика
	myCounter := Counter{}

	numGoroutines := 10

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			myCounter.Increment() // вызываем метод Increment в каждой горутине для инкрементации счетчика
		}()
	}

	wg.Wait() // ждем завершения всех горутин

	// Выводим итоговое значение счетчика
	finalValue := myCounter.GetCount()
	fmt.Println("Итоговое значение счетчика: ", finalValue)
}
