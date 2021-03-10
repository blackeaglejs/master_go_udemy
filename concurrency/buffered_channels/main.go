package main

import "fmt"

func main() {
	c := make(chan int, 3) // the capacity is set at 3, making this a buffered channel. unbuffered doesn't have capacity.
	fmt.Println("Channel's capacity:", cap(c))
	// sender only blocks when capacity not available in the channel.
	// receiver only  blocks when the channel is empty.

	// unbuffered channels can synchronize, while buffers cant for the blocking reason above.
	// unbuffered channels are sometimes called synchronized channels.

	// We use buffered channels typically when we know how many pieces of data we want to send down the channel at a time.
	//

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("func goroutine starts sending data into the channel")
			c <- i
			fmt.Println("func goroutine after sending data into the channel")
		}
		close(c)
	}(c)

	// receiving data from the channel
	for v := range c { // v is the value read from the channel, it's like using v := <- c2
		fmt.Println("main goroutine received value from channel:", v)

	}

	// After running the program  we notice that the goroutines start sending data
	// into the channel BEFORE the main goroutine had a chance
	// to receive data from the channel.

	// The sender of this buffered channel will block only when there is no empty slot in the channel, in this
	// case after 3 writing attempts because the channel has a capacity of 3.
	// The receiver will block on the channel when it's empty.

	// A receive operation on a closed channel will proceed without blocking
	// and yield the zero-value for the type that is sent through the channel.
	fmt.Println(<-c) // => 0

	// Sending a value into a closed channel will panic.
	// c1 <- 10 // => panic: send on closed channel
}
