package main

import "fmt"

func main() {
	demo3()
}
func demo1() { //      updating a value in maps
	emp := make(map[string]int)
	emp["Mark"] = 10
	emp["Sandy"] = 20

	emp["Mark"] = 50

	fmt.Println(emp)
}

func demo2() { //     iterating over a map

	emp := make(map[int]string)

	emp[01] = "Zebra"
	emp[02] = "Sheru"
	emp[03] = "Lion"
	emp[04] = "Ant"

	for key, value := range emp {
		fmt.Println("key is", key, "value is", value)
	}

}

func demo3() { //      checking if a value exists on the Map or not
	fruits := make(map[string]int)

	fruits["apple"] = 12
	fruits["grapes"] = 25
	fruits["banana"] = 44

	val, ok := fruits["apple"]
	fmt.Println(val, ok)

	val2, ok := fruits["orange"]

	fmt.Println(val2, ok)
}
