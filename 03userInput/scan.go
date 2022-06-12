package main

import "fmt"

func main() {
	fmt.Println("Please enter the rating")
	var input int
	fmt.Scanln(&input)
	if input < 1 || input > 5 {
		fmt.Println("Please enter a valid rating")
	} else {
		fmt.Printf("The type and value of input is %T,%v", input, input)
	}
}
