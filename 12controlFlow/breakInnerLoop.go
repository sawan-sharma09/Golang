package main

import "fmt"

func main() {
	for a := 0; a <= 3; a++ {
		for b := 0; b <= 3; b++ {
			if a == 2 && b == 2 {
				break

			}
			fmt.Println(a, " ", b)
		}
	}
}
