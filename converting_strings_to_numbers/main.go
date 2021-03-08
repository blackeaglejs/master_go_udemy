package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Works to convert number to string.
	s := string(99)
	fmt.Println(s)

	// s1:= string(44.2) would fail. Instead:
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	s1 := "3.123" // Type string
	fmt.Printf("%T\n", s1)

	// STring to Float gives the string and the numeric precision
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)

	// Atoi takes a base 10 string and returns an integer
	i, err := strconv.Atoi("-50")
	_ = err
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2)

}
