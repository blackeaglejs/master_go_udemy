package main

import "fmt"

func main() {
	// Similar to arrays in syntax, but have dynamic length
	// Length of slice belongs to runtime, not compiletime.
	// Uninitialized slice value is nil.
	declaringSlices()
	comparingSlices()
	appendingSlices()
}

func declaringSlices() {
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	// Elements cannot be assigned to a nil slice
	// cities[0] = "London"

	// Initialized slice
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best friend is ", myFriend)

	friends[0] = "Gabriel"

}

func comparingSlices() {
	// Unitialized slice
	var n []int
	fmt.Println(n == nil)

	// Initialized slice
	m := []int{}
	fmt.Println(m == nil)

	// Slices can only be compared to nil. To comapre slices, you must iterate over each slice element.
	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	// This will fail
	// fmt.Println(a == b)

	var eq bool = true

	if len(a) != len(b) {
		eq = false // check length first - if length is unequal, the slices are not equal.
	} else {
		for i, valueA := range a {
			if valueA != b[i] {
				eq = false
				break // break if there are unequal elements. In this check, order matters.
			}
		}
	}

	if eq {
		fmt.Println("Slices are equal.")
	} else {
		fmt.Println("Slices are unequal.")
	}
}

func appendingSlices() {
	numbers := []int{2, 3}

	// append function lets you add to slices - it removes the slice entirely and returns a new slice.
	numbers = append(numbers, 10)
	fmt.Println(numbers)

	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)

	// appending two slices
	n := []int{100, 200}
	numbers = append(numbers, n...) // note the ellipsis operator for the appended slice
	fmt.Println(numbers)

	// Copy function will pull in the values from a source slice into the destination slice.
	// If the source slice does not have any values, nothing will be added into the destination slice.
	// Returns the number of elements copied.
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)
}
