package main

import (
	"fmt"
	"sort"
)

func main() {
	slc := []int{23, 45, 77, 89, 00, 234, 679}
	slc2 := []string{"orange", "peach", "harbour", "apple", "pomegranate"}

	fmt.Println("Before sorting slc : ", slc)
	fmt.Println("Before sorting slc2 : ", slc2)

	sort.Ints(slc)
	fmt.Println("After sorting:", slc)

	sort.Strings(slc2)
	fmt.Println("After sorting:", slc2)

}
