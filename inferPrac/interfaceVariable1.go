package main

import "fmt"

type (
	ID     uint64
	SSN    string //----> Social Security Number
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Empty interface{}
)

func main() {
	//    review - declaring interface variables and assigning values
	// ------

	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = 11.04 //     ----------> here although the value and type of e is nil, but after assigning, the value field will record the new value and the type field will record the new type.
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = ID(142342524)
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = SSN("819-25-3443")
	fmt.Printf("e's value: %v, type: %T\n", e, e)
	e = Person{name: "Sawan", age: 24, ssn: SSN("819-25-3443")}
	fmt.Printf("e's value: %+v, type: %T\n", e, e)

	//e = &Person{name: "Sawan", age: 24, ssn: SSN("819-25-3443")}
	// fmt.Printf("e's value : %+v, type : %T\n", e, e)
	//          --> Result--> e's value : &{name:Sawan age:24 ssn:819-25-3443}, type : *main.Person
}
