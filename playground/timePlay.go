package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	demo4()
	demo5()
}

func demo1() {
	fmt.Println("Welcome to my Resturant")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate the pizza application")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the rating is ", input)
	}
	fmt.Printf("The type of the rating is %T\n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The rating has been revised to", numRating+1)
	}

}
func demo2() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(11)) //Here the value 11 in the Intn is the upper limit upto which  we want to generate the random no.
}
func demo3() {
	today := time.Now()
	fmt.Println("Today: ", today)
	year := today.Year()
	month := today.Month()
	day := today.Day()
	hour := today.Hour()
	minute := today.Minute()
	seconds := today.Second()
	fmt.Println("Present year is: ", year)
	fmt.Println("Present month is: ", month)
	fmt.Println("Present day is: ", day)
	fmt.Println("Present hour is: ", hour)
	fmt.Println("Present minute is: ", minute)
	fmt.Println("Present seconds is:", seconds)
}
func demo4() {
	today := time.Now().Format("01-02-2006")
	fmt.Println(today)
}
func demo5() {
	date := time.Date(2022, time.Month(5), 02, 05, 10, 0, 0, time.UTC)
	fmt.Println("Formated Date details: ", date)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	Demo3()
// }

// func Demo1() {
// 	fmt.Println("Hello everyone, welcome to my Pizza house")
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please enter the rating of the Pizza ")

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanks for the rating: ", input)
// 	fmt.Printf("The type of the rating is %T\n", input)

// 	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Added 1 to the rating of pizza: ", newRating+1)
// 	}
// }

// func Demo2() {
// 	rand.Seed(time.Now().Unix())
// 	fmt.Println(rand.Intn(5))
// }

// func Demo3() {
// 	today := time.Now()
// 	fmt.Println("Today time: ", today)

// 	Year := today.Year()
// 	// month := time.Month

// 	fmt.Println("Year is :", Year)

// 	fmt.Println("Formatted time is: ", time.Now().Format("01-02-2006"))
// }
