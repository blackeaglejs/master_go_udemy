package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	// Exercise 2
	f1Factorial, f1Sum := f1(4)
	fmt.Println(f1Factorial, f1Sum)

	// Exercise 3
	exerciseThree := myFunc("5")
	fmt.Println(exerciseThree)

	// Exercise 4 and 5
	exerciseFour := sum(2, 4, 6)
	fmt.Println(exerciseFour)

	// Exercise 6
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

	// Exercise 7
	animals = []string{"Lion", "tiger", "bear"}
	result = searchItem(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItem(animals, "lion")
	fmt.Println(result) // => true

	// Exercise 8
	exerciseEight()

	// Exercise 9
	a := exerciseNine
	fmt.Printf("%T\n", a)
	a(5)
}

// Exercise 1
func cube(x float64) {
	math.Pow(x, 3)
}

func f1(n uint) (int, int) {
	var x int = 1
	for i := int(n); i > 0; i-- {
		x *= i
	}
	var sum int = 0
	for i := 0; i <= int(n); i++ {
		sum += i
	}
	return x, sum
}

func myFunc(n string) int {
	integer, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	secondInteger := integer * 11
	thirdInteger := integer * 111
	sum := integer + secondInteger + thirdInteger
	return sum
}

func sum(a ...int) (total int) {
	for _, v := range a {
		total += v
	}
	return
}

func searchItem(a []string, b string) bool {
	for _, v := range a {
		if strings.EqualFold(v, b) {
			return true
		} else {
			continue
		}
	}
	return false
}

func print(msg string) {
	fmt.Println(msg)
}

func exerciseEight() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}

func exerciseNine(x int) {
	fmt.Println(x)
}
