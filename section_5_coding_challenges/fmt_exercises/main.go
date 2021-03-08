package main

import "fmt"

func main() {
	exerciseOne()
	exerciseTwo()
}

func exerciseOne() {
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

func exerciseTwo() {
	const x float64 = 1.422349587101

	fmt.Printf("%.4f\n", x)
}
