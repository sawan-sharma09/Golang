package main

import "fmt"

func main() {
	demo2()

}

func demo1() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Println(languages["RB"])

	languages["RB"] = "Ronaldo" //re-assigning different value to the key
	fmt.Println(languages["RB"])
	fmt.Println(languages)

	delete(languages, "PY") //deleting a key-value pair using delete func
	fmt.Println(languages)
}

func demo2() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"

	for key, value := range languages {
		// fmt.Printf("For key =>%v, value is:  %v\n", key, value)
		fmt.Println("value of", key, "is ", value)
	}
}
