package main

import (
	"fmt"
	"time"
)

func greet() {
	println("Hello, from Goroutine!")
}
func main() {
	//Goroutines in Go – Lightweight Concurrency
	//Goroutines are functions or methods that run concurrently with
	//other functions/methods in the same address space.
	//They're lightweight threads managed by the Go runtime, not the OS.

	//You start a goroutine with the go keyword:
	//go someFunction()
	//This launches someFunction() asynchronously, allowing your program to
	//  continue executing other code without waiting for it.

	go greet() //this runs in seperate goroutine
	fmt.Println("MAin function continues..")
	time.Sleep(1 * time.Second) //gives go routine time to finish

	// ⚠️ Without the time.Sleep,
	//  the program might exit before the goroutine runs.
	// Go doesn’t wait for goroutines to finish
	//  unless you tell it to (e.g., with sync.WaitGroup).

}
