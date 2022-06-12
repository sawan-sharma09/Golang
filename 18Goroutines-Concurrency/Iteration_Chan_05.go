package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	go throwingNinjaStar(channel)
	for message := range channel {
		fmt.Println(message)
	} //        even after the below function ends the for loop as per condition, this channel in the main  function
	//            will stil wait for receving the results from the below for loop,beacuse it will never know that the
	// below channel has been closed or the below function has ended its execution....for that reason we will use the
	// close() function for closing the channel

}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("Score is:%v,%T", score, score)
	}
	close(channel)
}
