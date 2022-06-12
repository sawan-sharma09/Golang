package main

import "fmt"

func swap(a, b int) {
	temp := a
	a = b
	b = temp

}
func main() {
	x, y := 20, 30
	fmt.Printf("Before swap: x=%d and y=%d\n", x, y)
	swap(x, y)
	fmt.Printf("After swap: x=%d and y=%d", x, y) // here the calue didn't change because we passed the copies of the variable not the actual value

}
