package main

import (
	"log"
	"errors"
)

func loadUser() error {
	return errors.New("user not found")
}

func main() {
	log.Println("This is a startup log message.")

	err := loadUser()
	if err != nil {
		log.Fatalf("Error loading user: %v", err)
		return
	}
	log.Println("Successfully loaded user.")
}
