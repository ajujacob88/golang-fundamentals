package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter a country in console")
		return
	}
	args := os.Args[1]

	switch args {
	case "India", "Bharath": //we can add multiple conditions via comma
		fmt.Println("My country")
		vip := true
		fmt.Println("vip availability", vip)
	case "UK":
		fmt.Println("United")
		fmt.Println("vip not available here, because its declared in previous block")
	default:
		fmt.Println("enter correctly,, default is not necessary")
	}

	n := 250
	//switch true {  //no need of the true value here, automatically true, because true is default, but for false, type false

	switch {
	case n > 100:
		fmt.Println("big positive number")
	case n > 0:
		fmt.Println("positive number")
	case n < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
