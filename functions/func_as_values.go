package main

import "fmt"

func greet(message string) {
	fmt.Println(message)
}

func compute(a, b float64 ) float64 {
	return a + b
}

func compute_sub(a, b float64) float64 {
	return a - b
}

func main() {
	// declaring a variable that accepts function as a value
	// define the function type using its signature (type of args it can accept and type of values it can return)
	var greeter func(msg string)
	greeter = greet
	greeter("Hello")


	var total func(v1 float64, v2 float64) float64
	// total can be assigned a function that must match its signature
	// total can take compute or compute_sub
	total = compute
	var result float64 = float64(total(5.4, 3.6))
	fmt.Println(result)
	fmt.Printf("type: %T", result)

	total = compute_sub
	var result_sub float64 = float64(total(10.0, 4.6))
	fmt.Println(result_sub)

}
