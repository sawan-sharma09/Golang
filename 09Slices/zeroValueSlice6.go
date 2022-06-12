package main

import "fmt"

func main() {
	slc := []string{} // we are allowed to create a slice which has nil value in it
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
}
