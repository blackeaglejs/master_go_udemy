package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	// you can use the strings package to make things easier.
	// The contain function will check if any of the test string is in the source string.
	result := strings.Contains("I love Go Programming!", "love")
	p(result) // -> True
	// The containsAny function will check if any of the characters in the test string are in the source string.
	result = strings.ContainsAny("success", "xy")
	p(result) // => false
	// The ContainsRune function checks if any of the test runes are in the test string.
	result = strings.ContainsRune("golang", 'g')
	p(result) // => true
	// The Count() function counts the number of instances of a character in a string.
	n := strings.Count("cheese", "e")
	p(n) // => 3
	// if the substr is an empty string Count() returns 1 + the number of runes in the string
	n = strings.Count("five", "")
	p(n) // => 5 (1 + 4 runes)
	// Note that any returned string is a new string since strings are immutable.
	// The ToLower() functino gives all lowercase.
	p(strings.ToLower("Go Python Java")) // -> go python java
	// The ToUpper() function gives all uppercase.
	p(strings.ToUpper("Go Python Java")) // -> GO PYTHON JAVA

	// comparing strings (case matters)
	p("go" == "go") // -> true
	p("Go" == "go") // -> false
	// You can test equality using the EqualFold() function when you want to check equality but not case.
	p(strings.ToLower("Go") == strings.ToLower("go")) // -> true // BAD
	p(strings.EqualFold("Go", "gO"))                  // -> true // GOOD

	// Repeat function repeats a string as many times as you tell it.
	myStr := strings.Repeat("ab", 10)
	p(myStr) // => abababababababababab
	// The Replace() function replaces characters in your string with other characters.
	// You must supply the number of occurrences you want to replace. -1 will replace all.
	myStr = strings.Replace("192.168.0.1", ".", ":", 2) // it replaces the first 2 occurrences
	p(myStr)                                            // -> 192:168:0.1

	// if the last argument is -1 it replaces all occurrences of old by new
	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr) // -> 192:168:0:1

	// The ReplaceAll() function replaces all character in your string with another character.
	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr) // -> 192:168:0:1

	// The Split() function will split a string based on a character and return a slice of strings.
	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)                  // -> []string
	fmt.Printf("strings.Split():%#v\n", s) // => strings.Split():[]string{"a", "b", "c"}

	// Not supplying a separator will split every character in the string, including spaces.
	s = strings.Split("Go for Go!", "")
	fmt.Printf("strings.Split():%#v\n", s) // -> []string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}
	// The Join() function will join slices of strings into a single element using whatever separator you want.
	s = []string{"I", "learn", "Golang"}
	j := strings.Join(s, "-")
	fmt.Printf("%T\n", j) // -> string
	p(j)                  // -> I-learn-Golang
	// The Fields() function will split based on whitespace and newlines.
	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) // it returns a slice of strings
	fmt.Printf("%T\n", fields)      // -> []string
	fmt.Printf("%#v\n", fields)     // -> []string{"Orange", "Green", "Blue", "Yellow"}
	// The TrimSpace() function will remove all leading tab and excessive whitespace.
	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")
	fmt.Printf("%q\n", s1) // "Goodbye Windows, Welcome Linux!"
	// You can also use the Trim() function to remove any leading and trailing characters you specify.
	s2 := strings.Trim("...Hello, Gophers!!!?", ".!?")
	fmt.Printf("%q\n", s2) // "Hello, Gophers"
}
