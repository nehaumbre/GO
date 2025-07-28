package main

import "fmt"

var number = 20

func main() {
	var a = "firstvar"
	b := 3
	var c, d, e int = 4, 5, 6
	var (
		greetings string = "hello"
		num       int    = 1
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d, e)
	fmt.Println(greetings)
	fmt.Println(num)
	fmt.Println(number)
}
