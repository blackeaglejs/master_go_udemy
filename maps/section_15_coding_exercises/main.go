package main

import "fmt"

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
}

func exerciseOne() {
	fmt.Println("Exercise 1")
	var m1 map[string]string
	fmt.Printf("%#v\n", m1)

	m2 := map[int]string{1: "Eminem", 2: "Dr. Dre"}
	m2[10] = "Abba"
	fmt.Println(m2[1])
	fmt.Println(m2[5])
	fmt.Println("===================")
}
func exerciseTwo() {
	fmt.Println("Exercise 2")
	var m1 = map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}
	_, _ = m2, m3

	// This below line fails because you cannot compare two maps to each other.
	// fmt.Println(m2 == m3)
	fmt.Println("===================")
}
func exerciseThree() {
	fmt.Println("Exercise 3")
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1)
	for k, v := range m {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
	fmt.Println("===================")
}
