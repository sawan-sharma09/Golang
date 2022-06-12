package main

import (
	"fmt"
	"time"
)

func main() {
	switch today := time.Now(); { //  here the ; symbol is the most imp other we will never know from te error message that wht its saying, so REMEMBER IT.
	case today.Day() < 5:
		fmt.Println("Its your day today")
	case today.Day() <= 10:
		fmt.Println("Buy some wine")

	case today.Day() > 15:
		fmt.Println("Lets Enjoy")

	case today.Day() == 25:
		fmt.Println("Buy some food")

	default:
		fmt.Println("What the hell was that number !")
	}
}
