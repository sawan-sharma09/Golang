package main

import (
	"bytes"
	"fmt"
)

func main() {
	demo2()
}

func demo1() {
	slice1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
	slice2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
	slice3 := []byte{'%', 'G', 'e', 'e', 'k', 's', '%'}

	fmt.Println("Original slice: ")
	fmt.Printf("Slice1 :  %s\n", slice1)
	fmt.Printf("Slice2 :  %s\n", slice2)
	fmt.Printf("Slice3 :  %s\n", slice3)

	result1 := bytes.Trim(slice1, "!#")
	result2 := bytes.Trim(slice2, "*^")
	result3 := bytes.Trim(slice3, "@")

	fmt.Println("Resulted new slice:  ")
	fmt.Printf("\nslice1:  %s", result1)
	fmt.Printf("\nslice2:  %s", result2)
	fmt.Printf("\nslice3:  %s", result3)
}

// I found this case as special because whenever we declare any array or slice, we us {} symbol for holding the elements
//but in the below example, the () sign is used instead of {}, so  this is very imp.
func demo2() {
	result1 := bytes.Trim([]byte("***Welcome to Geeks for Geeks***"), "*")
	result2 := bytes.Trim([]byte("!!!Learning how to trim a slice of bytes@@@"), "!@")
	result3 := bytes.Trim([]byte("^^Geeks$$"), "&")
	fmt.Printf("Slice1: %s\n", result1)
	fmt.Printf("Slice2: %s\n", result2)
	fmt.Printf("Slice3: %s\n", result3)
}
