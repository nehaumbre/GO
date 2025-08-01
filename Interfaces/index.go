package main

import "fmt"

// Define an interface 'Animal' with a method 'getInfo' that returns a string
type Animal interface {
	getInfo() string
}

// Define a struct 'Cat' with fields 'name' and 'color'
type Cat struct {
	name  string
	color string
}

// Define a struct 'Dog' with fields 'name' and 'color'
type Dog struct {
	name  string
	color string
}

// Implement the 'getInfo' method for 'Cat' type
// This satisfies the 'Animal' interface for 'Cat'
func (c Cat) getInfo() string {
	return "Cat: " + c.name + ", " + c.color
}

// Implement the 'getInfo' method for 'Dog' type
// This satisfies the 'Animal' interface for 'Dog'
func (d Dog) getInfo() string {
	return "Dog: " + d.name + "," + d.color
}

// A function that takes an Animal interface and prints the information
func printAnimalInfo(animal Animal) {
	fmt.Println(animal.getInfo())
}

// The main function where execution begins
func main() {
	// Create an instance of Cat
	cat1 := Cat{
		name:  "Tom",
		color: "Black",
	}

	// Create an instance of Dog
	dog1 := Dog{
		name:  "Jack",
		color: "White",
	}

	// Pass the Cat and Dog instances to printAnimalInfo
	printAnimalInfo(cat1)
	printAnimalInfo(dog1)
}
