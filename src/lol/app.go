package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	var m map[string]interface{}

	if _, err := toml.DecodeFile("config.toml", &m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
