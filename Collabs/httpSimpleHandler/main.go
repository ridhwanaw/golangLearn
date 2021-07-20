package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":4040", nil)
}
