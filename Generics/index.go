package main

import "fmt"

//Genrics
// allows devs to write functions, types, and data structures
//  that work with any data type without sacrificing type safety.

// Before generics, Go programmers often used:
// interface{} (empty interface) + type assertions (not type-safe)
// Code duplication (writing the same logic for different types)

func printNumber(item, defaultValue int) (int, int) {
	return item, defaultValue
}

func printString(item, defaultValue string) (string, string) {
	return item, defaultValue
}

//same function but one is printing int data type and other is printing string
// we dont wanna repeat same code just because we want to send diff data to same function
//hence we use generics

func returnItem[T any](item, defaultValue T) (T, T) {
	return item, defaultValue
}
func main() {
	num1, num2 := printNumber(10, 20)
	str3, str4 := printString("one", "two")
	fmt.Println(num1, num2)
	fmt.Println(str3, str4)

	// item1, item2 := returnItem("one", "one")
	// fmt.Println(item1, item2)
	item1, item2 := returnItem(1, 2)
	fmt.Println(item1, item2)
}
