package main

import "fmt"

const GlobalName = "Lets Learn"

func main() {
	 var username string = "Sawan"
	 fmt.Println(username)
	 fmt.Printf("variable is of type : %T\n", username)

	 var isLogged bool
	 fmt.Println(isLogged)
	 fmt.Printf("variable is of type : %T", isLogged)

	var smallVal uint8 = 255 // If we assign a value more than 255 to unit8, then it will throw an error (it overflows unit8)
	fmt.Println(smallVal)    // So in order to check the limits of any type search golang.org...>"The Go Programming Language Specification"
	fmt.Printf("variable is of type : %T", smallVal)

	var smallFloat float32 = 250.97257949959596
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T\n", smallFloat)

	var largeFloat float64 = 250.97257949959596
	fmt.Println(largeFloat)
	fmt.Printf("variable is of type : %T\n", largeFloat)

	fmt.Println(GlobalName) // the constant has been declared with capital letter to use it as global
	fmt.Printf("variable is of type : %T", GlobalName)

}
