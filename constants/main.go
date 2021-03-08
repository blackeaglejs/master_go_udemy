package main

import "fmt"

func main() {
	// Constants have to be set when declared.
	const days int = 7

	// Variables don't have to be declared.
	var i int
	fmt.Println(i)

	// Constants are checked at compile time, whereas variables are checked at runtime.
	// Group constants let you set more than one variable to the same value and type.
	const (
		min1 = -500
		min2 // this gets set by min1. outside of group syntax, this would be invalid.
		min3 // this gets set by min2. outside of group syntax, this would be invalid.
	)

	fmt.Println(min1, min2, min3)

	// Constant Rules
	// 1. Constants can't be changed after being set - they are frozen. It will fail at compile time.
	// 2. Constants cannot be initialized at run time - only at compile time.
	// const power = math.Pow(2,3) <- this is invalid because you're trying to use a runtime function at compile time.
	// 3. You can't use a variable to initialize a  constant because variables belong to runtime.
	// 4. You *can* set constants with compile time functions. For example:
	l1 := len("hello")
	fmt.Printf("The length of %q is %d \n", "hello", l1)
}
