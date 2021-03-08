package main

import (
	"fmt"
	"time"
)

func main() {
	exerciseOne()
	fmt.Println("==========")
	exerciseTwo()
	fmt.Println("==========")
	exerciseThree()
	fmt.Println("==========")
	exerciseFour()
	fmt.Println("==========")
	exerciseFive()
	fmt.Println("==========")
	exerciseSix()
}

func exerciseOne() {
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}

func exerciseTwo() {
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func exerciseThree() {
	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
		count++
		if count == 3 {
			break
		}
	}
}

func exerciseFour() {
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}
}

func exerciseFive() {
	for year := 1988; year <= time.Now().Year(); {
		fmt.Println(year)
		year++
	}
}

func exerciseSix() {
	switch age := -9; true {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
