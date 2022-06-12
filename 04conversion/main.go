package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pizza factory")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating between 1 and 5")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks a lot for the rating")
	fmt.Printf("The type of the rating is %T\n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("Added 1 to the rating", numRating+1)
	}
}
