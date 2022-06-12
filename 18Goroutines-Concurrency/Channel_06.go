package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(channel chan int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 4; i++ {
		randNum := rand.Intn(10)
		channel <- randNum

	}
	close(channel)
}

func main() {
	channel := make(chan int, 4)
	go generator(channel)
	for i := 0; i < 4; i++ {
		fmt.Println(<-channel)
	}

}
