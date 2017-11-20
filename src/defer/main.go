package main

import "fmt"

func explode() {
	panic("wrong")
}

func resolve() {
	if err := recover(); err != nil {
		fmt.Println("FIX")
		fmt.Println("ERR", err)
	}

}

func main() {

	defer resolve()

	explode()
}
