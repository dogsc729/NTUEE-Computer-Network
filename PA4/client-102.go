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
	conn, errc := net.Dial("tcp", ":12018")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	len, errw := writer.WriteString("hello world!\n")
	check(errw)
	fmt.Printf("send a string of %d bytes\n", len)
	writer.Flush()

	reader := bufio.NewReader(conn)
	message, errr := reader.ReadString('\n')
	check(errr)
	fmt.Printf("server replies: %s", message)
}
