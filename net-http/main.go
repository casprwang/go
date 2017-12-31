package main

import (
	"io"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", index)

	log.Print("Listening on localhost:3000")
	http.ListenAndServe(":13109", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "lol")
}
