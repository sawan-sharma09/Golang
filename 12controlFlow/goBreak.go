package main

import "fmt"

func main() {
	var input int
	fmt.Println("Please enter the number")
	fmt.Scanln(&input)

	for input < 10 {
		fmt.Printf("value is %v\n", input)
		input++

		if input > 5 {
			break // break can only be used in for, switch and select statements.
		}
	}
}
