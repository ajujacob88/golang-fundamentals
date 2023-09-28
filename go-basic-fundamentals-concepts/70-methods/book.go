package main

import "fmt"

type book struct {
	title string
	price float64
}

// func printBook(b book) {

// 	fmt.Println("Title:", b.title, "price:", b.price)
// }

//converting the above functional style of programming to methods style

// func (b book) printBook() { //Here (b book) is called the receiver...the method automatically receives the value of a type, b is the receiver variable name and book is the type

// 	fmt.Println("Title:", b.title, "price:", b.price)
// }

//also we can have the same function name for 2 methods of different types, so we can save rename this as well as game as print ,, no need of printbook printgame, because its understood that it is book or game from the method, also, we cna have same name

func (b book) print() { //Here (b book) is called the receiver...the method automatically receives the value of a type

	fmt.Println("Title:", b.title, "price:", b.price)
}

func (b book) discount(d float64) (dp float64) {
	dp = (b.price) - (b.price * d / 100)
	return
}
