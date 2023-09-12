package main

import (
	"fmt"
)

//fallthorugh executesthe next clause without checking for its condition

func main() {

	switch n := 250; { //short switch, decalring n together with switch statement, but need to use ; also.. actually true is there,,but no need to type true - switch n := 250;true
	case n > 200:
		fmt.Println("big ")
		fallthrough
	case n > 0:
		fmt.Println("positive ")
		fallthrough
	case n < 0:
		fmt.Println("negative ")
		fallthrough
	default:
		fmt.Println("number")
		fallthrough
	case n == 20:
		fmt.Println("hi")
	}

}
