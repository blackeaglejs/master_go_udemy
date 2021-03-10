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

	// The below series of functions is a race condition. The reason this is the case is because both functions are operating on the same data spce.
	// This is not goood because it means the final result of the variable n will be unpredictable and dependent on the order execution occurs.
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final value of n:", n)
}
