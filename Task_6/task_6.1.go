package main

import (
	"context"
	"fmt"
	"time"
)

/*
Использование контекста для отмены выполнение горутины.
*/

func worker2(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Выполненяется горутина.")
		case <-ctx.Done():
			fmt.Println("Горутина завершена.")
			return
		}
	}
}

func main() {
	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем горутину
	go worker2(ctx)

	// Ждем некоторое время
	time.Sleep(3 * time.Second)

	// Отменяем выполнение горутины
	cancel()

	// Ждем завршение горутины
	time.Sleep(1 * time.Second)
}
