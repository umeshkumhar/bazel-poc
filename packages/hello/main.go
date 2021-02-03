package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	count := Counting(4, 2)
	fmt.Println(count)
}

func Counting(a int, b int) int {
	return a + b
}
