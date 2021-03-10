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
