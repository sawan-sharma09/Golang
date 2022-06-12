package main

import "fmt"

func main() {
	arr := [7]int{44, 66, 83, 65, 78, 44, 38}
	slc := arr[1:5]

	fmt.Println("The array is:", arr) //If we make any changes in the slice then it will also affect the original array-
	fmt.Println("the slice is:", slc) //-it is referencing to and the changes will also be visible in the original array

	slc[1] = 300
	slc[3] = 200

	fmt.Println("The new slice is:", slc)
	fmt.Println("The new array is:", arr)

}
