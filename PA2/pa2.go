package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("pls enter input file name\n")
	input_file_name := ""
	fmt.Scanf("%s", &input_file_name)
	fmt.Printf("pls enter output file name\n")
	output_file_name := ""
	fmt.Scanf("%s", &output_file_name)
	f_in, err_in := os.Open(input_file_name)
	check(err_in)
	f_out, err_out := os.Create(output_file_name)
	check(err_out)
	defer f_out.Close()
	writer := bufio.NewWriter(f_out)
	scanner := bufio.NewScanner(f_in)
	for scanner.Scan() {
		_, errw := writer.WriteString(scanner.Text() + "\n")
		check(errw)
		writer.Flush()
	}

	f_in.Close()
}
