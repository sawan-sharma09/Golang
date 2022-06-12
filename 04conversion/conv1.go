package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pizza house")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating ")
	input, _ := reader.ReadString('\n')

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		if newRating < 1 || newRating > 5 {
			fmt.Println("Please enter a valid rating")
		} else {
			fmt.Println("The newRating is", newRating)
		}
	}

}
