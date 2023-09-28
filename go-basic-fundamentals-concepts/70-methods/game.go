package main

import "fmt"

type game struct {
	title string
	price float64
}

// func printGame(g game) {

// 	fmt.Println("Title:", g.title, "price:", g.price)
// }

//converting the above functional style of programming to methods style

// func (g game) printGame() {

// 	fmt.Println("Title:", g.title, "price:", g.price)
// }

func (g game) print() {

	fmt.Println("Title:", g.title, "price:", g.price)
}
