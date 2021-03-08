package main

import "fmt"

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
}
