package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}

func main() {
	fmt.Println("launching server...")

	hh := http.HandlerFunc(helloHandler)
	http.Handle("/hello", hh)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":12018", nil)
}
