package main

import "fmt"

func main() {
	var income int
	fmt.Println("Enter the annual income:")
	fmt.Scan(&income)
	if income <= 250000 {
		fmt.Println("No tax to pay")
	} else if income > 250000 && income <= 500000 {
		tax := (income * 5) / 100
		fmt.Println("Income tax Amount = ", tax)
	} else if income > 500000 && income <= 1000000 {
		tax := (income * 10) / 100
		fmt.Println("Income tax Amount = ", tax)
	} else if income > 100000 && income <= 5000000 {
		tax := (income * 30) / 100
		fmt.Println("Income tax Amount = ", tax)
	} else {
		fmt.Println("Contact income tax dept.")
	}

}
