package main

import "fmt"

func main() {
	var facNum int
	mul := 1
	fmt.Print("Please enter any number: ")
	fmt.Scanln(&facNum)
	fmt.Println("the number is: ", facNum)
	for i := facNum; i >= 1; i-- {
		func(i int) {
			mul *= i
		}(i)
	}
	fmt.Printf("The factorial of the number is: %v", mul)
}
