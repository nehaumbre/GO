package main

import "fmt"

func main() {
	// If statement
	num := 50

	if num > 5 {
		fmt.Println("num is greater than 5")
	} else if num < 5 {
		fmt.Println("num is less than 5")
	} else {
		fmt.Println("num is equal to 5")
	}

	password := ""
	if len(password) > 7 {
		fmt.Println("Password is strong")
	} else if password == "" {
		fmt.Println("Password is not entered. Please provide password!")
	} else {
		fmt.Println("Password is weak")
	}
}
