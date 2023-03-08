package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Reference video - Hitesh Choudary: https://youtu.be/zYIZtbyUIDY

	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	//comma ok syntax /error ok syntax  / err err / alternate for try catch error handling, in go try catch is not there

	input, _ := reader.ReadString('\n')      //while reading an input if an error arises, we can catch or store that error by using a variable istead of _ there, but if we dont need that error value then we can simply put underscore _ ,,, this is similar to try catch, we can say the first part(input) is try and the _ part is catch
	fmt.Println("Thanks for rating ", input) //\n means input will exit if an enter key is pressed

	fmt.Printf("Type of this rating is %T \n", input)

	fmt.Print("Enter another value (to study error): ")
	input1, err1 := reader.ReadString('\n') //here if an error arises, the value of erroe is stored in err1,
	fmt.Println("The input is ", input1, "and error is ", err1)

	fmt.Println("enter a value, this time if 'a' is entered instead of enter key, then exited")
	input2, _ := reader.ReadString('a')
	fmt.Println("ENtered value is ", input2)

	fmt.Println("If we need only the value of error(if arises), then we can simply put")
	_, erraju := reader.ReadString('\n') // but this is not the standard syntax
	fmt.Println("value of error if any is ", erraju)

}
