package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	money int = 2500
)

func John(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	money += value
	fmt.Printf("John is giving me: %v ,%v\n ", value, money)
	mutex.Unlock()
	wg.Done()

}
func Samuel(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	money -= value

	fmt.Printf("Samuel is taking from me: %v ,%v\n", value, money)
	mutex.Unlock()
	wg.Done()

}

func main() {
	fmt.Println("Simple example of mutex")
	var wg sync.WaitGroup
	wg.Add(2)
	go John(500, &wg)
	go Samuel(800, &wg)
	wg.Wait()
	fmt.Printf("Total money now I have is %v", money)

}
