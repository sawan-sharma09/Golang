package main

import "fmt"

func main() {
	demo2()
}
func demo1() {
	slc := []int{33, 44, 55, 66, 77, 78}
	slc2 := slc[2:5]
	fmt.Printf("The 1st slice is %v\n", slc)
	fmt.Printf("The 2nd slice is %d\n", slc2)

}
func demo2() {
	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
