package main

import "fmt"

type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("EMPTY SLICE")
	}
	for _, it := range l {
		it.print()
	}
}

type booklist []*book

func (l booklist) print() {
	if len(l) == 0 {
		fmt.Println("EMPTY SLICE")
	}
	for _, it := range l {
		it.print()
	}
}

func (l booklist) discount() {
	for _, it := range l {
		it.discount(10)
	}
}
