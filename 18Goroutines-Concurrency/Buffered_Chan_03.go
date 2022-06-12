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
		fmt.Println("Score is:", <-channel)

	}
}

func main() {
	channel := make(chan int, 4)
	generator(channel)

}
