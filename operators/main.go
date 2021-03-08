package main

import "fmt"

func main() {
	a, b := 4, 2

	r := (a + b) / (a - b) * 2
	fmt.Println(r) // -> 6

	// Modulo can only be used on integer. All other mathetmatical operators can be used on all numeric types.
	r = 9 % a // => 1

	// Reassing for next set
	a, b = 2, 3

	// increment assignment
	a += b

	//decrement assignment
	b -= 2

	// multiplication assignment
	b *= 10

	// division assignment
	b /= 5

	// modulus assignment
	a %= 3

	// increment statement adds one
	a++

	// decrement statement subtracts one
	b--

	// Comparison operators
	a, b = 6, 10
	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(a > 5, a >= 5)

	// Logical operators
	a, b = 5, 10
	// Logical AND operator will only check the second expression if the first one is TRUE
	fmt.Println(a > 1 && b == 10)
	// Logical OR operator will only check the second expression if the first one is FALSE.
	fmt.Println(a != 5 || b != 10)
	// Logical NOT operator flips to the opposite of what the expression evaluates to.
	fmt.Println(!(a == 5))
}
