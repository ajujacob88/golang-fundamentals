package main

import "fmt"

type book struct {
	title string
	price money
}

func (b *book) print() { //Here (b book) is called the receiver...the method automatically receives the value of a type

	//fmt.Println("Titlee:", b.title, "price:", b.price)
	fmt.Println("Titlee:", b.title, "price:", b.price.dollarstring())
}

func (b *book) discount(d float64) {
	b.price = (b.price) - (b.price * money(d) / 100)
	return
}

//while using a type as pointer , it is a convention that writing all the methods of that type as pointer,,, if a type is passing as values, then all the methods of that type should be written like passing values... its a convention
