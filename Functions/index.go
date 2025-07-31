package main

import "fmt"

// Simple function
func greet(name string) {
	fmt.Println("Name", name)
}

//function with multi arguments
func getUsers(user ...string) {
	fmt.Println(user)
}

//function but also add a diffrent data type
func getNums(name string, n ...int) {
	fmt.Println(name, n)
}

//how to iterate inside a function

func getCandidate(candidates ...string) {
	for index, value := range candidates {
		fmt.Println(index, value)
	}
}

//return statements in function
func getReturn(x int, y int) int {
	return x + y
}

func main() {
	greet("John")
	getUsers("John", "Jane", "Bob")
	getNums("John", 1, 2, 3, 4, 5)
	getCandidate("John", "Jane", "Bob")
	//return statement
	fmt.Println(getReturn(5, 6))
	//function expression : when we store a function inside a variable
	add := func(a int, b int) int {
		return a + b
	}
	//function expression in use
	res := add(10, 10)
	fmt.Println(res)

	//anonymous function
	//anonymous function is a function without a name
	//it is used when we need to pass a function as an argument to another function
	//or when we need to return a function from a function
	//anonymous function is also known as lambda function
	func() {
		fmt.Println("This is an anonymous functions")
	}() // () means executing the anonymous function right away
}
