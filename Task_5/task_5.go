package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

// countSecond определяет количество секунд, в течение которых программа будет выполняться.
var countSecond = 5

// wg - это объект WaitGroup, используемый для ожидания завершения горутин.
var wg = sync.WaitGroup{}

// Print получает значения из канала и выводит их на экран.
func Print(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
	// Уменьшаем счетчик горутин в WaitGroup после завершения работы функции Print.
	wg.Done()
}

// Send отправляет значения в канал.
func Send(ch chan<- int) {
	var i int
	for {
		ch <- i
		i++
	}
}

func main() {
	// Создаем канал для передачи целых чисел.
	ch := make(chan int)

	// Увеличиваем счетчик горутин в WaitGroup на 2 (для горутин Print и Send).
	wg.Add(2)
	// Запускаем горутину для функции Print, передавая канал в качестве аргумента.
	go Print(ch)
	// Запускаем горутину для функции Send, передавая канал в качестве аргумента.
	go Send(ch)

	// Ждем указанное количество секунд (countSecond).
	time.Sleep(time.Duration(countSecond) * time.Second)

	// Ожидаем завершения всех горутин в WaitGroup.
	wg.Wait()
	// Закрываем канал после завершения работы горутин.
	close(ch)
}

// Более простой вариант сложночитаемый
//func main() {
//
//	ch := make(chan int)
//
//	go func() {
//		for value := range ch {
//			fmt.Println(value)
//		}
//	}()
//
//	go func() {
//		var i int
//		for {
//			ch <- i
//			i++
//		}
//	}()
//
//	time.Sleep(time.Duration(countSecond) * time.Second)
//	close(ch)
//}
