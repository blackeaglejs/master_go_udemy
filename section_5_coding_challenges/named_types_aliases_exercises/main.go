package main

import "fmt"

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
}

func exerciseOne() {
	type duration int
	var hour duration
	fmt.Printf("Value of hour is %v, Type of hour is %T\n", hour, hour)
	hour = 3600
	fmt.Printf("Value of hour is %v, Type of hour is %T\n", hour, hour)
}

func exerciseTwo() {
	type duration int
	var hour duration = 3600
	minute := 60
	fmt.Println(int(hour) != minute)
}

func exerciseThree() {
	type mile float64
	type kilometer float64
	const m2km = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer
	kmBerlinToParis = kilometer(float64(mileBerlinToParis) * m2km)
	fmt.Println("Distance from Berlin to Paris in km is", kmBerlinToParis)
}
