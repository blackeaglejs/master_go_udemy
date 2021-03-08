package main

import (
	"fmt"
	"strconv"
)

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
}

func exerciseOne() {
	var i int = 3
	var f float64 = 3.2

	iFloat := float64(i)
	fInt := int(f)
	fmt.Printf("iFloat is value %f, type %T; fInt is value %d, type %T\n", iFloat, iFloat, fInt, fInt)
}

func exerciseTwo() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// Convert i to string
	iString := strconv.Itoa(i)
	// convert s2 to int
	s2Int, _ := strconv.Atoi(s2)
	// Convert f to string
	fString := strconv.FormatFloat(f, 'f', -1, 64)
	// Convert s2 fo float64
	s1Float64, _ := strconv.ParseFloat(s1, 64)

	// Print out the value and type of each.
	fmt.Printf("iString value is %s, type is %T\n", iString, iString)
	fmt.Printf("s2Int value is %d, type is %T\n", s2Int, s2Int)
	fmt.Printf("fString value is %s, type is %T\n", fString, fString)
	fmt.Printf("s1Float64 value is %f, type is %T\n", s1Float64, s1Float64)
}

func exerciseThree() {
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5
	b := 6.2 * float64(a)
	fmt.Println(b)
}
