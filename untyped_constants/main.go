package main

import "fmt"

func main() {
	// This is a typed  constant.
	const a float64 = 5.1

	// This is an untyped constant.
	const b = 6.7

	// If setting an untyped constant by a compile time function, you *can* mix types, but only if you are
	// not mixing two typed constants.
	// For example:
	// Incorrect
	// const x int = 5
	// const y float64 = 2.2 * x
	// Correct
	const x = 5
	const y = 2.2 * 5
	fmt.Printf("%T\n", y)

	// Variables set by constants can have their type changed.
	var i int = x     // x will become type int.
	var j float64 = x // x will become type float64.
	var p byte = x    // x will become type byte.

	fmt.Println(i, j, p)

	// Constants (when converted) will have their type inferred.
}
