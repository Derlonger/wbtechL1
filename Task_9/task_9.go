package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем каналы для обмена данными между горутинами
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Используем WaitGroup для ожидания завершения вех горутин
	var wg sync.WaitGroup

	// Запускаем первую горутину, которая пишет числа в первый канал
	wg.Add(1)
	go writeNumbers(inputChannel, &wg)

	// Запускаем вторую горутину, которая умножает числа на 2 и пишет во второй канал
	wg.Add(1)
	go multiplyByTwo(inputChannel, outputChannel, &wg)

	// Запускаем третью горутину, которая выводит результат в stdout
	wg.Add(1)
	go printResult(outputChannel, &wg)

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// Горутина для записи чисел в первый канал
func writeNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	numbers := []int{3, 4, 5, 6, 7, 8, 2}

	for _, num := range numbers {
		ch <- num
	}

	// Закрываем канал, чтобы сообщить второй горутине, что данныех больше не будет
	close(ch)
}

// Горутина для умножения чисел на 2 и записи во второй канал
func multiplyByTwo(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range input {
		result := num * 2
		output <- result
	}

	// Закрываем второй канал, чтобы сообщить третьей горутине, что данных больше не будет
	close(output)
}

func printResult(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for result := range ch {
		fmt.Println(result)
	}
}
