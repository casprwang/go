package main

import (
	"github.com/wangsongiam/go/src"
	"fmt"
	"reflect"
)

func main() {
	count := bar.Bar(10)
	// list := [3]int{1, 2, 3}

	// for _, val := range list {
	fmt.Println(reflect.TypeOf(int(count)))
	// }
	// fmt.Println(count)
}
