# Overview
Learning Go by 
following the course [Go: The complete Developer's Guide](https://www.udemy.com/go-the-complete-developers-guide/).<br>
(Section 4)

## Go Documentations
- Go Documentation: https://golang.org
- Effective Go: https://golang.org/doc/effective_go.html

# What I Learned
## Structs
A struct is short for structure.<br>
It is a data structure in Go and you can think of it 
as being like a collection of different properties
that are somehow related together or have some type of common purpose.

### 1. Defining a struct
Whenever we make a struct we have to first define all of the different properties that a struct might have.

##### Example
```
 Tell Go what fileds the person struct has.
 
          [ person struct ]
 +----------------------------------+
 |                                  |
 |  +-----------+     +----------+  |
 |  | firstName | --> | <string> |  |
 |  +-----------+     +----------+  |
 |
 |  +-----------+     +----------+  |
 |  | lastName  | --> | <string> |  |
 |  +-----------+     +----------+  |
 +----------------------------------+
 
 Create a new value of type person.
 +----------------------------------+
 |                                  |
 |  +-----------+     +----------+  |
 |  | firstName | --> |  "Taro"  |  |
 |  +-----------+     +----------+  |
 |
 |  +-----------+     +----------+  |
 |  | lastName  | --> |  "Sato"  |  |
 |  +-----------+     +----------+  |
 +----------------------------------+
```

##### Code
```go
package main

import "fmt"

// Tell Go what fileds the person struct has.
// Let's say a person struct should have a first name and a last name
type person struct {
	firstName string
	lastName string
}

func main() {
	// Define the person struct
	// person{"Alex", "Anderson"}, this syntax works too
	alex := person{firstName:"Alex", lastName:"Anderson"}
	fmt.Println(alex)
}
```
test