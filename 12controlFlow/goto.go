package main

import "fmt"

func main() {
	a := 0
loop:
	for a < 10 {
		if a == 5 {
			a += 1
			goto loop //   this Program took me more than 20 minutes and it was the toughest till now

		} else {
			fmt.Println(a)
			a++
		}
	}

}
