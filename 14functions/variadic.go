package main

import "fmt"

func calculator(val ...int) int {
	sum := 0
	for _, v := range val {
		sum += v

	}
	return sum
}

func main() {
	a := 23
	b := 45
	fmt.Println(calculator(a, b, 45, 54, 66, 7))
	fmt.Println(calculator([]int{a, b, 45, 54, 66, 7}...)) // passing slice as an  variadic argument
}
