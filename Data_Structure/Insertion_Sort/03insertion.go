package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := generateSlice(20)
	fmt.Println("Before sort:\n", list)
	fmt.Println(insertionSort(list))
}
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
func insertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	fmt.Println("After sort:")
	return list
}
