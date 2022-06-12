package main

import "fmt"

type (
	ID     uint64
	SSN    string
	Person struct {
		name string
		age  uint8
		ssn  SSN
	}
	Empty interface{}
)

func main() {

	var e Empty

	printInfo(e)
	printInfo(11.04)
	printInfo("String")      //-----> Here the string type was not defined in the previous program but we are able
	printInfo(ID(142342524)) //        to assign it to the empty  interface variable because it can take any value and type.
	printInfo(SSN("819-25-3443"))
	printInfo(Person{name: "Sawan", age: 24, ssn: SSN("819-25-3443")})
	printInfo(&Person{name: "Sawan", age: 24, ssn: SSN("819-25-3443")})

}

func printInfo(e Empty) {
	fmt.Printf("e's value : %+v, type : %T\n", e, e)
}
