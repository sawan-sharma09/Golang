package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	channel := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja, channel)
	fmt.Println(<-channel)
}
func attack(target string, attacked chan bool) {
	fmt.Println("Throwing ninjaStar at: ", target)

	attacked <- true
}
