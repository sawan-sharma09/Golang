package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(20)
	fmt.Println("Before sort:\n", list)
	fmt.Println(insertionSort(list))

}
func insertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}

		}

	}
	fmt.Println("after sort:")
	return list
}
