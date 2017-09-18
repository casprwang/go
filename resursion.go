package main

import "fmt"

func addBy(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	addBy5 := addBy(5)
	res := addBy5(3)
	fmt.Println("addby5", res)
}
