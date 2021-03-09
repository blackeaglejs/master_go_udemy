package main

import (
	"fmt"
	"math"
)

func main() {
	// A function is a small piece of code that does a particular task based on input values
	// Go prefer simple words or CamelCase
	// Go can return multiple values from a function.
	f1()
	f2(5, 7)
	f3(4, 5, 6, 7.7, 8.8, "golang")
	p := f4(5.1)
	fmt.Println(p)
	a, b := f5(8, 9)
	fmt.Printf("a:%d, b:%d\n", a, b) // => a:42, b:13

	ss := sum(4, 5)
	fmt.Println(ss) // -> 9

}

func f1() {
	fmt.Println("This is f1() function")
}

func f2(a int, b int) {
	// a and b are required to get types in the function declaration.
	// a and b are local to the function and cannot be used anywhere else.
	fmt.Println("Sum:", a+b)
}

// This is a shorthand notation that you can use to declare a single type for several inputs
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// Return variables must have their type declared.
func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

// This uses a named parameter. You can use naked returns with this.
func sum(a, b int) (s int) {
	s = a + b
	fmt.Println(s) // s=> 0
	return         // naked return
}
