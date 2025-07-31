package main

import "fmt"

type Person struct {
	firstName  string
	LastName   string
	age        int
	profession string
	isMarried  bool
}

func (p Person) getInfo() {
	fmt.Println("Name:", p.firstName, "LastName:", p.LastName, "age:", p.age, "profession", p.profession, "isMarried:", p.isMarried)
}

//============================================callback function
func addName(name string, callback func(string)) {
	callback(name)
}

func main() {
	// Methods in Go
	Joe := Person{
		firstName:  "Joe",
		LastName:   "Doe",
		age:        30,
		profession: "Software Engineer",
		isMarried:  false,
	}
	Joe.getInfo()

	//callback functions
	// callback function is a function that is passed as an argument to another function
	// Also called First class function in Go language
	addName("Nea", func(nm string) {
		fmt.Printf("Hi Im %v", nm)
	})
}
