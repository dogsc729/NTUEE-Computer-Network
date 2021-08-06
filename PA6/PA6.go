package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(conn net.Conn) {
	//===========read the file size first==================
	f_out, err_out := os.Create("whatever.txt")
	check(err_out)
	defer f_out.Close()
	writer := bufio.NewWriter(f_out)
	scanner := bufio.NewScanner(conn)
	filecontent := ""
	message := ""
	//println("FUCKKK")
	i := 1
	index := 0
	old_file_size := ""
	for scanner.Scan() {
		if index < 1 {
			//println("FUCK")
			//println(scanner_1.Text())
			message = scanner.Text()
			fmt.Printf("Upload file size: %s\n", message)
			old_file_size = strings.TrimSuffix(message, "\n")
		} else {
			prepend := strconv.Itoa(i)
			filecontent = filecontent + scanner.Text()
			filecontentsize := strconv.Itoa(len(filecontent) + i)
			_, errw := writer.WriteString(prepend + " " + scanner.Text() + "\n")
			check(errw)
			i++
			writer.Flush()
			if filecontentsize == old_file_size {
				break
			}
		}
		index = index + 1
	}
	//==============================================================
	//===========message of the original file and the new file size===========
	file, err := os.Stat("whatever.txt")
	check(err)
	filesize := file.Size()
	filesizestring := strconv.FormatInt(filesize, 10)
	fmt.Printf("Output file size: %s\n", filesizestring)
	writer_filesize := bufio.NewWriter(conn)
	newline := old_file_size + " bytes received, " + filesizestring + " bytes file generated"
	_, errwf := writer_filesize.WriteString(newline)
	check(errwf)
	writer_filesize.Flush()
	conn.Close()
	time.Sleep(5 * time.Second)
	//========================================================================
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12018")
	defer ln.Close()
	i := 1
	for {
		conn, _ := ln.Accept()
		fmt.Printf("%d ", i)
		go handleConnection(conn)
		i++
	}
}
