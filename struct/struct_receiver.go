package main

import "fmt"

type Addr struct {
	houseNo string
	streetName string
	city string
	zip int
}

type WebUser struct {
	username string
	email string
	address Addr	// embedded struct
}

// There are two ways to pass a receiver into a function
// 1) Value Receiver - Go passes a copy of the struct to the function
//			When you only need to read data from the struct, or for small structs that don't change. 
// 			Any changes you make to fields inside the function will not affect the original struct.
// 2) Pointer Receiver - Go passes a pointer(memory address) to the original struct
//			When you need to modify/mutate the fields of the original struct, 
// 			or when the struct is very large (to avoid the performance cost of copying it in memory).

// function that receives a struct of type WebUser
// Value Receiver
func (w WebUser) printInfo() {
	fmt.Printf("%+v", w)
}

// receiver function that updates WebUser address
// note that function args are passed by value in Go
func (w WebUser) updateAddress(a Addr) {
	w.address = a
}

// Pointer Receiver function
func (wPtr *WebUser) updateOrigAddress(aPtr *Addr) {
	(*wPtr).address = *aPtr
}
// address-of operator (&)
//		When you place it before a variable, it tells Go to find that variable's location in computer memory and 
// 		return its memory address (a pointer) instead of the actual data value stored inside it.

// dereference operator (*) roles
//		1. The Pointer Type Literal: When placed before a type name (like *int or *string), 
// 			it means "this variable holds a memory address".
//		2. The Dereference Operator: When placed before an existing pointer variable name (like *myPointer), 
// 			it means "go to this address and fetch/change the value inside".
func main() {
	usr := WebUser{
		username: "akumar",
		email: "akumar@example.com",
		address: Addr{
			houseNo: "100A",
			streetName: "Walmart Dr",
			city: "BedRock",
			zip: 12344,
		},
	}

	newAddr := Addr{
		houseNo: "200B",
		streetName: "Lombard Dr",
		city: "Boston",
		zip: 43222,
	}
	usr.updateAddress(newAddr)	// this will update new address on the copy of the "usr"
	// but, not the original. hence, the new address update is not shown when printed

	usr.printInfo()

	fmt.Println("\n--- Correct way to update address using Pointer Receiver function ----")
	usrPtr := &usr
	newAddrPtr := &newAddr

	(*usrPtr).updateOrigAddress(newAddrPtr)
	fmt.Println("New address updated. Printing the user details ...")
	usr.printInfo()

}

