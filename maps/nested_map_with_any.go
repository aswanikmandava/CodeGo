package main

import "fmt"

func main() {
	c := map[string]any{
		"user": map[string]any {
			"Name": "Ben",
			"Age": 32,
			"City": "London",
		},
		"group": "myGroup",
	}

	// access the nested map safely
	if um, ok:= c["user"].(map[string]any); ok {
		fmt.Println(um["City"])
	}
}