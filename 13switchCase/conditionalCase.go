package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch { //            Here after the Switch I tried writing today, but it took  it as a struct
	case today.Day() < 5:
		fmt.Println("Clean your house")

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
