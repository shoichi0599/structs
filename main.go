package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Define the person struct
	// person{"Alex", "Anderson"}, this syntax works too
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
