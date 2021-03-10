package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) //unbuffered channel

	c2 := make(chan int, 3) // buffered channel
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	// we sleep for a second to give time to the goroutine to finish
	time.Sleep(time.Second)

	// Unbuffered channel blocks the channel until main wakes up.
}
