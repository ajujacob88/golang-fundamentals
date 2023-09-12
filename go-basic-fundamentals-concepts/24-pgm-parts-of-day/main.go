package main

import (
	"fmt"
	"os"
	"time"
)

//use switch statements instead of longer if statements to make the code more readable

// task is to check thecurrent time and print the greetings according to that
func main() {
	t := time.Now()
	hour := t.Hour()

	fmt.Println(hour)

	//but we can directly get h like below

	h := time.Now().Hour()
	fmt.Println(h)
	switch {
	case h >= 18:
		fmt.Println("Good Evening")
	case h >= 12:
		fmt.Println("Good afternoon")
	case h > 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}

	//or we can use short switch like below
	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("Good Evening")
	case h >= 12:
		fmt.Println("Good afternoon")
	case h > 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}

	//Another pgm

	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Println("Not a month")

	}

}
