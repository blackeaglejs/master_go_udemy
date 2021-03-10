package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	// Fixing a race condition using a mutex.
	//1.
	var m sync.Mutex

	// The below series of functions is a (fixed) race condition. The reason this is the case is because both functions are operating on the same data spce.
	// This is not goood because it means the final result of the variable n will be unpredictable and dependent on the order execution occurs.
	// You can use race detector from golang to find race conditions.
	// It is called by adding the -race condition to your go run command.
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			//2.
			m.Lock() // locks this memory space temporarily so that only one goroutine can access a memory space at a time.
			defer m.Unlock()
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			//2.
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final value of n:", n)
}
