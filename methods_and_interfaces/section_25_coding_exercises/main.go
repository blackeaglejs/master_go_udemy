package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("$%.2f\n", m)
}

func (m money) printStr() string {
	shortened, _ := fmt.Printf("$%.2f\n", m)
	return string(shortened)
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * .09
}

func (b *book) discount() {
	(*b).price = b.price * .9
}

func (b *book) changePrice(p float64) {
	(*b).price = p
}

func main() {
	methodsExercises()
	interfacesExercises()
}

func methodsExercises() {
	// Exercise 1
	var blah money = 2.03245
	blah.print()

	// Exercise 2
	blah.printStr()

	// Exercise 3
	ihygtm := book{price: 10.99, title: "I Hope You Get This Message"}
	fmt.Printf("VAT of %s is $%.2f\n", ihygtm.title, ihygtm.vat())
	ihygtm.discount()
	fmt.Printf("Discounted price of %s is $%.2f\n", ihygtm.title, ihygtm.price)

	// Exercise 4
	ihygtm.changePrice(10.99)
	fmt.Printf("Price of %s is $%.2f\n", ihygtm.title, ihygtm.price)
}

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func interfacesExercises() {
	// Exercise 1
	var sportsCar vehicle
	s2K := car{licenceNo: "9000", brand: "Honda"}
	sportsCar = s2K

	fmt.Println(sportsCar.License())
	fmt.Println(sportsCar.Name())

	// Exercise 2
	var empty interface{}
	fmt.Printf("%T\n", empty)
	empty = 5
	fmt.Printf("%T\n", empty)
	empty = 5.5
	fmt.Printf("%T\n", empty)
	empty = []int{}
	empty = append(empty.([]int), 5)
	fmt.Println(empty)

	// Exercise 3
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}
