package main

import "fmt"

func main() {
	// basicsOfChannels()
	goroutinesAndChannels()
}

func basicsOfChannels() {
	// Channels allow goroutines to communicate with each other.
	// Channels must have all data of the same point - the type declared.
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	c := make(chan int)

	// Send statements let you send data.
	// Receive data let you receive data.

	// <- channel operator.

	// SEND
	c <- 10

	// RECEIVE
	num := <-c

	_ = num

	close(c) // closes the channel.
}

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	c <- f
}

func goroutinesAndChannels() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	f := <-ch // this is a wake-up call for the channel
	fmt.Println(f)

	// 20 goroutines for the channel.
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	// anonymoucs function instead.
	for i := 5; i <= 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}
			c <- f
		}(i, ch)
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}
}
