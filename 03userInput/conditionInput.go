package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating")
	rating, _ := reader.ReadString('\n')
	fmt.Printf("%T,%v", rating, rating)
	newRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else if newRating <= 1 || newRating > 5 {
		fmt.Println("Please enter a valid rating")

	} else {
		fmt.Printf("%T,%v", newRating, newRating+1)
	}
}
