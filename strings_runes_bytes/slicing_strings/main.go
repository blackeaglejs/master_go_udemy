package main

import "fmt"

func main() {
	// Slicing strings is very efficient because the string is the backing array.
	// It always stays the same.
	s1 := "I love Golang!"
	fmt.Println(s1[2:5])

	// For slicing a rune, you must convert the string to rune, then slice the rune.
	// Please note that this inefficient since you need a new backing array each time.l
	s2 := "中文维基是世界上"
	fmt.Println(s2[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.
	// returning a slice of runes
	// 1st step: converting string to rune slice
	rs := []rune(s2)
	fmt.Printf("%T\n", rs) // => []int32

	// 2st step: slicing the rune slice
	fmt.Println(string(rs[0:3])) // => 中文维
}
