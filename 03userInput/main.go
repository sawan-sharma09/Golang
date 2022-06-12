package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of Pizza:")

	//Comma Ok sysntax, or, Comma Err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating", input)
	fmt.Printf("Type of this rating is %T", input)
}
