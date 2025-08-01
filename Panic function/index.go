package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	fmt.Println("Result", a/b)
}

//another example
func safeAccess() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}
	}()

	arr1 := []int{1, 2, 3}
	fmt.Println("accessing element at index 6:")
	fmt.Println(arr1[6])
}

func main() {
	//Panic in GO:
	//panic is a built-in function that stops execution of the current function
	// and begins unwinding the stack.
	//It can be caught and handled using recover (typically inside a defer).

	// fmt.Println("Start")
	// panic("Something went wrong")
	// fmt.Println("This won't print") //this wont execute

	// When to Use panic?
	// Use panic when:
	// You hit an unrecoverable error (e.g., corrupt state, impossible condition).
	// You want to fail fast.
	// You're writing low-level libraries where errors can't be propagated easily.
	safeDivide(10, 0)
	fmt.Println("after panic")
	fmt.Println("safeAccess Panic code")
	safeAccess()
	fmt.Println("after panic")
}
