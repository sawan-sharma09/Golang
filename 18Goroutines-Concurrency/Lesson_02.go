package main

import ( //                    goroutine with channel
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	smokeSignal := make(chan bool) //   smokeSignal-------> channel
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)

	fmt.Println(<-smokeSignal) //                     ---->receiver channel

}
func attack(target string, attacked chan bool) {
	fmt.Println("Throwing Ninjastar at: ", target)
	attacked <- true //                               ---> sender channel

}
