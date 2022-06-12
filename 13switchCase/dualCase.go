package main

import "fmt"

func main() {

	var a int
	fmt.Println("Please enter any number")
	fmt.Scanln(&a)

	switch a {
	case 1, -1:
		fmt.Println("one_one")

	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	}
}
