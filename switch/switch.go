package main

import "fmt"

func main() {
	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("EWW!! Monday")

	case "Friday":
		fmt.Println("Finaalyy!!")

	case "Saturday", "Sunday":
		fmt.Println("YAY weekend!!")

	default:
		fmt.Println("Meh! Just a regular day.")

	}
}
