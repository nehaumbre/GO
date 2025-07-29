package main

import "fmt"

func main() {
	//for loops
	for i := 0; i < 5; i++ {
		fmt.Println("Hello World", i)
	}
	fmt.Println("--------------------------------------")
	for j := 1; j < 10; j++ {
		if j == 5 {
			continue //5 wont get printed
		} else {
			fmt.Println("Hello World", j)
		}
	}
	fmt.Println("--------------------------------------")
	for j := 1; j < 10; j++ {
		if j == 5 {
			break //execution stops at 5
		} else {
			fmt.Println("Hello World", j)
		}
	}

	//nested loops
	fmt.Println("-----------------nested loops---------------------")
	for outer := 0; outer < 3; outer++ {
		for inner := 0; inner < 3; inner++ {
			fmt.Println("Outer:", outer, "Inner:", inner)
		}
	}

	//while loop: theres no while loops in go but u can achieve it with for loop
	x := 1

	for x <= 6 {
		fmt.Println(x)
		x++
	}
}
