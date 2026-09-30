package main

// mutex protect shared variables by coordinating access among multiple goroutines
// prevent race conditions
// ensure consistent results in concurrent programs


import (
	"fmt"
	"sync"
)


func main() {
	var counter int = 0
	var wg sync.WaitGroup
	var mt sync.Mutex

	for i:=0; i<5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mt.Lock()
			counter++
			mt.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Counter: ", counter)
}