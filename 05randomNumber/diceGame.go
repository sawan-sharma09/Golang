package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dice := rand.Intn(8)
	fmt.Println("The random number is ", dice)

	switch dice {
	case 1:
		fmt.Println("move 1 step ")
	case 2:
		fmt.Println("move 2 step ")
	case 3:
		fmt.Println("move 3 step ")
	case 4:
		fmt.Println("move 4 step ")
	case 5:
		fmt.Println("move 5 step ")
	case 6:
		fmt.Println("Move 6 steps and play again ")
	default:
		fmt.Println("That was out of box!!")
	}

}
