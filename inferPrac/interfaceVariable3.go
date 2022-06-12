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
	//   interface variable as function variadic parameter
	var e Empty

	printInfo(e, 11.04, "String",
		ID(142342524),
		SSN("819-25-3443"),
		Person{name: "Sawan", age: 24, ssn: SSN("819-25-3443")},
		&Person{name: "Sawan", age: 24, ssn: SSN("819-25-3443")})

}

func printInfo(e ...Empty) {
	for i := range e {
		fmt.Printf("parameter[%v]'s value: %v, type: %T\n", i, e[i], e[i])
	}

}
