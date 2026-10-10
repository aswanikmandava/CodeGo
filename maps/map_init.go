package main

import "fmt"

func main() {

	// nil map
	var colors map[string]string

	// reading is fine
	_ = colors["blue"]

	// writing to a nil map will not work. program crashes
	// colors["red"] = "#ff0000"
	fmt.Printf("%+v", colors)
	fmt.Println()

	// valid initialized map with no data
	var countryCodes = make(map[string]string)
	
	// writing works
	countryCodes["USA"] = "us"
	fmt.Printf("%+v", countryCodes)
	fmt.Println()

	// another way to declare and init map
	colors2 := map[string]string{
		"green": "#00ff00",
	}

	// add key later
	colors2["blue"] = "#0000ff"
	fmt.Printf("%+v", colors2)
}