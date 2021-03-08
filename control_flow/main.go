package main

import (
	"fmt"
	"os"
)

func main() {
	// You can only use booleans in control flow statements, contrary to java and python, as
	// there is no conception of "truthiness" which you may know from ruby.
	// You do not need to use parasentheses.
	// Otherwise if/else if/else flow is simialr to other languages.

	userInput()
}

func userInput() {
	// Typically use the os package.
	fmt.Println("os.Args", os.Args)
	// os.Args is indexed from zero. The path is always going to be the zero. The first argument is indexed at 1.
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("No. of items inside os.Args:", len(os.Args))
}
