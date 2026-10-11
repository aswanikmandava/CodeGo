package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("HTTP ERR: ", err)
		os.Exit(1)
	}

	// make a byte slice of size 9999
	bs := make([]byte, 9999)

	fmt.Println("Response code: ", resp.StatusCode)
	fmt.Println("Response status: ", resp.Status)

	// read the response body into the byte slice
	resp.Body.Read(bs)

	// print the byte slice
	fmt.Println(string(bs))
	// fmt.Println(len(string(bs)))
}
