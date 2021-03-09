package main

import "fmt"

func main() {
	// Anonymous functions do not have a name and are generally returned for closures.
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!")

	// Functions can return other functions - this is normal in golang.
	// Functions are first class in golang because they can be assigned to other functions.
	a := increment(10)
	fmt.Printf("%T\n", a)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
