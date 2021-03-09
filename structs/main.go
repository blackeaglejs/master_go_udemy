package main

import "fmt"

func main() {
	// Structs are a sequence of named elements called typed.
	// Can be thought of a class.
	// Struct fields are like instance attributes.
	// Fields can only be changed at compiletime, not at runtime.
	// Golang uses structs as a blueprint - they are not really that much like a class-object structure.

	title1, author1, year1 := "The Divine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)
	// The above is disorganized and doesn't have relationships etween the data.
	// So lets create a struct

	type book struct {
		title  string
		author string
		year   int
	}

	// This is also valid.
	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Divine Comedy", "Dante Aligheri", 1320}                 // order matters. This is not a recommended way to create a struct.
	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945} // order no longer matter. this is recommend.
	fmt.Println(myBook)
	fmt.Println(bestBook)

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)
	// Fields that are not defined will throw an error in golang.
	// It will use the blank fields when they are not supplied.
	lastBook.year = 1878
	lastBook.author = "Leo Tolstoy"
	fmt.Println(lastBook)

	// STructs can be compared if they are of hte same type. It can compare field by field.
	fmt.Println(myBook == lastBook) // returns false

	// copies of structs have their own memory address
	aBook := myBook
	aBook.year = 2021
	fmt.Println(aBook, myBook)

	// anonymous structs
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Miller",
		age:       30,
	}
	// These structs are great when you don't plan on reusing a struct later and only need to use it once.
	fmt.Printf("%#v\n", diana)
	// If you create an anoymous struct and don't specify field names, golang will supply for them for you using the type.

	// You can mix named and unnamed fields in an anoymous struct.
}

func embeddedStruct() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Stream 20, London",
			phone:   00323224241414,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee's emai:%s\n", john.contactInfo.email)

	// The above is an example of embeded structs.
}
