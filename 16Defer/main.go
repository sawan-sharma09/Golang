package main

import "fmt"

func mul(a, b int) int {
	res := a * b
	fmt.Println(res)
	return res
}
func show() {
	fmt.Println("Hello Welcome to Sheru's house")
}
func main() {
	defer mul(20, 34)
	show()
	fmt.Println(mul(10, 35))
}
