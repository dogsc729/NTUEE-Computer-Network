package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
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
			//file, _ := os.Stat(requestfile[1:])
			//filesize := file.Size()
			//filesizeString := strconv.FormatInt(filesize, 10)
			// Read entire file content, giving us little control but
			// making it very simple. No need to close the file.
			content, err := ioutil.ReadFile(requestfile[1:])
			if err != nil {
				log.Fatal(err)
			}
			text := string(content)
			//fmt.Printf("File size = %s\n", filesizeString)
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Date: ...\r\n")
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, text+"\r\n")
			fmt.Fprintf(conn, "\r\n")
		} else {
			//fmt.Printf("File not found\n")
			fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")
			fmt.Fprintf(conn, "Date: ...\r\n")
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, "File not found\r\n")
			fmt.Fprintf(conn, "\r\n")
		}
		conn.Close()
	}
}
