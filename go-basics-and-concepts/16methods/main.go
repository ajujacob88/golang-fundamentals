package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in go")
	//if we write function in structs then we can call it as methods.
	//ref: https://youtu.be/GhYIKwMxz_Y

	//calling a struct/object
	aju := Useraju{"Aju", "aju@gm.com", false, 34}
	fmt.Println(aju)
	fmt.Printf("Aju details are: %v \n", aju)
	fmt.Printf("Aju details are with +v syntax: %+v \n", aju)

	fmt.Printf("Name is %v and email is %v\n", aju.Name, aju.Email)

	aju.GetStatus()
	aju.NewMail()
	fmt.Printf("Now checking wether NewMail changed the original email or not: Name is %v and email is %v\n", aju.Name, aju.Email)
}

//define a struct

type Useraju struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//creating a function as method

func (u Useraju) GetStatus() { //we can pass an entire struct or things inside struct only, here iam passing entire struct
	fmt.Println("Is user active: ", u.Status)
}

//now doing something similar to getter  ans setter

func (u Useraju) NewMail() {
	u.Email = "aju@go.dev" //iam going to manipulate the property
	fmt.Println("Email of this user is: ", u.Email)
	//this clears that, whenever we are pasing the objects/structs it actually passes on a copy only,,, so here comes the concept of pointers,,if we pass on the original values, not copy, then use pointers.
	//that passing with pointers will be dealt up later.
}
