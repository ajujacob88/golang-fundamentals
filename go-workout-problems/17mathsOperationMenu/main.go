package main

import (
	"fmt"
	"os"
)

//just used structs

func main() {
	var num1, num2, operation float64
	var repeat string
	fmt.Println("Enter the 2 numbers")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	enteredValues := mathsOperation{num1, num2}
onceAgain:
	fmt.Println("Enter the code for the operations to perform: \n 1 - Addition\n 2 - Substraction\n 3 - Multiplication\n 4 - Division\n 5 - Exit from the program")
	fmt.Scan(&operation)

	switch operation {
	case 1:
		enteredValues.addition()
	case 2:
		enteredValues.substraction()
	case 3:
		enteredValues.multiplication()
	case 4:
		enteredValues.division()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Wrong operation Entered")
	}

	fmt.Println("Press # to perform another operation or press any other key to to exit")
	fmt.Scan(&repeat)
	if repeat == "#" {
		goto onceAgain
	} else {
		os.Exit(0)
	}
}

type mathsOperation struct {
	value1 float64
	value2 float64
}

func (numbers mathsOperation) addition() {
	sum := numbers.value1 + numbers.value2
	fmt.Println("The result is ", sum)
}

func (numbers mathsOperation) substraction() {
	sub := numbers.value1 - numbers.value2
	fmt.Println("The result is ", sub)
}

func (numbers mathsOperation) multiplication() {
	multiply := numbers.value1 * numbers.value2
	fmt.Println("The result is ", multiply)
}

func (numbers mathsOperation) division() {
	divide := numbers.value1 / numbers.value2
	fmt.Println("The result is ", divide)
}

/*
package main

import (
	"fmt"
	"os"
)

//This is just used structs, other way
func main() {
	fmt.Println("Enter the numbers:")
	var numb1, numb2 int
	fmt.Scan(&numb1)
	fmt.Scan(&numb2)
	numbers := mathsOperation{numb1, numb2, func(numb1 int, numb2 int) int { return numb1 + numb2 }}
	fmt.Println(numbers.sum(numbers.num1, numbers.num2))

}

type mathsOperation struct {
	num1 int
	num2 int
	sum  addition
}
type addition func(int, int) int

/*
func (numbers mathsOperation) addition() {
	sum := numbers.num1 + numbers.num2
	fmt.Println(sum)
}
*/
