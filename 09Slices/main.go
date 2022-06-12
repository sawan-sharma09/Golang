package main

import "fmt"

func main() {
	var slc = []string{"Apple", "banana", "peach", "mango", "grapes"}
	fmt.Println("The size of the slc is", len(slc))

	slc2 := slc[1:4]
	fmt.Printf("the new size of the slc is : %d\n", len(slc2))
	fmt.Println("the capacity of the slc is :", cap(slc2))

}
