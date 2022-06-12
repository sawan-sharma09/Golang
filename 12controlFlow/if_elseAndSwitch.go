package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	demo()

}
func demo() {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(8)
	fmt.Println(randNum)
	if randNum < 1 || randNum > 5 {
		fmt.Println("Please enter a valid input")

	} else {
		fmt.Printf("NewRating is %v\n", randNum)
		switch randNum {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
			fallthrough
		default:
			fmt.Println("Please enter anything under 5")
		}
	}
}
