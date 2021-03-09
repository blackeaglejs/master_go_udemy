package main

import (
	"fmt"
	"time"
)

// declare a type
type names []string

// you can declare functions for a type
func (n names) print() {
	// n is a the self. reference from OOP for a type.
	// It's best to choose soemthing short and easiest way.
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	// Go does not have classes, but you CAN assign methods on types you've either imported or defined yourself.
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)

	// call a method on a type
	friends := names{"Dan", "Mary"}
	friends.print()

	// another way to call a method on a type
	names.print(friends) // not very common

	// Difference between methods and functions is that methods are tied to a type, and functions are tied to a package.

	// Methods with Pointers
	methodsWithPointers()
}

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func methodsWithPointers() {
	myCar := car{brand: "Audi", price: 40000}
	// Does not change value because funtion works on copies.
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	// Method call for type also fails because working on copy.
	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar)

	// Using the pointer instead changes the value.
	myCar.changeCar2("Seat", 25000)
	fmt.Println(myCar)

	// pointer types
	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	// valid ways to change
	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar)
	// This will change too
	fmt.Println(myCar)

	// Note that method declarations on pointer types are not allowed.
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
	// the changes are not seen to the outside world (pass by value)
}

// declaring a method with a pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
	// the changes are seen the outside world

}
