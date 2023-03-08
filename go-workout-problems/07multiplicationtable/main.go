package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number")
	var input int
	fmt.Scan(&input)
	fmt.Println("The multiplication table of ", input, " is: ")

	for i := 1; i <= 10; i++ {
		fmt.Println(i, " * ", input, " = ", i*input)
	}

	//other way of printing
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v \n", i, input, i*input)
	}

}
