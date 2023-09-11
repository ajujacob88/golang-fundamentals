package main

import "fmt"

// constants always have to be initilized to a value
// If a value doesn't change through out the program, use constant
// while using the constants the values are assigned to constant at the compile time, while for variables, the values are assigned during runtime, so constant allows you to catch the error earlier at compile time
func main() {
	const meters int = 100
	cm := 100
	m := cm / meters
	fmt.Printf("%d cm is %d m\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%d cm is %d m\n", cm, m)

	//tot := 5
	//num := 0
	//fmt.Println(tot / num) //this error will catch only at runtime

	const total int = 5
	const numberOf int = 0
	//fmt.Println(total / numberOf)   //but this error will catch during compile time

	const min int = 1
	const min2 = 1
	const min3 = 1 + 2
	const tex = "AJu Jacob"

	fmt.Printf("%T type  %T\n", min, min2)

	//for better readability
	const (
		mini = 1
		max  = 1000
	)

	fmt.Println(mini, max)

	//another feature of constant
	const (
		num1 = 20
		num2 //go automaticallyrepeats the previous constant's type and expression
	)
	fmt.Println(num1, num2)

}
