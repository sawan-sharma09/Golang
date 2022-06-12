package main

import "fmt"

type Duck interface {
	Waddle()
	Quack()
}

type Person struct {
	name string
	age  int
}
type IndianRunner struct {
	name string
}

func (p Person) Quack() {
	fmt.Printf("%v qaucking", p)
}
func (p Person) Waddle() {
	fmt.Printf("%v waddling", p)
}

func (i IndianRunner) Quack() {
	fmt.Printf("%v qaucking", i)
}
func (i IndianRunner) Waddle() {
	fmt.Printf("%v waddling", i)
}

func main() {
	fmt.Println("Type Assertion")

	Sheru := Person{"Sheru Baba", 5}
	lucky := IndianRunner{"Lucky Baba"}

	myPrinter(Sheru)
	myPrinter(lucky)
}

func myPrinter(m Duck) { //  --> here m is the static type
	fmt.Printf("MyPrinter:  -%T -> %v\n", m, m) //   --> here m is the dynamic type

	if val, ok := m.(IndianRunner); ok {
		fmt.Println("I have an IndianRunner as the type of value in 'm' ", val)
	} else if val1, isOk := m.(Person); isOk {
		fmt.Println("I have a Person as the type of value in 'm' ", val1)
	} else {
		fmt.Println("Unsupported type")
	}

}
