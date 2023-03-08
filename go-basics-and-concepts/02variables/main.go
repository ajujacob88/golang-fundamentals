package main

import "fmt"

var myname2global string = "we can declare anything outside of any methods(global) also, but using walrus declaration(:=) is not possible outside of methods"
var ajuj string = "this is also allowed"
var no = 25

const example string = "this is constant and cant be changed"

//variable name first character captal(here E is capital in ExamplePublic) is having significant importance in golang. reason for that is this is now a public variable. this is almost equivilant to' Public const example string= "aju" ' in java. ie its equivqlent of putting public before that. in golang we make a public variable by putting the first letter as capital.
// so this ExamplePublic is accessible by any other file in this program and we can use this anywhere
const ExamplePublic string = "this is public because first letter E is capital"

func main() {
	fmt.Println(myname2global)

	var myname string = "Aju Jacob"
	fmt.Println(myname)
	fmt.Print("variable ", myname, " is of type: %T \n", myname)
	fmt.Printf("variable myname is of type: %T \n", myname)

	var isLoggedIn bool = true //for boolean its good convention to save the variable name starting as 'is loggdin or is verified etc..
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255 // refer the language speciification  for the value range and types, but usally we use int... https://go.dev/ref/spec#Numeric_types
	fmt.Println(smallValue)
	fmt.Printf("variable is of type: %T \n", smallValue)

	var value int = 12548
	fmt.Println(value)
	fmt.Printf("variable is of type: %T \n", value)

	var smallFloat float32 = 12548.1234567890123456789
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 12548.1234567890123456789
	fmt.Println(largeFloat)
	fmt.Printf("variable is of type: %T \n", largeFloat)

	// default values and some aliases

	var anotherVariable int //the default value will be always zero while declaring a thing, no garbage values are added, and this is a good thing in go.
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var anotherString string //the default value will be always a blank for string while declaring a thing, no garbage values are added, and this is a good thing in go.
	fmt.Println(anotherString)
	fmt.Printf("variable is of type: %T \n", anotherString)

	//another way of declaring variables -  implicit type or implicit way of declaring the variables.

	var website = "ajujacob" //actualy we need to declare the type of var also, but if we not filled, then lexa will automatically filled it for you based on the values you entered, so here it is treated as string.
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)
	website = "aju" //now if we try to insert an int value, then it will gives error also..

	// no var style  -- here ignowring the keyword war
	// by using walrus operator (:=), we can avoid the keyword var and also data type, and the pgm will works fine. This is fine,and go accepts is a completely fine, and this is the most common syntax and the easiest syntax
	// but walrus operator can be used only inside of any methods.
	numberOfUser := 3000
	aju2 := 256.3
	aju3 := "mynamee"
	fmt.Println(numberOfUser)
	fmt.Println(aju2)
	fmt.Println(aju3)
	fmt.Printf("variable is of type: %T \n", aju2)

	fmt.Println(ExamplePublic)
	fmt.Printf("variable is of type: %T \n", ExamplePublic)

}
