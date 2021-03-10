package main

import (
	"fmt" // packages have to be imported in each go file. 
	"master_go_udemy/packages/numbers" // needs to be defined in relation to $GOPATH - there is an issue for this. 
	// goimports command will automatically import packages. 
)

func init() {
	// This is a special reserved function for a package. 
	// This does not take any arguments or return anything. 
	// It returns before the main functions. 
	fmt.Println("this is init")
}

func main() {
	// A package is a collection of source files each ending in .go and in a single folder.
	// Everything in a package will have the same package name.
	// Executable packages generate executables can be turned into executables and have a main() function.
	// Non-Executable packages contain useful functions but can't be executed.
	var x uint = 40
	fmt.Printf("%d is even: $t \n", numbers.Even(x))
}

func PublicVsPrivateNames{} {
	// Uppercase function names mean that a function in a package is publicly available if you import a package. 
	// Lowercase function names mean that a function in a package is only available to that package. 
}

func understandingScope() {
	// if a function is in the same package but not the same file, you can still call it in another file. 
	// constants and variables declared outside function scope are available to the entire package. 
	// You cannot declare a function more than once inside a package. 
}
