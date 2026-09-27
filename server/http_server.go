package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("program started")

	// registers a route "/" to a request handler that handles the incoming HTTP request from the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// formatted text is written to HTTP Response body
		fmt.Fprintln(w, "Welcome page")
	})

	// registers a route "/home" to another request handler
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// data is written to HTTP Response body as a slice of bytes
		w.Write([]byte("Home page"))
	})

	// start a http server listening on local port 8080
	http.ListenAndServe(":8080", nil)
}
