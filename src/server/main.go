package main

import (
	"fmt"
	"net/http"
)

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome</h1>")
}

func main() {
	http.HandleFunc("/", Default)
	http.ListenAndServe("localhost:3000", nil)

}
