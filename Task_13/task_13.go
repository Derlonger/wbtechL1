package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	sequence := []int{1, 2}

	sequence[1], sequence[0] = sequence[0], sequence[1]

	fmt.Println(sequence)
}
