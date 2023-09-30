package main

import "fmt"

type game struct {
	title string
	//price float64
	price money
}

func (g *game) print() {

	//fmt.Println("Title:", g.title, "price:", g.price)
	fmt.Println("Title:", g.title, "price:", g.price.dollarstring())
}

func (g *game) discount(d float64) {
	g.price = (g.price) - (g.price * money(d) / 100)
	return
}
