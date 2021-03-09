package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
	exerciseFour()
	exerciseFive()
	exerciseSix()
}

func exerciseOne() {
	fmt.Println("Exercise 1")
	var name string = "Zoheb"
	country := "United States"
	fmt.Printf("Your Name: %v \nCountry: %v\n", name, country)
	fmt.Printf("He says: %q\n", "Hello")
	fmt.Println(`C:\Users\a.txt`)
	fmt.Println("====================")
}
func exerciseTwo() {
	fmt.Println("Exercise 2")
	r := 'ă'
	fmt.Printf("Type of r is %T\n", r)
	str1, str2 := "ma", "m"
	newStr := str1 + str2 + string(r)
	fmt.Println(newStr)
	fmt.Println("====================")
}
func exerciseThree() {
	fmt.Println("Exercise 3")
	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c", s1[i])
	}
	fmt.Println("")
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("Rune: %c, Index: %v\n", r, i)
		i += size
	}
	fmt.Println("")
	fmt.Println("====================")
}
func exerciseFour() {
	fmt.Println("Exercise 4")
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// Strings are immutable, the line below will fail.
	// s1[0] = 'x'

	// printing the number of runes in the string
	fmt.Println(len(s1))
	// Use the utf8 package to count runes in string.
	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Println("====================")
}
func exerciseFive() {
	fmt.Println("Exercise 5")
	s := "你好 Go!"
	// Convert to byte slice.
	sAsSlice := strings.Split(s, "")
	fmt.Println(sAsSlice)

	// Iterate over the byte slice.
	for i, v := range sAsSlice {
		fmt.Printf("Index is %v, Byte Size is %v\n", i, len(v))
	}
	fmt.Println("====================")
}
func exerciseSix() {
	fmt.Println("Exercise 6")
	s := "你好 Go!"
	runeSlice := []rune(s)
	for i, v := range runeSlice {
		fmt.Printf("Index is %v, Rune is %c\n", i, v)
	}
	fmt.Println("====================")
}
