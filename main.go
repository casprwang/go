package main

import (
	"github.com/wangsongiam/go/src/foobar"
	"fmt"
)

func main() {
	count := foobar.Count(10)
	fmt.Println(count)
}
