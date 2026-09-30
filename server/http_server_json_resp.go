package main

import (
	"net/http"
	"encoding/json"
)

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			// headers must be set before writing the content
			// set HTTP response header
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-MyAppHeader", "localapp")
			// set HTTP response code
			w.WriteHeader(http.StatusCreated)

			// init User
			user_1 := User{Username: "amandava", Email: "amandava@example.com"}

			// encode User struct into JSON and write to HTTP response writer
			err := json.NewEncoder(w).Encode(user_1)
			if err != nil {
				http.Error(w, "JSON encoding failed", http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	http.ListenAndServe(":8080", nil)
}
