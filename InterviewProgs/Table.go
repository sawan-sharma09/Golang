package main

import "fmt"

func main() {

	for a := 2; a <= 10; a++ {
		fmt.Printf("%v's table is:", a)
		for i := 1; i <= 10; i++ {
			fmt.Print(a*i, " ")
		}
		fmt.Println()
		if a == 4 {
			break
		}
	}
}
