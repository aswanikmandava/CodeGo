package main

import (
	"fmt"
	"errors"
)

func getHours(secs int) (float32, error) {
	if secs <= 0 {
		return 0, errors.New("invalid input: seconds must be a positive integer")
	}
	return float32(secs) / 3600, nil
}

func main() {
	hours, err := getHours(7200)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Hours: %.2f\n", hours)
}