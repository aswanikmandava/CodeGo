package main

import (
	"log"
)

func main() {
	logger := logging.NewConsoleLogger()
	logger.Log("Hello, Logger!")
}
