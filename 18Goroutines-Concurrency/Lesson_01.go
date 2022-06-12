package main

import ( //           Simple goroutine program
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Jonny", "Bobby", "Andy"}
	for i := range evilNinjas {
		go attack(evilNinjas[i])
		time.Sleep(time.Microsecond) // using sleep() to make the main goroutine sleep for sometime so that the
		//                         attack goroutine can finish its execution before the main goroutine ends
	}
}
func attack(target string) {
	fmt.Println("Throwing ninja stars at: ", target)
}
