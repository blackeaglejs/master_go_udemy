package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Defer statements postpone execution of a statement until the return of the function around them.
	defer foo()
	bar()

	fmt.Println("Just a string after deferring foo() and calling bar()")

	// If you defer multiple functions, it will go in reverse orders from the most recently deferred to the first deferred.

	defer foobar()

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func foobar() {
	fmt.Println("This is foobar()!")
}
