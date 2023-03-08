package main

import "fmt"

func main() { //main act as an entry point to the program
	//you are not allowed to write functions inside a function.
	fmt.Println("Welcome to functions in go")
	//ref: https://youtu.be/rcUST3QvVOQ
	greeter()

	result := adder(3, 5, 8)
	fmt.Println("Result is: ", result)

	//sending more nos
	proResult, myMessage := proAdder(2, 5, 8, 7, 6) // also we can use ,ok syntax, if we put _ instead of mymessage, then that value we can omit.
	fmt.Println("Pro Result is: ", proResult)
	fmt.Println("Pro Message is: ", myMessage)

}

func greeter() {
	fmt.Println("Namasthe from golang")
}

func adder(val1 int, val2 int, val3 int) int { //function signatures type to return should be written after function parameters to receive,, we can return and receive as much as values, 2 values, 3 values , any values etc.
	return val1 + val2 + val3

}

//no suppose i have no idea about how many values are come in, but i want to add all of them, so the syntax for that is below
func proAdder(values ...int) (int, string) { //saying so many values are receiving and all values are integers,, now the values receiving are a slice. So ... are the variadic functions and we can expect any values there.
	total := 0
	//so the values receiving are a slice and we need to loop with that

	for _, val := range values {
		total += val
	}
	return total, "Hi pro result"

}
