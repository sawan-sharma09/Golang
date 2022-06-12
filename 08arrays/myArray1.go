package main

import "fmt"

func main() {
	demo2()
}

func demo1() {
	arr := [3]int{4, 5, 6}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func demo2() {
	arr := [4]int{1, 2, 57, 78}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Total sum is: ", sum)

}
