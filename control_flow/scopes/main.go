package main

import (
	"fmt"
	f "fmt"
)

// Scope is where a variable can be seen.
// You cannot reuse variable names in the same scope, but you can declare it in a different scope.

// Three scope levels.
// 1. File Scope.
//// a. import statements are scoped to the file.
// 2. Package Scope.
//// a. any constant or variable declare outside a function is scoped to the package.
// 3. Block Scope.
//// a. a constant or variable can be declared in a function and is scoped to that function only.

const done = false

func main() { // package scoped

	// block scoped: visible until the end of the block "}"
	var task = "Running:"
	fmt.Println(task, done) // => Running: false (this is done from package scope)
	f.Println("Bye bye!")

	// names must be unique only within the same scope
	const done = true                        // local scoped
	fmt.Printf("done in main(): %v\n", done) // => done in main(): true
	f1()
}

func f1() {
	fmt.Printf("done in f(): %v\n", done) //this is done from package scope
}
