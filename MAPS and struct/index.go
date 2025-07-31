package main

import "fmt"

//structs (structures)
//create a collection of members of different data types into a single variable

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

//nested struct
type Programmer struct {
	Person
	isProgrammer bool
}

func main() {
	// MAP : stores data in key value pairs
	// key : unique identifier
	// value : value associated with the key
	// default value of map is nil
	//maps are unordered and changable collection
	//no duplicates allowed
	userinfo := map[string]int{
		"John":  25,
		"Emma":  26,
		"Tom":   27,
		"James": 23,
	}
	fmt.Println(userinfo)
	// Accessing map values
	fmt.Println(userinfo["James"])
	// Adding new key value pair
	userinfo["Kon"] = 24
	fmt.Println(userinfo)
	//changing existing value
	userinfo["John"] = 20
	fmt.Println(userinfo)

	//running for loop on Maps
	for key, value := range userinfo {
		fmt.Println(key, value)
	}

	// creating an instance of strcture

	var userOne Person
	userOne.name = "Joe"
	userOne.age = 30
	userOne.job = "Software Engineer"
	userOne.salary = 50000
	fmt.Println(userOne)
	fmt.Println("My name is", userOne.name, "I'm ", userOne.age, "Years old", ". I am a ", userOne.job, " and my salary is:", userOne.salary)

	// ✅ Initialize using struct literal
	hyx := Person{
		name:   "Hannah",
		age:    28,
		job:    "Doctor",
		salary: 60000,
	}

	fmt.Println(hyx) //{Hannah 28 Doctor 60000}

	var someone Programmer
	someone.name = "Shizu"
	someone.age = 25
	someone.job = "Programmer"
	someone.salary = 40000
	someone.isProgrammer = true

	fmt.Println(someone) // {{Shizu 25 Programmer 40000} true}

	//anonymous struct
	h := struct {
		name string
		age  int
	}{
		name: "ZOe",
		age:  30,
	}
	fmt.Println(h.name, h.age)
}
