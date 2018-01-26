package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
	// or if you don't want to have the name implied from the type
	// contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//  *type here tells us the receiver is a pointer to a person
func (p *person) updateName(newFirstName string) {
	//  this won't work cuz without a receiver of (p *person) a new variable of p is copied from the caller. This means the update will not happen
	// p.firstName = newFirstName

	// to fix this we use pointer in the reciver type and the value of it in teh update call
	// *p is the _value_ at the pointer's address
	(*p).firstName = newFirstName

	// KEY TO WORKING WITH POINTERS
	// turn an _address_ into a _value_ with *address
	// turn a _value_ into an _address_ with &value ie a pointer
}

func main() {
	// create a new person obj

	//  1, order of fileds in the strut (must do all fields)
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
		contactInfo: contactInfo{
			email:    "jim.jones@jj.com",
			postCode: "ABC-123",
		},
	}
	// &variable is a reference to the memory address of a varibale
	jimPointer := &jim
	// *pointer is the _value_ at the address being pointed to
	jimPointer.print()

	jim.print()
	jim.updateName("James")
	jim.print()

	fmt.Println("\n****")

}
