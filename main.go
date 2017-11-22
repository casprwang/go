package main

import (
	"fmt"

	"github.com/wangsongiam/go/src"
)

func main() {
	test := []int{1, 2, 3}
	res := play.Shuffle(test)

	fmt.Println(res)
}
