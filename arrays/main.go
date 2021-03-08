package main

import (
	"fmt"
	"strings"
)

func main() {
	// Arrays are a composite, indexable type.
	// All array elements must be the same type.
	// Arrays have a fixed length.
	// Array length is determined at compile time.
	// Arrays are stored in continuous memoery locations, making them efficient.
	declaringArrays()
	arrayOperations()
	keyedElementArrays()
}

func declaringArrays() {
	var numbers [4]int
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%+v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%v\n", a4)

	// ellipsis operator determines the length of an array by how many elements we add initially.
	a5 := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("%v\n", a5)
	fmt.Printf("the length of a5 is %d\n", len(a5))

	// for multiple line arrays, you need to add a comma on the last line.
	a6 := [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Printf("%v\n", a6)
}

func arrayOperations() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%v\n", numbers)

	// You access arrays by index number, so an index number can be out of range based on the size of the array

	// iteration
	for i, v := range numbers {
		fmt.Println("index:", i, " value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index", i, " value:", numbers[i])
	}

	// Multidimensional array
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	// Equality is when two arrays have the same length, content, and *order*
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	// When creating one array from another, the two arrays are not linked. modifying one does not modify the other.
	m[0] = -1
	fmt.Println("n is equal to m:", n == m)
}

func keyedElementArrays() {
	// You can specify which value goes into which position in the array, even out of order.
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // -> [5 10 7]

	// You don't have to specify all the elemments in an array.
	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts) //[0 0 50]

	// If using ellipses syntax, the length of the array will be one longer than the last element you specify.
	names := [...]string{
		4: "Dan",
	}
	fmt.Println(len(names)) // -> 5

	// If you assign some elements with a key and oethers without a key, then the unkeyed one will be filled
	// on the last index in the array.
	cities := [...]string{
		5:        "Paris",
		"London", // this is at index 6
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // => [false false false false false true true]
}
