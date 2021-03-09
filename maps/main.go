package main

import "fmt"

func main() {
	// Maps are like dictionaries in Python, or hashes in Ruby.
	// All keys are statically type.
	// Keys must be unique but values can be duplicated.
	// Can use any comparable type as a KeyMap
	// Do not use float as a key -> It causes issues.
	// Maps are unordered.
	// Maps can only be compared to nil, not to other map.

	var employees map[string]string
	// Zero value of a map is nil.
	fmt.Printf("%#v\n", employees)                      // -> map[string]string(nil).
	fmt.Printf("No. of elements: %d\n", len(employees)) // => No. of elements: 0

	// The len() function will return the number of key/pairs.
	fmt.Printf("No. of elements: %d\n", len(employees)) // => No. of elements: 0
	// The value of an empty key will return the default value for that type.
	fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"]) // => The value for key "Dan" is ""

	// Slices cannot be used for a key beause they're not comparable.
	// var m1 map[[]int]string // error -> invalid map key type []int (slices are not comparable)
	// Arrays can be used for a key because they *are* comparable.

	// Maps need to be initialized before they can have values. The {} at the end of a map initializes it.
	// employees["Dan"] = "Programmer" // error -> panic: assignment to entry in nil map
	people := map[string]float64{} // empty map

	// inserting key:value pairs in a map
	people["John"] = 30.5
	people["Marry"] = 22

	fmt.Printf("%v\n", people) // => map[John:30.5 Marry:22]

	// declaring and initializing a map using the make() function:
	map1 := make(map[int]int)
	fmt.Printf("map1: %#v\n", map1) // -> map1: map[int]int{}
	map1[4] = 8
	// Last comma mandatory if you declare a map on multiple lines. Not required on single-line maps.
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
		//50: "ABC"  //illegal, all keys have the same type which is string and all values have the same type which is float64
		"CHF": 600, //this last comma (,) is mandatory when declaring the map on multiple lines
	}
	fmt.Println(balances) // => map[CHF:600 EUR:555.11 USD:233.11]

	m := map[int]int{1: 3, 4: 5, 6: 8}
	_ = m

	// If a key exists, it updates the value.
	// If a key does not exist, it adds the key and value.
	balances["USD"] = 500.5
	balances["GBP"] = 800.8
	fmt.Println(balances) // => map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]

	// If you search for a key and it does not exist, it'll return the default value for the values.
	// if you assign two variables to a key lookup, the first return will be the value, and the second will be whether the key exists in the map.
	v, ok := balances["RON"]

	//v is the key's corresponding value
	// ok is bool type value which is true if the key exists and false otherwise

	if ok {
		fmt.Println("The RON Balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	// You can iterate with k,v := range, but Golang is optimized for fast lookups and not fast looping.
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	// To remove a key value pair, you can use the delete() function with the map and the key.
	delete(balances, "USD")

	// Comparing maps is not possible. You can only compare them to nil.
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "X"}
	_, _ = a, b
	// fmt.Println(a == b) // error -> invalid operation: a == b (map can only be compared to nil)

	// If you REALLY want to compare them, you can convert them to string and use the == operator.
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
	// How Maps Work under the hood.
	// The map header is a golang construct that contains memory pointers to the key-value pairs, which are stored separately.
	// If you create a copy of a map with a differet name, altering the copy will also alter the original because the map header is the same.
	// When creating a map variable Go creates a pointer to a map header value in memory.
	// The key: value pairs of the map are not stored directly into the map.
	// They are stored in memory at the address referenced by the map header.

	friends := map[string]int{"Dan": 40, "Maria": 35}

	// neighbors is not a copy of friends.
	// both maps reference the same data structure in memory
	neighbors := friends

	// modifying friends AND neighbors
	friends["Dan"] = 30

	fmt.Println(neighbors) // -> map[Dan:30 Maria:35]

	// True cloning requires manually creating a new map.
	// 1. initialize a new map
	colleagues := make(map[string]int)
	// 2. use a for loop to copy each element into the new map
	for k, v := range friends {
		colleagues[k] = v
	}
	// 3. friends and colleagues are different because they were not initialized off each other and thus have different headers and memory pointers.
}
