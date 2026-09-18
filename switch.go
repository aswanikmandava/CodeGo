package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var os string = runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Operating system is Linux")
	case "windows":
		fmt.Println("Operating system is Windows")
	default:
		fmt.Println("Operating system is not recognized")
	}

	// switch with no condition, acts like if-else
	var hour int = time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
