package main //         *********SLICE CREATION USING MAKE FUNCTION*********

import "fmt"

func main() {
	demo3()
}
func demo1() {
	mySlice := make([]int, 4)
	fmt.Println("slice is:", mySlice)
	fmt.Println("size of slice: ", len(mySlice))
	fmt.Println("capacity of slice: ", cap(mySlice))
}
func demo2() {
	mySlice := make([]int, 5, 7)      //here 5 is the size or the length of the slice, and 7 is  the capacity
	fmt.Println("slice is:", mySlice) //the capacity will always be equal to  or greater than the length
	fmt.Println("size of slice: ", len(mySlice))
	fmt.Println("capacity of slice: ", cap(mySlice))
}

func demo3() {
	mySlice := make([]int, 4, 6)
	fmt.Printf("Slice is : %v\n slice length is : %d\n slice capacity is : %d", mySlice, len(mySlice), cap(mySlice))
}
