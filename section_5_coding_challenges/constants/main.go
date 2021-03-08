package main

import (
	"fmt"
)

func main() {
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	fmt.Printf("Days In Week: %v, Speed of Light (m/s): %v, Value of Pi: %v\n", daysWeek, lightSpeed, pi)

	const secPerDay = 24 * 60 * 60
	const daysYear = 365
	fmt.Printf("Number of seconds in a year is %v\n", secPerDay*daysYear)

	// Excercise 4 - fix below code.
	const x int = 10

	// declaring a const of type slice int([]int) The fix is to comment this out. there isn't such a thing as an
	// array or slice literal which means no array or slice constants.
	// const m = []int{1: 3, 4: 5, 6: 8}
	// _ = m

	// Exercise 5 - fix below code.

	const a int = 7
	const b float64 = 5.6
	// The const c setting can be fixed by fixing mismatched types.
	const c = float64(a) * b

	// This code can be fixed by commenting it out OR setting xc explicitly to 8, as you can't set a constant
	// with a variable.
	// x = 8
	const xc int = 8

	// This code has to be commented out because you can't use a runtime function to set a constant.
	// const noIPv6 = math.Pow(2, 128)
}
