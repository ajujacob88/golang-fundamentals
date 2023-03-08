package main

import "fmt"

func main() {

	//ref: https://youtu.be/-BFJ0dZyxHg
	fmt.Println("Welcome to pointers")

	var myptr *int //here a pointer is just ccreated
	fmt.Println("Default value of pointer is: ", myptr)

	myNumber := 23
	var ptr = &myNumber // here a pointer is created which is also referencing to some memory. so whenever a referencing is there use &
	fmt.Println("Value of actual pointer(memory location of my no) is: ", ptr)
	fmt.Println("Value of actual pointer(value inside ptr) is: ", *ptr)

	*ptr = *ptr + 8
	fmt.Println("New value is: ", myNumber)
	fmt.Println("New value is: ", ptr)
	fmt.Println("New value is: ", *ptr)

	//So pointers makes it a guarantee, that no matter what you are doing the operation with those values, those operations are actually performed on the values, not on the copies of those values.

}
