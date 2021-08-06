package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	conn, errc := net.Dial("tcp", "140.112.42.221:12000")
	check(errc)
	defer conn.Close()

	fmt.Printf("pls enter the name of the file you want to upload\n")
	filename := ""
	fmt.Scanf("%s", &filename)
	//fIn, _ := os.Open(filename)

	//send the file size
	file, err := os.Stat(filename)
	check(err)
	filesize := file.Size()
	filesizeString := strconv.FormatInt(filesize, 10)
	writer1 := bufio.NewWriter(conn)
	filesizeString2 := filesizeString + "\n" + "hello\r\n"
	_, errw := writer1.WriteString(filesizeString2)
	check(errw)
	writer1.Flush()
	fmt.Println(filesizeString2)
	//send the file content
	//writer2 := bufio.NewWriter(conn)

	/*scanner1 := bufio.NewScanner(f_in)
	for scanner1.Scan() {
		_, errw := writer2.WriteString(scanner1.Text() + "\n")
		check(errw)
	}
	_, errw2 := writer2.WriteString("hello\r\n")
	check(errw2)
	writer2.Flush()*/

	//get the reply from the server
	scanner2 := bufio.NewScanner(conn)
	if scanner2.Scan() {
		fmt.Printf("server replies: %s\n", scanner2.Text())
	}
	//fmt.Println(filesizeString)
}
