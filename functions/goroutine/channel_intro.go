package main

import "fmt"

func main() {
	// create a messages channel that can carry string values
	// this channel is unbuffered which is used to communicate synchronous string data between goroutines
	var messages chan string = make(chan string)

	// start a goroutine that sends a message
	go func() {
		messages <- "Hello, go from a channel"
	}()

	// receiver (variable and string data coming from a channel together)
	// receive a string value from a channel and store it in msg
	// this operation is blocking until the data is sent by the goroutine
	msg := <-messages
	fmt.Println(msg)
}
