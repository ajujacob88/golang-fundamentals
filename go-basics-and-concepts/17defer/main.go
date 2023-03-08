package main

import "fmt"

func main() {
	//ref: https://youtu.be/jiy584-vv38

	fmt.Println("This is defer, line which use the keyword defer will be executed lastly")
	defer fmt.Println("Hello") //using defer means imagine that line is written just before the } in main, which means it will be executed as the last line in the program
	fmt.Println("world")
	fmt.Println("This is go programming")

	defer fmt.Println("one") // if we use multiple defere, then go will use the concept first in last out or in other words last in(defer) first out
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
