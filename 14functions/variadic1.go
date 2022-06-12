package main

import (
	"fmt"
	"sort"
)

func main() {
	slc := []int{66, 549, 548, 364, 18}

	fmt.Println(demo1(slc...))
	fmt.Println(demo2(slc))
}
func demo1(slc ...int) []int {
	for index, _ := range slc {
		fmt.Printf("the index is %d and the value is %d\n", index, slc[index])
	}
	return slc

}

func demo2(slc []int) []int {
	for _, val := range slc {
		fmt.Printf("values are %d\n", val)
	}
	fmt.Println("Before sorting: ", slc)
	sort.Ints(slc)
	return slc
}
