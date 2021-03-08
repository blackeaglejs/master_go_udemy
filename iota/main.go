package main

import "fmt"

func main() {
	// Iota is a number generator - by default it adds by one each time, but this can be manipulated.
	const (
		c1 = iota // sets to zero
		c2 = iota // sets to one
		c3 = iota // sets to two
	)

	fmt.Println(c1, c2, c3)

	// This is also valid.
	const (
		c11 = iota // sets to zero
		c22        // sets to one
		c33        // sets to two
	)

	fmt.Println(c11, c22, c33)

	// You can perform mathematical operations as well.
	const (
		c111 = iota * 2 // sets to zero
		c222            // sets to two
		c333            // sets to four
	)
	fmt.Println(c111, c222, c333)

	// You can also skip using blank operator.
	const (
		c1111 = iota // sets to zero
		_            // skips one
		c2222        // sets to two
		c3333        //sets to three
	)
	fmt.Println(c1111, c2222, c3333)
}
