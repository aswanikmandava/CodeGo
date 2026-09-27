package main

import (
	"fmt"
	// "errors"
)

// func loadConfig() error {
// 	return errors.New("Missing property")
// }

func loadConfig() error {
	return nil
}

func main() {
	if appError := loadConfig(); appError != nil {
		fmt.Println("Error: ", appError)
		return
	}
	fmt.Println("Config loaded successfully")
}
