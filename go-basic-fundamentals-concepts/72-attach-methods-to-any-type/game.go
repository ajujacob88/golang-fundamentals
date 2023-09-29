package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) print() {

	fmt.Println("Title:", g.title, "price:", g.price)
}
