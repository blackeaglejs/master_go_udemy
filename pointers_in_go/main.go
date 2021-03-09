package main

import "fmt"

func main() {
	// Memory is a sequence of boxes or cells. Each cell is has an address for that memory location.
	// A variable is a convenient alphanumeric nickname for a memory location.
	// A pointer is the address of a variable. Pointers can change or mutate the data they're pointing to.
	// Go does not have pointer arithmetic.

	// The & operator is run in front of a variable name and makes it a pointer.
	name := "Andrei"
	fmt.Println(&name)

	// Iniitalization
	var x int = 2
	ptr := &x // type *int

	fmt.Printf("ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	p := new(int)
	x = 100
	p = &x
	fmt.Printf("p is of type %T with a value of %v and address %p\n", p, p, &p)
	fmt.Printf("address of x is %p\n", &x)

	// Dereferencing operator.
	*p = 90 // equal to x = 90 because they are the same.
	fmt.Println(x, *p)
	fmt.Println("*p==x:", *p == x)

	// &value => pointer // if you have a value you turn it into an address or pointer by using the ampersand operator
	// *pointer => value // if you have pointer you turn it into value value by using the star operator

	//pointers to pointers
	pointerToPointer()

	passingAndReturningPointers()
	passingAndReturningPointersOfOtherTypes()

}

func pointerToPointer() {
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)

	// pointers can be compared using the equality operator.

	// two pointers are equal if they point to the same variable or if they are nil.

}

func passingAndReturningPointers() {
	x := 8
	p := &x
	fmt.Println("value of x before calling change():", x)
	// calling change on a pointer changes the value of the variable that is being pointed to.
	change(p)
	fmt.Println("value of x before calling change():", x)
}

func change(a *int) *float64 {
	*a = 100
	b := 5.5
	return &b
}

func passingAndReturningPointersOfOtherTypes() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	// when you pass an original value, it does not change in memoery because functions work on copies.
	fmt.Println("BEFORE calling changevalues():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("AFTER calling changevalues():", quantity, price, name, sold)

	// they do however change if you use pointers.
	fmt.Println("BEFORE calling changevalues():", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER calling changevalues():", quantity, price, name, sold)

	// Note that slices and maps are just pointers to memory addresses as well - because of this, we do not use
	// pointers for these. calling a function on a slice or map means that you're directly changing the data.
	// Arrays are not quite like slices - you can pass pointers for them, but it is considered bad form to do so.
	// Only pass pointers when absolutely necessary - they put load on the garbage collector.

	// declaring a struct value
	present := Product{
		price:       100,
		productName: "Watch",
	}

	// invoking the function has no effect on the struct value.
	// the function works on and modifies a copy, not the original.
	changeProduct(present)
	fmt.Println(present) // => {100 Watch}

	// the function modifies the struct value.
	changeProductByPointer(&present)
	fmt.Println("AFTER calling changeProductByPointer:", present)
	// => AFTER calling changeProductByPointer: {300 Bicycle}

	// declaring a slice
	prices := []int{10, 20, 30}

	// When a function changes a slice or a map the actual data is changed.
	changeSlice(prices)
	fmt.Println("prices slice after calling changeSlice():", prices)
	// => prices slice after calling changeSlice(): [11 21 31]

	// declaring a map
	myMap := map[string]int{"a": 1, "b": 2}
	// When a function changes a slice or a map the actual data is changed.

	changeMap(myMap)
	fmt.Println("myMap after calling changeMap():", myMap)
	// => myMap after calling changeMap(): map[a:10 b:20 x:30

}

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

// declaring struct type
type Product struct {
	price       float64
	productName string
}

// declaring a function that takes in a struct value and modifies it
func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
	// the changes are not seen to the outside world
}

// declaring a function that takes in a pointer to struct value and modifies the value
func changeProductByPointer(p *Product) {
	(*p).price = 300
	p.productName = "Bicycle"
	// the changes are seen to the outside world

}

// declaring a function that takes in a slice
func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
	// the changes are seen to the outside world
}

// declaring a function that takes in a map
func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["x"] = 30
	// the changes are seen to the outside world
}
