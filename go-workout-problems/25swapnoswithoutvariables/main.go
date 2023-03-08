package main

//swapping numbers without using third variable

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter the first number")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number")
	fmt.Scan(&num2)
	// num1 = num1 + num2
	// num2 = num1 - num2
	// num1 = num1 - num2
	// fmt.Println("After swapping,\nNumber 1: ", num1, "\nNumber 2: ", num2)

	//This is the way
	num1, num2 = num2, num1
	fmt.Println("After swapping,\nNumber 1: ", num1, "\nNumber 2: ", num2)

}
