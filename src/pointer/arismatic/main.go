package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0] // an adress
	c := &a[1] // an adress

	fmt.Println("%v %p %p", a, b, c)
	fmt.Println("%d %d", *b, *c)
}
