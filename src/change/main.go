package main

// func main() {
// 	var b int
// 	b = 33
// 	fmt.Println(b)
// 	change(&b) // take value as an argument
// 	fmt.Println(b)
// }

// func change(a *int) {
// 	*a = 44
// }

func main() {
	a := 33
	b := &a

	*b = 3

	println(*b)
	println(a)

	*b = 4

	println(*b)
	println(a)

	a = 5

	println(*b)
	println(a)
}
