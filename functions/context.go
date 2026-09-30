package main

import (
	"context"
	"fmt"
	"time"
)

// slowFn now accepts a context and respects its cancellation/timeout
func slowFn(ctx context.Context) (string, error) {
	fmt.Println("Long running task started...")

	// Simulate work using a select statement to listen for timeouts
	select {
	// simulating long running task longer than the time given
	case <-time.After(7 * time.Second):
		return "Success data", nil
	case <-ctx.Done():
		return "", ctx.Err() // Returns context.DeadlineExceeded if it times out
	}
}

func slowOperation() (string, error) {
	// Create a timeout context. 
	// Note: It's best practice to derive from an existing context if available, 
	// but context.Background() works fine for a top-level call.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return slowFn(ctx)
}

func main() {
	result, err := slowOperation()
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		return
	}

	fmt.Printf("Operation succeeded: %s\n", result)
}
