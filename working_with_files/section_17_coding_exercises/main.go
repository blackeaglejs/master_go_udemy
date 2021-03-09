package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
	exerciseFour()
	exerciseFive()
	exerciseSix()
}

func exerciseOne() {
	fmt.Println("Exercise 1")
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	fmt.Println("====================")
}
func exerciseTwo() {
	fmt.Println("Exercise 2")
	err := os.Rename("info.txt", "data.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File Does Not Exist")
		}
	}
	fmt.Println("====================")
}
func exerciseThree() {
	fmt.Println("Exercise 3")
	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("====================")
}
func exerciseFour() {
	fmt.Println("Exercise 4")
	file, err := os.OpenFile("info.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("blah blah blah\n zoheb")
	fileAsBytes, err := ioutil.ReadFile("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileAsString := fmt.Sprintf("%s", fileAsBytes)
	fmt.Println(fileAsString)
	fmt.Println("====================")
}
func exerciseFive() {
	fmt.Println("Exercise 5")
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println("====================")
}
func exerciseSix() {
	fmt.Println("Exercise 6")
	file, err := os.OpenFile("info.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("The Go gopher is an iconic mascot!")
	fmt.Println("====================")
}
