package main

import "fmt"


type User struct {
	FirstName string
	LastName string
	// Embedded
	Address		// variable of type User will be able to access Address fields directly
}

type User2 struct {
	FirstName string
	LastName string
	Addr	Address	// Named field. variable of type User2 must access fields inside Address struct explicity using dot
					// Example: <var_name>.Addr.City
}

type Address struct {
	HouseNumber int
	StreetName string
	City string
	Zip int
}
func main() {
	var u User = User{
		FirstName: "Aswani", 
		LastName: "Mandava", 
		Address: Address{HouseNumber: 123, StreetName: "Douglas St", City: "Philadelphia", Zip: 12345},
	}

	var u2 User2 = User2{
		FirstName: "Aswani2", 
		LastName: "Mandava2", 
		Addr: Address{HouseNumber: 123, StreetName: "Douglas St", City: "Philadelphia", Zip: 12345},
	}
	fmt.Println("Firstname: ", u.FirstName, "City: ", u.City)
	fmt.Println("Firstname: ", u2.FirstName, "City: ", u2.Addr.City)

}
