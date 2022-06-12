package main

import "fmt"

func main() {
	list := []int{2, 5, 7, 9, 12, 34, 56, 87} // In binary search we always need a sorted list
	binarySearch(list, 2)
}
func binarySearch(list []int, value int) {
	min := 0
	max := len(list) - 1
	mid := 0

	for min < max {
		mid = (min + max) / 2
		if value == list[mid] {
			break
		} else if value > list[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	if min >= len(list) {
		fmt.Println("Not found")
	} else {
		fmt.Println("Element found at:", mid)
	}
}
