package main

import "fmt"

func newCounter() func() int { // here func()int is used as a return type.
	GFG := 0
	return func() int {
		GFG++
		return GFG
	}
}

func main() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

}
