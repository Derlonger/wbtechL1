package main

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const numWorkers = 5

func main() {
	// Создаем канал для передачи данных
	ch := make(chan string)

	// Создаем WaitGroup для ожидания завершения работы всех воркеров
	var wg sync.WaitGroup

	// Создаем N воркеров, с помощью горутины
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// Запускаем горутину для записи данных в канал
	go func() {
		defer close(ch)
		for {
			data := generateData()
			ch <- data
			time.Sleep(time.Second) // Пауза между записью данных
		}
	}()

	// Ожидаем сигнал завершения программы (Ctrl+C)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	// Завершаем работу программы
	fmt.Println("\nПолучен сигнал прерывания работы(Ctr+C).Завершение работы worker...")
	close(ch) // Закрываем канал для завершения работы воркеров
	wg.Wait() // Ожидаем завершения всех воркеров
	fmt.Println("Все worker были закрыты. Выход...")
}

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select { // Структура select необходима для выбора одного из нескольких условий, как структура if else
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: Канал закрыт. Выход...\n", id)
				return
			}
			fmt.Printf("Worker %d получил данные: %s\n", id, data)
		}
	}
}

func generateData() string {
	return time.Now().Format(time.Stamp)
}

// В данном примере использован механизм закрытия канала для информирования воркеров о необходимости завершения
