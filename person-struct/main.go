package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// create a new person obj

	//  1, order of fileds in the strut
	// alex := person{"Alex", "Anderson"}

	// 2, Key value setting. Not relying on order of struct
	alice := person{firstName: "Alice", lastName: "Anderson"}

	// 3, create type with Zero Value for the properties contained eg
	// string = ""
	// int = 0
	// float = 0
	// bool = false
	var alan person
	alan.firstName = "Alan"
	alan.lastName = "Anderson"

	var albert person

	// fmt.Println(alex)
	fmt.Println(alice)
	fmt.Println(alan)
	fmt.Println(albert)
	// or print values and keys to show the Zero Values
	fmt.Printf("%+v", albert)

	fmt.Println("\n****")

	jim := person{
		firstName: "Jim",
		lastName:  "Jones",
		contact: contactInfo{
			email:    "jim.jones@jj.com",
			postCode: "ABC-123",
		},
	}
	fmt.Printf("%+v", jim)

}
