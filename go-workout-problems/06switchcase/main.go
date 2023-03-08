package main

import "fmt"

func main() {
onceAgain:
	fmt.Println("Enter the input no: 1-Sun, 2-Mon, 3-Tuesday, 4-Wednesday, 5-Thursday, 6-Friday, 7-Saturday")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid Entry, Enter the correct code once again")
		goto onceAgain
	}
}
