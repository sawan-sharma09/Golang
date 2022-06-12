package main

import "fmt" //      ******	SLICE CREATION USNIG A SLICE******

func main() { //				  0         1     2      3         4
	original_Slice := []string{"Welcome", "to", "my", "golang", "grind"}
	fmt.Println("The original is: ", original_Slice)

	slice1 := original_Slice[0:1]
	slice2 := original_Slice[2:4]
	slice3 := original_Slice[:]
	slice4 := original_Slice[:3]
	slice5 := original_Slice[2:2]

	fmt.Println("The slice1 is: ", slice1)
	fmt.Println("The slice2 is: ", slice2)
	fmt.Println("The slice3 is: ", slice3)
	fmt.Println("The slice4 is: ", slice4)
	fmt.Println("The slice5 is: ", slice5)
}
