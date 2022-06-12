package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		// b = a + b //  incorrect result  will come if we write this line separately.
		return b - a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f()) // it means that the latest value of a and b are getting retained each time
	}
}
