package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			w.Write([]byte("GET method is received"))
			return
		}
		if r.Method == http.MethodPost {
			w.Write([]byte("POST method is received"))
			return
		}
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	})

	http.ListenAndServe(":8080", nil)
}