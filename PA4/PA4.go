package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":12018")
	fmt.Println("Launching server...")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()
	//===========read the file size first==================
	reader := bufio.NewReader(conn)
	message, errr := reader.ReadString('\n')
	check(errr)
	fmt.Printf("Upload file size: %s", message)
	old_file_size := strings.TrimSuffix(message, "\n")
	//=====================================================
	//===========reads from the socket one line at a time===========
	f_out, err_out := os.Create("whatever.txt")
	check(err_out)
	defer f_out.Close()
	writer := bufio.NewWriter(f_out)
	scanner := bufio.NewScanner(conn)
	filecontent := ""
	check(errr)
	i := 1
	for scanner.Scan() {
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
	//========================================================================

}
