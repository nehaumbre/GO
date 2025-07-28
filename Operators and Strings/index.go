package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var a, b = 3, 3

func main() {
	//Arithmetic operators
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	//Increment operator
	x := 10
	x++
	fmt.Println(x)

	//decrement operator
	y := 10
	y--
	fmt.Println(y)

	//assignment operators
	t := 6
	fmt.Println(t)
	t += 4
	fmt.Println(t)
	t -= 2
	fmt.Println(t)
	t *= 8
	fmt.Println(t)
	t /= 1
	fmt.Println(t)
	t = 8
	t %= 3
	fmt.Println(t)

	//strings
	name := "John Doe"
	var str string = "Hello, World!"
	fmt.Println(str)
	fmt.Println(name)

	//multiline srings can be in ``
	fmt.Println(`My name is 
	            John `)

	// string concatenation
	intro := "hi " + "My name is" + " " + "Camila Doe"
	fmt.Println(intro)

	//getting character from strings
	episode := "Episode 01:"
	fmt.Println(episode[0]) // 69 is output here
	//* useful link  for fmt package: https://pkg.go.dev/fmt
	fmt.Printf("%c", episode[9])

	//length of a string
	fruit := "lychees"
	fmt.Println(len(fruit))                    //17 : not correct
	fmt.Println(utf8.RuneCountInString(fruit)) // 7 : correct

	//comparing strings
	msg := "one"
	msg2 := "one"
	msg3 := "two"
	fmt.Println(strings.Compare(msg, msg2)) // 0 : strings are equal
	fmt.Println(strings.Compare(msg, msg3)) // -1 : msg is less than msg3

	//Check if string contains any character
	stringOne := "GO Programming Language"
	//use of strings.Contains() function
	fmt.Println(strings.Contains(stringOne, "GO"))   // true
	fmt.Println(strings.Contains(stringOne, "GOOz")) //false
	//use of strings.ContainsAny() function
	fmt.Println(strings.ContainsAny(stringOne, "GO")) // true
	fmt.Println(strings.ContainsAny(stringOne, "PO")) //true
	fmt.Println(strings.ContainsAny(stringOne, "$$")) //false

	//use of strings.ContainsRune() function
	fmt.Println(strings.ContainsRune(stringOne, 'G')) // true
	fmt.Println(strings.ContainsRune(stringOne, 'b')) //false

	// to uper and lower cases
	leWords := "HELLO World"
	fmt.Println(strings.ToLower(leWords)) //hello world
	fmt.Println(strings.ToUpper(leWords)) //HELLO WORLD

	//Comparison Operators
	// == : Equal to
	// != : Not Equal to
	// < : Less than
	// > : Greater than
	// <= : Less than or Equal to
	// >= : Greater than or Equal to
	num1 := 3
	num2 := 7
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)

	//Logical operators
	// && : AND
	// || : OR
	// ! : NOT
	fmt.Println("------------")
	fmt.Println(true && true)
	fmt.Println(true && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

}

// string values must be in double quotes
