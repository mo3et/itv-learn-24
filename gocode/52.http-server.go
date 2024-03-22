package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello http server!\n")
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
