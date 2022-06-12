package main //       ***********SLICE CREATION USING AN ARRAY**************

import "fmt"

func main() { //         0       1		2			3		   4	  5
	arr := [6]string{"Welcome", "to", "google", "language", "let's", "go"}
	slice1 := arr[:5]
	slice2 := arr[3:4]
	slice3 := arr[1:5]
	slice4 := arr[0:3]
	slice5 := arr[:]
	slice6 := arr[5:]
	slice7 := arr[5:5] //special attention to be given

	fmt.Println("The 1st slice is", slice1)
	fmt.Println("The 2nd slice is", slice2)
	fmt.Println("The 3rd slice is", slice3)
	fmt.Println("The 4th slice is", slice4)
	fmt.Println("The 5th slice is", slice5)
	fmt.Println("The 6th slice is", slice6)
	fmt.Println("The 7th slice is", slice7)

}
