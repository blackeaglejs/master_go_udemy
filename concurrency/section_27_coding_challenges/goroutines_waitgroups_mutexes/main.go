package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	ExerciseOne()
	ExerciseTwo()
	ExerciseThree()
	ExerciseFour()
	ExerciseFive()
}
func ExerciseOne() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("Mr. Wick", &wg)
	wg.Wait()
}

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func ExerciseTwo() {
	var wg sync.WaitGroup
	wg.Add(3)
	go sum(1., 2., &wg)
	go sum(3., 4., &wg)
	go sum(5., 6., &wg)

	wg.Wait()
}

func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("%v + %v = %.2f\n", a, b, a+b)
	wg.Done()
}

func ExerciseThree() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(x float64) {
		fmt.Printf("Square root of %v is %v\n", x, math.Sqrt(x))
		wg.Done()
	}(3.)
	wg.Wait()
}

func ExerciseFour() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100.; i < 150.; i++ {
		go func(x float64, wg *sync.WaitGroup) {
			fmt.Printf("Square root of %v is %v\n", x, math.Sqrt(x))
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

func ExerciseFive() {
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	*b += n
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	*b -= n
	wg.Done()
}
