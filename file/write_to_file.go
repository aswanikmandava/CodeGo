package main

import (
	"os"
	"fmt"
)

func main() {
	// creates or replaces a file
	fh, err := os.Create("testfile.txt")
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return
	}
	defer fh.Close()
	fh.WriteString("Hello, world !\n")
	fh.WriteString("How are you?\n")
	fh.WriteString("Finished writing !\n")
	fmt.Println("Completed")
}
