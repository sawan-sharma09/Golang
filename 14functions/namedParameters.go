package main

import "fmt"

func calculate(x, y int) (div, mul, remainder int) {
	div = x / y
	mul = x * y
	remainder = x % y
	return

}

func main() {
	fmt.Println("Lets calculate")

	a, b, c := calculate(100, 20) /*Special attention to be given on the way m and d has been used..
	Although we passed 2 actual parameters but in return we have 3 value in it and we need to store it in
	3 different variables*/
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// fmt.Println(calculate(74, 21))
	/***If we use this line then we can print the return directly
	and we dont need to store it in multiple variables but the Problem is that it will print the return
	in single line***/

}
