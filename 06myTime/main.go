package main

import (
	"fmt"
	"time"
)

func main() {
	demo4()
}

func demo1() {
	today := time.Now()
	fmt.Println(today)
	year := today.Year()
	month := today.Month()
	day := today.Day()
	hour := today.Hour()
	minute := today.Minute()
	second := today.Second()
	fmt.Println("Present year is:", year)
	fmt.Println("Present month is:", month)
	fmt.Println("Present day is:", day)
	fmt.Println("Present hour is:", hour)
	fmt.Println("Present miniute is:", minute)
	fmt.Println("Present second is:", second)
}

func demo2() {
	today := time.Now()
	fmt.Println("Format date:", today.Format("01-02-2006"))
	fmt.Println("Format time:", today.Format("15:04:05"))
	fmt.Println("Format date and time:", today.Format(" 01-02-2006 15:04:05 "))
}

func demo3() {
	s := "10-03-2034"
	value, err := time.Parse("01-02-2006", s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New date: ", value.Format("2006-01-02"))
	}
}

func demo4() {
	date1 := time.Date(2022, time.Month(01), 03, 04, 05, 06, 07, time.UTC)
	fmt.Println(date1)
}
