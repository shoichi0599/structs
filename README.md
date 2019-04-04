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

### 1. Defining and Declaring a Struct
Whenever we make a struct we have to first define all of the different properties that a struct might have.

```
 Tell Go what fileds the person struct has.
 
          [ person struct ]
 +----------------------------------+
 |                                  |
 |  +-----------+     +----------+  |
 |  | firstName | --> | <string> |  |
 |  +-----------+     +----------+  |
 |                                  |
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
 |                                  |
 |  +-----------+     +----------+  |
 |  | lastName  | --> |  "Sato"  |  |
 |  +-----------+     +----------+  |
 +----------------------------------+
```

```go
package main

import "fmt"

// Define a person struct
// Tell Go what fileds the person struct has.
// Let's say a person struct should have a first name and a last name
type person struct {
    firstName string
    lastName string
}

func main() {
    // Declare the person struct
    // person{"Taro", "Sato"}, this syntax works too
    taro := person{firstName:"Taro", lastName:"Sato"}
    fmt.Println(taro)
}
```

### 2. Updating Struct Values
```go
package main

import "fmt"

type person struct {
    firstName string
    lastName  string
}

func main() {
    // Declare a person (note that usually declare with ":=")
    var taro person
    
    // Update the struct
    taro.firstName = "Taro"
    taro.lastName = "Sato"

    fmt.Println(taro)
    // "%+v" will print out all the different field names and their values from 'taro'
    fmt.Printf("%+v", taro)
}
```

### 3. Embedding Structs
```
          [ person struct ]
 +---------------------------------------+
 |                                       |
 |  +-----------+     +---------------+  |
 |  | firstName | --> |   <string>    |  |
 |  +-----------+     +---------------+  |
 |  +-----------+     +---------------+  |
 |  | lastName  | --> |   <string>    |  |
 |  +-----------+     +---------------+  |
 |  +-----------+     +---------------+  |
 |  |  contact  | --> | <contactInfo> |  |
 |  +-----------+     +---------------+  |
 +---------------------------------------+
 
      [ contactInfo struct ]
 +------------------------------+
 |                              |
 |  +-------+     +----------+  |
 |  | email | --> | <string> |  |
 |  +-------+     +----------+  |
 |                              |
 |  +-------+     +----------+  |
 |  |  zip  | --> |  <int>   |  |
 |  +-------+     +----------+  |
 +------------------------------+
```
```go
package main

import "fmt"

type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    contact   contactInfo
}
// Or defining like below works too
type person2 struct {
    firstName string
    lastName  string
    contactInfo
}

func main() {
    jim := person{
        firstName: "Jim",
        lastName:  "Party",
        contact: contactInfo{
            email:   "jim@gmail.com",
            zipCode: 94000,
        },
    }
    jim2 := person2{
        firstName: "Jim",
        lastName:  "Party",
        contactInfo: contactInfo{
            email:   "jim@gmail.com",
            zipCode: 94000,
        },
    }
    fmt.Printf("%+v", jim)
    fmt.Printf("%+v", jim2)
}
```