package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Go was released after the first dual core processor was released by Intel.
	// Go was the first major language built with concurrency as a built-in feature.
	// Concurrency *is not* the same as parallelism.
	// With concurrency, you add a new goroutine if one goroutine is blocking. The next goroutine gets picked up when the previous one finishes.
	// Parallelism is being able to run multiple goroutines at the same time.

	// A goroutine is a function that can be called at the same time as other functins.
	// You add the go keyword before a function to call that function in a goroutine.                         O
	// Goroutine are smaller than threads, take minimal stack space, and can run concurrently with other functions.
	// The stack size of a goroutine stays as small or large as necessary.
	// Goroutines are scheduled by their own Go Scheduler, rather than the OS kernel. This is much faster.
	// Goroutines have no identity.

	fmt.Println("main execution started")
	fmt.Println("No. of CPUs: ", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)

	fmt.Println("GOMACPROCS:", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup // waits for all goroutines to finish exeuction before closing.

	wg.Add(1) // Adds the number of goroutines to wait for.

	// Main does not wait for this to finish. The whole program will execute in main(). We normally deal with this using wait groups.
	// wg needs to be passed as a pointer in order to work correctly.
	go f1(&wg)

	f2()

	wg.Wait() // wait until this is complete.
	fmt.Println("main execution stopped")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// sleep for a second to simulate an expensive task.
		time.Sleep(time.Second)

	}
	fmt.Println("f1 execution finished")
	wg.Done() // tells the waitgroup that this is complete.
}

func f2() {
	fmt.Println("f2 execution started.")
	fmt.Println("f2 execution finished.")
}
