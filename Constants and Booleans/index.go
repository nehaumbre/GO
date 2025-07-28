package main

import "fmt"

const pi = 3.14

const (
	a = 2
	b = 3
)

func main() {
	isHTMLPL := false
	isGOPL := true
	fmt.Println(pi)
	fmt.Println(a * b)
	fmt.Println(isGOPL)
	fmt.Println(isHTMLPL)
}

// Constants have values that cannot be changed
// const PI = 3.14
// write constants in UPPERCASE letters as shown above
// can be declared both inside or outside of a function
// Booleans : true and false
