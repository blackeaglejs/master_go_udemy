package main

import (
	"fmt"
	"math"
)

// Any type that implements these methods can work with this interface.
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	interfaceBasics()
	interfaceDynamicTypes()
	typeAssertions()
	embeddedInterfaces()
	emptyInterfaces()
}

func interfaceBasics() {
	// An interfaces is a collection of method signatures that an object can implement.
	// They are not generic types like they are in other languages.

	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 2.1}

	// Interfaces are a contract betwen methods (?)

	// the signature of a method is its name, input, and return values

	// Here's an interface at work. We call the shape interface print() method on our circle and rectangle from before.
	print(c1)
	print(r1)

	a := c1.area()
	fmt.Println("Circle Area:", a)
}

func interfaceDynamicTypes() {
	var s shape // Does not have any methods or values associated iwth it. calling a method on it wil throw an error.
	fmt.Printf("%T\n", s)

	// concrete types, however, can call interface types.
	// an interface holds a value of a dynamic type.
	ball := circle{radius: 2.5}
	s = ball

	print(s)
	fmt.Printf("Type of s: %T\n", s) // type would be circle.

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s) // type would be rectangle.

	// Interfaces hold references to the unederlying types.

}

func typeAssertions() {
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s) //interface dynamic type is circle
	// If you call a method that belows to an underlying type for an interface,
	// that method will only be callable from the interface if you add that method to the interface.
	// s.volume() -> error - fails because interface shape does not implement circle method volume.

	// Otherwise, you can call the underlying type using an assertion and then the type method.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// ball is the circle object, ok is the assertion check - if false, assertion failed.
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

}

func embeddedInterfaces() {
	// An interface cannot extend or implement or another interface - but it can embed another interface.
	// when we add an interface to an interface, we get the methods from the embedded interface.
	// circular embedding of interfaces is not allowed. that's when you embed interface x in interface y, then embed interface y in interface x.
}

func emptyInterfaces() {
	// declaration looks like this
	type emptyInterface interface{}

	// or in a struct
	type person struct {
		info interface{}
	}

	// can hold any value.
	var empty interface{}
	empty = 5
	fmt.Println(empty) // => 5

	// storing a string in the empty interface
	empty = "Go"
	fmt.Println(empty) // => Go

	// storing a slice in the empty interface
	empty = []int{2, 34, 4}
	fmt.Println(empty) // => [2 34 4]

	you := person{}

	// assigning any value to empty interface field
	you.info = "You name"
	you.info = 40
	you.info = []float64{4.5, 6., 8.1}

	fmt.Println(you.info)

	// very important because it's used everywhere in golang.
	// the fmt.Println() function actually is a variadic function that implements as many empty interfaces as you want.
	// in order to call a method on an empty interface though, you need to impose a type.

	// useful when you're working with data of unknown type.

	// The downside is that it hurts maintanability. Empty interfaces mean that you can bypass type-checking, which is against the central tenets of golang.
	// Golang prefers explicit interfaces or declared types for code stability.
}
