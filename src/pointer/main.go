package main

import "fmt"

func main() {
	i := 10

	p := &i

	fmt.Println(*p)
	fmt.Println(i)

	i = 1
	fmt.Println(*p)
	fmt.Println(i)

	*p = 3

	fmt.Println(*p)
	fmt.Println(i)
}
