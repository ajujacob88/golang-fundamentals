package main

import "fmt"

func main() {
	fmt.Println("Structus in golan - Alternative to classes")
	//ref: https://youtu.be/NMTN543WVQY

	//There is NO INHERITANCE in golang; No super or parent or child etc.

	//calling a struct/object
	aju := Useraju{"Aju", "aju@gm.com", false, 34}
	fmt.Println(aju)
	fmt.Printf("Aju details are: %v \n", aju)
	fmt.Printf("Aju details are with +v syntax: %+v \n", aju) //%+v syntax will not only gives the values, but also the structure and naming conversions used there.

	fmt.Printf("Name is %v and email is %v\n", aju.Name, aju.Email) //%+v syntax will not only gives the values, but also the structure and naming conversions used there.

}

//define a struct

type Useraju struct { // U is capital, because it is class/struct/template that need to be exported out
	Name   string //here also N is capital, becaus it can be used by anybody(public).
	Email  string
	Status bool
	Age    int
}
