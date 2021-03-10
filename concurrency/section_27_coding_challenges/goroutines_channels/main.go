package main

import "fmt"

func main() {
	ExerciseOne()
	ExerciseTwo()
	ExerciseThree()
	ExerciseFour()
	ExerciseFive()
}

func ExerciseOne() {
	var c1 chan float64

	c2 := make(chan<- rune)
	c3 := make(<-chan rune)

	c4 := make(chan int, 10)

	fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)
	fmt.Println("=====================")
}

func ExerciseTwo() {
	var stringChannel chan string
	go func(s string, c chan string) {
		c <- s
	}("Hello Zoheb!", stringChannel)
}

func ExerciseThree() {
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}

func ExerciseFour() {
	c4 := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, c4)
	}

	for i := 1; i <= 50; i++ {
		go fmt.Println(<-c4)
	}
}

func power(n int, c chan int) {
	c <- n * n
}

func ExerciseFive() {
	c5 := make(chan int)

	for i := 1; i <= 50; i++ {
		go func(a int, c chan int) {
			c <- a * a
		}(i, c5)
	}

	for i := 1; i <= 50; i++ {
		go fmt.Println(<-c5)
	}
}
