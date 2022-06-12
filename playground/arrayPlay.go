package main

import "fmt"

func main() {
	arr := [2][3]int{{1, 2, 3}, {3, 3, 4}}
	fmt.Println(arr[1][2]) //here [1] is for accessing the second array, and [2] is used for
	//accessing the 3rd element of that array
	fmt.Println(arr[1:1])
}
