// Structs are collection of properties that are related together
package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
