package main

import "fmt"

func main() {
	// Strings are defined in double quotes. AND NOT single quotes.
	// They are UTF8 formatted bye default.
	s1 := "Hi there  Go!"
	fmt.Printf("%s\n", s1) // => Hi there  Go!
	fmt.Printf("%q\n", s1) // => "Hi there  Go!"

	// Double quotes inside quotes must be escaped with a backslash.
	fmt.Println("He say: \"Hello!\"")
	// Backticks can also be used for strings and avoid escaping.
	fmt.Println(`He say: "Hello!"`)
	// In raw strings (backticked strings), things like \n have no meaning.
	s2 := `Hi there Go!`
	fmt.Println(s2)

	// In fmt, things like \n and backlashes do have meaning. They can be used to create multiline strings.
	fmt.Println("Price: 100 \nBrand: Nike")
	//is the same as:
	fmt.Println(`
			Price: 100
			Brand: Nike`)
	// Backslash itself has to be escaped with a Slash.
	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei")
	// concatenation can be done using + or fmt.
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!") // -> I love Go Programming!
	// A string is also a slice of bites in golang. You can use the bracket notation from arrays to get the ASCII value.
	fmt.Println("Element at index zero:", s3[0]) // => 73 (ascii code for I)
	// Strings are immutable and cannot be changed.
	// s3[5] = 'x' // => error: Cannon assign to s3[5].
}
