package main

import (
	"fmt"
	"time"
)

func main() {
	// Select statement lets goroutine wait on multiple communications.
	// Allows us to wait on multiple operations.
	// Select is only used with channels.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salut!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: // waits for messages from channel 1
			fmt.Println("Received:", msg1)
		case msg2 := <-c2: // waits for messages from channel 2
			fmt.Println("Received:", msg2)
		default: // normally a select statment is blocking, but the default stops that because it is the default routine.
			fmt.Println("no activity")
		}
	}
}
