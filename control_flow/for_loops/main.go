package main

import "fmt"

func main() {
	// For loops have 3 parts.
	// 1. Initialization Statement.
	// 2. Conditional to break the loop.
	// 3. (Optional) Follow on statement. This can also be placed inside the curley braces.
	// The code inside the curly braces will execute on each iteration of the loop.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Golang does not have an explicit while code, but it is possible to create the same effect by:
	// 1. Move the initialization statement outside the "for" loop.
	// 2. Stage for [conditional]
	// 3. add your code for each iteraiton.
	// Excluding the conditional from #2 will create an infinite loop.
	j := 10
	for j > 0 {
		j--
	}

	// Continue statements. It moves us to the next iteration of the for loop if a condition is met.
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skipping the remaining code in this iteration
		}
		fmt.Println(i)
	}

	// Break statements will end the loop if a condition will met. This can resolve infinite loops.
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 { //if 10 numbers were found, break!
			break //it breaks the current loop (inner loop if there are more loops)
		}
	}

	// the break statement is not terminating the program entirely;
	fmt.Println("Just a message after the for loop")
}
