package main

import "fmt"
import "math"
import "reflect"

func main() {
	res := math.Sqrt(15)
        fmt.Println(res == math.Floor(res))
        fmt.Println(math.Floor(res))
        fmt.Println(reflect.TypeOf(math.Floor(res)))
}
