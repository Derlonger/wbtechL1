package main

import (
	"fmt"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

func createHugeString(size int) string {
	str := make([]byte, size)
	return string(str)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	sliceBytes := []byte(v[:100])
	justString = string(sliceBytes)
}

func main() {
	someFunc()
	fmt.Println("Length of justString:", len(justString))
	fmt.Println(justString)
}
