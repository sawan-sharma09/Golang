package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := generateSlice(10)
	fmt.Println("Before sorting:- ", slice)
	fmt.Println("After sorting")
	fmt.Println(bubbleSort(slice))
}
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999) // here we can say that the random number is generating between -999 to 999
	} // rand.Intn(999) - rand.Intn(999) , this statement generates only one value at a time, suppose if we want 4 random numbers
	//then we need to run that statement 4 times.
	return slice
}
func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
