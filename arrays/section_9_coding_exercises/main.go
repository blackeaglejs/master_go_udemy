package main

import "fmt"

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
}

func exerciseOne() {
	var cities [2]string
	fmt.Println(cities)

	grades := [3]float64{
		0: 95,
		1: 98,
	}
	fmt.Println(grades)

	taskDone := [...]bool{}
	_ = taskDone

	for i := 0; i < len(cities); i++ {
		fmt.Printf("Index is %d, Value is %s\n", i, cities[i])
	}

	for i, v := range grades {
		fmt.Printf("Index is %d, value is %f\n", i, v)
	}
}

func exerciseTwo() {
	nums := [...]int{30, -1, -6, 90, -6}
	evenCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
		}
	}
	fmt.Printf("The number of even numbers in array nums is %d\n", evenCount)
}

func exerciseThree() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	// Issue here is that a is not a float. Modifying fixes that.
	myArray[0] = float64(a)

	// Issue here is that myArray[3] would correspond to fourth element - this has to be excluded.
	// myArray[3] = 10.10

	fmt.Println(myArray)
}
