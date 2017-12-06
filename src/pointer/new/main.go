package main

import "fmt"

func main() {
	temp := foo{a: 1}

	var b *foo

	b = new(foo)
	b.a = 1

	fmt.Println(temp)
	fmt.Println(b)
}

type foo struct {
	a int
}
