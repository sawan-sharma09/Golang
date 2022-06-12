package main

import "fmt"

func main() {
	slc1 := []int{22, 24, 66}
	slc2 := []string{"Sawan", "Kumar", "Sharma"}

	fmt.Println(slc1 == nil)
	fmt.Println(slc2 == nil)

	// slc3 := []int{33, 66, 77}
	// fmt.Println(slc3 == slc1)  //  (slice can only be compared to nil), it will throw error uf we try to compare 2 slices

	//If we want to compare 2 slices, then use Range For loop to match each element or we can use  DEEPEQUAL function

}
