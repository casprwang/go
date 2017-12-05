package main

func main() {
	c := make(chan int)

	go func() {
		c <- 10
	}()

	val := <-c

	println(val)
}
