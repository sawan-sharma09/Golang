package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	input := rand.Intn(8)
	fmt.Println("The number is", input)

	switch input {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		fallthrough
	case 4:
		fmt.Println("four")
		break
		fmt.Println("Break the rules") // here this statement will not be printer because of break statment
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("Please try again")
	}

}
