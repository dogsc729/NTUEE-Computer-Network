package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching Server...")
	httphandler := http.FileServer(http.Dir("./"))
	http.Handle("/", httphandler)
	file, err := os.Stat("/server-test.html")
	check(err)
	filesize := file.Size()
	filesizeString := strconv.FormatInt(filesize, 10)
	fmt.Println(filesizeString)
	log.Fatal(http.ListenAndServe(":12018", nil))
}
