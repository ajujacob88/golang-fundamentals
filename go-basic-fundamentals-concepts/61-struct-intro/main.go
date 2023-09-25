package main

import "fmt"

func main() {
	// var aju struct {
	// 	firstName, lastName string
	// 	age                 int
	// }
	// fmt.Printf("%+v\n", aju)
	// aju.firstName = "Anu"
	// aju.lastName = "Ja"
	// aju.age = 52

	// fmt.Println(aju)

	// fmt.Printf("%+v\n", aju)

	// var anu struct {
	// 	firstName, lastName string
	// 	age                 int
	// }
	// fmt.Printf("%+v\n", anu)

	type person struct {
		firstName, lastName string
		age                 int
	}

	var rahul person
	fmt.Printf("%+v\n", rahul)

	rahul.firstName = "Rahul"
	rahul.lastName = "aasda"
	rahul.age = 35
	fmt.Printf("%+v\n", rahul)

	nik := person{
		firstName: "Nik",
		lastName:  "asdsa",
		age:       25,
	}
	fmt.Printf("%+v\n", nik)

	fmt.Println(rahul == nik)

	nik2 := nik
	fmt.Println(nik == nik2)

	nik2.firstName = "NIKK2"
	fmt.Println(nik)
	fmt.Println(nik2)
}
