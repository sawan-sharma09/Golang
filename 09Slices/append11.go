package main

import "fmt"

func main() {
	demo1()
}
func demo1() {
	slc := []string{"Sawan", "kumar", "sharma"}
	slc = append(slc, "Sunil", "Sharma")
	fmt.Println(slc)
}
func demo2() {
	slc := []int{22, 54, 65, 78}
	fmt.Println(append(slc, 55, 77))
}
