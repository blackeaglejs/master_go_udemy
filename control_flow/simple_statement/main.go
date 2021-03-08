package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	// Initialization statements are useful when a block returns an error - in this scenario the error
	// is scoped to the if block. These statements can be chained with else if statements if needed.
	if i, err := strconv.Atoi("20x"); err == nil {
		fmt.Println("No error, i is:", i)
	} else {
		fmt.Println(err)
	}

	// Chaining example.
	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		fmt.Println("KM is", km)
	}
}
