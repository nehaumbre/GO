package main

import "fmt"

func main() {
	//arrays

	var numbers [7]int
	numbers[0] = 10
	fmt.Println(numbers) //[10 0 0 0 0 0 0]

	//or
	number := [3]int{1, 2, 3}
	fmt.Println(number) //[1 2 3]
	fmt.Println(len(numbers))
	fmt.Println(cap(number))

	//array of strings

	arr1 := [3]string{"one", "two", "three"}

	for item := 0; item < len(arr1); item++ {
		fmt.Println(arr1[item])
	}

	//Slices
	//used to store multiple values of same type in a variable
	//length of slice can shrink or grow as you see fit
	//3 ways to create slices
	//1. using make function

	days := make([]string, 5, 20)
	days = append(days, "Monday", "Thursday") //adds new data to end
	days[0] = "Monday"                        //adds data to empty existing areas
	days[1] = "Thursday"
	days[2] = "friday"
	fmt.Println(days)
	//2.creating slice from an array

	//3. using []datatype{values} format
	myslice1 := []int{}
	fmt.Println(myslice1)
	myslice2 := []string{"Go", "Slices", "are", "powerful"}
	fmt.Println(myslice2)
	fmt.Println(myslice2[1:2])
	myslice2 = append(myslice2, "and", "useful")
	fmt.Println(myslice2)

	//another way of printing values
	for index, value := range myslice2 {
		fmt.Println(index, value)
	}

	for index, value := range days {
		fmt.Println(index, value)
	}
}
