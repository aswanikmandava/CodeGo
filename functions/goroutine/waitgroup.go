package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("-------- Main thread started ---------")
	var wg sync.WaitGroup
	// add the number of tasks to wait group. it increments the internal waitgroup counter by this number  
	wg.Add(1)
	// main thread doesn't wait for this goroutine to finish
	// to fix it, using WaitGroup
	go func() {
		// decrements the WaitGroup task counter by one
		// equivalent to wg.Add(-1)
		defer wg.Done()
		fmt.Println("----- Go routime -----")
	}()
	// block until all goroutines finish
	wg.Wait()
	fmt.Println("-------- Main thread finished ---------")
}
