package main

import "fmt"

func main() {
	// Exercise One
	ExerciseOne()

	// Exercise Two
	ExerciseTwo()

	// Exercise Three
	x, y := 5.5, 8.8
	ExerciseThree(&x, &y)
	fmt.Printf("x is %v, y is %v\n", x, y)
}

// ExerciseOne has me declare a pointer to a variable and then show the address of the pointer, its vaLue, and the original variable's address.
func ExerciseOne() {
	x := 10.10
	fmt.Printf("Address of x: %v\n", &x)
	ptr := &x
	fmt.Printf("ptr type is %T and value is %v\n", ptr, ptr)
	fmt.Printf("ptr address is %v, value of x through pointer is %v\n", &ptr, *ptr)
}

// ExerciseTwo has me divide two variables through their pointers and then store that value in a new variable.
func ExerciseTwo() {
	x, y := 10, 2
	ptrX, ptrY := &x, &y
	z := *ptrX / *ptrY
	fmt.Println("Value of z is", z)
}

// ExerciseThree Swaps the values of the variables using pointers.
func ExerciseThree(x *float64, y *float64) {
	*x = *y
	*y = *x
}
