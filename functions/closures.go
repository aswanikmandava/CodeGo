package main

import "fmt"

// In Go, a closure is a function value that references variables from outside its own body. 
// The function is "bound" to these variables, meaning it can access, read, and modify them even after 
// the outer scope in which they were created has finished executing

func main() {
	var cnt int = 0

	res := func() int {
		// this anonymous function is accessing the variable outside the function and modifying it
		// hence this variable cnt together with this function called as a closure
		for i:=0; i<5; i++ {
			cnt++
		}
		return cnt
	}()

	fmt.Println(res)

}
