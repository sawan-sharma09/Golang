package main

import "fmt"

func main() {
	emp := make(map[string]int)
	emp["Sawan"] = 10
	emp["Sheru"] = 20
	emp["Khusi"] = 30

	for key := range emp {
		delete(emp, key) // --> We can use this statement to delete a single key from tha map also  Ex-delete(emp, "Sawan")

	}
	fmt.Println(emp)
}
