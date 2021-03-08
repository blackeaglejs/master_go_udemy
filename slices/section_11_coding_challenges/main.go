package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
	exerciseFour()
	exerciseFive()
	exerciseSix()
	exerciseSeven()
	exerciseEight()
}

func exerciseOne() {
	exampleSlice := []string{"one", "two", "three"}
	fmt.Println("Exercise 1")
	for i, v := range exampleSlice {
		fmt.Printf("Index is %d, Value is %s\n", i, v)
	}
	fmt.Println("========================")
}

func exerciseTwo() {
	fmt.Println("Exercise 2")
	mySlice := []float64{1.2, 5.6}

	// This line fails because the you need to append instead.
	// mySlice[2] = 6
	mySlice = append(mySlice, 6)

	a := 10
	// This line fails because you're trying to add an int to a float64 slice. Fix is to convert type.
	mySlice[0] = float64(a)

	// This line fails because it is out of range for the slice. Append instead.
	// mySlice[3] = 10.10
	mySlice = append(mySlice, 10.10)

	// This line fails because we're appending an int to a float.
	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)
	fmt.Println("========================")
}

func exerciseThree() {
	fmt.Println("Exercise 3")
	nums := []float64{1, 2, 3}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)
	n := []float64{11, 12}
	nums = append(nums, n...)
	fmt.Println(nums)
	fmt.Println("========================")
}

func exerciseFour() {
	fmt.Println("Exercise 4")
	arguments := os.Args[1:]
	if len(arguments) > 10 || len(arguments) < 2 {
		panic("Number of Arguments must be between 2 and 10 inclusive.")
	}
	var argumentsSum float64
	var argumentsProduct float64 = 1
	for _, v := range arguments {
		argumentAsFloat, _ := strconv.ParseFloat(v, 64)
		argumentsSum += argumentAsFloat
		argumentsProduct *= argumentAsFloat
	}
	fmt.Printf("Sum: %v, Product: %v\n", argumentsSum, argumentsProduct)
	fmt.Println("========================")
}

func exerciseFive() {
	fmt.Println("Exercise 5")
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	smallerSlice := nums[1 : len(nums)-2]
	sum := 0
	for _, v := range smallerSlice {
		sum += v
	}
	fmt.Println(smallerSlice)
	fmt.Printf("Sum: %v\n", sum)
	fmt.Println("========================")
}

func exerciseSix() {
	fmt.Println("Exercise 6")
	friends := []string{"Marry", "John", "Paul", "Diana"}
	var friends2 = make([]string, len(friends))
	friendsCopied := copy(friends2, friends)
	fmt.Println(friendsCopied)
	friends2[0] = "Katherine"
	fmt.Printf("Original Slice: %+v\n", friends)
	fmt.Printf("New Slice: %+v\n", friends2)
	fmt.Println("========================")
}

func exerciseSeven() {
	fmt.Println("Exercise 7")
	friends := []string{"Marry", "John", "Paul", "Diana"}
	friends2 := append(friends, "Zoheb")
	friends2[4] = "Christopher"
	fmt.Printf("Original Slice: %+v\n", friends)
	fmt.Printf("New Slice: %+v\n", friends2)
	fmt.Println("========================")
}

func exerciseEight() {
	fmt.Println("Exercise 8")
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := years[0:3]
	newYears = append(newYears, years[len(years)-3:]...)
	fmt.Println(newYears)
	fmt.Println("========================")
}
