package main

import "fmt"

func main() {
	// Switches compile to if staTements, but are cleaner for longer statements.

	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python! You don't use curly braces but indentation!!")
	case "Go", "golang": // Can have multiple potential conditions in a case.
		fmt.Println("Good, go for Go! You are using curly braces {}!")
	default: // Gets executed when no case is matched.
		fmt.Println("Any other programming language is good!")
	}

	n := 5
	// comparing the result of an expression which is bool to another bool value
	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
