package main

import (
	"fmt"
	"sync"
)

func myFunc2(wg *sync.WaitGroup) {
	fmt.Println("Hello from myFunc2")
	wg.Done()
}
func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Hello from myFunc")
	wg.Done() // each time it decreases 1 from the incrementer----STEP-2

}
func main() {
	var wg sync.WaitGroup
	wg.Add(2) //   adds the goroutines in the incrementer-----STEP-1

	go myFunc(&wg)
	go myFunc2(&wg)
	wg.Wait() // it waits until the incrementer becomes 0-----STEP-3
	fmt.Println("Bye")
}
