package main

import "fmt"

func main() {
	arr := [5]int{33, 55, 77, 70, 88}
	fmt.Println(arr)
	fmt.Println(arr[1:4])

	arr1 := arr[1:4]
	fmt.Println(arr1)
	arr1[1] = 100
	fmt.Println(arr)  //Once we do any changes to any part of the array then it reflects in the original array as well
	fmt.Println(arr1) // and the changes are permanent.

	fmt.Println(len(arr), cap(arr))
	fmt.Println(len(arr1), cap(arr1))
	fmt.Printf("%T,%v", arr1, arr1) //even if we think it as an array, but internally it is a slice and the output also shows it as a slice
}
