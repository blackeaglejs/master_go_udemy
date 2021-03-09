package main

import (
	"fmt"
	"strings"
)

func main() {
	// Variadic functions are functions that varying numbers of functins.
	// Use them when the number of arguments is unknown.
	// Also use them as an alternative (sometimes) to passing a slice.
	f1(1, 2, 3, 4)
	f1()

	// Append is an example of a varidaic operator.
	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	// You can pass in slices to variadic funtions as well.
	s, p := sumAndProduct(2., 5., 10.)
	fmt.Println(s, p) // -> 17 100

	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart
}

// Pass any number of Ints
func f1(a ...int) {
	fmt.Printf("%T\n", a) // => []int, slice of int
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	// When passign a slice, you'll be able to change the value of a slice.
	a[0] = 50
}

func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.
	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

// You can also mix variadic and non-variadic parameters.

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name : %s", age, fullName)
	return returnString
}
