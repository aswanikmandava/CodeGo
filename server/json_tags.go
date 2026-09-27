package main

import (
	"fmt"
	"encoding/json"
)


type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	// fieldname in lowercase cannot be encoded
	// fieldname must start with uppercase letter
	// address string `json:"address`
	Address string `json:"address"`
}

func main() {
	var p Person = Person{Name: "Aswani", Age: 42, Address: "Wall St"}

	// encode the person struct into JSON using Marshal()
	data, _ := json.Marshal(p)

	fmt.Println(string(data))
}