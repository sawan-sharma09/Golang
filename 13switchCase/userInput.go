package main

import (
	"fmt"
)

func main() {

	var input int
	fmt.Println("Please enter the number")
	fmt.Scanln(&input) //        I tried using the fmt.scanf("%d",input) but it didn't give the desired result.

	switch input {
	case 1:
		fmt.Println("Take one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four") // **Fallthrough means, if the input is 4 then it will print the 4th statement as well as the 5th statement
		fallthrough
	case 5:
		fmt.Println("five")
	}

}
