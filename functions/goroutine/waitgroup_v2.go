package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("-------- Main thread started ---------")
	var wg sync.WaitGroup
	var tasks = []string{"T1", "T2", "T3"}
	for _, tsk := range tasks {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			fmt.Printf("Task: %s\n", t)
		}(tsk)
	}
	// block until all goroutines finish
	wg.Wait()
	fmt.Println("-------- Main thread finished ---------")

}
