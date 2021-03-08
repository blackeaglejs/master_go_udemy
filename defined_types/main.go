package main

import "fmt"

func main() {
	// the same underlying type means that types can be converted and share a common sizing, but they *are*
	// different types.
	type age int        // int is the underlying type
	type oldAge age     // age is the underlying type
	type veryOldAge age // age is its underlying type

	type speed uint //unit is the underlying type
	var s1 speed = 10
	var s2 speed = 20

	// Works.
	fmt.Println(s2 - s1)

	// Does not work.
	// var x uint
	// x = s1 <- error different types

	x := uint(s1) // Works because the types match now.
	fmt.Println(x)

	var s3 speed = speed(x)
	_ = s3

	type mile float64
	type km float64
	var paristoLondon km = 465
	var distanceinMiles mile

	distanceinMiles = mile(paristoLondon) / 0.621
	fmt.Println(distanceinMiles)

	// Aliases has the form type T1 = T2. It's the same type as it's underlying type, just a new name for it.
	// byte and uint8 are aliases of the same type with different types. Rune is an alias for int32
	// Aliases should be used with caution - they are mainly used for refactoring.

	var a uint8 = 10
	var b byte

	b = a
	_ = b

	type second = uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d \n", hour/60)
}
