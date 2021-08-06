package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12018") //socket handle, works just like a file handler or the standard I/O
	conn, _ := ln.Accept()               //socket, dedicated to data transmission between the client and server
	defer ln.Close()
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	message := ""
	if scanner.Scan() {
		message = scanner.Text()
		fmt.Println(message)
	}

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}
