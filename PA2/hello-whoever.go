package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf(("who's there?\n"))
	text := ""
	fmt.Scanf("%s", &text)

	fmt.Printf("hello, %s\n", text)
	fmt.Println("hello,", text)
	fmt.Fprintf(os.Stdout, "hello, %s\n", text)
}
