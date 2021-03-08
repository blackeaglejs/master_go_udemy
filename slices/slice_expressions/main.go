package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	// a[start:stop] includes start element but not stop element

	b := a[1:3] // returns elements 2,3
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	// Not including a stop will go all the way to the  end of the slice.
	fmt.Println(s1[2:])       // same as s1[2:len(s1)]
	fmt.Println(s1[:3])       // same as s1[0:3]
	fmt.Println(s1[:])        // returns entire slice
	fmt.Println(s1[:len(s1)]) // returns entire slice

	// Combine this with append!
	s1 = append(s1[:4], 100) // will return the 0,1,2,3 elements of the slice, then add 100 as a final element
	fmt.Println(s1)

	s1 = append(s1[:4], 200) // will do same as previous, but  will replace 100 with 200
	fmt.Println(s1)

}
