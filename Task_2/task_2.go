package main

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

var numbers = [5]int{2, 4, 6, 8, 10}

var wg sync.WaitGroup

func main() {
	// Создание канала для обработки чисел
	ch := make(chan int)

	fmt.Printf("Начальный список чисел: %v\n", numbers)

	// Увеличиваем счетчик горутин в WaitGroup
	wg.Add(1)

	// Запускаем новую горутину(main является основной горутиной)
	go squares(ch)

	// Запускаем итерацию по массиву numbers
	for _, number := range numbers {
		ch <- number // Каждое число отправляется в канал
	}

	// Закрываем канал после отправки всех чисел
	close(ch)

	// Ждем завершения работы горутины
	wg.Wait()
}

// Функция для получения чисел из канала возведения их в квадрат и вывод в stdout
func squares(ch <-chan int) { // ch <-chan сделано  ограничения 'только для чтения'
	// Уменьшаем счетчик горутин в WaitGroup при завершении
	defer wg.Done()

	for val := range ch {
		sq := val * val
		fmt.Printf("Квадрат числа %v равен %v\n", val, sq)
	}
}
