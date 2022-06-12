package main

import "fmt"

func main() {
	toChange := "hello"
	willChange := 5
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)

	changeValue(&toChange)
	fmt.Println(toChange)

	//experimentation with int type

	changeValue3(willChange)
	fmt.Println(willChange)

	changeValue4(&willChange)
	fmt.Println(willChange)

}

func changeValue(str *string) {
	*str = "Changed"
}

func changeValue2(str string) {
	str = "changed!"
}

func changeValue3(str int) {
	str = 14
}
func changeValue4(str *int) {
	*str = 16
}
