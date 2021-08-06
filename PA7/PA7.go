package main

import (
	"bufio"
	"fmt"
	"net"
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
	fmt.Println("launching server...")
	ln, _ := net.Listen("tcp", ":12018")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		reader := bufio.NewReader(conn)
		req, err := http.ReadRequest(reader)
		requestfile := req.URL.String()
		check(err)
		if _, err := os.Stat(requestfile[1:]); err == nil {
			file, _ := os.Stat(requestfile[1:])
			filesize := file.Size()
			filesizeString := strconv.FormatInt(filesize, 10)
			fmt.Printf("File size = %s\n", filesizeString)
		} else {
			fmt.Printf("File not found\n")
		}
		conn.Close()
	}
}
