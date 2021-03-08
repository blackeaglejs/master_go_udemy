package main

import "fmt"

func main() {
	// fmt.Println() lets you pass as many arguments as you like, and prints to a line.
	name := "Zoheb"
	fmt.Println("Hello, playground. I am", name)

	a, b := 4, 6
	// This will automatically whitespace.
	fmt.Println("Sum:", a+b, "Mean:", (a+b)/2)

	// fmt.Printf() lets you  custom format your text and annotations - it does not automatically newline.
	// %d formats to integer.
	fmt.Printf("Your age is %d. \n", 32)
	// %f formats to float.
	fmt.Printf("x is %d, y is %f \n", 5, 6.8)
	// Adding .3 to the %f will show a certain number of decimal places.
	fmt.Printf("x * y is %.03f", 5.2*6.8)
	// Use %s for strings.
	// Use %v for all types to conversion.
	// %T is for types.
	// %t is for booleans.
	// %b will convert the integer/float to base-2.
	// %x will convert to hexadecimal.
}
