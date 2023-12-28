package main

import (
	"fmt"
	"sort"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	// Заданная последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения групп температур
	temperatureGroups := make(map[int][]float64)

	// Сортируем последовательность температур
	sort.Float64s(temperatures)

	// Определяем шаг в 10 градусов
	step := 10

	// Группируем температуры по шагу
	for _, temperature := range temperatures {
		groupKey := int(temperature) / step * step
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temperature)
	}

	// Выводим результат
	for key, value := range temperatureGroups {
		fmt.Printf("%d: %v\n", key, value)
	}
}
