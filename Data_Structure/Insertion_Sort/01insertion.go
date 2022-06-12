package main

import (
	"fmt"
)

func main() {
	list := []int{56, 465, 16, 256, 74, 1}
	fmt.Println("Before sort:\n", list)
	fmt.Println(insertionSort(list))
}
func insertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}

		}

	}
	fmt.Println("after sort:")
	return list
}
