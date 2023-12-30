package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func getType(x interface{}) {
	switch x := x.(type) {
	case bool:
		fmt.Println("Boolean", x)
	case string:
		fmt.Println("String", x)
	case int:
		fmt.Println("Integer", x)
	case chan int:
		fmt.Println("Channel", x)
	default:
		fmt.Println("Unknown type", x)
	}
}

func main() {
	stringVar := "string"
	intVar := 1
	var chanVar chan int
	getType(stringVar)
	getType(intVar)
	getType(true)
	getType(chanVar)
}
