package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}() // to check the time difference with and without subtracting the i from len(list)-1

	list := []int{51, 84, 25, 79, 1, 24, 97, 36, 48, 72, 18}
	fmt.Println("Before Bubble sorting:\n", list)
	fmt.Println("After Bubble sorting:")
	fmt.Println(bubbleSort(list))
}
func bubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1-i; j++ { // as with every loop, the last index gets the largest number,
			// so we dont need to compare that index anymore, for that reason we have subtracted the i from the len(cards)-1
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]

			}
		}
	}
	return list
}
