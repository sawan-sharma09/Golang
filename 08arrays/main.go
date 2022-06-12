package main

import "fmt"

func main() {
	demo1()
}

func demo1() {
	var arr1 [5]string
	fmt.Println(arr1)
}

func demo2() {
	var arr1 [5]int
	fmt.Println(arr1)
}

func demo3() {
	var arr [5]int
	fmt.Println(arr)
	arr[0] = 100
	arr[4] = 500
	fmt.Println(arr)
}
