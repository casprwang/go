package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addOne(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(addOne(12, 123))
}
