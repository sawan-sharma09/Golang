package main

import (
	"fmt"
)

func main() {
	list := []int{23, 4, 5, 678, 784, 534, 667}
	fmt.Println(linear(list, 784))
}

func linear(list []int, a int) int {
	k := 0 // to find the index of that number
	for i, val := range list {
		if a == val {
			k = i
		}
	}
	fmt.Println("number found at:")
	return k
}
