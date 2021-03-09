package main

import "fmt"

type Grades struct {
	grade  int
	course string
}

func main() {
	fmt.Println("Exercise 1")
	type person struct {
		name      string
		age       int
		colorList []string
		grades    Grades
	}
	me := person{name: "Zoheb", age: 32, colorList: []string{"Black", "Red", "Orange", "Green"}}
	you := person{name: "john", age: 33, colorList: []string{"Crimson", "Orange", "Blue"}}
	fmt.Printf("%+v\n%+v\n", me, you)
	fmt.Println("====================")

	fmt.Println("Exercise 2")
	me.name = "Andrei"
	youColors := you.colorList
	fmt.Printf("John's Colors: %v, Type: %T", youColors, youColors)
	you.colorList = append(you.colorList, "Black")
	fmt.Println("====================")

	fmt.Println("Exercise 3")
	for _, v := range me.colorList {
		fmt.Println(v)
	}
	fmt.Println("====================")

	fmt.Println("Exercise 4")

	me.grades = Grades{grade: 100, course: "calculus"}
	you.grades = Grades{grade: 100, course: "writing"}
	me.grades = Grades{grade: 98, course: "Golang"}
	fmt.Printf("%+v\n%+v\n", me, you)
	fmt.Println("====================")
}
