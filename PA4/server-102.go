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
	fmt.Println("launching server...")
	ln, _ := net.Listen("tcp", ":12018")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()

	reader := bufio.NewReader(conn) //works just like bufio.NewScanner and bufio.NewWriter
	message, errr := reader.ReadString('\n')
	check(errr)
	fmt.Printf("%s", message)

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}
