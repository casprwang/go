package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "I am Song"
	result := WordCount(str)

	fmt.Println("Number of cats:", result["I"])

}

func WordCount(str string) map[string]int {
	counts := map[string]int{}

	for _, word := range strings.Fields(str) {
		counts[word]++
	}

	return counts
}
