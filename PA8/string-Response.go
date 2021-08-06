package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
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
	conn, _ := ln.Accept()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	check(err)
	fmt.Printf("method: %s\n", req.Method)

	fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(conn, "Date: ...\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "File Not Found\r\n")
	fmt.Fprintf(conn, "\r\n")
}
