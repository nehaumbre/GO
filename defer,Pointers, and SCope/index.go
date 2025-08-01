package main

import (
	"fmt"
)

// Package Scope:
// Identifiers declared outside of functions,
// typically at the top level of a .go file, have package scope.
// They are accessible to all functions within the same package.
var PS = "GO-PAckage Scope" //PAckage scope accessible anywhere see line 45

func hello() {
	msg := "hello iam a function scope" //function scope
	fmt.Println(msg)
}

func main() {
	fmt.Println("HELLO")
	defer fmt.Println("hello")
	fmt.Println("world")
	//"world" is printed after main() finishes everything else
	// Defer follows LIFO (Last In, First Out).

	name := "Alice"
	defer fmt.Println("Hello", name)
	name = "Bob"
	// ➡️ defer evaluates arguments at the moment the defer is called, not when it's executed.

	//scope refers to the region of the code where a variable, constant, type, or function is accessible.
	// Go uses lexical (static) scoping, which means that scope is determined by the structure of the code blocks.

	// Block Scope
	// A block is defined by {}.
	//  Variables declared in a block
	// (such as if, for, switch, etc.) are only visible within that block.
	if true {
		x := 10
		fmt.Println(x) //block scope
	}
	//fmt.Println(x) // Error : undefined: x :x is not visible here

	// Function Scope (Local Scope)
	//Identifiers declared within a function are local to that function.
	//They cannot be accessed outside the function.
	hello()
	fmt.Println(PS)
	// fmt.Println(greet.Greeting)
}

// defer is a keyword used to postpone the execution of a function until the surrounding function returns.
// It's commonly used for:
// Releasing resources (e.g., closing files)
// Unlocking mutexes
// Logging exit points
// Ensuring cleanup happens, even during a panic

// 4. File Scope (via init function)
// init() functions can access all package-level variables
// and are automatically called before main().

// 5. Global Scope (Across Packages)
// If a variable, constant, or function name starts with an uppercase letter,
//  it is exported and can be accessed from other packages (similar to public in other languages).
