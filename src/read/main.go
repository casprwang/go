package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	srcFile, err := os.Open("~/.zsh_history")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	count := map[string]int{}

	scanner := bufio.NewScanner(srcFile)

	counter := 0
	for scanner.Scan() {
		line := string(scanner.Text())

		counter++
		count[line]++
		// fmt.Println(len(count))
		// fmt.Println(">>>", string(line))
	}

	type kv struct {
		Key string
		Val int
	}

	ss := []kv{}

	for k, v := range count {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i int, j int) bool {
		return ss[i].Val < ss[j].Val
	})

	for _, kv := range ss {
		fmt.Printf("%s %d %v%% \n", kv.Key, kv.Val, GetPercentage(kv.Val, counter))
	}

	fmt.Printf("Total executed commands: %d", counter)
}

func GetPercentage(a int, total int) float32 {
	return (float32(a) / float32(total)) * 100
}
