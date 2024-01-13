package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func MySleep(seconds int) {
	// Мы используем time.After который возвращает канал и отправляет в него значение после заданного времени
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало работы программы...")

	timeSleep := 1

	MySleep(timeSleep)

	fmt.Printf("Прошло %d секунд, продолжаем выполнение программы...", timeSleep)
}
