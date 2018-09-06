package main

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
)

type Node struct {
	Id string `bson:"_id"`
}

func main() {
	node := Node{Id: "lol"}
	b, e := bson.Marshal(node)

	if e != nil {
		panic(e)
	}

	fmt.Println(string(b))
}
