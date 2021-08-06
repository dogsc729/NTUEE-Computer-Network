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
	conn, errc := net.Dial("tcp", "127.0.0.1:12018") //connect a client to the server, returns the socket handle
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	len, errw := writer.WriteString("hello world!\n")
	check(errw)
	fmt.Printf("send a string of %d bytes\n", len)
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("server replies: %s\n", scanner.Text())
	}
}
