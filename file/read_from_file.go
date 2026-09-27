package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	fh, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer fh.Close()
	// read all bytes from a file
	content, err := io.ReadAll(fh)
	// convert bytes into text using string()
	fmt.Println(string(content))
}