package main

import (
	"bytes"
	"fmt"
)

func main() {
	demo2()
}

func demo1() {
	slc1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'G', 'e', 'e', 'k', '#', '#'}
	slc2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
	slc3 := []byte{'%', 'G', 'e', 'e', 'k', 's', '%'}

	res1 := bytes.Split(slc1, []byte("eek")) // here []byte..> this symbol is used because "eek" and string and we are converting it to byte otherwise it will throw error
	res2 := bytes.Split(slc2, []byte(""))
	res3 := bytes.Split(slc3, []byte("%"))

	fmt.Println("After splitting : ")
	fmt.Printf("slice1 is: %s\n", res1)
	fmt.Printf("slice2 is: %s\n", res2)
	fmt.Printf("slice3 is: %s\n", res3)
}

func demo2() {
	res1 := bytes.Split([]byte("***Geeks for geeks***"), []byte("*"))
	res2 := bytes.Split([]byte("Learning X how X to X split X a X slice"), []byte("X"))
	res3 := bytes.Split([]byte("Geeks for geeks, Geeks"), []byte(""))
	res4 := bytes.Split([]byte(""), []byte(""))

	fmt.Printf("slice1: %s\n", res1)
	fmt.Printf("slice2: %s\n", res2)
	fmt.Printf("slice3: %s\n", res3)
	fmt.Printf("slice4: %s\n", res4)
}
