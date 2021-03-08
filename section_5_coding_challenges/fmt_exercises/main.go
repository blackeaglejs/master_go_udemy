package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// Print each variable
	fmt.Printf("x is %d, y is %f, z is %s, score is %v\n", x, y, z, score)
	// Reprint "gophers" with quotes
	fmt.Printf("%q\n", z)
	// Print each variable using the same verb.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)
	// Print the types of y and score.
	fmt.Printf("y is type %T, score is type %T\n", y, score)
}
