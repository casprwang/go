package main

import "fmt"

func logSpread(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	arr := []int{1, 2, 3}
	logSpread(arr...)
}
