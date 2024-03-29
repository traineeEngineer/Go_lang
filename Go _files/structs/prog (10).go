// Structs are collection of properties that are related together
package main

import (
	"fmt"
)

// Embedding Strcuts
type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //here embedding strcuts use
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	//updating struct
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Amber"

	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Joya",
		contact: contactInfo{
			email: "jim.j@gmail.com",
			zip:   223445,
		},
	}
	jimPointer := &jim
	jimPointer.updateFirstName("Jimmy")
	jim.print()

}

// struct with reciever
//struct with pointer
func (pointer *person) updateFirstName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
