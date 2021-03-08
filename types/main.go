package main

import "fmt"

func main() {
	//8-bit int
	var i1 int8 = 100
	fmt.Printf("%T\n", i1)

	//16-bit unsigned int
	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// FLOAT (32 or 64 depending on OS)
	var f1, f2, f3 float64 = 1.1, -0.2, .5
	fmt.Println(f1, f2, f3)

	// RUNE Type
	var r rune = 'f'
	fmt.Printf("%T\n", r)

	// BOOL TYPE
	var b bool = true
	fmt.Printf("%T\n", b)

	// Strings need to be declared with double quotes.
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	// Arrays and Slices
	// Arrays are of fixed length, slices are of dynamic length.
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// Slice
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	// Map
	// Unordered group of elements.
	// The type in brackets is the type of the key, the type after the brackets is the type of the value.
	balances := map[string]float64{"USD": 2332.2, "EUR": 511.11}
	fmt.Printf("%T\n", balances)

	// Struct is a user defined type.
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// Pointer is a variable that stores the memory address of another variable -
	// pointers have a default value of nil. They are defined using the ampersand syntax attachd to variable
	// that they are pointing at.
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// Function type
	fmt.Printf("%T\n", f)
}

func f() {
	// write your code here
}
