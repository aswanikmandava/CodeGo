package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// myStdoutWrite is a custom type that implements the io.Writer interface
type myStdoutWrite struct {}

func main() {
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("HTTP ERR: ", err)
		os.Exit(1)
	}
	// Create an instance of myStdoutWrite and use it to copy the response body to standard output
	myWriter := myStdoutWrite{}
	// Use io.Copy to copy the response body to the custom writer
	io.Copy(myWriter, res.Body)

}

// this receiver function implements the Write method of the io.Writer interface for the myStdoutWrite type. 
// It takes a byte slice as input, converts it to a string, and prints it to the standard output. 
func (myStdoutWrite) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Written ", len(b), " bytes")
	return len(b), nil
}