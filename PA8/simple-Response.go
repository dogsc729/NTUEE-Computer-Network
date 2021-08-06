package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("launching server...")

	http.ListenAndServe(":12018",
		http.FileServer(http.Dir(".")))
}
