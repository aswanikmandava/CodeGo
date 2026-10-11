package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// read cmd line args
	cmd_args := os.Args
	if len(cmd_args) <= 1 {
		fmt.Println("Missing cmdline argument: filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	fmt.Println(filename)
	// get a file handler
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("Err: ", err)
		os.Exit(1)
	}
	// print the file contents
	io.Copy(os.Stdout, fh)
}