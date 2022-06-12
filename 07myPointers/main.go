package main

import "fmt"

func main() {
	demo2()
}

func demo1() {
	var ptr *int
	fmt.Println(ptr)
}

func demo2() {
	myNumber := 23
	ptr := &myNumber
	fmt.Println("Value of the pointer in: ", ptr)  // here the pointer is pointing to the actual memory location.
	fmt.Println("Value of the pointer in: ", *ptr) // here the pointer is pointing to the value of the variable it is referencing to
}
