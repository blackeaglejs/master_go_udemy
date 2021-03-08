package main

import "fmt"

func main() {
	var x = 3   // type Int
	var y = 3.1 // type float64

	// x = x * y // fails because these are of different types.
	x = x * int(y) // succeeds because both are ints.
	fmt.Println(x)

	// Multiply two floats and convert to int.
	x = int(float64(y) * y)

	var a = 5        // int
	var b int64 = 22 // int64
	// These are not the same type
	fmt.Printf("%v: %T, %v: %T \n", a, a, b, b)

}
