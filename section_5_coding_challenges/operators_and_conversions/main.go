package main

import "fmt"

func main() {
	exerciseOne()
}

func exerciseOne() {
	var i int = 3
	var f float64 = 3.2

	iFloat := float64(i)
	fInt := int(f)
	fmt.Printf("iFloat is value %f, type %T; fInt is value %d, type %T\n", iFloat, iFloat, fInt, fInt)
}
