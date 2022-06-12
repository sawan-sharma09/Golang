package main

import "fmt" // iterating over a slice using range

func main() {
	slc := []string{"Sawan", "Kumar", "Sharma"}
	for index, _ := range slc {
		fmt.Println("Index is:", index, ":and value is:", slc[index])
	}
}
