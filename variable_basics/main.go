package main

import "fmt"

func main() {
	// Declaring a variable
	// var declares a variable.
	// The second term is the variable name.
	// The third term is the variable type.
	// The fourth (optional) term is the value of the variable.
	var age int = 30
	fmt.Println("Age:", age)

	// Note that in some cases it is possible to not declare the type as Go can infer it.
	var name = "Zoheb"
	fmt.Println("Your name is: ", name)

	// You can also use the short declaration operator. Go will infer the type.
	// If calling a function to return a value, Go will get the type from the function return.
	// You MUST use a variable or the compilation will fail.
	s := "Learning Golang"
	fmt.Println(s)

	// You cannot use short declaration operator with existing variables.
	// This is wrong.
	// name := "David"

	// You can declare multiple variables simultaneously using a comma.
	car, cost := "Audi", 50000
	fmt.Printf("The cost of a %v is %v. \n", car, cost)

	// You can also declare multiple variables using the var() syntax.
	// It will assign default values if not supplied.
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Printf("%v makes %v and their gender is %v. \n", firstName, salary, gender)
	// Use short declaraion for this when you know types and values.
	// Use classic declaration for multiple assignment when you do not use the values.

	// Note that Golang (like C and C++) is statically typed, meaning that you must explicitly declare
	// the variable type, or the compiler must be able to infer it. Otherwise the program will fail at
	// compile time. In dynamically typed languages (like Ruby), the program will instead fail at runtime
	// if it cannot determine a type.

	// Separately, if you you declare one value with one type and another with another type, you must perform
	// type conversion before those values can work together.
	var a = 4   // This declaration makes variable "a" an int.
	var b = 5.2 // This declaration makes variable "b" a float64.

	// a = b would fail because of mismatched types.
	a = int(b) // This would succeed because of proper type conversion.
	fmt.Println(a)

	// A "literal" is the explicit value of a variable.
	stringLiteral := "5" // This is a string literal
	intLiteral := 5      // This is an int literal.
	floatLiteral := 5.01 // This is a float literal.
	fmt.Println(stringLiteral, intLiteral, floatLiteral)

	// Golang declares zero values for most types to keep code robust.
	// Int => Initialized as 0
	// Float64 => Initialized as 0.0
	// String => Initialized as ""
	// Bool => Initialized as false

	// Naming conventions
	/*
		Start with a letter or underscore.
		Case matters - lower-case vs upper-case is two different variables.
		Can't use go keywords.
		Use just the first letters.
		Smaller variable names in small scopes, full words in larger scopes.
		Don't use snake case - camel case is idiomatic in golang.
		Acronmys should be capitalized.
	*/
}
